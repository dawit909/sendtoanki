package handler

import (
	"archive/zip"
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	_ "modernc.org/sqlite"

	"tuto.sqlc.dev/app/go/constants"
	"tuto.sqlc.dev/app/go/dictparser"
	"tuto.sqlc.dev/app/go/oxforddicthandler"
	"tuto.sqlc.dev/app/go/sendtoanki"
	"tuto.sqlc.dev/app/go/tutorial"
)

// Define a function map to allow "safe" HTML rendering in templates
var funcMap = template.FuncMap{
	"safe": func(s string) template.HTML {
		return template.HTML(s)
	},
}

// Parse templates with the function map attached
var templates = template.Must(template.New("").Funcs(funcMap).ParseFiles("html/upload.html", "html/view.html"))

var processedWords []oxforddicthandler.OxfordWord

// 1. HELPER: Convert slice to map for O(1) lookups
// We do this once so we don't loop through 1000 words for every single check.
var basicWordMap map[string]bool

func init() {
	basicWordMap = make(map[string]bool)
	for _, w := range constants.WIKIPEDIA_1000_BASIC_ENG {
		basicWordMap[w] = true
	}
}

// 2. WRAPPER STRUCT
// This is what the HTML template will actually receive.
type ViewWord struct {
	oxforddicthandler.OxfordWord
	IsBasic bool
}

// Update ViewData to use the wrapper
type ViewData struct {
	Words       []ViewWord
	DownloadURL string
}

// UploadHandler serves the upload page on GET and processes the file on POST.
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := templates.ExecuteTemplate(w, "upload.html", nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error rendering template: %v", err), http.StatusInternalServerError)
		}
		return
	}

	if r.Method == "POST" {
		file, header, err := r.FormFile("vocab.db")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Create a temporary file to store the upload.
		tmpFilePath := filepath.Join(constants.ROOT, "resources", header.Filename)
		tmpFile, err := os.Create(tmpFilePath)
		if err != nil {
			http.Error(w, "Could not save file", http.StatusInternalServerError)
			return
		}
		defer tmpFile.Close()

		_, err = io.Copy(tmpFile, file)
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not save file content: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		// dbName := header.Filename

		// 0. Read db
		rawRows, err := readEveryRowFromDB(tmpFilePath)
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not read database: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		log.Printf("%d rows found from %s\n", len(rawRows), tmpFilePath)

		// 1. Write a stems text file to be shared with python codes
		stemDst := filepath.Join(constants.ROOT, "resources", constants.STEM_FILENAME)
		if err := writeStems(rawRows, stemDst); err != nil {
			http.Error(w, "Could not write stems", http.StatusInternalServerError)
			return
		}

		// 2. create a raw JSON containing definitions of stems
		cmd := exec.Command(filepath.Join(constants.ROOT, "venv/bin/python3"), "./python/extract.py", stemDst, "./resources/"+constants.ZIP_FILENAME)
		cmd.Dir = constants.ROOT
		if err := cmd.Run(); err != nil {
			log.Println("Python script error")
			log.Fatal(err)
		}

		// 3. unzip and read the raw JSON
		jsonDirLoc := strings.TrimSuffix(filepath.Join(constants.ROOT, "resources", constants.ZIP_FILENAME), filepath.Ext(constants.ZIP_FILENAME))
		if err := Unzip(filepath.Join(constants.ROOT, "resources", constants.ZIP_FILENAME), jsonDirLoc); err != nil {
			log.Fatal(err)
		}
		b, err := os.ReadFile(filepath.Join(jsonDirLoc, constants.JSON_FILENAME))
		if err != nil {
			panic(err)
		}
		rawJSON := string(b)

		dictParsed, err := dictparser.ParseDictJson(rawJSON)
		if err != nil {
			log.Panic(err)
		}

		// 3.5 read 'word_key's with multiple usages
		wordKeysWithMultipleUsages, err := readWordKeysWithMultipleUsagesFromDB(tmpFilePath)

		// 4. parse the raw JSON into golang structures
		oxWords := []oxforddicthandler.OxfordWord{}
		for _, r := range rawRows {
			found := false
			for _, p := range dictParsed {
				if r.Stem.String == p.Word {
					found = true
					if idx := getIdxOfWordInEntry(r.Stem.String, oxWords); isLookedUpMoreThanOnce(r.WordKey.String, wordKeysWithMultipleUsages) && idx > -1 {
						oxWords[idx].AppendUsageAndBook(r.Usage.String, r.Title.String)
					} else {
						oxWords = append(oxWords, *oxforddicthandler.CreateWord(r, p))
					}
					break
				}
			}
			if !found {
				fmt.Printf("\"%s\" not found in a dictionary\n", r.Stem.String)
			}
		}
		processedWords = oxWords

		log.Println("All's well!")
		http.Redirect(w, r, "/view", http.StatusSeeOther)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// ViewHandler generates the deck and displays the list of words.
func ViewHandler(w http.ResponseWriter, r *http.Request) {
	if len(processedWords) == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Generate deck (unchanged)
	filePath := filepath.Join(constants.ROOT, "resources/output.apkg")
	if err := sendtoanki.GenerateDeck(processedWords, filePath); err != nil {
		http.Error(w, fmt.Sprintf("Error generating deck: %v", err), http.StatusInternalServerError)
		return
	}

	// PREPARE DATA: Wrap the words and check if they are "Basic"
	viewWords := make([]ViewWord, len(processedWords))
	for i, word := range processedWords {
		viewWords[i] = ViewWord{
			OxfordWord: word,
			IsBasic:    basicWordMap[word.Word()],
		}
	}

	data := ViewData{
		Words:       viewWords,
		DownloadURL: "/download",
	}

	err := templates.ExecuteTemplate(w, "view.html", data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error rendering template: %v", err), http.StatusInternalServerError)
	}
}

// DownloadHandler specifically serves the generated .apkg file
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join(constants.ROOT, "resources/output.apkg")

	// Check if file exists before serving
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found. Please upload a database first.", http.StatusNotFound)
		return
	}

	// Force the browser to download the file instead of trying to open it
	w.Header().Set("Content-Disposition", "attachment; filename=output.apkg")
	w.Header().Set("Content-Type", "application/octet-stream")

	http.ServeFile(w, r, filePath)
}

