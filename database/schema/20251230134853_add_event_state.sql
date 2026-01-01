-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE events
ADD COLUMN state STATES,
ADD COLUMN registration_link TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE events
DROP COLUMN IF EXISTS registration_link,
DROP COLUMN IF EXISTS state;
-- +goose StatementEnd
