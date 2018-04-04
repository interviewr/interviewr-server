-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE organization
(
  id VARCHAR(250) PRIMARY KEY,
  name VARCHAR(20) NOT NULL,
  email VARCHAR(25) NOT NULL,
  description VARCHAR(150),
  location VARCHAR(50)
);


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE organization;
