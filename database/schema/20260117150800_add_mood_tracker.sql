-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE mood_tracker (
    id         SERIAL       PRIMARY KEY,
    user_uid   VARCHAR(255) NOT NULL REFERENCES users (firebase_uid),
    mood       INTEGER      NOT NULL,
    created_at TIMESTAMPTZ    DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS mood_tracker;
-- +goose StatementEnd
