package handler

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"tuto.sqlc.dev/app/go/constants"
	"tuto.sqlc.dev/app/go/dbreader"
	"tuto.sqlc.dev/app/go/dictparser"
	"tuto.sqlc.dev/app/go/oxforddicthandler"
)

// Holds the parsed HTML templates.
var templates = template.Must(template.ParseFiles("html/upload.html", "html/view.html"))

var processedWords []oxforddicthandler.OxfordWord

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
			http.Error(w, "Could not save file content", http.StatusInternalServerError)
			return
		}

		dbName := header.Filename
		// 1. Write stems.txt
		if err := dbreader.WriteStems(filepath.Join(constants.ROOT, "resources", dbName), filepath.Join(constants.ROOT, "resources/stems.txt")); err != nil {
			log.Fatal(err)
		}

		// 2. create a raw JSON containing definitions of stems
		fmt.Println("part 2 begun")
		fmt.Println(filepath.Join(constants.ROOT, "venv/bin/python3"), "./python/extract.py", "./resources/stems.txt", "./resources/"+constants.ZIP_FILENAME)
		cmd := exec.Command(filepath.Join(constants.ROOT, "venv/bin/python3"), "./python/extract.py", "./resources/stems.txt", "./resources/"+constants.ZIP_FILENAME)
		cmd.Dir = constants.ROOT
		out, err := cmd.Output()
		if err != nil {
			log.Println("here here")
			log.Fatal(err)
		}
		println(string(out))

		fmt.Println("part 3 begun")
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

		// 4. parse the raw JSON into golang structures
		wordsDB, err := dbreader.ReadDB(filepath.Join(constants.ROOT, "resources", dbName))
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%d words found\n", len(wordsDB))

		oxWords := []oxforddicthandler.OxfordWord{}
		for _, v := range wordsDB {
			found := false
			for _, p := range dictParsed {
				if v.Stem.String == p.Word {
					oxWords = append(oxWords, oxforddicthandler.CreateWord(v, p))
					found = true
					break
				}
			}
			if !found {
				fmt.Printf("\"%s\" not found in a dictionary\n", v.Stem.String)
			}
		}
		processedWords = oxWords

		// // 4. create anki deck
		// err = sendtoanki.GenerateDeck(oxWords, filepath.Join(constants.ROOT, "resources/output.apkg"))
		// if err != nil {
		// 	log.Panic(err)
		// }

		log.Println("All's well!")
		http.Redirect(w, r, "/view", http.StatusSeeOther)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// ViewHandler displays the list of processed words.
func ViewHandler(w http.ResponseWriter, r *http.Request) {
	if len(processedWords) == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		// Check for ZipSlip (Directory traversal)
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
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