// Handle deletion of a specific word
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 1. Get the word to delete from the form
	wordToDelete := r.FormValue("word_to_delete")
	if wordToDelete == "" {
		http.Redirect(w, r, "/view", http.StatusSeeOther)
		return
	}

	// 2. Remove it from the slice
	// We iterate backwards or creating a new slice is often safer/easier
	// to avoid index out of range issues if multiple items matched (unlikely here).
	newWords := []oxforddicthandler.OxfordWord{}
	for _, item := range processedWords {
		if item.Word() != wordToDelete {
			newWords = append(newWords, item)
		}
	}
	processedWords = newWords

	// 3. Regenerate the Anki deck immediately
	// If the list is now empty, we might not want to generate a deck,
	// but usually, an empty deck or skipping generation is fine.
	if len(processedWords) > 0 {
		filePath := filepath.Join(constants.ROOT, "resources/output.apkg")
		err := sendtoanki.GenerateDeck(processedWords, filePath)
		if err != nil {
			log.Printf("Error regenerating deck after delete: %v", err)
			http.Error(w, "Error updating deck", http.StatusInternalServerError)
			return
		}
	} else {
		// Optional: Remove the old file if list is empty
		os.Remove(filepath.Join(constants.ROOT, "resources/output.apkg"))
	}

	// 4. Redirect back to the view page
	http.Redirect(w, r, "/view", http.StatusSeeOther)
}

// Unzip helper
func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	os.MkdirAll(dest, 0755)

	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		path := filepath.Join(dest, f.Name)
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		if err := extractAndWriteFile(f); err != nil {
			return err
		}
	}
	return nil
}

func readEveryRowFromDB(dbLoc string) ([]tutorial.GetRowsRow, error) {
	ctx := context.Background()

	db, err := sql.Open("sqlite", dbLoc)
	if err != nil {
		return []tutorial.GetRowsRow{}, err
	}

	queries := tutorial.New(db)

	rows, err := queries.GetRows(ctx)
	if err != nil {
		return []tutorial.GetRowsRow{}, err
	}
	for _, r := range rows {
		if !(r.Word.Valid && r.Stem.Valid && r.Title.Valid && r.Usage.Valid && r.WordKey.Valid) {
			return []tutorial.GetRowsRow{}, fmt.Errorf("\"%s\" is invalid in DB", r.Word.String)
		}
	}

	return rows, nil
}

func readWordKeysWithMultipleUsagesFromDB(dbLoc string) ([]tutorial.GetWordKeysWithMultipleUsagesRow, error) {
	ctx := context.Background()

	db, err := sql.Open("sqlite", dbLoc)
	if err != nil {
		return []tutorial.GetWordKeysWithMultipleUsagesRow{}, err
	}

	queries := tutorial.New(db)

	rows, err := queries.GetWordKeysWithMultipleUsages(ctx)
	if err != nil {
		return []tutorial.GetWordKeysWithMultipleUsagesRow{}, err
	}
	for _, r := range rows {
		if !(r.WordKey.Valid) {
			return []tutorial.GetWordKeysWithMultipleUsagesRow{}, fmt.Errorf("a row with the following word_key is invalid: %s", r.WordKey.String)
		}
	}

	return rows, nil
}

func writeStems(rows []tutorial.GetRowsRow, dst string) error {
	text := ""
	for _, r := range rows {
		text += r.Stem.String + "\n"
	}
	err := os.WriteFile(dst, []byte(text), 0644)
	return err
}

// getIdxOfWordInEntry(r.Stem, oxWords); isLookedUpMoreThanOnce(r.WordKey, wordKeysWithMultipleUsages) && idx > -1 {

func isLookedUpMoreThanOnce(wordKey string, wordKeysWithMultipleUsages []tutorial.GetWordKeysWithMultipleUsagesRow) bool {
	for _, v := range wordKeysWithMultipleUsages {
		if v.WordKey.String == wordKey {
			return true
		}
	}
	return false
}

func getIdxOfWordInEntry(stem string, wordEntry []oxforddicthandler.OxfordWord) int {
	for i := range wordEntry {
		if wordEntry[i].Word() == stem {
			return i
		}
	}
	return -1
}
