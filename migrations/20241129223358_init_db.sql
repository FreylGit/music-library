-- +goose Up
-- +goose StatementBegin
CREATE TABLE songs (
    id SERIAL PRIMARY KEY,
    group_name TEXT NOT NULL,
    song_name TEXT NOT NULL,
    release_date DATE NOT NULL,
    text TEXT NOT NULL,
    link TEXT NOT NULL
);
CREATE INDEX idx_group_name ON songs(group_name);
CREATE INDEX idx_song_name ON songs(song_name);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX idx_group_name;
DROP INDEX idx_song_name;
DROP TABLE songs;
-- +goose StatementEnd
