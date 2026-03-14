package sendtoanki

import (
	"strings"

	"github.com/dawit909/sendtoanki/src/constants"
	"github.com/iancoleman/strcase"
	"github.com/npcnixel/genanki-go"
)

type WordData interface {
	Stem() string
	Book() string
	Definition() string
	Usage() string
}

func GenerateDeck[W WordData](w map[string]*W, fileName string) error {
	// Create a basic model and deck
	model := genanki.NewModel(constants.ANKI_MODEL_ID, constants.ANKI_DECK_NAME)

	// Customize the model (each method call separately)
	model.SetCSS(constants.ANKI_MODEL_CSS)

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
	deck := genanki.StandardDeck(constants.ANKI_DECK_NAME, "A deck made from your Kindle Vocabulary Builder")
	var noteTag string
	// Add notes to the deck using method chaining
	for _, v := range w {
		if v == nil {
			continue
		}
		noteTag = constants.ANKI_DECK_NAME + "::" + strcase.ToCamel(strings.ReplaceAll((*v).Book(), ",", ""))
		deck.AddNote(genanki.NewNote(
			model.ID,
			[]string{(*v).Stem(), (*v).Usage(), (*v).Definition()},
			[]string{noteTag},
		),
		)
	}

	// Create and write package
	pkg := genanki.NewPackage([]*genanki.Deck{deck}).AddModel(model)
	if err := pkg.WriteToFile(fileName); err != nil {
		return err
	}
	return nil
}
