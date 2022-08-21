CREATE DATABASE IF NOT EXISTS kudaki_rental DEFAULT COLLATE = utf8_general_ci;
CREATE USER IF NOT EXISTS 'kudaki_rental_repo' @'%' IDENTIFIED BY 'kudakirentalreporocks';
GRANT ALL PRIVILEGES ON kudaki_rental.* TO 'kudaki_rental_repo' @'%' WITH GRANT OPTION;
USE kudaki_rental;
CREATE TABLE IF NOT EXISTS carts(
  `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `uuid` VARCHAR(255) NOT NULL UNIQUE,
  `user_uuid` VARCHAR(255),
  `total_price` INT(20) UNSIGNED,
  `total_items` INT(20) UNSIGNED,
  `open` TINYINT(1),
  `created_at` BIGINT UNSIGNED
);
CREATE TABLE IF NOT EXISTS cart_items(
  `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `uuid` VARCHAR(255) NOT NULL UNIQUE,
  `cart_uuid` VARCHAR(255),
  `item_uuid` VARCHAR(255),
  `total_item` INT(20),
  `total_price` INT(20) UNSIGNED,
  `unit_price` INT(20),
  `duration_from` BIGINT,
  `duration_to` BIGINT,
  `created_at` BIGINT UNSIGNED,
  FOREIGN KEY(cart_uuid) REFERENCES carts(uuid) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS returnment_confirmations(
  `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `order_uuid` VARCHAR(255) NOT NULL UNIQUE,
  `tenant_user_uuid` VARCHAR(255) NOT NULL,
  `owner_user_uuid` VARCHAR(255) NOT NULL,
  `tenant_confirmed` TINYINT(1),
  `owner_confirmed` TINYINT(1),
  `created_at` BIGINT UNSIGNED
);
CREATE TABLE IF NOT EXISTS  owner_returnment_confirmations(
  `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `uuid` VARCHAR(64) NOT NULL UNIQUE
);