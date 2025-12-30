-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TYPE states AS ENUM (
    'Johor', 'Kedah', 'Kelantan', 'Melaka', 'Negeri Sembilan', 'Pahang',
    'Perak', 'Perlis', 'Penang', 'Sabah', 'Sarawak', 'Selangor', 'Terengganu'
);

ALTER TABLE users
ADD COLUMN state states,
ADD COLUMN age int;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE users
DROP COLUMN IF EXISTS state,
DROP COLUMN IF EXISTS age;
-- +goose StatementEnd
