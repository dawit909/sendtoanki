// package main

// import (
// 	"archive/zip"
// 	_ "embed"
// 	"fmt"
// 	"html/template"
// 	"io"
// 	"log"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// 	"strings"

// 	_ "modernc.org/sqlite"

// 	"tuto.sqlc.dev/app/dbreader"
// 	"tuto.sqlc.dev/app/dictparser"
// 	"tuto.sqlc.dev/app/oxforddicthandler"
// )

// const (
// 	UsageInFront int = iota
// 	UsageInBack
// )

// const (
// 	OriginOn int = iota
// 	OriginOff
// )

// type OxWords []oxforddicthandler.OxfordWord

// func (words OxWords) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	for _, v := range words {
// 		t, _ := template.ParseFiles("oneword.html")
// 		t.Execute(w, struct {
// 			Word string
// 			Book string
// 		}{v.Word(), v.Book()})
// 	}
// }

// //go:embed schema.sql
// var ddl string

// type Page struct {
// 	Title string
// 	Body  []byte
// }

// func (p *Page) save() error {
// 	filename := p.Title + ".txt"
// 	return os.WriteFile(filename, p.Body, 0600)
// }

// func uploadHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/upload/"):]
// 	p, err := loadPage(title)
// 	if err != nil {
// 		p = &Page{Title: title}
// 	}
// 	t, _ := template.ParseFiles("upload.html")
// 	t.Execute(w, p)
// }

// func loadPage(title string) (*Page, error) {
// 	filename := title + ".txt"
// 	body, err := os.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Page{Title: title, Body: body}, nil
// }

// func main() {
// 	log.Println("started...")
// 	dbName := "./vocab.db"
// 	// // 1. Write stems.txt
// 	// if err := dbreader.WriteStems(dbName, "/Users/zibabshivan/Desktop/parse_dict/stems.txt"); err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// // 2. create a raw JSON containing definitions of stems
// 	// cmd := exec.Command("/Users/zibabshivan/Desktop/parse_dict/venv/bin/python3", "./extract.py", "./stems.txt", "./PATH_TO_OUTPUT3.zip")
// 	// cmd.Dir = "/Users/zibabshivan/Desktop/parse_dict"
// 	// out, err := cmd.Output()
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// println(string(out))

// 	// 3. unzip and read the raw JSON
// 	if err := Unzip("../parse_dict/PATH_TO_OUTPUT3.zip", "../parse_dict/PATH_TO_OUTPUT3"); err != nil {
// 		log.Fatal(err)
// 	}
// 	b, err := os.ReadFile("../parse_dict/PATH_TO_OUTPUT3/master.json")
// 	if err != nil {
// 		panic(err)
// 	}
// 	rawJSON := string(b)

// 	dictParsed, err := dictparser.ParseDictJson(rawJSON)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	// 4. parse the raw JSON into golang structures
// 	wordsDB, err := dbreader.ReadDB(dbName)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("%d words found\n", len(wordsDB))

// 	oxWords := []oxforddicthandler.OxfordWord{}
// 	for _, v := range wordsDB {
// 		found := false
// 		for _, p := range dictParsed {
// 			if v.Stem.String == p.Word {
// 				oxWords = append(oxWords, oxforddicthandler.CreateWord(v, p))
// 				found = true
// 				break
// 			}
// 		}
// 		if !found {
// 			fmt.Printf("\"%s\" not found in a dictionary\n", v.Stem.String)
// 		}
// 	}
// 	http.Handle("/view/", OxWords(oxWords))
// 	http.HandleFunc("/upload/", uploadHandler)
// 	log.Fatal(http.ListenAndServe(":8080", nil))

// 	// // 4. create anki deck
// 	// err = sendtoanki.GenerateDeck(oxWords, "output.apkg")
// 	// if err != nil {
// 	// 	log.Panic(err)
// 	// }

// 	log.Println("All's well!")
// }

// // Source - https://stackoverflow.com/a
// // Posted by Astockwell, modified by community. See post 'Timeline' for change history
// // Retrieved 2026-01-11, License - CC BY-SA 4.0

// func Unzip(src, dest string) error {
// 	r, err := zip.OpenReader(src)
// 	if err != nil {
// 		return err
// 	}
// 	defer func() {
// 		if err := r.Close(); err != nil {
// 			panic(err)
// 		}
// 	}()

// 	os.MkdirAll(dest, 0755)

// 	// Closure to address file descriptors issue with all the deferred .Close() methods
// 	extractAndWriteFile := func(f *zip.File) error {
// 		rc, err := f.Open()
// 		if err != nil {
// 			return err
// 		}
// 		defer func() {
// 			if err := rc.Close(); err != nil {
// 				panic(err)
// 			}
// 		}()

// 		path := filepath.Join(dest, f.Name)

// 		// Check for ZipSlip (Directory traversal)
// 		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
// 			return fmt.Errorf("illegal file path: %s", path)
// 		}

// 		if f.FileInfo().IsDir() {
// 			os.MkdirAll(path, f.Mode())
// 		} else {
// 			os.MkdirAll(filepath.Dir(path), f.Mode())
// 			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
// 			if err != nil {
// 				return err
// 			}
// 			defer func() {
// 				if err := f.Close(); err != nil {
// 					panic(err)
// 				}
// 			}()

// 			_, err = io.Copy(f, rc)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 		return nil
// 	}

// 	for _, f := range r.File {
// 		err := extractAndWriteFile(f)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WordTemp struct {
	Word  string
	Book  string
	Usage string
	Def   string
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	words := []WordTemp{
		{"something", "Book1", "something's gonna change", "something"},
		{"you", "Book2", "I am the warlus you are the warlus", "not me"},
	}

	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}
