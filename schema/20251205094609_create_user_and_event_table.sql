-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE users (
    id           SERIAL       PRIMARY KEY,
    firebase_uid VARCHAR(255) UNIQUE   NOT NULL,
    created_at   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_name    VARCHAR(100) NOT NULL,
    email        VARCHAR(100) UNIQUE   NOT NULL
);

CREATE TABLE events (
    id            SERIAL    PRIMARY KEY,
    created_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    organiser     VARCHAR(255),
    is_online     BOOLEAN   NOT NULL,
    location_name VARCHAR(255),
    start_time    TIMESTAMP,
    end_time      TIMESTAMP,
    details       TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS events;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
