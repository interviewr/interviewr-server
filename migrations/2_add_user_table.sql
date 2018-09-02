-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE user
(
  id VARCHAR(250) PRIMARY KEY,
  login VARCHAR(64) NOT NULL,
  name VARCHAR(64) NOT NULL,
  email VARCHAR(255) NOT NULL,
  location VARCHAR(64),
  bio VARCHAR(255),
  avatarUrl VARCHAR(255),
);


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE user;
