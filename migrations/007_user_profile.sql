CREATE TABLE user_education (
  id            INT AUTO_INCREMENT PRIMARY KEY,
  user_id       INT          NOT NULL,
  institution   VARCHAR(200) NOT NULL,
  degree        VARCHAR(100) NOT NULL,
  field_of_study VARCHAR(150) NOT NULL,
  start_year    SMALLINT     NOT NULL,
  end_year      SMALLINT     DEFAULT NULL,
  description   TEXT         DEFAULT NULL,
  created_at    TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE user_projects (
  id          INT AUTO_INCREMENT PRIMARY KEY,
  user_id     INT          NOT NULL,
  title       VARCHAR(200) NOT NULL,
  description TEXT         NOT NULL,
  url         VARCHAR(500) DEFAULT NULL,
  year        SMALLINT     DEFAULT NULL,
  created_at  TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE user_courses (
  id          INT AUTO_INCREMENT PRIMARY KEY,
  user_id     INT          NOT NULL,
  title       VARCHAR(200) NOT NULL,
  institution VARCHAR(200) NOT NULL,
  year        SMALLINT     DEFAULT NULL,
  url         VARCHAR(500) DEFAULT NULL,
  created_at  TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB;
