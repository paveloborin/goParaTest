CREATE SCHEMA  IF NOT EXISTS test;
USE test;

DROP TABLE IF EXISTS payments;
CREATE TABLE payments
(
  id         INT AUTO_INCREMENT
    PRIMARY KEY,
  student_id INT               NOT NULL,
  datetime   DATETIME          NOT NULL,
  amount     FLOAT DEFAULT '0' NULL
);

CREATE INDEX student_id
  ON payments (student_id);

DROP TABLE IF EXISTS student;
CREATE TABLE student
(
  id      INT AUTO_INCREMENT
    PRIMARY KEY,
  name    VARCHAR(20)                                          NOT NULL,
  surname VARCHAR(20) DEFAULT ''                               NOT NULL,
  gender  ENUM ('male', 'female', 'unknown') DEFAULT 'unknown' NULL
);

CREATE INDEX gender
  ON student (gender);

DROP TABLE IF EXISTS student_status;
CREATE TABLE student_status
(
  id         INT AUTO_INCREMENT
    PRIMARY KEY,
  student_id INT                                                                   NOT NULL,
  status     ENUM ('new', 'studying', 'vacation', 'testing', 'lost') DEFAULT 'new' NOT NULL,
  datetime   DATETIME                                                              NOT NULL
);

CREATE INDEX datetime
  ON student_status (datetime);

CREATE INDEX student_id
  ON student_status (student_id);

DROP TABLE IF EXISTS teacher;
CREATE TABLE teacher
(
  id      INT AUTO_INCREMENT
    PRIMARY KEY,
  name    VARCHAR(20)                                          NOT NULL,
  surname VARCHAR(20) DEFAULT ''                               NOT NULL,
  gender  ENUM ('male', 'female', 'unknown') DEFAULT 'unknown' NULL
);

CREATE INDEX gender
  ON teacher (gender);

DROP TABLE IF EXISTS teacher_student;
CREATE TABLE teacher_student
(
  id         INT AUTO_INCREMENT
    PRIMARY KEY,
  student_id INT NOT NULL,
  teacher_id INT NOT NULL
);

DROP TABLE IF EXISTS user;
CREATE TABLE user
(
  id      INT AUTO_INCREMENT
    PRIMARY KEY,
  name    VARCHAR(20)                                                    NOT NULL,
  surname VARCHAR(20) DEFAULT ''                                         NOT NULL,
  phone   VARCHAR(20) DEFAULT ''                                         NULL,
  status  ENUM ('new', 'joined', 'refused', 'unavailable') DEFAULT 'new' NULL,
  date    TIMESTAMP DEFAULT CURRENT_TIMESTAMP                            NOT NULL
);

