package sendtoanki

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/npcnixel/genanki-go"
	"tuto.sqlc.dev/app/go/constants"
)

type WordData interface {
	Word() string
	Book() string
	Definition() string
	Usage() string
}

func GenerateDeck[W WordData](w []W, fileName string) error {
	// Create a basic model and deck
	model := genanki.NewModel(constants.ANKI_MODEL_ID, "sendtokindle")

	// Customize the model (each method call separately)
	model.SetCSS(`h1 {
 text-align: center;
}

/* Container for the whole card */
.entry-container {
    text-align: left;
    max-width: 100%;
}

/* The Main Word */
.word-header {
    font-size: 24px;
    font-weight: bold;
    margin-bottom: 10px;
    color: #2c3e50; /* Dark blue-grey */
}

.superscript {
    font-size: 0.6em;
    vertical-align: super;
    color: #7f8c8d;
}

/* Part of Speech Header (I. noun) */
.pos-group {
    margin-bottom: 15px;
}

.pos-header {
    font-weight: bold;
    font-style: italic;
    color: #2980b9; /* Nice blue */
    margin-bottom: 5px;
}

/* The Definition List */
ol.def-list {
    margin: 0;
    padding-left: 20px; /* Much smaller indent than blockquote */
}

ol.def-list li {
    margin-bottom: 4px;
    line-height: 1.4;
    color: #34495e;
}

blockquote small:before {
 content: " -";
}

.nightMode .word-header { color: #ecf0f1; }
.nightMode .pos-header { color: #3498db; }
.nightMode .def-list li,.nightMode h1 { color: #bdc3c7; }
`)

	model.AddField(genanki.Field{
		Name: "Word",
		Font: "Arial",
		Size: 20,
	})

	model.AddField(genanki.Field{
		Name: "Usage",
		Font: "Arial",
		Size: 20,
	})

	model.AddField(genanki.Field{
		Name: "Definition",
		Font: "Arial",
		Size: 20,
	})

	model.Templates = []genanki.Template{
		{
			Name: "Card 1",
			Qfmt: "<h1>{{Word}}</h1>\n<hr>\n{{Usage}}",
			Afmt: "{{FrontSide}}\n\n<hr>\n{{Definition}}",
		},
	}
	deck := genanki.StandardDeck("sendtoanki", "A deck for testing")

	// Add notes to the deck using method chaining
	for _, v := range w {
		deck.AddNote(genanki.NewNote(
			model.ID,
			[]string{v.Word(), v.Usage(), v.Definition()},
			[]string{"sendtokindle::" + strcase.ToCamel(v.Book())},
		),
		)
	}

	// Create and write package
	pkg := genanki.NewPackage([]*genanki.Deck{deck}).AddModel(model)
	if err := pkg.WriteToFile(fileName); err != nil {
		return err
	}
	fmt.Printf("Successfully created Anki deck: %s\n", fileName)
	return nil
}
