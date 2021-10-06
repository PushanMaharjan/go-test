
-- +migrate Up
CREATE TABLE IF NOT EXISTS `users` (
  `id` BINARY(16) PRIMARY KEY NOT NULL, 
  `fname` VARCHAR(100) NOT NULL,
  `lname` VARCHAR(100) NOT NULL,
  `role_id` BINARY(16) NOT NULL,
    FOREIGN KEY (`role_id`)
    REFERENCES `roles` (`id`)
);

-- +migrate Down
DROP TABLE IF EXISTS `users`; 
