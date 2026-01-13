package sendtoanki

import (
	"fmt"

	"github.com/npcnixel/genanki-go"
)

type WordData interface {
	Word() string
	Book() string
	Definition() string
	Usage() string
}

func GenerateDeckOrig[W WordData](w []W, fileName string) error {
	// Create a basic model and deck
	model := genanki.NewModel(141592653, "Basic Kindle2")

	// Customize the model (each method call separately)
	model.SetCSS(`html {
    font-size: 62.5%; /* 1rem equals 10px now */
    line-height: 1.5;
    text-align: start;
    overflow-y: scroll;
}

.card {
    font-size: 2rem;
    margin-inline: auto;
    padding: 0 2rem;
    max-width: 75ch;
}

h1 {
 text-align: center;
}
blockquote small:before {
 content: " -";
}

blockquote {
  margin: 0 0 1rem 1.5rem; 
}


.win        .card { font-family: "Segoe UI" }
.mac        .card { font-family: BlinkMacSystemFont }
.linux      .card { font-family: sans-serif; }
.ios        .card { font-family: -apple-system }

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
	deck := genanki.StandardDeck("My Deck2", "A deck for testing")

	// Add notes to the deck using method chaining
	for _, v := range w {
		deck.AddNote(genanki.NewNote(
			model.ID,
			[]string{v.Word(), v.Usage(), v.Definition()}, // `<div>inset </div> <div><blockquote data-depth="0" data-listtype="ol" data-value="0">I. noun —  inset /ˈɪnsɛt /  <div><blockquote data-depth="1" data-listtype="ol" data-value="0">1.  a thing that is put in or inserted <div><div><blockquote data-depth="2" data-listtype="ul" data-value="0"> •  a pair of doors with their original stained-glass insets.  </blockquote></div></div></blockquote><blockquote data-depth="1" data-listtype="ol" data-value="1">2.  a small picture or map inserted within the border of a larger one. </blockquote><blockquote data-depth="1" data-listtype="ol" data-value="2">3.  a section of cloth or needlework inserted into a garment <div><div><blockquote data-depth="2" data-listtype="ul" data-value="0"> •  elastic insets in the waistband.  </blockquote></div></div></blockquote><blockquote data-depth="1" data-listtype="ol" data-value="3">4.  an insert in a magazine or other publication. </blockquote></div> </blockquote><blockquote data-depth="0" data-listtype="ol" data-value="1">II. verb —  inset /ɪnˈsɛt /  <div><blockquote data-depth="1" data-listtype="ol" data-value="0">1.  put in (something) as an inset <div><div><blockquote data-depth="2" data-listtype="ul" data-value="0"> •  washbasins are usually inset into a toilet table to form a vanity unit.  </blockquote></div></div></blockquote><blockquote data-depth="1" data-listtype="ol" data-value="1">2.  decorate with an inset <div><div><blockquote data-depth="2" data-listtype="ul" data-value="0"> •  tables inset with ceramic tiles.  </blockquote></div></div></blockquote></div> </blockquote><blockquote data-depth="0" data-listtype="ol" data-value="2">III. derivatives <div><blockquote data-value="1"> <div>insetter </div> <div><blockquote data-value="2">noun <div><blockquote data-value="3"> </blockquote></div> </blockquote></div> </blockquote></div> </blockquote></div>`},

			[]string{},
		),
		)
	}

	// Create and write package
	pkg := genanki.NewPackage([]*genanki.Deck{deck}).AddModel(model)
	if err := pkg.WriteToFile(fileName); err != nil {
		return err
	}
	fmt.Printf("Successfully created Anki deck: %s", fileName)
	return nil
}

func GenerateDeck[W WordData](w []W, fileName string) error {
	// Create a basic model and deck
	model := genanki.NewModel(141592653, "Basic Kindle2")

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

/* OPTIONAL: Night Mode Support for Anki */
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
	deck := genanki.StandardDeck("My Deck2", "A deck for testing")

	// Add notes to the deck using method chaining
	for _, v := range w {
		deck.AddNote(genanki.NewNote(
			model.ID,
			[]string{v.Word(), v.Usage(), v.Definition()}, // `<div>inset </div> <div><blockquote data-depth="0" data-listtype="ol" data-value="0">I. noun —  inset /ˈɪnsɛt /  <div><blockquote data-depth="1" data-listtype="ol" data-value="0">1.  a thing that is put in or inserted <div><div><blockquote data-depth="2" data-listtype="ul" data-value="0"> •  a pair of doors with their original stained-glass insets.  </blockquote></div></div></blockquote><blockquote data-depth="1" data-listtype="ol" data-value="1">2.  a small picture or map inserted within the border of a larger one. </blockquote><blockquote data-depth="1" data-listtype="ol" data-value="2">3.  a section of cloth or needlework inserted into a garment <div><div><blockquote data-depth="2" data-listtype="ul" data-value="0"> •  elastic insets in the waistband.  </blockquote></div></div></blockquote><blockquote data-depth="1" data-listtype="ol" data-value="3">4.  an insert in a magazine or other publication. </blockquote></div> </blockquote><blockquote data-depth="0" data-listtype="ol" data-value="1">II. verb —  inset /ɪnˈsɛt /  <div><blockquote data-depth="1" data-listtype="ol" data-value="0">1.  put in (something) as an inset <div><div><blockquote data-depth="2" data-listtype="ul" data-value="0"> •  washbasins are usually inset into a toilet table to form a vanity unit.  </blockquote></div></div></blockquote><blockquote data-depth="1" data-listtype="ol" data-value="1">2.  decorate with an inset <div><div><blockquote data-depth="2" data-listtype="ul" data-value="0"> •  tables inset with ceramic tiles.  </blockquote></div></div></blockquote></div> </blockquote><blockquote data-depth="0" data-listtype="ol" data-value="2">III. derivatives <div><blockquote data-value="1"> <div>insetter </div> <div><blockquote data-value="2">noun <div><blockquote data-value="3"> </blockquote></div> </blockquote></div> </blockquote></div> </blockquote></div>`},

			[]string{},
		),
		)
	}

	// Create and write package
	pkg := genanki.NewPackage([]*genanki.Deck{deck}).AddModel(model)
	if err := pkg.WriteToFile(fileName); err != nil {
		return err
	}
	fmt.Printf("Successfully created Anki deck: %s", fileName)
	return nil
}
