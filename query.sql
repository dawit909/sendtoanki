-- name: GetRows :many
select WORDS.stem, WORDS.word, BOOK_INFO.title, LOOKUPS.usage, LOOKUPS.word_key
from LOOKUPS left join WORDS
on WORDS.id = LOOKUPS.word_key
left join BOOK_INFO
on BOOK_INFO.id = LOOKUPS.book_key
order by WORDS.stem, LOOKUPS.timestamp;

-- name: GetWordKeysWithMultipleUsages :many
select LOOKUPS.word_key, COUNT(LOOKUPS.word_key) from LOOKUPS
GROUP BY LOOKUPS.word_key
HAVING COUNT(LOOKUPS.word_key) > 1;