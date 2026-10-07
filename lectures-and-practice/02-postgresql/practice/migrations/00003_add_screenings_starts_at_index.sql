-- +goose Up
CREATE INDEX screenings_starts_at_idx ON screenings (starts_at);

-- +goose Down
DROP INDEX screenings_starts_at_idx;
