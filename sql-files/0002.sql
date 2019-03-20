CREATE TABLE IF NOT EXISTS comments
(
  comment_id INT         NOT NULL AUTO_INCREMENT,
  user_id    INT         NOT NULL,
  post_id    INT         NOT NULL,
  dept_name  VARCHAR(40) NOT NULL,
  PRIMARY KEY (comment_id),
  FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE,
  FOREIGN KEY (post_id) REFERENCES posts (post_id) ON DELETE CASCADE
) ENGINE = InnoDB
  CHARACTER SET utf8;

INSERT INTO db_versions (version, executed_time) VALUES (2, NOW());
