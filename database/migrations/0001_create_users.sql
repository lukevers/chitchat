-- +migrate Up
CREATE TABLE users
(

);

-- +migrate Down
DROP TABLE users;
