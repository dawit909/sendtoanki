package constants

import (
	"path/filepath"
)

const JSON_FILENAME = "oxf_amrc_eng_entries.json"
const ROOT = "."
const ANKI_MODEL_ID int64 = 1558522999 // randomly generated
const DECK_FILENAME string = "sendtoanki.apkg"

var RSC_DIR_PATH string = filepath.Join(ROOT, "resources")
var TMP_DIR_PATH string = filepath.Join(ROOT, "tmp")
var DECK_PATH string = filepath.Join(TMP_DIR_PATH, DECK_FILENAME)
var JSON_PATH string = filepath.Join(ROOT, "resources", JSON_FILENAME)
var ANKI_DECK_NAME string = "sendtoanki"

var WIKIPEDIA_ENG_1000_BASIC = map[string]bool{
	"a": true, "about": true, "above": true, "across": true, "act": true, "active": true, "activity": true, "add": true, "afraid": true, "after": true, "again": true, "age": true, "ago": true, "agree": true, "air": true, "all": true, "alone": true, "along": true, "already": true, "always": true, "am": true, "amount": true, "an": true, "and": true, "angry": true, "another": true, "answer": true, "any": true, "anyone": true, "anything": true, "anytime": true, "appear": true, "apple": true, "are": true, "area": true, "arm": true, "army": true, "around": true, "arrive": true, "art": true, "as": true, "ask": true, "at": true, "attack": true, "aunt": true, "autumn": true, "away": true,
	"baby": true, "back": true, "bad": true, "bag": true, "ball": true, "bank": true, "base": true, "basket": true, "bath": true, "be": true, "bean": true, "bear": true, "beautiful": true, "bed": true, "bedroom": true, "beer": true, "before": true, "begin": true, "behave": true, "behind": true, "bell": true, "below": true, "besides": true, "best": true, "better": true, "between": true, "big": true, "bird": true, "birth": true, "birthday": true, "bit": true, "bite": true, "black": true, "bleed": true, "block": true, "blood": true, "blow": true, "blue": true, "board": true, "boat": true, "body": true, "boil": true, "bone": true, "book": true, "border": true, "born": true, "borrow": true, "both": true, "bottle": true, "bottom": true, "bowl": true, "box": true, "boy": true, "branch": true, "brave": true, "bread": true, "break": true, "breakfast": true, "breathe": true, "bridge": true, "bright": true, "bring": true, "brother": true, "brown": true, "brush": true, "build": true, "burn": true, "bus": true, "business": true, "busy": true, "but": true, "buy": true, "by": true,
	"cake": true, "call": true, "can": true, "candle": true, "cap": true, "car": true, "card": true, "care": true, "careful": true, "careless": true, "carry": true, "case": true, "cat": true, "catch": true, "central": true, "century": true, "certain": true, "chair": true, "chance": true, "change": true, "chase": true, "cheap": true, "cheese": true, "chicken": true, "child": true, "children": true, "chocolate": true, "choice": true, "choose": true, "circle": true, "city": true, "class": true, "clean": true, "clear": true, "clever": true, "climb": true, "clock": true, "close": true, "cloth": true, "clothes": true, "cloud": true, "cloudy": true, "coat": true, "coffee": true, "coin": true, "cold": true, "collect": true, "color": true, "comb": true, "come": true, "comfortable": true, "common": true, "compare": true, "complete": true, "computer": true, "condition": true, "contain": true, "continue": true, "control": true, "cook": true, "cool": true, "copper": true, "corn": true, "corner": true, "correct": true, "cost": true, "count": true, "country": true, "course": true, "cover": true, "crash": true, "cross": true, "cry": true, "cup": true, "cupboard": true, "cut": true,
	"dance": true, "dangerous": true, "dark": true, "daughter": true, "day": true, "dead": true, "decide": true, "decrease": true, "deep": true, "deer": true, "depend": true, "desk": true, "destroy": true, "develop": true, "die": true, "different": true, "difficult": true, "dinner": true, "direction": true, "dirty": true, "discover": true, "dish": true, "do": true, "dog": true, "door": true, "double": true, "down": true, "draw": true, "dream": true, "dress": true, "drink": true, "drive": true, "drop": true, "dry": true, "duck": true, "dust": true, "duty": true,
	"each": true, "ear": true, "early": true, "earn": true, "earth": true, "east": true, "easy": true, "eat": true, "education": true, "effect": true, "egg": true, "eight": true, "either": true, "electric": true, "elephant": true, "else": true, "empty": true, "end": true, "enemy": true, "enjoy": true, "enough": true, "enter": true, "entrance": true, "equal": true, "escape": true, "even": true, "evening": true, "event": true, "ever": true, "every": true, "everybody": true, "everyone": true, "exact": true, "examination": true, "example": true, "except": true, "excited": true, "exercise": true, "expect": true, "expensive": true, "explain": true, "extremely": true, "eye": true,
	"face": true, "fact": true, "fail": true, "fall": true, "false": true, "family": true, "famous": true, "far": true, "farm": true, "fast": true, "fat": true, "father": true, "fault": true, "fear": true, "feed": true, "feel": true, "female": true, "fever": true, "few": true, "fight": true, "fill": true, "film": true, "find": true, "fine": true, "finger": true, "finish": true, "fire": true, "first": true, "fish": true, "fit": true, "five": true, "fix": true, "flag": true, "flat": true, "float": true, "floor": true, "flour": true, "flower": true, "fly": true, "fold": true, "food": true, "fool": true, "foot": true, "football": true, "for": true, "force": true, "foreign": true, "forest": true, "forget": true, "forgive": true, "fork": true, "form": true, "four": true, "fox": true, "free": true, "freedom": true, "freeze": true, "fresh": true, "friend": true, "friendly": true, "from": true, "front": true, "fruit": true, "full": true, "fun": true, "funny": true, "furniture": true, "further": true, "future": true,
	"game": true, "garden": true, "gate": true, "general": true, "gentleman": true, "get": true, "gift": true, "give": true, "glad": true, "glass": true, "go": true, "goat": true, "god": true, "gold": true, "good": true, "goodbye": true, "grandfather": true, "grandmother": true, "grass": true, "grave": true, "gray": true, "great": true, "green": true, "ground": true, "group": true, "grow": true, "gun": true,
	"hair": true, "half": true, "hall": true, "hammer": true, "hand": true, "happen": true, "happy": true, "hard": true, "hat": true, "hate": true, "have": true, "he": true, "head": true, "healthy": true, "hear": true, "heart": true, "heaven": true, "heavy": true, "height": true, "hello": true, "help": true, "hen": true, "her": true, "here": true, "hers": true, "hide": true, "high": true, "hill": true, "him": true, "his": true, "hit": true, "hobby": true, "hold": true, "hole": true, "holiday": true, "home": true, "hope": true, "horse": true, "hospital": true, "hot": true, "hotel": true, "hour": true, "house": true, "how": true, "hundred": true, "hungry": true, "hurry": true, "hurt": true, "husband": true,
	"I": true, "ice": true, "idea": true, "if": true, "important": true, "in": true, "increase": true, "inside": true, "into": true, "introduce": true, "invent": true, "invite": true, "iron": true, "is": true, "island": true, "it": true, "its": true,
	"jelly": true, "job": true, "join": true, "juice": true, "jump": true, "just": true,
	"keep": true, "key": true, "kill": true, "kind": true, "king": true, "kitchen": true, "knee": true, "knife": true, "knock": true, "know": true,
	"ladder": true, "lady": true, "lamp": true, "land": true, "large": true, "last": true, "late": true, "lately": true, "laugh": true, "lazy": true, "lead": true, "leaf": true, "learn": true, "leave": true, "left": true, "leg": true, "lend": true, "length": true, "less": true, "lesson": true, "let": true, "letter": true, "library": true, "lie": true, "life": true, "light": true, "like": true, "lion": true, "lip": true, "list": true, "listen": true, "little": true, "live": true, "lock": true, "lonely": true, "long": true, "look": true, "lose": true, "lot": true, "love": true, "low": true, "lower": true, "luck": true,
	"machine": true, "main": true, "make": true, "male": true, "man": true, "many": true, "map": true, "mark": true, "market": true, "marry": true, "matter": true, "may": true, "me": true, "meal": true, "mean": true, "measure": true, "meat": true, "medicine": true, "meet": true, "member": true, "mention": true, "method": true, "middle": true, "milk": true, "million": true, "mind": true, "minute": true, "miss": true, "mistake": true, "mix": true, "model": true, "modern": true, "moment": true, "money": true, "monkey": true, "month": true, "moon": true, "more": true, "morning": true, "most": true, "mother": true, "mountain": true, "mouth": true, "move": true, "much": true, "music": true, "must": true, "my": true,
	"name": true, "narrow": true, "nation": true, "nature": true, "near": true, "nearly": true, "neck": true, "need": true, "needle": true, "neighbor": true, "neither": true, "net": true, "never": true, "new": true, "news": true, "newspaper": true, "next": true, "nice": true, "night": true, "nine": true, "no": true, "noble": true, "noise": true, "none": true, "nor": true, "north": true, "nose": true, "not": true, "nothing": true, "notice": true, "now": true, "number": true,
	"obey": true, "object": true, "ocean": true, "of": true, "off": true, "offer": true, "office": true, "often": true, "oil": true, "old": true, "on": true, "once": true, "one": true, "only": true, "open": true, "opposite": true, "or": true, "orange": true, "order": true, "other": true, "our": true, "out": true, "outside": true, "over": true, "own": true,
	"page": true, "pain": true, "paint": true, "pair": true, "pan": true, "paper": true, "parent": true, "park": true, "part": true, "partner": true, "party": true, "pass": true, "past": true, "path": true, "pay": true, "peace": true, "pen": true, "pencil": true, "people": true, "pepper": true, "per": true, "perfect": true, "period": true, "person": true, "petrol": true, "photograph": true, "piano": true, "pick": true, "picture": true, "piece": true, "pig": true, "pin": true, "pink": true, "place": true, "plane": true, "plant": true, "plastic": true, "plate": true, "play": true, "please": true, "pleased": true, "plenty": true, "pocket": true, "point": true, "poison": true, "police": true, "polite": true, "pool": true, "poor": true, "popular": true, "position": true, "possible": true, "potato": true, "pour": true, "power": true, "present": true, "press": true, "pretty": true, "prevent": true, "price": true, "prince": true, "prison": true, "private": true, "prize": true, "probably": true, "problem": true, "produce": true, "promise": true, "proper": true, "protect": true, "provide": true, "public": true, "pull": true, "punish": true, "pupil": true, "push": true, "put": true,
	"queen": true, "question": true, "quick": true, "quiet": true, "quite": true,
	"radio": true, "rain": true, "rainy": true, "raise": true, "reach": true, "read": true, "ready": true, "real": true, "really": true, "receive": true, "record": true, "red": true, "remember": true, "remind": true, "remove": true, "rent": true, "repair": true, "repeat": true, "reply": true, "report": true, "rest": true, "restaurant": true, "result": true, "return": true, "rice": true, "rich": true, "ride": true, "right": true, "ring": true, "rise": true, "road": true, "rob": true, "rock": true, "room": true, "round": true, "rubber": true, "rude": true, "rule": true, "ruler": true, "run": true, "rush": true,
	"sad": true, "safe": true, "sail": true, "salt": true, "same": true, "sand": true, "save": true, "say": true, "school": true, "science": true, "scissors": true, "search": true, "seat": true, "second": true, "see": true, "seem": true, "sell": true, "send": true, "sentence": true, "serve": true, "seven": true, "several": true, "sex": true, "shade": true, "shadow": true, "shake": true, "shape": true, "share": true, "sharp": true, "she": true, "sheep": true, "sheet": true, "shelf": true, "shine": true, "ship": true, "shirt": true, "shoe": true, "shoot": true, "shop": true, "short": true, "should": true, "shoulder": true, "shout": true, "show": true, "sick": true, "side": true, "signal": true, "silence": true, "silly": true, "silver": true, "similar": true, "simple": true, "since": true, "sing": true, "single": true, "sink": true, "sister": true, "sit": true, "six": true, "size": true, "skill": true, "skin": true, "skirt": true, "sky": true, "sleep": true, "slip": true, "slow": true, "small": true, "smell": true, "smile": true, "smoke": true, "snow": true, "so": true, "soap": true, "sock": true, "soft": true, "some": true, "someone": true, "something": true, "sometimes": true, "son": true, "soon": true, "sorry": true, "sound": true, "soup": true, "south": true, "space": true, "speak": true, "special": true, "speed": true, "spell": true, "spend": true, "spoon": true, "sport": true, "spread": true, "spring": true, "square": true, "stamp": true, "stand": true, "star": true, "start": true, "station": true, "stay": true, "steal": true, "steam": true, "step": true, "still": true, "stomach": true, "stone": true, "stop": true, "store": true, "storm": true, "story": true, "strange": true, "street": true, "strong": true, "structure": true, "student": true, "study": true, "stupid": true, "subject": true, "substance": true, "successful": true, "such": true, "sudden": true, "sugar": true, "suitable": true, "summer": true, "sun": true, "sunny": true, "support": true, "sure": true, "surprise": true, "sweet": true, "swim": true, "sword": true,
	"table": true, "take": true, "talk": true, "tall": true, "taste": true, "taxi": true, "tea": true, "teach": true, "team": true, "tear": true, "telephone": true, "television": true, "tell": true, "ten": true, "tennis": true, "terrible": true, "test": true, "than": true, "that": true, "the": true, "their": true, "then": true, "there": true, "therefore": true, "these": true, "thick": true, "thin": true, "thing": true, "think": true, "third": true, "this": true, "though": true, "threat": true, "three": true, "tidy": true, "tie": true, "title": true, "to": true, "today": true, "toe": true, "together": true, "tomorrow": true, "tonight": true, "too": true, "tool": true, "tooth": true, "top": true, "total": true, "touch": true, "town": true, "train": true, "tram": true, "travel": true, "tree": true, "trouble": true, "true": true, "trust": true, "try": true, "turn": true, "twice": true, "two": true, "type": true,
	"ugly": true, "uncle": true, "under": true, "understand": true, "unit": true, "until": true, "up": true, "use": true, "useful": true, "usual": true, "usually": true,
	"vegetable": true, "very": true, "village": true, "visit": true, "voice": true,
	"wait": true, "wake": true, "walk": true, "want": true, "warm": true, "was": true, "wash": true, "waste": true, "watch": true, "water": true, "way": true, "we": true, "weak": true, "wear": true, "weather": true, "wedding": true, "week": true, "weight": true, "welcome": true, "well": true, "were": true, "west": true, "wet": true, "what": true, "wheel": true, "when": true, "where": true, "which": true, "while": true, "white": true, "who": true, "why": true, "wide": true, "wife": true, "wild": true, "will": true, "win": true, "wind": true, "window": true, "wine": true, "winter": true, "wire": true, "wise": true, "wish": true, "with": true, "without": true, "woman": true, "wonder": true, "word": true, "work": true, "world": true, "worry": true,
	"yard": true, "yell": true, "yesterday": true, "yet": true, "you": true, "young": true, "your": true,
	"zero": true, "zoo": true,
}

