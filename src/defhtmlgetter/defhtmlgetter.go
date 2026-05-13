package defhtmlgetter

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const closingTag = "</d:entry>"

const htmlFooter = `</body>
</html>


`

// entryLoc stores exactly where in the file a definition lives on the disk.
type entryLoc struct {
	offset int64
	length int
}

// Global Index: Maps a word (string) to a list of file locations.
// We use a list []entryLoc because one word can have multiple definitions.
var (
	index     map[string][]entryLoc
	cachePath string
)

func init() {
	cachePath = "./resources/noad.cache"
	index = make(map[string][]entryLoc)

	// Build the lightweight index at startup
	if err := buildIndex(cachePath); err != nil {
		fmt.Printf("Warning: Could not build dictionary index: %v\n", err)
	}
}

// buildIndex scans the file line-by-line without keeping the content in memory.
func buildIndex(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	var currentOffset int64 = 0

	for {
		// ReadBytes is safer and more memory-efficient for long lines
		lineBytes, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				// Process the last line if the file doesn't end with a newline
				if len(lineBytes) > 0 {
					processLine(lineBytes, currentOffset)
				}
				break
			}
			return err
		}

		processLine(lineBytes, currentOffset)
		currentOffset += int64(len(lineBytes))
	}
	return nil
}

// processLine extracts the title and calculates the byte offset of the body
func processLine(lineBytes []byte, currentOffset int64) {
	// Find the Tab delimiter
	tabIdx := bytes.IndexByte(lineBytes, '\t')
	if tabIdx == -1 {
		return // Malformed line, skip
	}

	// Extract Title (Only memory allocation here)
	title := string(lineBytes[:tabIdx])

	// Calculate Body Location (Do NOT allocate body memory)
	bodyStart := currentOffset + int64(tabIdx) + 1

	// Check for trailing newline to exclude it from the length
	trim := 0
	if len(lineBytes) > 0 && lineBytes[len(lineBytes)-1] == '\n' {
		trim = 1
	}

	bodyLength := len(lineBytes) - tabIdx - 1 - trim

	if bodyLength > 0 {
		index[title] = append(index[title], entryLoc{
			offset: bodyStart,
			length: bodyLength,
		})
	}
}

// Get looks up the word in the index, reads directly from disk, and returns HTML.
func Get(word string) (string, error) {
	// 1. Check Index
	locs, found := index[word]
	if !found {
		// Word not in dictionary, return just the footer (matches old behavior)
		return CleanAppleDictHTML(htmlFooter)
	}

	// 2. Open File (On Demand)
	f, err := os.Open(cachePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var sb strings.Builder

	// 3. Read and Concatenate all definitions for this word from the disk
	for _, loc := range locs {
		bodyBuf := make([]byte, loc.length)
		_, err := f.ReadAt(bodyBuf, loc.offset)
		if err != nil && err != io.EOF {
			return "", err
		}

		sb.Write(bodyBuf)
	}

	sb.WriteString(htmlFooter)

	dirtyHtml := sb.String()
	return CleanAppleDictHTML(dirtyHtml)
}

// CleanAppleDictHTML processes the raw Apple Dictionary HTML to fix layout issues and remove clutter.
func CleanAppleDictHTML(rawXML string) (string, error) {
	doc, err := html.Parse(strings.NewReader(rawXML))
	if err != nil {
		return "", err
	}

	var cleanNode func(*html.Node)
	cleanNode = func(n *html.Node) {
		if n.Type == html.ElementNode {
			// 1. Transform <d:entry> to <div>
			if n.Data == "d:entry" {
				n.Data = "div"
				n.DataAtom = atom.Div
			}

			// 2. Attribute Processing
			var cleanAttrs []html.Attribute
			for _, attr := range n.Attr {
				key := strings.ToLower(attr.Key)

				// A. Blocklist: Remove metadata
				if strings.HasPrefix(key, "d:") ||
					strings.HasPrefix(key, "xmlns") ||
					key == "id" ||
					key == "prlexid" ||
					key == "soundfile" ||
					key == "media" ||
					key == "dialect" ||
					key == "role" ||
					key == "onmouseover" ||
					key == "onmouseout" {
					continue
				}

				// B. Fix The "Gap" (Collusion vs Commend)
				if key == "class" {
					// We split the classes into a slice to check them individually
					classes := strings.Fields(attr.Val)

					hasPosg := false
					hasXdH := false

					for _, c := range classes {
						if c == "posg" {
							hasPosg = true
						}
						if c == "x_xdh" {
							hasXdH = true
						}
					}

					// TARGET: If an element has BOTH 'posg' and 'x_xdh' (like in 'collusion'),
					// it means the POS header is empty of other content (like [with object]).
					// We MUST remove 'x_xdh' to stop it from acting like a block with margins.
					if hasPosg && hasXdH {
						var newClasses []string
						for _, c := range classes {
							if c != "x_xdh" { // Skip the block class
								newClasses = append(newClasses, c)
							}
						}
						attr.Val = strings.Join(newClasses, " ")
					}
				}

				cleanAttrs = append(cleanAttrs, attr)
			}
			n.Attr = cleanAttrs
		}

		// 3. Process Children
		var next *html.Node
		for c := n.FirstChild; c != nil; c = next {
			next = c.NextSibling
			cleanNode(c)
		}

		// 4. Remove Empty Tags (Cleanup)
		if n.Type == html.ElementNode {
			if strings.HasPrefix(n.Data, "d:") && n.FirstChild == nil {
				removeNode(n)
				return
			}
			// Remove empty structural spans
			if n.Data == "span" && n.FirstChild == nil {
				for _, a := range n.Attr {
					if a.Key == "class" && (a.Val == "hsb" || a.Val == "tg_pos" || a.Val == "tg_df") {
						removeNode(n)
						return
					}
				}
			}
		}
	}

	// Extract body content
	var body *html.Node
	var finder func(*html.Node)
	finder = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "body" {
			body = n
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			finder(c)
		}
	}
	finder(doc)

	if body != nil {
		cleanNode(body)
		var buf bytes.Buffer
		for c := body.FirstChild; c != nil; c = c.NextSibling {
			html.Render(&buf, c)
		}
		return buf.String(), nil
	}

	return "", fmt.Errorf("failed to parse HTML structure")
}

func removeNode(n *html.Node) {
	if n.Parent != nil {
		n.Parent.RemoveChild(n)
	}
}
