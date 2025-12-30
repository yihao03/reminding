-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE events
ADD COLUMN state STATES;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE events
DROP COLUMN IF EXISTS state;
-- +goose StatementEnd
