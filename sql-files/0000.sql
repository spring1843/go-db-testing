CREATE DATABASE IF NOT EXISTS blog;
USE blog;

CREATE TABLE IF NOT EXISTS users
(
  user_id    INT            NOT NULL AUTO_INCREMENT,
  first_name VARCHAR(14)    NOT NULL,
  last_name  VARCHAR(16)    NOT NULL,
  email      VARCHAR(150)    NOT NULL,
  PRIMARY KEY (user_id),
  UNIQUE KEY (email)
) ENGINE = InnoDB
  CHARACTER SET utf8;

CREATE TABLE IF NOT EXISTS posts
(
  post_id INT         NOT NULL AUTO_INCREMENT,
  user_id INT         NOT NULL,
  title   VARCHAR(50) NOT NULL,
  body    TEXT        NOT NULL,
  PRIMARY KEY (post_id),
  FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE
) ENGINE = InnoDB
  CHARACTER SET utf8;
