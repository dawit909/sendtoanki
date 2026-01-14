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
	usage     string
	book      string
}

func (w OxfordWord) Word() string {
	return w.wordEntry.Word
}
func (w OxfordWord) Book() string {
	return w.book
}

func (w OxfordWord) Definition() string {
	return generateHTML(w.wordEntry)
	// Step A: Group definitions by EntryID, then by PartOfSpeech
	// Structure: entries[id][pos] -> []definitions
	entries := make(map[int]map[string][]string)

	// Track unique IDs to sort them later
	var entryIDs []int
	seenIDs := make(map[int]bool)

	for _, def := range w.wordEntry.Definitions {
		id := def.EntryID
		pos := def.PartOfSpeech

		// Initialize maps if nil
		if _, exists := entries[id]; !exists {
			entries[id] = make(map[string][]string)
		}
		if !seenIDs[id] {
			entryIDs = append(entryIDs, id)
			seenIDs[id] = true
		}

		entries[id][pos] = append(entries[id][pos], def.Definition)
	}

	sort.Ints(entryIDs) // Ensure Entry 1 comes before Entry 2

	var sb strings.Builder

	// Step B: Iterate over Sorted Entry IDs
	for _, id := range entryIDs {
		posMap := entries[id]

		sb.WriteString("<blockquote>\n")

		// Word Header
		// Logic: Add superscript only if there is more than 1 distinct entry ID
		if len(entryIDs) > 1 {
			sb.WriteString(fmt.Sprintf("    <div>%s<sup>%d</sup></div>\n", w.wordEntry.Word, id))
		} else {
			sb.WriteString(fmt.Sprintf("    <div>%s</div>\n", w.wordEntry.Word))
		}

		// Sort Parts of Speech for consistent output (e.g., Noun before Verb)
		var posKeys []string
		for k := range posMap {
			posKeys = append(posKeys, k)
		}
		sort.Strings(posKeys)

		// Step C: Iterate over Parts of Speech (I. noun, II. verb...)
		for i, pos := range posKeys {
			roman := toRoman(i + 1)
			defs := posMap[pos]

			sb.WriteString(fmt.Sprintf("    <blockquote>%s. %s\n", roman, pos))
			sb.WriteString("        <div>\n")

			// Step D: Iterate over Definitions (1. ..., 2. ...)
			for j, d := range defs {
				sb.WriteString(fmt.Sprintf("            <blockquote>%d. %s</blockquote>\n", j+1, d))
			}

			sb.WriteString("        </div>\n")
			sb.WriteString("    </blockquote>\n")
		}

		sb.WriteString("</blockquote>")
	}

	return sb.String()
}

func (w OxfordWord) Usage() string {
	return fmt.Sprintf("<blockquote>%s<small>%s</small></blockquote>\n", w.usage, w.book)
}

func CreateWord(fromDB tutorial.GetWordsByTitleRow, fromPython dictparser.WordEntry) OxfordWord {
	return OxfordWord{wordEntry: fromPython, usage: fromDB.Usage.String, book: fromDB.Title.String}
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
