package dbreader

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"os"

	_ "modernc.org/sqlite"

	"tuto.sqlc.dev/app/tutorial"
)

func ReadDB(dbName string) ([]tutorial.GetWordsByTitleRow, error) {
	ctx := context.Background()

	db, err := sql.Open("sqlite", dbName)
	if err != nil {
		return []tutorial.GetWordsByTitleRow{}, err
	}

	queries := tutorial.New(db)

	words, err := queries.GetWordsByTitle(ctx)
	if err != nil {
		return []tutorial.GetWordsByTitleRow{}, err
	}
	for _, w := range words {
		if !isWordValid(w) {
			return []tutorial.GetWordsByTitleRow{}, fmt.Errorf("\"%s\" is invalid in DB", w.Word.String)
		}
	}
	return words, nil
}

func WriteStems(dbName, dstName string) error {
	ctx := context.Background()

	db, err := sql.Open("sqlite", dbName)
	if err != nil {
		return err
	}

	queries := tutorial.New(db)

	words, err := queries.GetWordsByTitle(ctx)
	if err != nil {
		return err
	}

	text := ""
	for _, v := range words {
		text += v.Stem.String + "\n"
	}
	err = os.WriteFile(dstName, []byte(text), 0644)
	return err
}

func isWordValid(w tutorial.GetWordsByTitleRow) bool {
	if !(w.Word.Valid && w.Stem.Valid && w.Title.Valid && w.Usage.Valid) {
		return false
	}
	return true
}
