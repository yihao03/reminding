-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE users
ADD COLUMN updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE events
ADD COLUMN event_name VARCHAR(255) NOT NULL;

CREATE OR REPLACE FUNCTION UPDATE_TIMESTAMP_FUNC()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_timestamp
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION UPDATE_TIMESTAMP_FUNC();

CREATE TRIGGER update_timestamp
BEFORE UPDATE ON events
FOR EACH ROW
EXECUTE FUNCTION UPDATE_TIMESTAMP_FUNC();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TRIGGER IF EXISTS update_timestamp ON users;
DROP TRIGGER IF EXISTS update_timestamp ON events;
DROP FUNCTION IF EXISTS update_timestamp_func();

ALTER TABLE users
DROP COLUMN IF EXISTS updated_at;

ALTER TABLE events
DROP COLUMN IF EXISTS event_name;
-- +goose StatementEnd
