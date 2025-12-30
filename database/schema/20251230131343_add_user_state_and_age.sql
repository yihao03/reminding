-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TYPE STATES AS ENUM (
    'Johor', 'Kedah', 'Kelantan', 'Melaka', 'Negeri Sembilan', 'Pahang',
    'Perak', 'Perlis', 'Penang', 'Sabah', 'Sarawak', 'Selangor', 'Terengganu'
);

ALTER TABLE USERS
ADD COLUMN STATE STATES,
ADD COLUMN AGE INT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE USERS
DROP COLUMN IF EXISTS STATE,
DROP COLUMN IF EXISTS AGE;

DROP TYPE IF EXISTS STATES;
-- +goose StatementEnd
