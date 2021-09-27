
-- +migrate Up
CREATE TABLE IF NOT EXISTS `users` (
  `id` BINARY(16) NOT NULL AUTO_INCREMENT, 
  `fname` VARCHAR(100) NOT NULL,
  `lname` VARCHAR(100) NOT NULL,
);

-- +migrate Down
DROP TABLE IF EXISTS `users`;
