CREATE DATABASE IF NOT EXISTS kudaki_mountain DEFAULT COLLATE = utf8_general_ci;
CREATE USER IF NOT EXISTS 'kudaki_mountain_repo' @'%' IDENTIFIED BY 'kudakimountainreporocks';
GRANT ALL PRIVILEGES ON kudaki_mountain.* TO 'kudaki_mountain_repo' @'%' WITH GRANT OPTION;
USE kudaki_mountain;
CREATE TABLE IF NOT EXISTS mountains (
  `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `uuid` VARCHAR(64) NOT NULL UNIQUE,
  `name` VARCHAR(64),
  `height` INT(20),
  `latitude` DECIMAL(10, 8),
  `longitude` DECIMAL(11, 8),
  `difficulty` DECIMAL(4.3),
  `description` TEXT,
  `created_at` BIGINT UNSIGNED,
  FULLTEXT(`name`, `description`)
);
CREATE TABLE IF NOT EXISTS mountain_files (
  `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `mountain_uuid` VARCHAR(64) NOT NULL UNIQUE,
  `file_path` VARCHAR(255),
  `created_at` BIGINT UNSIGNED,
  FOREIGN KEY(mountain_uuid) REFERENCES mountains(uuid) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS mountain_reviews (
  `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `uuid` VARCHAR(64) NOT NULL UNIQUE,
  `user_uuid` VARCHAR(64) NOT NULL,
  `mountain_uuid` VARCHAR(64) NOT NULL,
  `difficulty` DECIMAL(4, 3),
  `review` TEXT,
  `created_at` BIGINT UNSIGNED,
  FOREIGN KEY(mountain_uuid) REFERENCES mountains(uuid) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS recommended_gears (
  `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `uuid` VARCHAR(64) NOT NULL UNIQUE,
  `user_uuid` VARCHAR(64) NOT NULL,
  `mountain_uuid` VARCHAR(64) NOT NULL,
  `upvote` INT(20),
  `downvote` INT(20),
  `seen` INT(20),
  `created_at` BIGINT UNSIGNED,
  FOREIGN KEY(mountain_uuid) REFERENCES mountains(uuid) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS recommended_gear_items (
  `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `uuid` VARCHAR(64) NOT NULL UNIQUE,
  `recommended_gear_uuid` VARCHAR(64) NOT NULL,
  `item_type` VARCHAR(255),
  `total` INT(20),
  `created_at` BIGINT UNSIGNED,
  FOREIGN KEY(recommended_gear_uuid) REFERENCES recommended_gears(uuid) ON DELETE CASCADE
);