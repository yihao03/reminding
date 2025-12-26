-- +goose Up
-- +goose StatementBegin
ALTER TABLE users RENAME COLUMN user_name TO display_name;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users RENAME COLUMN display_name TO user_name;
-- +goose StatementEnd
