CREATE DATABASE IF NOT EXISTS kudaki_user DEFAULT COLLATE = utf8_general_ci;
CREATE USER IF NOT EXISTS 'kudaki_user_repo' @'%' IDENTIFIED BY 'kudakiuserreporocks';
GRANT ALL PRIVILEGES ON kudaki_user.* TO 'kudaki_user_repo' @'%' WITH GRANT OPTION;
USE kudaki_user;
CREATE TABLE IF NOT EXISTS users (
  `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `uuid` VARCHAR(64) NOT NULL UNIQUE,
  `email` VARCHAR(255) NOT NULL UNIQUE,
  `password` VARCHAR(255),
  `token` TEXT,
  `role` ENUM('ADMIN', 'USER', 'KUDAKI_TEAM', 'ORGANIZER'),
  `phone_number` VARCHAR(255),
  `account_type` ENUM('NATIVE', 'GOOGLE', 'FACEBOOK'),
  `created_at` BIGINT UNSIGNED
);
CREATE TABLE IF NOT EXISTS profiles(
  `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `uuid` VARCHAR(64) NOT NULL UNIQUE,
  `user_uuid` VARCHAR(64) NOT NULL UNIQUE,
  `full_name` VARCHAR(255),
  `photo` VARCHAR(255),
  `created_at` BIGINT UNSIGNED,
  FOREIGN KEY(user_uuid) REFERENCES users(uuid) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS unverified_users(
  `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `user_uuid` VARCHAR(64) NOT NULL UNIQUE,
  `created_at` BIGINT UNSIGNED,
  FOREIGN KEY(user_uuid) REFERENCES users(uuid) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS reset_passwords(
  `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `user_uuid` VARCHAR(64) NOT NULL UNIQUE,
  `token` TEXT,
  `created_at` BIGINT UNSIGNED,
  FOREIGN KEY(user_uuid) REFERENCES users(uuid) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS addresses(
  `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `uuid` VARCHAR(64) NOT NULL UNIQUE,
  `profile_uuid` VARCHAR(64) NOT NULL,
  `full_address` TEXT,
  `receiver_name` VARCHAR(255),
  `receiver_phone_number` VARCHAR(255),
  `zip_code` VARCHAR(255),
  `latitude` DECIMAL(10, 8),
  `longitude` DECIMAL(11, 8),
  `created_at` BIGINT UNSIGNED,
  FOREIGN KEY(profile_uuid) REFERENCES profiles(uuid) ON DELETE CASCADE
);