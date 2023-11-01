-- SET SQL_MODE='NO_AUTO_VALUE_ON_ZERO';

-- START TRANSACTION;

-- SET
-- 	time_zone = "+00:00";

-- --
-- -- Database: cheifoon
-- --

-- CREATE DATABASE IF NOT EXISTS cheifoon DEFAULT CHARACTER SET
-- utf8mb4 COLLATE utf8mb4_general_ci;

-- USE cheifoon;

-- --
-- -- Table structure for table 'admin'
-- --

-- CREATE TABLE admin (
-- 	id INT UNSIGNED NOT NULL AUTO_INCREMENT,
-- 	seasoning_name VARCHAR(255) NOT NULL UNIQUE,
-- 	tea_second DOUBLE NOT NULL,
-- 	bottle_image VARCHAR(255) NOT NULL,
-- 	PRIMARY KEY(id)
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --
-- -- Table structure for table user
-- --

-- CREATE TABLE user (
-- 	id INT UNSIGNED NOT NULL AUTO_INCREMENT,
-- 	seasoning_id INT UNSIGNED NOT NULL,
-- 	PRIMARY KEY(id),
-- 	FOREIGN KEY (seasoning_id) REFERENCES admin (id)
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --
-- -- Table structure for table 'recipe'
-- --

-- CREATE TABLE recipe (
-- 	id INT UNSIGNED NOT NULL AUTO_INCREMENT,
-- 	recipe_name VARCHAR(255) NOT NULL,
-- 	menu_image VARCHAR(255) NOT NULL,
-- 	PRIMARY KEY(id)
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --
-- -- Table structure for table 'menu'
-- --

-- CREATE TABLE menu (
-- 	id INT UNSIGNED NOT NULL AUTO_INCREMENT,
-- 	recipe_id INT UNSIGNED NOT NULL,
-- 	seasoning_id INT UNSIGNED NOT NULL,
-- 	table_spoon INT NOT NULL,
-- 	tea_spoon INT NOT NULL,
-- 	PRIMARY KEY(id,recipe_id),
-- 	FOREIGN KEY(recipe_id) REFERENCES recipe (id),
-- 	FOREIGN KEY(seasoning_id) REFERENCES admin (id)
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;