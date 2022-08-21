CREATE DATABASE IF NOT EXISTS kudaki_event DEFAULT COLLATE = utf8_general_ci;
CREATE USER IF NOT EXISTS 'kudaki_event_repo' @'%' IDENTIFIED BY 'kudakieventreporocks';
GRANT ALL PRIVILEGES ON kudaki_event.* TO 'kudaki_event_repo' @'%' WITH GRANT OPTION;
USE kudaki_event;
CREATE TABLE IF NOT EXISTS kudaki_events (
  `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `uuid` VARCHAR(64) NOT NULL UNIQUE,
  `organizer_user_uuid` VARCHAR(64) NOT NULL,
  `name` VARCHAR(255),
  `latitude` DECIMAL(10, 8),
  `longitude` DECIMAL(11, 8),
  `venue` VARCHAR(255),
  `description` TEXT,
  `ad_duration_from` BIGINT,
  `ad_duration_to` BIGINT,
  `duration_from` BIGINT,
  `duration_to` BIGINT,
  `seen` INT(20),
  `status` ENUM('UNPUBLISHED', 'PUBLISHED', 'TAKEN_DOWN'),
  `created_at` BIGINT UNSIGNED,
  FULLTEXT(`name`, `description`, `venue`)
);
CREATE TABLE IF NOT EXISTS doku_invoices (
  `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `uuid` VARCHAR(64) NOT NULL UNIQUE,
  `kudaki_event_uuid` VARCHAR(64) NOT NULL,
  `mall_id` INT(20),
  `chain_merchant` INT(20),
  `amount` INT(20),
  `purchase_amount` INT(20),
  `transaction_id_merchant` VARCHAR(255),
  `words` VARCHAR(255),
  `request_date_time` BIGINT UNSIGNED,
  `currency` INT(20),
  `purchase_currency` INT(20),
  `session_id` VARCHAR(255),
  `name` VARCHAR(128),
  `email` VARCHAR(254),
  `basket` TEXT,
  `created_at` BIGINT UNSIGNED,
  `status` ENUM('NEW', 'SUCCESS', 'FAILED'),
  FOREIGN KEY(kudaki_event_uuid) REFERENCES kudaki_events(uuid) ON DELETE CASCADE
);