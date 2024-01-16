CREATE USER 'web'@'localhost';
CREATE USER 'web'@'172.17.0.1';



GRANT SELECT, INSERT, UPDATE, DELETE ON goush.* TO 'web'@'localhost'
GRANT SELECT, INSERT, UPDATE, DELETE ON goush.* TO 'web'@'172.17.0.1'

CREATE TABLE short_links ( id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT, short_code VARCHAR(10) NOT NULL,  original_url TEXT NOT NULL, created D
ATETIME NOT NULL);

CREATE DATABASE goush CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;