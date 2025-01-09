USE tasker_db;

BEGIN;

CREATE TABLE IF NOT EXISTS `users` (
  `user_id` INT NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `email` (`email`)
) Engine=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `todo_lists` (
  `todo_list_id` INT NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  PRIMARY KEY (`todo_list_id`)
) Engine=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `members`(
  `user_id` INT NOT NULL,
  `todo_list_id` INT NOT NULL,
  `membership` varchar(30) NOT NULL,
  PRIMARY KEY (`user_id`, `todo_list_id`),
  FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`),
  FOREIGN KEY (`todo_list_id`) REFERENCES `todo_lists`(`todo_list_id`)
) Engine=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `todos` (
  `todo_id` INT NOT NULL AUTO_INCREMENT,
  `todo_list_id` INT NOT NULL,
  `description` varchar(255) NOT NULL,
  `status` varchar(30) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_by` INT NOT NULL,
  PRIMARY KEY (`todo_id`),
  FOREIGN KEY (`todo_list_id`) REFERENCES `todo_lists`(`todo_list_id`),
  FOREIGN KEY (`created_by`) REFERENCES `users`(`user_id`)
) Engine=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `todo_events`(
  `todo_event_id` INT NOT NULL AUTO_INCREMENT,
  `todo_id` INT NOT NULL,
  `old_value` varchar(255) NOT NULL,
  `new_value` varchar(255) NOT NULL,
  `event_type` varchar(30) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT NOT NULL,
  PRIMARY KEY (`todo_event_id`),
  FOREIGN KEY (`todo_id`) REFERENCES `todos`(`todo_id`),
  FOREIGN KEY (`created_by`) REFERENCES `users`(`user_id`)
) Engine=InnoDB DEFAULT CHARSET=utf8;

COMMIT;
