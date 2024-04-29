CREATE DATABASE readinglist;

CREATE ROLE readinglist WITH LOGIN PASSWORD '*';

CREATE TABLE IF NOT EXISTS books (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    author text,
    published integer NOT NULL,
    pages integer NOT NULL,
    genres text[] NOT NULL,
    rating real NOT NULL, 
    isbn text,
    version integer NOT NULL DEFAULT 1
);

GRANT SELECT, INSERT, UPDATE, DELETE ON books TO readinglist;

GRANT USAGE, SELECT ON SEQUENCE books_id_seq TO readinglist;

/*Sample Book*/
INSERT INTO books (title, author, published, pages, genres, rating, isbn)
VALUES ('Sample Book', 'Author Name', 2021, 300, ARRAY['Fiction'], 4.5, '000-00-00000-00-1');