-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE journals (
    id              SERIAL       PRIMARY KEY,
    user_uid        VARCHAR(255) NOT NULL REFERENCES users (firebase_uid),
    created_at      TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    title           VARCHAR(255) NOT NULL,
    journal_content TEXT         NOT NULL
);

CREATE TRIGGER update_timestamp
BEFORE UPDATE ON journals
FOR EACH ROW
EXECUTE FUNCTION UPDATE_TIMESTAMP_FUNC();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TRIGGER IF EXISTS update_timestamp ON journals;

DROP TABLE IF EXISTS journals;
-- +goose StatementEnd
