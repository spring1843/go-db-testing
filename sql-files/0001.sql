USE blog;

CREATE TABLE IF NOT EXISTS db_versions
(
  db_version_id INT       NOT NULL AUTO_INCREMENT,
  version       INT       NOT NULL,
  executed_time TIMESTAMP NOT NULL,
  PRIMARY KEY (db_version_id)
) ENGINE = InnoDB
  CHARACTER SET utf8;

INSERT INTO db_versions (version, executed_time) VALUES (1, NOW());
