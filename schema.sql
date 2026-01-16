CREATE TABLE BOOK_INFO (
    id      text PRIMARY KEY,
    asin    text,
	guid    text,
	lang    text,
	title   text,
	authors text
);

CREATE TABLE LOOKUPS (
    id          text PRIMARY KEY,
    word_key    text,
	book_key    text,
	dict_key    text,
	pos         text,
	usage       text,
    timestamp   INTEGER
);

CREATE TABLE WORDS (
    id      text PRIMARY KEY,
    word    text,
	stem    text,
	lang    text,
	category   INTEGER,
	timestamp INTEGER,
    profileid   text
);