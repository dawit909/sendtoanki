package oxforddicthandler

import (
	"fmt"
	"sort"
	"strings"

	"tuto.sqlc.dev/app/go/dictparser"
	"tuto.sqlc.dev/app/go/tutorial"
)

type OxfordWord struct {
	wordEntry dictparser.WordEntry
	usage     []string
	book      []string
}

func (w OxfordWord) Word() string {
	return w.wordEntry.Word
}
func (w OxfordWord) Book() string {
	return w.book[0]
}

func (w OxfordWord) Definition() string {
	return generateHTML(w.wordEntry)
}

func (w OxfordWord) Usage() string {
	rst := ""
	for i := range w.usage {
		rst += fmt.Sprintf("<blockquote>%s<small>%s</small></blockquote>", w.usage[i], w.book[i])
	}
	return rst
}

func (w *OxfordWord) AppendUsageAndBook(usage, book string) {
	w.usage = append(w.usage, usage)
	w.book = append(w.book, book)

}

func CreateWord(fromDB tutorial.GetRowsRow, fromPython dictparser.WordEntry) *OxfordWord {
	w := OxfordWord{wordEntry: fromPython}
	w.AppendUsageAndBook(fromDB.Usage.String, fromDB.Title.String)
	return &w
}

func toRoman(n int) string {
	numerals := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	if n > 0 && n < len(numerals) {
		return numerals[n]
	}
	return fmt.Sprintf("%d", n) // Fallback
}

func generateHTML(w dictparser.WordEntry) string {
	entries := make(map[int]map[string][]string)
	var entryIDs []int
	seenIDs := make(map[int]bool)

	// Grouping Logic (Same as before)
	for _, def := range w.Definitions {
		id := def.EntryID
		if _, exists := entries[id]; !exists {
			entries[id] = make(map[string][]string)
		}
		if !seenIDs[id] {
			entryIDs = append(entryIDs, id)
			seenIDs[id] = true
		}
		entries[id][def.PartOfSpeech] = append(entries[id][def.PartOfSpeech], def.Definition)
	}
	sort.Ints(entryIDs)

	var sb strings.Builder
	sb.WriteString("<div class=\"entry-container\">\n")

	for _, id := range entryIDs {
		posMap := entries[id]

		// Word Header
		if len(entryIDs) > 1 {
			sb.WriteString(fmt.Sprintf("    <div class=\"word-header\">%s<span class=\"superscript\">%d</span></div>\n", w.Word, id))
		} else {
			sb.WriteString(fmt.Sprintf("    <div class=\"word-header\">%s</div>\n", w.Word))
		}

		// Sort Parts of Speech
		var posKeys []string
		for k := range posMap {
			posKeys = append(posKeys, k)
		}
		sort.Strings(posKeys)

		// Loop Parts of Speech
		for i, pos := range posKeys {
			roman := toRoman(i + 1)
			defs := posMap[pos]

			sb.WriteString("    <div class=\"pos-group\">\n")
			// We print the Roman numeral here manually
			sb.WriteString(fmt.Sprintf("        <div class=\"pos-header\">%s. %s</div>\n", roman, pos))

			// Use standard Ordered List <ol> for definitions
			sb.WriteString("        <ol class=\"def-list\">\n")
			for _, d := range defs {
				sb.WriteString(fmt.Sprintf("            <li>%s</li>\n", d))
			}
			sb.WriteString("        </ol>\n")
			sb.WriteString("    </div>\n")
		}
	}
	sb.WriteString("</div>")

	return sb.String()
}
