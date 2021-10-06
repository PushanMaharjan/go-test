
-- +migrate Up
CREATE TABLE IF NOT EXISTS `roles` (
  `id` BINARY(16) PRIMARY KEY NOT NULL, 
  `role` VARCHAR(100) NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS `roles`;