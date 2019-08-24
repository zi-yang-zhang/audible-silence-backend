CREATE TABLE `photo_meta` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id of photo',
  `url` varchar(1024) NOT NULL COMMENT 'url of photo',
  `thumbnail_url` varchar(1024) NOT NULL,
  `title` varchar(128) DEFAULT NULL COMMENT 'title of photo',
  `likes` int(11) DEFAULT '0' COMMENT 'number of likes',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created date of photo\n',
  `created_by` varchar(32) DEFAULT NULL COMMENT 'created by who\n',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'modified date of photo',
  `updated_by` varchar(32) NOT NULL COMMENT 'modified by who',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
