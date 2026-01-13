package dictparser

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

// --- Structs for Input Parsing ---

// InputData represents the top-level input JSON structure
type InputData struct {
	Definitions map[string]string `json:"definitions"`
}

// --- Structs for Output Generation ---

// Definition represents a single definition with context
type Definition struct {
	EntryID      int    `json:"entry_id"`       // Distinguishes homographs (1, 2, 3...)
	PartOfSpeech string `json:"part_of_speech"` // e.g., "noun", "verb"
	Definition   string `json:"definition"`     // The text meaning
}

// WordEntry represents the final JSON structure for a word
type WordEntry struct {
	Word        string       `json:"word"`
	Definitions []Definition `json:"definitions"`
}

func ParseDictJson(rawJSON string) ([]WordEntry, error) {
	var input InputData
	if err := json.Unmarshal([]byte(rawJSON), &input); err != nil {
		return []WordEntry{}, err
	}

	// 2. Process data and build output slice
	var output []WordEntry

	for word, xmlContent := range input.Definitions {
		defs, err := parseDetailedXML(xmlContent)
		if err != nil {
			fmt.Printf("Error processing %s: %v\n", word, err)
			continue
		}

		output = append(output, WordEntry{
			Word:        word,
			Definitions: defs,
		})
	}

	return output, nil
}

func WriteJSON(fileName string, words []WordEntry) error {
	// We use MarshalIndent so the output file is human-readable
	fileContent, err := json.MarshalIndent(words, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(fileName, fileContent, 0644); err != nil {
		return err
	}

	fmt.Printf("Successfully saved definitions to %s\n", fileName)
	return nil
}

func ParseAndWriteJSON(rawJSON string) error {
	var input InputData
	if err := json.Unmarshal([]byte(rawJSON), &input); err != nil {
		return err
	}

	// 2. Process data and build output slice
	var output []WordEntry

	for word, xmlContent := range input.Definitions {
		defs, err := parseDetailedXML(xmlContent)
		if err != nil {
			fmt.Printf("Error processing %s: %v\n", word, err)
			continue
		}

		output = append(output, WordEntry{
			Word:        word,
			Definitions: defs,
		})
	}

	// 3. Serialize to JSON
	// We use MarshalIndent so the output file is human-readable
	fileContent, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return err
	}

	// 4. Write to File
	fileName := "dictionary_output.json"
	if err := os.WriteFile(fileName, fileContent, 0644); err != nil {
		return err
	}

	fmt.Printf("Successfully saved definitions to %s\n", fileName)
	return nil
}

// parseDetailedXML extracts structure from the XML string
func parseDetailedXML(xmlStr string) ([]Definition, error) {
	decoder := xml.NewDecoder(strings.NewReader(xmlStr))
	var results []Definition

	currentEntryIndex := 0
	currentPOS := "unknown" // Default

	for {
		t, err := decoder.Token()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		switch se := t.(type) {
		case xml.StartElement:
			// New Entry
			if se.Name.Local == "entry" {
				currentEntryIndex++
				currentPOS = "unknown"
			}

			// Part of Speech
			if se.Name.Local == "span" && hasAttr(se, "class", "gp tg_pos") {
				var posText string
				if err := decoder.DecodeElement(&posText, &se); err == nil {
					currentPOS = strings.TrimSpace(posText)
				}
			}

			// Definition
			if se.Name.Local == "span" && hasAttr(se, "class", "df") {
				var defText string
				if err := decoder.DecodeElement(&defText, &se); err == nil {
					results = append(results, Definition{
						EntryID:      currentEntryIndex,
						PartOfSpeech: currentPOS,
						Definition:   strings.TrimSpace(defText),
					})
				}
			}
		}
	}
	return results, nil
}

// Helper: check attribute existence
func hasAttr(se xml.StartElement, name, value string) bool {
	for _, attr := range se.Attr {
		if attr.Name.Local == name && attr.Value == value {
			return true
		}
	}
	return false
}
