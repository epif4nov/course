-- +goose Up
CREATE TABLE bookings (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    screening_id bigint NOT NULL REFERENCES screenings(id),
    user_id bigint NOT NULL,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT bookings_screening_user_unique UNIQUE (screening_id, user_id)
);

-- +goose Down
DROP TABLE bookings;
