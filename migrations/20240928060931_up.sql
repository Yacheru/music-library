-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS songs (
    id SERIAL PRIMARY KEY,
    "group" VARCHAR NOT NULL,
    title VARCHAR NOT NULL,
    text TEXT NOT NULL,
    link VARCHAR NOT NULL UNIQUE,
    release_date TIMESTAMP
);

CREATE INDEX IF NOT EXISTS songs_title_idx ON songs(title);
CREATE INDEX IF NOT EXISTS songs_group_idx ON songs("group");
CREATE INDEX IF NOT EXISTS songs_text_idx ON songs(text);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS songs;

DROP INDEX IF EXISTS songs_title_idx;
DROP INDEX IF EXISTS songs_group_idx;
DROP INDEX IF EXISTS songs_text_idx;
-- +goose StatementEnd
