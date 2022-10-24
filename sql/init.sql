CREATE DATABASE IF NOT EXISTS go_mysql;

CREATE TABLE IF NOT EXISTS customer (
  id VARCHAR(100) NOT NULL,
  name VARCHAR(100) NOT NULL, 
  email VARCHAR(255) NOT NULL,
  balance INT NOT NULL DEFAULT 0,
  rating DOUBLE NOT NULL DEFAULT 0.0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  birth_date DATE,
  married BOOLEAN DEFAULT false,
  PRIMARY KEY(id)
);

CREATE  TABLE IF NOT EXISTS user(
  username VARCHAR(100) NOT NULL ,
  password VARCHAR(100) NOT NULL ,
  PRIMARY KEY (username)
)ENGINE InnoDB;

CREATE TABLE IF NOT EXISTS comment (
  id INT NOT NULL AUTO_INCREMENT,
  email VARCHAR(100) NOT NULL,
  comment TEXT,
  PRIMARY KEY(id)
);
