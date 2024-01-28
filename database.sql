CREATE USER 'web'@'localhost';
CREATE USER 'web'@'172.17.0.1';



GRANT SELECT, INSERT, UPDATE, DELETE ON goush.* TO 'web'@'localhost'
GRANT SELECT, INSERT, UPDATE, DELETE ON goush.* TO 'web'@'172.17.0.1'

CREATE TABLE short_links (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT, 
    short_code VARCHAR(10) NOT NULL,
    user_id INTEGER NOT NULL,
    original_url TEXT NOT NULL,
    created DATETIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE DATABASE goush CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE goush;
CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created DATETIME NOT NULL
);
ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);
