-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE organization
(
  id VARCHAR(255) PRIMARY KEY,
  name VARCHAR(64) NOT NULL,
  email VARCHAR(64) NOT NULL,
  description VARCHAR(255),
  location VARCHAR(64)
);


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE organization;