const ANKI_MODEL_CSS string = `html {
    scrollbar-gutter: stable;
}

h1 {
 text-align: center;
}


.w-usage{
 line-height: 1.2;
margin-bottom: 5px;
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


div small:before {
 content: " -";
}



/* belows are from apple-dictionary-parser */

/*==== HTML ====*/

body
{
font-size: 20px !important;
    line-height: 1.6 !important;
    color: CanvasText;
    
    /* The Fix for PC and Mobile */
    max-width: 800px;           /* Prevents text from stretching too wide on PC */
    margin: 1.5em auto !important; /* Auto left/right margin centers the 800px block */
    padding: 0 1rem !important;    /* Keeps a small, safe boundary on mobile screens */
}


html.apple_client-panel body
{
	margin-top: 0em;
}

@media (prefers-color-scheme: dark)
{
	img
	{
		filter: invert(100%);
	}
}


a
{
	color: -webkit-link;
	text-decoration: none;
}

a:hover
{
	color: -webkit-link;
	text-decoration: underline;
}

html.apple_client-panel a,
html.apple_client-panel a:hover
{
	color: CanvasText;
	text-decoration: none;
}

table
{
	table-layout: fixed;
	width: 90%;
	margin-top: 1em;
	margin-left: 14px;	/*if indent 30px*/
	border-style: solid;
	border-width: 1px;
	font-weight: normal;
}

caption
{
	font-size: 80%;
	margin-bottom: 0.5em;
	text-align: left;
	text-transform: uppercase;
}

tr
{
	/* font-size: 80%; */
	width: 30%;
	-webkit-hyphens: auto;
}


/*=====================================*/
/* content                             */
/*=====================================*/

span.hsb:before
{
	/* '•' */
	content: '\2022';
	font-size: 80%;
	vertical-align: 3%;
}

span.wbb:before
{
	/* ' | ' */
	content: '\a0\7c\a0';
	font-size: 70%;
	vertical-align: 5%;
	color: -apple-system-secondary-label;
	margin-left: -0.15em;
	margin-right: -0.15em;
}

/*=====================================*/
/* stress                              */
/*=====================================*/

span.ph.t_stress span.str.t_l
{
	text-decoration: underline;
}
span.ph.t_stress span.str.t_s
{
	text-decoration: none;
}



/*=====================================*/
/* layout                              */
/*=====================================*/

*.entry
{
	display: block;
	font-size: 84%;
	margin-bottom: 1em;
	line-height: 130%;
}
*.entry + *.entry
{
	margin-top: 4em;
}


/*=====================================*/
/* xh                                  */
/*=====================================*/

*.entry > span.x_xh0
{
	display: block;
	margin-bottom: 0.6em;
}

html.apple_client-panel *.entry > span.x_xh0
{
	margin-bottom: 0.1em;
}

span.x_xh1
{
	display: block;
	/*
	margin-left: 2.2em;
	margin-top: 1em;
	*/
	margin-top: 0.6em;
	margin-bottom: 0.4em;
}
html.apple_client-panel span.x_xh1
{
	margin-top: 0.1em;
}
span.x_xh1 + span.x_xh1
{
	margin-top: 0.4em;
	margin-left: 1.0em;
}
html.apple_client-panel span.x_xh1 + span.x_xh1
{
	margin-top: 0.1em;
}

*.entry > span.x_xh0 + span.x_xd0,
*.entry > span.x_xh0 + span.sg
{
	margin-top: 0.6em;
}
html.apple_client-panel *.entry > span.x_xh0 + span.x_xd0,
html.apple_client-panel span.x_xd0
{
	margin-top: 0.1em;
}


/*=====================================*/
/* xd                                  */
/*=====================================*/

span.x_xd0
{
	display: block;
	margin-top: 0.3em;
	margin-bottom: 2.0em;
}
html.apple_client-panel span.x_xd0
{
	margin-top: 0.1em;
	margin-bottom: 0.4em;
}

/*
span.x_xd0.hasSn,
*/
span.x_xd0
{
	margin-left: 1.0em;
}
span.x_xd0 > span.x_xdh
{
	display: block;
	margin-left: -0.8em;
	margin-bottom: 0.5em;
	font-size: 120%;
	line-height: 120%;
}
html.apple_client-panel span.x_xd0 > span.x_xdh
{
	font-size: 100%;
	line-height: 100%;
	margin-bottom: 0.1em;
}

span.x_xd0.hasSn > span.x_xdh > span.sn
{
	/*
	margin-right: 0.0em;
	*/
	margin-right: 0.2em;
}

span.x_xd1
{
	display: block;
	margin-top: 0.3em;
	margin-bottom: 0.7em;
	
	margin-left: 0em;
	/* bilingual */
	/*
	margin-left: 1.0em;
	*/
}
html.apple_client-panel span.x_xd1
{
	margin-top: 0.1em;
	margin-bottom: 0.3em;
}
span.x_xd1.hasSn
{
	margin-left: 1.0em;
	text-indent: 0em;
}
span.x_xd1 > span.sn,
span.x_xd1 > span.x_xdh > span.sn
{
	margin-left: -1em;
	margin-right: 0.0em;
}

span.x_xd1 > span.x_xd2:first-child	/* bilingual */
{
	/*margin-left: -1.2em;*/
	/*
	text-indent: -1.0em;
	*/
}
span.x_xd1 > span.x_xd2
{
}

span.x_xd1 > span.x_xdh
{
	/*display: block;*/
	margin-bottom: 0.5em;
}
span.x_xd1 > span.x_xdh.sn,
span.x_xd1 > span.x_xdh > span.sn
{
	display: inline;
	margin-right: 0.1em;
}


span.x_xd1sub
{
	display: block;
}
/* first x_xd1sub (just after sn) */
span.x_xd1.hasSn span.x_xdh.sn + span.x_xd1sub
{
	display: inline;
	margin-left: 0em;
}
/* first x_xd1sub.hasSn (just after sn) */
span.x_xd1.hasSn span.x_xdh.sn + span.x_xd1sub.hasSn.t_first
{
	display: block;
	margin-left: 1em;
}
span.x_xd1sub.hasSn
{
	margin-left: 1em;
}
span.x_xd1sub.hasSn span.sn
{
	margin-left: -1em;
	margin-right: 0.1em;
}


/*
	xd2
*/
span.x_xd2
{
	display: block;
	margin-left: 0em;
	margin-top: 0.3em;
	margin-bottom: 0.3em;
}
html.apple_client-panel span.x_xd2
{
	margin-top: 0.1em;
	margin-bottom: 0.3em;
}
span.x_xd2.hasSn
{
	margin-left: 1.0em;
}
span.x_xd2.hasSn span.sn
{
	margin-left: -1.0em;
	margin-right: 0.2em;
}
/*
span.x_xd2 span.sn
{
	margin-left: -1em;
	margin-right: 0.0em;
}
*/

/* first x_xd2 (just after sn) */
span.x_xd1.hasSn span.x_xdh.sn + span.x_xd2
{
	display: inline;
	margin-left: 0em;
}
span.x_xd1.hasSn span.x_xdh.sn + span.x_xd2.hasSn
{
	display: block;
	margin-left: 1.0em;
}
span.x_xd2.hasSn span.x_xdh.sn + span.x_xd3
{
	display: inline;
	margin-left: 0em;
}
span.x_xd2.hasSn span.x_xdh.sn + span.x_xd3.hasSn
{
	display: block;
	margin-left: 1.0em;
}
span.x_xd3.hasSn span.x_xdh.sn + span.x_xd4
{
	display: inline;
	margin-left: 0em;
}
span.x_xd3.hasSn span.x_xdh.sn + span.x_xd4.hasSn
{
	display: block;
	margin-left: 1.0em;
}
span.x_xd4.hasSn span.x_xdh.sn + span.x_xd5
{
	display: inline;
	margin-left: 0em;
}
span.x_xd4.hasSn span.x_xdh.sn + span.x_xd5.hasSn
{
	display: block;
	margin-left: 1.0em;
}
span.x_xd5.hasSn span.x_xdh.sn + span.x_xd6
{
	display: inline;
	margin-left: 0em;
}
span.x_xd5.hasSn span.x_xdh.sn + span.x_xd6.hasSn
{
	display: block;
	margin-left: 1.0em;
}


span.x_xo3.hasSn span.x_xdh.sn + span.x_xo4
{
	display: inline;
	margin-left: 0em;
}
span.x_xo3.hasSn span.x_xdh.sn + span.x_xo4.hasSn
{
	display: block;
	margin-left: 1.0em;
}

span.x_xo2.hasSn span.x_xoh.sn + span.x_xo3
{
	display: inline;
	margin-left: 0em;
}
span.x_xo2.hasSn span.x_xoh.sn + span.x_xo3.hasSn
{
	display: block;
	margin-left: 1.0em;
}
span.x_xo3.hasSn span.x_xoh.sn + span.x_xo4
{
	display: inline;
	margin-left: 0em;
}
span.x_xo3.hasSn span.x_xoh.sn + span.x_xo4.hasSn
{
	display: block;
	margin-left: 1.0em;
}



/*
span.x_xd1 > span.x_xd2 > span.x_xdh
{
	color: red;
	margin-left: -1em;
}
*/

/*
	xd0, xd1 - de
	Cascading sn
*/
span.x_xd0.cHasSn.gcHasSn
{
	margin-left: 2.0em;
}
span.x_xd1.cHasSn.gcHasSn
{
	margin-left: 1.0em;
}

span.x_xd0.cHasSn.gcHasSn span.x_xd1.hasSn.cHasSn,
span.x_xd1.cHasSn.gcHasSn span.x_xd2.hasSn.cHasSn
{
	margin-left: 1.2em;
}
span.x_xd0.cHasSn.gcHasSn span.x_xd1.hasSn,
span.x_xd1.cHasSn.gcHasSn span.x_xd2.hasSn
{
	margin-left: 0em;
}
/*
span.x_xd0.cHasSn.gcHasSn span.x_xd1.hasSn:not(.cHasSn),
span.x_xd1.cHasSn.gcHasSn span.x_xd2.hasSn:not(.cHasSn)
{
	margin-left: 0em;
}
*/
span.x_xd0.cHasSn.gcHasSn span.x_xd1.hasSn.cHasSn > span.sn,
span.x_xd1.cHasSn.gcHasSn span.x_xd2.hasSn.cHasSn > span.sn
{
	margin-left: -2.2em;
	margin-right: 2.3em;
}

span.x_xd1 span.x_xdh span.sn
{
	margin-right: 0.2em;
}


/*
span.x_xd2.hasSn
span.x_xd3.hasSn
*/
span.x_xd0.cHasSn.gcHasSn span.x_xd2.hasSn,
span.x_xd0.cSn.gcSn span.x_xd2.hasSn,
span.x_xd1.cHasSn.gcHasSn span.x_xd3.hasSn,
span.x_xd1.cSn.gcSn span.x_xd3.hasSn
{
	margin-left: 0em;
}
/*
span.x_xd2.hasSn span.sn
span.x_xd3.hasSn span.sn
*/
span.x_xd0.cHasSn.gcHasSn span.x_xd2.hasSn span.sn,
span.x_xd0.cSn.gcSn span.x_xd2.hasSn span.sn,
span.x_xd1.cHasSn.gcHasSn span.x_xd3.hasSn span.sn,
span.x_xd1.cSn.gcSn span.x_xd3.hasSn span.sn
{
	margin-left: -0.8em;
	margin-right: 0.0em;
}
/*
span.x_xd1.hasSn span.x_xd2.hasSn
span.x_xd2.hasSn span.x_xd3.hasSn
*/
span.x_xd0.cHasSn.gcHasSn span.x_xd1.hasSn span.x_xd2.hasSn,
span.x_xd0.cSn.gcSn span.x_xd1.hasSn span.x_xd2.hasSn,
span.x_xd1.cHasSn.gcHasSn span.x_xd2.hasSn span.x_xd3.hasSn,
span.x_xd1.cSn.gcSn span.x_xd2.hasSn span.x_xd3.hasSn
{
}

/*
span.x_xd1.hasSn span.x_xd2.hasSn span.sn
span.x_xd2.hasSn span.x_xd3.hasSn span.sn
*/
span.x_xd0.cHasSn.gcHasSn span.x_xd1.hasSn span.x_xd2.hasSn span.sn,
span.x_xd0.cSn.gcSn span.x_xd1.hasSn span.x_xd2.hasSn span.sn,
span.x_xd1.cHasSn.gcHasSn span.x_xd2.hasSn span.x_xd3.hasSn span.sn,
span.x_xd1.cSn.gcSn span.x_xd2.hasSn span.x_xd3.hasSn span.sn
{
	margin-left: -1.2em;
	margin-right: 0.1em;
}
/*
span.x_xd1.hasSn > span.sn + span.x_xd2.hasSn > span.sn
span.x_xd2.hasSn > span.sn + span.x_xd3.hasSn > span.sn
*/
span.x_xd0.cHasSn.gcHasSn span.x_xd1.hasSn > span.sn + span.x_xd2.hasSn > span.sn,
span.x_xd0.cSn.gcSn span.x_xd1.hasSn > span.sn + span.x_xd2.hasSn > span.sn,
span.x_xd1.cHasSn.gcHasSn span.x_xd2.hasSn > span.sn + span.x_xd3.hasSn > span.sn,
span.x_xd1.cSn.gcSn span.x_xd2.hasSn > span.sn + span.x_xd3.hasSn > span.sn
{
	margin-left: -3.6em;
	margin-right: 0.2em;
}
/* first x_xd2, x_xd3 (just after sn) */
/*
span.x_xd1.hasSn span.x_xdh.sn + span.x_xd2
span.x_xd2.hasSn span.x_xdh.sn + span.x_xd3
*/
span.x_xd0.cHasSn.gcHasSn span.x_xd1.hasSn span.x_xdh.sn + span.x_xd2,
span.x_xd0.cSn.gcSn span.x_xd1.hasSn span.x_xdh.sn + span.x_xd2,
span.x_xd1.cHasSn.gcHasSn span.x_xd2.hasSn span.x_xdh.sn + span.x_xd3,
span.x_xd1.cSn.gcSn span.x_xd2.hasSn span.x_xdh.sn + span.x_xd3
{
	display: inline;
	margin-left: 0.0em;
}
/*
span.x_xd1.hasSn span.x_xdh.sn + span.x_xd2.hasSn
span.x_xd2.hasSn span.x_xdh.sn + span.x_xd3.hasSn
*/
span.x_xd0.cHasSn.gcHasSn span.x_xd1.hasSn > span.x_xdh.sn + span.x_xd2.hasSn,
span.x_xd0.cSn.gcSn span.x_xd1.hasSn > span.x_xdh.sn + span.x_xd2.hasSn,
span.x_xd1.cHasSn.gcHasSn span.x_xd2.hasSn > span.x_xdh.sn + span.x_xd3.hasSn,
span.x_xd1.cSn.gcSn span.x_xd2.hasSn > span.x_xdh.sn + span.x_xd3.hasSn
{
	display: inline;
	margin-left: 1.4em;
}
/*
*/
span.x_xd1.hasSn span.x_xdh.sn + span.x_xd2.hasSn
{
	display: block;
	margin-left: 1.0em;
}
span.x_xo3.hasSn span.x_xdh.sn + span.x_xo4.hasSn
{
	display: block;
	margin-left: 1.0em;
}


/* de: lg + msDict */
/*
span.x_xd1 > span.x_xdh + span.x_xd2
*/
/*
span.x_xd1 > span.x_xdh + span.x_xd2.t_first
{
	display: inline;
}
*/
span.x_xd1 > span.sn + span.x_xd2.t_first
{
	margin-left: 1.2em;
}
span.x_xd1 > span.x_xdh > span.sn + span.x_xo0,
span.x_xd2 > span.x_xdh > span.sn + span.x_xo0
{
	display: inline;
}


/*
	x_xd2sub
*/
span.x_xd2sub
{
	display: block;
}
span.x_xd2sub.hasSn
{
	margin-left: 1em;
}
span.x_xd2sub.hasSn span.sn
{
	margin-left: -0.9em;
	margin-right: 0.1em;
}
/* first x_xd2sub (just after sn) */
span.x_xd2.hasSn span.x_xdh.sn + span.x_xd2sub
{
	display: inline;
	margin-left: 0.2em;
}


/*
	xd3
*/
span.x_xd3
{
	display: block;
	margin-left: 0em;
	margin-top: 0.2em;
	margin-bottom: 0.2em;
}
span.x_xd3.hasSn
{
	margin-left: 1.0em;
}
span.x_xd3.hasSn span.sn
{
	margin-left: -1.0em;
}

/*
	xd4
*/
span.x_xd4
{
	display: block;
	margin-left: 0em;
	margin-top: 0.2em;
	margin-bottom: 0.2em;
}
span.x_xd4.hasSn
{
	margin-left: 1.0em;
}
span.x_xd4.hasSn span.sn
{
	margin-left: -1.0em;
}

/*
	xd5
*/
span.x_xd5
{
	display: block;
	margin-left: 0em;
	margin-top: 0.2em;
	margin-bottom: 0.2em;
}
span.x_xd5.hasSn
{
	margin-left: 1.0em;
}
span.x_xd5.hasSn span.sn
{
	margin-left: -1.0em;
}

/*
	xd6
*/
span.x_xd6
{
	display: block;
	margin-left: 0em;
	margin-bottom: 0.2em;
}
span.x_xd6.hasSn
{
	margin-left: 1.0em;
}
span.x_xd6.hasSn span.sn
{
	margin-left: -1.0em;
}





/*=====================================*/
/* xo                                  */
/*=====================================*/

span.x_xo0
{
	display: block;
	margin-top: 0.5em;
	margin-bottom: 0.5em;
	margin-left: 0em;
}
html.apple_client-panel span.x_xo0
{
	margin-top: 0.1em;
	margin-bottom: 0.2em;
}
*.entry > span.x_xo0
{
	margin-top: 1em;
	margin-bottom: 1em;
}
html.apple_client-panel *.entry > span.x_xo0
{
	margin-top: 0.3em;
	margin-bottom: 0.3em;
}

span.x_xdt > span.x_xo0
{
	margin-bottom: 1em;
}

span.x_xo0 > span.x_xoLblBlk
{
	display: block;
	font-variant: small-caps;
	font-size: 90%;
	display: block;
	padding-bottom: 0.3em;
	border-bottom: solid thin -apple-system-tertiary-label;
	color: -apple-system-secondary-label;
	margin-top: 2em;
	margin-bottom: 0.5em;
}
html.apple_client-panel span.x_xo0 > span.x_xoLblBlk
{
	margin-top: 1em;
}

div.xrgBlock.x_xoLblBlkTop,
span.x_xo0.x_xoLblBlkTop
{
	margin-top: 3em;
	border-top: solid thin -apple-system-tertiary-label;
	margin-bottom: 1em;
}
html.apple_client-panel div.xrgBlock.x_xoLblBlkTop,
html.apple_client-panel span.x_xo0.x_xoLblBlkTop
{
	margin-top: 1em;
}
div.xrgBlock.x_xoLblBlkTop > span.label,
span.x_xo0.x_xoLblBlkTop > span.label
{
	display: block;
	margin-top: 1em;
}
html.apple_client-panel div.xrgBlock.x_xoLblBlkTop > span.label,
html.apple_client-panel span.x_xo0.x_xoLblBlkTop > span.label
{
	margin-top: 0.5em;
}

div.xrgBlock span.xrgs
{
	display: block;
	margin-left: 2em;
}
div.xrgBlock span.xrgs span.xrg
{
	display: block;
}

span.x_xo1
{
	display: block;
	margin-top: 0.5em;
	margin-bottom: 0.5em;
	margin-left: 0em;
}
html.apple_client-panel span.x_xo1
{
	margin-top: 0.1em;
	margin-bottom: 0.1em;
}

span.x_xd0 > span.x_xo0
{
	margin-left: -1em;
}

span.x_xo0.note
{
	margin-top: 1.3em;
	margin-bottom: 1.3em;
}


/* IT */
span.x_xd0 > span.x_xo1
{
	margin-left: 0em;
}
span.x_xd0 > span.x_xo1.hasSn
{
	padding-left: 1em;
}
span.x_xd0 > span.x_xo1.hasSn span.sn
{
	margin-left: -0.9em;
}

/* ES */
span.x_xd2.msDict > span.synGroup
{
	display: block;
}

/* NL */
span.x_xd1sub + span.x_xo0
{
	margin-top: 0.8em;
}


/* DE : idiom in sense block */
/*
span.x_xd0 span.x_xd2 span.x_xo1
{
	margin-left: 0em;
	margin-top: 0.4em;
}
span.x_xd0 span.x_xd2 span.x_xo1 span.x_xo2
{
	display: inline;
}
span.x_xd0 span.x_xd2 span.x_xo1 span.x_xoh + span.x_xo2
{
	margin-left: 0.6em;
}
*/

span.x_xo1
{
	display: block;
	margin-left: 1em;
	/*margin-bottom: 0.8em;*/
	margin-top: 0.5em;
	margin-bottom: 0.5em;
}
html.apple_client-panel span.x_xo1
{
	margin-bottom: 0.3em;
}
/* subEntry in Definition */
span.x_xd1 span.x_xo1
{
	margin-left: 0em;
}

span.x_xo2
{
	display: block;
	margin-left: 1em;
	margin-top: 0.4em;
	margin-bottom: 0.4em;
}
/*
span.x_xd1 span.x_xo2
{
	margin-left: 0em;
}
*/
span.x_xo2.hasSn
{
	margin-left: 2em;
}
span.x_xo2.hasSn > span.sn,
span.x_xo2.hasSn >span.x_xdh > span.sn,
span.x_xo2.hasSn >span.x_xoh > span.sn
{
	margin-left: -1em;
	margin-right: 0.2em;
}

span.x_xo2sub
{
	display: block;
}
/* first x_xo2sub (just after sn) */
span.x_xo2.hasSn span.x_xoh.sn + span.x_xo2sub
{
	display: inline;
	margin-left: 0.1em;
}
span.x_xo2sub.hasSn
{
	margin-left: 1em;
}
span.x_xo2sub.hasSn span.sn
{
	margin-left: -1em;
	margin-right: 0.1em;
}

span.x_xo2.hasSn span.x_xdh.sn + span.x_xo3,
span.x_xo2.hasSn span.x_xdh.sn + span.x_xo3 > span.x_xo4.trg:first-child
{
	/* AR-EN */
	display: inline;
	margin-left: 0em;
}

span.x_xo3
{
	display: block;
	/*
	margin-left: 1em;
	*/
	margin-top: 0.3em;
	margin-bottom: 0.3em;
}
span.x_xo3.hasSn
{
	margin-left: 1em;
}
span.x_xo3.hasSn > span.sn,
span.x_xo3.hasSn > span.x_xoh > span.sn,
span.x_xo3.hasSn > span.x_xdh > span.sn
{
	margin-left: -1em;
	margin-right: 0.2em;
}

span.x_xo4
{
	display: block;
	/*
	margin-left: 1em;
	*/
	margin-top: 0.3em;
	margin-bottom: 0.3em;
}
span.x_xo4.hasSn
{
	margin-left: 1em;
}
span.x_xo4.hasSn span.sn
{
	margin-left: -1em;
	margin-right: 0.1em;
}

span.x_xo5
{
	display: block;
	margin-top: 0.1em;
	margin-bottom: 0.1em;
}
span.x_xo5.hasSn
{
	margin-left: 1em;
}

span.x_xo5.hasSn span.sn
{
	margin-left: -1em;
	margin-right: 0.1em;
}

/* first x_xo5 (just after sn) */
span.x_xo4.hasSn > span.x_xoh.sn + span.x_xo5
{
	display: inline;
	margin-left: 0em;
}

span.x_xo6
{
	display: block;
	margin-top: 0.1em;
	margin-bottom: 0.1em;
}
span.x_xo6.hasSn
{
	margin-left: 1em;
}
span.x_xo6.hasSn span.sn
{
	margin-left: -1em;
	margin-right: 0.1em;
}



/*=====================================*/
/* Hide                                */
/*=====================================*/

/* ▶ */
span.x_xd1 > span.x_xdh > span.gp.tg_se1
{
	display: none;
}

span.construction	{ display: none; }

span.cdf	{ display: none; }

span.domClass	{ display: none; }
span.semClass	{ display: none; }
span.encClass	{ display: none; }

span.morphSet	{ display: none; }

span.meta		 { display: none; }

/* 
span.hg span.gg.t_wordbreak
{
	display: none;
}

span.hg > span.lg 
{
	display: none;
}

span.hg > span.hw + span.posg
{
	display: none;
}
*/



/*=====================================*/
/* Style                               */
/*=====================================*/


/*=====================================*/
/* hw                                  */
/*=====================================*/

span.hg span.hw,
span.hg span.vg,
span.hwg span.hw,
span.ehw,
span.vhw,
/* span.hwg span.hvg span.hv, */
span.hwg span.hvg,
span.hg span.gp.t_hw
{
	font-size: 170%;
	font-weight: normal;
	line-height: 150%;
	/*
	line-height: 130%;
	*/
	-webkit-hyphens: auto;
}
span.hvg > span.hm
{
	font-size: 59%;
}
html.apple_client-panel span.hvg > span.hm
{
	font-size: 80%;
}
span.hg > span.hgSub1 span.vg
{
	font-size: 100%;
}
span.hg span.vg.x_weak,
span.hg span.vg.x_weak span.v
{
	font-size: 100%;
	color: -apple-system-secondary-label;
}

span.gramb span.hvg span.hv
{
	font-weight: 600;
}
span.x_xh0 span.vg
{
	font-weight: 200;
	color: -apple-system-secondary-label;
}

span.x_xh0 span.vg span.v
{
	font-weight: normal;
	color: CanvasText;
}
span.x_xh0 span.vg span.t_label
{
	font-weight: normal;
}

span.hwg span.frmg span.frm
{
	font-weight: normal;
	font-size: 170%;
}
span.hwg span.hvg span.frmg span.frm
{
	font-size: 100%;
}
span.frm_type
{
	font-weight: normal;
}
span.hwg span.frmg span.gr,
span.hwg span.frmg span.gp.tg_frmg
{
	font-size: 150%;
}
span.hwg span.hvg span.frmg span.gr,
span.hwg span.hvg span.frmg span.gp.tg_frmg
{
	font-size: 100%;
}
span.hwg span.frmg span.gp.tg_frmg,
span.hwg span.infg span.gp.tg_infg,
span.hwg span.infg span.gp.tg_inf,
span.hw span.gp.tg_hw,
span.hvg span.hv span.gp.tg_hv
{
	color: -apple-system-secondary-label;
}

span.x_xd1 > span.posg.x_xdh,
span.x_xd1 > span.x_xdh > span.posg:first-child
{
	font-size: 120%;
}
html.apple_client-panel span.x_xd1 > span.posg.x_xdh
{
	font-size: 110%;
}

html.apple_client-panel span.hg span.hw,
html.apple_client-panel span.hg span.vg,
html.apple_client-panel span.hwg span.hw,
html.apple_client-panel span.ehw,
html.apple_client-panel span.vhw,
/* html.apple_client-panel span.hwg span.hvg span.hv, */
html.apple_client-panel span.hwg span.hvg,
html.apple_client-panel span.hg span.gp.t_hw,
html.apple_client-panel span.hwg span.frmg span.frm,
html.apple_client-panel span.hwg span.frmg span.gr,
html.apple_client-panel span.hwg span.frmg span.gp.tg_frmg
{
	font-size: 120%;
	line-height: 130%;
}

span.hg span.fg
{
	font-size: 150%;
}
html.apple_client-panel span.hg span.fg
{
	font-size: 100%;
}

span.hg span.fg span.f	{ font-weight: normal; }
span.hg span.fg span.sy	{ font-style: italic; }

span.fhw
{
	font-weight: 500;
	font-style: italic;
}

span.hm
{
	vertical-align: super;
}
span.gp.ty_hom,
span.xr span.hm
{
	vertical-align: super;
	font-size: 70%;
}
span.hg > span.gp.ty_hom
{
	font-size: 120%;
	vertical-align: 50%;
}
html.apple_client-panel span.hg > span.gp.ty_hom
{
	font-size: 100%;
	vertical-align: super;
}

/* hgSub1 */

span.hg > span.hgSub1
{
	font-size: 110%;
	margin-left: 0.1em;
}
html.apple_client-panel span.hg > span.hgSub1
{
	font-size: 100%;
}

span.hg > span.hgSub1 span.alt-hg span.hw
{
	display: block;
	font-size: 144%;
	margin-top: 0.3em;
}
html.apple_client-panel span.hg > span.hgSub1 span.alt-hg span.hw
{
	margin-top: 0.3em;
	font-size: 111%;
}
span.hg > span.hgSub1 span.posg,
span.hwg > span.hgSub1 span.ps
{
	font-size: 110%;
}
span.hg > span.hgSub1 span.pr,
span.hg > span.hgSub1 span.prx
{
	font-size: 100%;
}
span.hgSub1 > span.etym,
span.hgSub1 > span.grnote
{
	display: block;
}
span.hgSub1 > span.etym.x_nblk
{
	display: inline;
}

/*
span.hg span.hgSub1 span.etym
{
	display: inline;
	margin-right: 0.3em;
}
*/

span.hg > span.hgSub2
{
	display: block;
	font-size: 100%;
	/* margin-left: 2em; */
}

span.hg > span.hgSub2 span.fg,
span.hg > span.hgSub2 span.infg,
span.hg > span.hgSub2 span.vg
{
	display: block;
	font-size: 100%;
	margin-top: 0.3em;
}


/*=====================================*/
/* inf                                 */
/*=====================================*/

span.InfGrp,
span.iGrp
{
	font-weight: normal;
}
/*
span.infg,
span.gg
{
	vertical-align: 1%;
}
*/
span.inf
{
	font-weight: 600;
	/*
	font-style: italic;
	*/
}
span.infg span.sy + span.inf
{
	margin-left: 0.1em;
}

span.hg span.hw + span.infg,
span.hg span.hw + span.gp.ty_hom + span.infg
{
	font-size: 170%;
	font-weight: normal;
}
span.hvg span.infg
{
	font-size: 58.8%;
}
span.hg span.hw + span.infg span.inf,
span.hg span.hw + span.gp.ty_hom + span.infg span.inf
{
	font-weight: normal;
}
html.apple_client-panel span.hg span.hw + span.infg,
html.apple_client-panel span.hg span.hw + span.gp.ty_hom + span.infg
{
	font-size: 130%;
}
html.apple_client-panel span.hvg span.infg
{
	font-size: 76.9%;
}

/*=====================================*/
/* f                                   */
/*=====================================*/

span.formGrp,
span.fGrp
{
	font-weight: normal;
}
span.f,
span.form
{
	font-weight: 600;
}

span.x_xdh > span.sn + span.fg + span.posg
{
	margin-left: 0.3em;
}


/*=====================================*/
/* POS                                 */
/*=====================================*/

/*
span.x_xh0 span.posg span.pos,
span.x_xd0 span.x_xdh span.posg span.pos
{
	font-size: 120%;
}
*/
/*
span.hg span.pos,
span.hwg span.pos
*/
span.hg span.posg,
span.hwg span.posg
{
	margin-right: 0.3em;
}
span.gp.tg_pos
{
	margin-right: 0.3em;
}

span.pos.x_rr,
span.ps.x_rr,
span.hwg > span.hgSub1 span.ps.x_rr,
span.lbl.x_rr,
span.ph.x_rr,
span.sn.x_rr,
span.gr.x_rr
{
	border: solid 1px -apple-system-secondary-label;
	/* margin-left: 0.12em; */
	margin-right: 0.18em;
	margin-top: 0.18em;
	margin-bottom: 0.18em;
	padding: 1px 1px 1px 1px;
	font-size: 70%;
	-webkit-border-radius: 2px;
	vertical-align: 5%;
}
span.lbl.x_rr
{
	margin-right: 0.4em;
	vertical-align: 10%;
}
span.note span.lbl.x_rr
{
	vertical-align: 5%;
}
span.subEntry span.pos.x_rr
{
	font-size: 70%;
}
span.ph.x_rr
{
	padding: 2px 3px 0px 3px;
	margin-left: 0.2em;
	margin-right: 0.2em;
}
span.ps.x_rr,
span.gr.x_rr
{
	margin-right: 0.3em;
}

span.lbl.x_rrbd
{
	border: solid -apple-system-secondary-label;
	border-width: 2px;
	padding: 2px 4px 1px 4px;
	-webkit-border-radius: 3px;
	color: -apple-system-secondary-label;
	font-weight: 600;
	font-size: 70%;
	margin-right: 0.5em;
	vertical-align: 20%;
}
html.apple_client-panel span.lbl.x_rrbd
{
	vertical-align: 15%;
}

span.msDict > span.posg:first-child,
span.se2 > span.posg:first-child,
span.x_xdh > span.sn + span.posg
{
	margin-right: 0.2em;
}

span.x_xd0 > span.ps.x_xdh.x_rr
{
	display: inline;
}

span.x_xd0 > span.x_xdh.ps.x_rr
{
	font-size: 84%;
}

/*=====================================*/
/* pr                                  */
/*=====================================*/

span.pr,
span.prx
{
	font-weight: normal;
	color: -apple-system-secondary-label;
	/*
	font-family: 'Times New Roman';
	*/
}

span.x_xh0 span.gp.tg_pr,
span.x_xh0 span.gp.tg_prx,
span.x_xd0 span.x_xdh span.gp.tg_pr,
span.x_xd0 span.x_xdh span.gp.tg_prx
{
	color: -apple-system-tertiary-label;
}

span.hg span.pr,
span.hg span.prx,
span.hwg span.pr,
span.hwg span.prx
{
	display: inline-block;
	/*
	font-size: 130%;
	*/
	font-size: 120%;
	line-height: 130%;
	/* vertical-align: 12%; */
	margin-left: 0.4em;
	margin-right: 0.4em;
}
span.prx + span.prx
{
	margin-left: 0.0em;
}
span.hg span.hw span.pr,
span.hg span.v span.pr
{
	font-size: 80%;
	margin-left: 0em;
	margin-right: 0em;
}
span.hwg span.hvg span.pr,
span.hwg span.hvg span.prx,
span.hg span.vg span.pr,
span.hg span.vg span.prx
{
	font-size: 70%;
}
html.apple_client-panel span.hwg span.hvg span.pr,
html.apple_client-panel span.hg span.vg span.pr
{
	font-size: 100%;
}


/*
span.hg > span.pr + span
{
	margin-left: 0.6em;
}
*/

span.ph
{
/*
	margin-left: 0.3em;
	margin-right: 0.3em;
*/
}

span.cnt
{
	font-weight: 200;
	color: -apple-system-secondary-label;
}
span.cnt.x_rr
{
	margin-left: 0.2em;
	margin-right: 0.2em;
	padding-left: 0.2em;
	padding-right: 0.2em;
	color: CanvasText;
	border: solid 1px -apple-system-secondary-label;
	-webkit-border-radius: 2px;
	font-size: 70%;
	/* vertical-align: 10%; */
}


/*=====================================*/
/* sn                                  */
/*=====================================*/

span.sn
{
	font-weight: 600;
	color: -apple-system-secondary-label;
}
span.sn.x_rrbg
{
	background: -apple-system-secondary-label;
	border: solid -apple-system-secondary-label;
	border-width: 0px;
	padding: 3px 2px 1px 2px;
	-webkit-border-radius: 2px;
	color: -apple-system-text-background;
	font-weight: 600;
	font-size: 70%;
	vertical-align: 10%;
	margin-right: 0.5em;
}


/*=====================================*/
/* ex                                  */
/*=====================================*/

span.eg span.ex
{
	/* font-weight: normal; */
	font-weight: 200;
	font-style: italic;
}
span.ex span.bold
{
	font-weight: 600;
	font-style: italic;
}
span.eg span.ex > span.italic
{
	font-style: normal;
}
span.eg span.ref
{
	font-weight: 200;
}

/* TH */
span.eg span.general-text.t_eg_pr
{
	font-weight: 200;
	color: -apple-system-secondary-label;
}
/* TR */
span.eg span.general-text.t_eg_pri
{
	font-style: italic;
	font-weight: 200;
	color: -apple-system-secondary-label;
}
span.gp.t_arrow,
span.general-text.t_arrow
{
	font-size: 90%;
	vertical-align: -2%;
	color: -apple-system-secondary-label;
}


/* bilingual */
span.exg span.ex
{
	font-style: normal;
	font-weight: 500;
	color: -apple-system-secondary-label;
}
span.exg span.ex span.bold
{
	font-weight: normal;
}
span.reg,
span.hwg span.reg,
span.semb span.reg,
span.semb span.ge,
span.semb span.sj,
span.exg span.ex span.reg,
span.exg span.ex span.ge,
span.exg span.ex span.sj
{
	font-weight: normal;
	color: -apple-system-secondary-label;
}
span.hwg span.reg
{
	font-size: 170%;
	color: -apple-system-tertiary-label;
}
html.apple_client-panel span.hwg span.reg
{
	font-size: 120%;
}
span.hwg span.hvg span.reg,
span.hwg span.pr span.reg,
span.hwg span.prx span.reg
{
	font-size: 100%;
	color: -apple-system-tertiary-label;
}

span.lbl
{
	color: -apple-system-secondary-label;
}
span.pronBlock span.lbl
{
	font-weight: 500;
	margin-right: 0.3em;
}

span.lbl.x_sqb
{
	background: -apple-system-secondary-label;
	border: solid -apple-system-secondary-label;
	border-width: 0px;
	padding: 2px 4px 2px 4px;
	font-size: 80%;
	color: -apple-system-text-background;
	font-weight: 600;
	margin-right: 1em;
	vertical-align: 5%;
}


/*=====================================*/
/* note                                */
/*=====================================*/

span.note,
span.cntxt
{
	display: block;
	/*
	background-color: window;
	padding: 0.33em;
	margin-top: 0.3em;
	*/
	border: solid 1px -apple-system-tertiary-label;
	padding: 0.13em 0.49em 0.13em 0.49em;
	-webkit-border-radius: 3px;
	margin-top: 0.1em;
}
span.note span.lbl,
span.cntxt span.lbl
{
	color: -apple-system-secondary-label;
}
span.note.x_nb
{
	display: inline;
	background-color: transparent;
	padding: 0em;
	margin-top: 0em;
	margin-right: 0em;
}
/*
span.paragraph,
span.para
{
	display: block;
	text-indent: 1em;
	margin-top: 0.25em;
}
*/

/*=====================================*/
/* subEntry                            */
/*=====================================*/

span.l,
span.subEntry span.l ~ span.vg span.v
{
	font-weight: 600;
}
span.drv
{
	font-weight: 600;
}
span.drvg span.drv
{
	margin-right: 0.5em;
}

span.subEntry span.pos
{
	font-size: 100%;
}

span.subEntry > span.l + span.pr,
span.subEntry > span.pr + span.posg
{
	margin-left: 0.6em;
}

/*=====================================*/
/* etym                                */
/*=====================================*/

/*
span.etym	{font-weight:normal;}
*/
span.etym > span.lbl
{
	color: -apple-system-secondary-label;
}
span.etym span.dgg
{
	display: block;
	margin-top: 0.2em;
	margin-right: 1em;
}


/*=====================================*/
/* Other                               */
/*=====================================*/

span.lg,
span.gg
{
	font-style: italic;
	color: -apple-system-secondary-label;
	font-size: 94%;
	/*
	margin-right: 0.3em;
	*/
}
/*
span.gp,
span.lg,
span.lbl
{
	font-style: normal;
}
*/

span.Name	 {font-style: italic;}
span.altName	 { font-weight: 600; }

span.cite[highlight="true"]
{
	font-style: italic;
}
span.cite[highlight="false"]
{
	font-weight: normal;
}

span.co
{
	font-style: italic;
}
span.co.x_ni
{
	font-style: normal;
}
span.com
{
	font-weight: normal;
}
span.con
{
	font-style: normal;
	font-weight: 600;
	color: -apple-system-secondary-label;
}


span.date			{font-weight: normal;}
span.date + span.italic	{font-style: italic;}
span.lang			{font-weight: normal;}

span.ex > span.lbl	{ font-weight: normal; }

span.ff
{
	font-style: italic;
	font-weight: 600;
}
span.rel
{
	font-weight: 600;
}

span.sp
{
	font-weight: 600;
}
span.sp span.gp
{
	font-weight: normal;
}

span.trans
{
	font-weight: 500;
}
span.trg span.trans
{
	font-weight: normal;
}
span.trg.x_xd2 > span.trans,
span.trgg.x_xd2 > span.trg > span.trans
{
	font-weight: 500;
}
span.trg.x_xd2 > span.trans span.tgr,
span.trgg.x_xd2 > span.trg > span.trans span.tgr
{
	font-weight: normal;
}
span.x_xdh.sn:first-child + span.exg.x_xd2 > span.trg.x_xd3 > span.trans
{
	font-weight: 500;
}
span.x_xdh.sn:first-child + span.exg.x_xd2 > span.trg.x_xd3 > span.trans span.tgr
{
	font-weight: normal;
}

span.semb.x_xd2 > span.trg.x_xd3 > span.trans
{
	font-weight: 500;
}

span.efg span.ef
{
	font-weight: normal;
}

span.finf	 {font-style: italic;}
span.fl		 {font-weight: 600;}
span.lp[highlight="true"] {font-weight: 600;}
span.lp[highlight="false"] { font-style: normal;}

span.hwtrans	 { font-weight: normal;}

span.oup_label
{
	font-style: italic;
	color: -apple-system-secondary-label;
}
span.oup_label.x_ni
{
	font-style: normal;
}

/*
span.df > span.lg > span.oup_label,
span.gp + span.lg > span.oup_label
{
	margin-left: 0.2em;
}
*/

/*
span.reg,
span.ge,
span.sj
{
	font-size: 80%;
}
*/

span.xr
{
	font-weight: 500;
}
span.xr span.gp
{
	font-variant: normal;
}


span.hg span.hw span.sy_underline,
span.hg span.vg span.v span.sy_underline,
span.sy_underline,
span.pr span.ph span.str
{
	text-decoration: underline;
}



/*
span.lg + span.xrg
{
	margin-left: 0.4em;
}
span.xrg + span.xrg
{
	margin-left: 0.3em;
}
*/

span.xu	{font-weight: 600;}


/*
span.def
{
	font-weight: normal;
}
*/

/*
span.MS > span.lbl {
	font-weight: 500;
	font-size: 80%;
}
*/

/*
span.gramGrp,
span.gGrp
{
	font-weight: normal;
}
span.gg,
span.gp.tg_gg
{
	font-weight: normal;
}
*/

/*
span.techGrp>span.lbl, span.techMS>span.lbl { font-weight: 500; font-size: 80%;}

span.txn[highlight="true"] {font-style:italic;}
span.txn[highlight="false"] { font-weight:normal;}

span.txnBin	 {font-style:italic;}
span.txnFam	 {font-weight:normal;}
span.txnGenus	 { font-weight:normal;}



span.render	{font-style:italic;}
span.render[as="sans-serif"]	{ font-family:'HelveticaNeue';}
span.render[as="bold"] 		{font-weight: 600;}
span.render[as="italic"] 		{font-style:italic;} 
*/


/*=====================================*/
/* General                             */
/*=====================================*/

span.bold
{
	font-weight: 600;
}
span.italic
{
	font-style: italic;
}
span.underline
{
	text-decoration: underline;
	/*
	font-style: italic;
	font-weight: normal;
	*/
	/* color: #444; */
}
span.r
{
	font-weight: normal;
	font-style: normal;
}
span.sup
{
	vertical-align: super;
	font-size: 70%;
}
span.sub,
span.subEnt
{
	vertical-align: sub;
	font-size: 70%;
}
span.sc
{
	font-variant: small-caps;
}

span.bi			 {font-weight: 500; font-style: italic;}
span.in			 { vertical-align: sub; font-size: 70%}
span.ini		 { vertical-align: sub; font-size: 70%; font-style:italic;}
span.su			 { vertical-align: super; font-size: 70%}
span.sub		 { vertical-align: sub; font-size: 70%}
span.sui		 { vertical-align: super; font-size: 70%; font-style:italic;}
span.sup		 { vertical-align: super; font-size: 70%}
span.subEnt		 { vertical-align: sub; font-size: 70%}
span.nu			 { vertical-align: super; font-size: 70%}
span.dn			 { vertical-align: sub; font-size: 70%}


span.usageLbl
{
	font-weight: normal;
	font-size: 90%;
	color: -apple-system-secondary-label;
}
span.hw span.usageLbl
{
	font-size: 80%;
}


span.general-text.t_label,
span.general-text.t_label_block
{
	font-variant: small-caps;
	margin-right: 0.3em;
	color: -apple-system-secondary-label;
}
span.general-text.t_label_block
{
	display: block;
	/* margin-top: 0.4em; */
}
span.gp.t_label.tg_synList,
span.gp.t_label.tg_antList,
span.gp.t_label.tg_see,
span.gp.t_label.tg_REMISIO,
span.gp.t_label.tg_etym,
span.etym span.gp.t_label,
span.dgg span.gp.t_label
{
	margin-right: 0.3em;
	color: -apple-system-secondary-label;
}

span.hw span.general-text.en_er,
span.ehw span.general-text.en_er,
span.vhw span.general-text.en_er,
span.sg span.general-text.en_er
{
	font-size: 80%;
}



span.frmg span.frm
{
	font-weight: 600;
	/*
	font-size: 120%;
	*/
}
span.frmg span.frm.x_nb
{
	font-weight: normal;
}

/* span.idmb span.idmsec span.idm */
span.idmsec span.idm
{
	font-weight: 600;
}

span.pvb span.pvsec span.pvg span.pv,
span.pvb span.pvsec span.pvg span.pvv
{
	font-weight: 600;
	/* margin-top: 0.5em; */
}

span.pvsec span.gramb span.semb span.pgr,
span.pvsec span.gramb span.pgr
{
	font-weight: 600;
	color: -apple-system-secondary-label;
}

span.cb span.csec span.cwg span.cw,
span.cb span.csec span.cwg span.cvg span.cv,
span.csec span.cwg span.cw
{
	font-weight: 600;
}

span.ind,
span.gr
{
	color: -apple-system-secondary-label;
}

span.lev,
span.fld
{
	font-size: 90%;
	color: -apple-system-secondary-label;
}

/*
span.exg span.ex:before
{
	content: "▸ ";
}
*/


span.trans.ty_pinyin,
span.trg.x_xd2 > span.trans.ty_pinyin,
span.trg.x_xd2 > span.trans span.trans.ty_pinyin,
span.trgg.x_xd2 > span.trg > span.trans.ty_pinyin,
span.trgg.x_xd2 > span.trg > span.trans span.trans.ty_pinyin
{
	/* So that Pinyin match with the one shown on the index column of the dictionary.app */
	/*font-family: 'STKaitiSC-Regular';*/
	font-weight: normal;
	margin-left: 0.2em;
	margin-right: 0.2em;
	color: -apple-system-secondary-label;
}


span.box
{
	display: block;
	/*
	background-color: window;
	padding: 0.33em 0.49em 0.33em 0.49em;
	*/
	border: solid 1px -apple-system-tertiary-label;
	padding: 0.13em 0.49em 0.13em 0.49em;
	-webkit-border-radius: 3px;
	margin-top: 0.1em;
}
html.apple_client-panel span.box
{
	margin-left: 0em;
	margin-right: 0em;
}
/*
html.apple_client-panel span.box
{
	margin-top: 0.3em;
	margin-bottom: 0.3em;
}
*/


/*=====================================*/
/* Thesaurus                           */
/*=====================================*/

span.synGroup > span.syn.t_core
{
	font-variant: small-caps;
	font-size: 120%;
	vertical-align: -1%;
}
span.synGroup > span.syn.t_core > span.gp
{
	font-variant: normal;
	font-size: 90%; 
}

span.antList > span.gp.ty_label.tg_antList
{
	font-variant: small-caps;
	color: -apple-system-secondary-label;
	font-size: 80%;
}	/* ANTONYMS */


/*=====================================*/
/* Features                            */
/*=====================================*/

span.feature
{
	display: block;
	margin-top: 2.5em;
	/* margin-right: 1em; */
}
html.apple_client-panel span.feature
{
	margin-top: 1em;
}

span.feature-title
{
	display: block;

	color: -apple-system-secondary-label;
	font-weight: 500;
	font-size: 90%;
	line-height : 23px;
	padding-top : 1px;
	padding-left : 9px;
	padding-right : 9px;
	margin-top : 2em;
	
	border-style : solid;
	border-width : 1px 1px 1px 1px;
	border-color : -apple-system-tertiary-label;
	-webkit-border-top-left-radius: 4px;
	-webkit-border-top-right-radius: 4px;
	-webkit-border-bottom-left-radius: 0px;
	-webkit-border-bottom-right-radius: 0px;
}

html.apple_client-panel span.feature-title
{
	margin-top : 1em;
}

span.feature-body
{
	display: block;

	color: -apple-system-label;

	padding-left : 10px;
	padding-right : 10px;
	padding-top : 5px;
	padding-bottom : 9px;

	border-style : solid;
	border-width : 0px 1px 1px 1px;
	border-color : -apple-system-tertiary-label;
	-webkit-border-top-left-radius: 0px;
	-webkit-border-top-right-radius: 0px;
	-webkit-border-bottom-left-radius: 4px;
	-webkit-border-bottom-right-radius: 4px;
}

span.feature-body table
{
	table-layout: fixed;
	/* table-layout: auto; */
	width: 100%;
	margin: 0 0 0 0;
	border: none;
}

span.feature-caption
{
	display: block;

	color: -apple-system-secondary-label;
	font-weight: normal;
	font-size: 90%;

	padding-left : 3px;
	padding-right : 3px;
	padding-top : 3px;
	padding-bottom : 9px;
}




span.feature span.list
{
	display: block;
}

span.feature span.list + span.list
{
	margin-top: 0.25em;
}

span.feature.t_wordLinks span.list span.item.t_related
{
	font-weight: 600;
	font-variant: small-caps;
}
span.feature.t_wordLinks span.list span.item.t_related + span.item.t_description
{
	margin-left: 0.5em;
}

span.para,
span.list span.item,
span.np
{
	display: block;
	margin-top: 0.25em;
}

span.feature.t_rightWord span.heading,
span.feature.t_usage span.heading,
span.feature.t_wordNote span.heading
{
	font-weight: 500;
}
span.feature.t_wordNote span.heading
{
	display: block;
}

span.ll
{
	font-weight: 500;
}
span.ll span.gp
{
	font-weight: normal;
}

span.feature.t_toolkit > span.feature-body > table tbody tr td.lemma
{
	font-weight: 600;
	font-size: 110%;
}

span.feature.t_quotes span.para span.eg span.ref
{
	display: block;
}

span.feature.t_encyclopaedic
{
	display: block;
	/*
	background-color: window;
	padding: 0.33em 0.49em 0.33em 0.49em;
	*/
	margin-right: 1em;
	margin-top: 1em;
	margin-bottom: 2em;
	border: solid 1px -apple-system-tertiary-label;
	padding: 0.13em 0.49em 0.13em 0.49em;
	-webkit-border-radius: 3px;
	/* margin-top: 0.1em; */
}


/*=====================================*/
/* Panel                               */
/*=====================================*/

/*
html.apple_client-panel body
{
	margin-top: 0em;
}

html.apple_client-panel *.entry
{
	line-height: 120%;
}

html.apple_client-panel span.hw, span.ehw, span.vhw 	{ font-size: 120%;}
html.apple_client-panel span.hg span.fg	{ font-size: 100%; }

html.apple_client-panel span.hg span.pr
{
	font-size: 100%;
}
html.apple_client-panel span.pos
{
	font-size: 100%;
}

*/


/* POS after the headword */
/*
html.apple_client-panel span.sg,
html.apple_client-panel span.hg,
html.apple_client-panel span.sg span.se1:first-child,
html.apple_client-panel span.sg span.se1:first-child span.pr
{
	display: inline;
	margin-left: 0em;
	padding-left: 0em;
}
html.apple_client-panel span.sg span.se1:first-child span.posg
{
	display: inline;
	margin-left: 1em;
}
html.apple_client-panel span.sg span.se1 span.posg
{
	margin-left: 0em;
}
*/

/*=====================================*/
/*                                     */
/*=====================================*/


span.eg.x_blk,
span.xrg.x_blk,
span.lbl.x_blk
{
	display: block;
}
span.eg span.ex.x_ni
{
	font-style: normal;
	font-weight: 200;
}
span.eg span.ex.x_ni span.bold
{
	font-style: normal;
	font-weight: 600;
}
span.lg.x_ni,
span.gg.x_ni,
span.ff.x_ni
{
	font-style: normal;
}


img.gi
{
	width: 1.0em;
	height: 1.0em;
}

img.image.th
{
	display: block;
	width: 96%;
	margin-top: 1em;
}



/*=====================================*/
/* Panel                               */
/* Move POS                            */
/* x_xd0:not(.hasSn)                   */
/* :first-child or x_xh0 + x_xd0       */
/*=====================================*/
html.apple_client-panel *.entry > span.x_xh0:first-child,
html.apple_client-panel span.x_xd0:not(.hasSn):first-child,
html.apple_client-panel span.x_xd0:not(.hasSn):first-child > span.x_xdh,
html.apple_client-panel span.x_xh0 + span.x_xd0:not(.hasSn),
html.apple_client-panel span.x_xh0 + span.x_xd0:not(.hasSn) > span.x_xdh
{
	display: inline;
}
html.apple_client-panel span.x_xd0:not(.hasSn):first-child > span.x_xd1,
html.apple_client-panel span.x_xd0:not(.hasSn):first-child > span.x_xd0,
html.apple_client-panel span.x_xh0 + span.x_xd0:not(.hasSn) > span.x_xd1,
html.apple_client-panel span.x_xh0 + span.x_xd0:not(.hasSn) > span.x_xd0
{
	margin-left: 1em;
}
html.apple_client-panel span.x_xd0:not(.hasSn):first-child > span.x_xd1.hasSn,
html.apple_client-panel span.x_xd0:not(.hasSn):first-child > span.x_xd0.hasSn,
html.apple_client-panel span.x_xh0 + span.x_xd0:not(.hasSn) > span.x_xd1.hasSn,
html.apple_client-panel span.x_xh0 + span.x_xd0:not(.hasSn) > span.x_xd0.hasSn
{
	margin-left: 2em;
}
html.apple_client-panel span.hg > span.hgSub1,
html.apple_client-panel span.hwg > span.hgSub1
{
	display: inline;
	margin-left: 1em;
}

/* EN_US */


span.brb:before
{
	/* ' ¦ ' */
	content: '\a0\a6\a0';
	font-size: 70%;
	vertical-align: 5%;
	color: -apple-system-secondary-label;
	margin-left: -0.15em;
	margin-right: -0.15em;
}

span.wbb:before,
span.brb:before
{
	font-weight: 200;
}

span.syl_txt
{
	display: none;
}

span.lbrk_txt
{
	display: none;
}

span.eg span.ex span.r
{
	font-weight: 400;
	font-style: italic;
}

span.gp.tg_pos
{
	margin-right: 0em;
}

	`
