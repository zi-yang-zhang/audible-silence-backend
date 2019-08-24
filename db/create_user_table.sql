CREATE TABLE user(
  id VARCHAR(64) not null,
  name VARCHAR(128) not null,
  email VARCHAR(128) DEFAULT '',
  wechat_open_id VARCHAR(64) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
