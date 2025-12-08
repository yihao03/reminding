-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE event_registrations (
    id                SERIAL    PRIMARY KEY,
    event_id          INT       NOT NULL REFERENCES events (id),
    user_id           INT       NOT NULL REFERENCES users (id),
    registration_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS event_registrations;
-- +goose StatementEnd
