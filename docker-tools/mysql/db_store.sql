CREATE DATABASE IF NOT EXISTS kudaki_store DEFAULT COLLATE = utf8_general_ci;
CREATE USER IF NOT EXISTS 'kudaki_store_repo' @'%' IDENTIFIED BY 'kudakistorereporocks';
GRANT ALL PRIVILEGES ON kudaki_store.* TO 'kudaki_store_repo' @'%' WITH GRANT OPTION;
USE kudaki_store;
CREATE TABLE IF NOT EXISTS storefronts(
  `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `uuid` VARCHAR(64) NOT NULL UNIQUE,
  `user_uuid` VARCHAR(64),
  `total_item` INT(20),
  `rating` DECIMAL(4, 3),
  `total_raw_rating` DECIMAL(65, 3),
  `created_at` BIGINT UNSIGNED
);
CREATE TABLE IF NOT EXISTS items(
  `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `uuid` VARCHAR(64) NOT NULL UNIQUE,
  `storefront_uuid` VARCHAR(64),
  `name` VARCHAR(255),
  `amount` INT(20),
  `unit` VARCHAR(255),
  `price` INT(20),
  `price_duration` ENUM('DAY', 'WEEK', 'MONTH', 'YEAR'),
  `description` TEXT,
  `photo` VARCHAR(255),
  `rating` DECIMAL(4, 3),
  `total_raw_rating` DECIMAL(65, 3),
  `length` INT(20),
  `width` INT(20),
  `height` INT(20),
  `color` VARCHAR(255),
  `unit_of_measurement` ENUM('MM', 'CM', 'DM', 'M', 'DAM', 'HM', 'KM'),
  `created_at` BIGINT UNSIGNED,
  FULLTEXT(`name`, `description`),
  FOREIGN KEY(storefront_uuid) REFERENCES storefronts(uuid) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS item_reviews(
  `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `uuid` VARCHAR(64) NOT NULL UNIQUE,
  `user_uuid` VARCHAR(64) NOT NULL,
  `item_uuid` VARCHAR(64) NOT NULL,
  `review` TEXT,
  `rating` DECIMAL(4, 3),
  `created_at` BIGINT UNSIGNED,
  FOREIGN KEY(item_uuid) REFERENCES items(uuid) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS review_comments(
  `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `uuid` VARCHAR(64) NOT NULL UNIQUE,
  `item_review_uuid` VARCHAR(64) NOT NULL,
  `user_uuid` VARCHAR(64) NOT NULL,
  `comment` TEXT,
  `created_at` BIGINT UNSIGNED,
  FOREIGN KEY(item_review_uuid) REFERENCES item_reviews(uuid) ON DELETE CASCADE
);