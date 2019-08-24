create table markdowns (
	id int(11) auto_increment,
    title varchar(128) not null,
    content text not null,
		likes int(11) not null default 0,
		`likes` int(11) DEFAULT '0' COMMENT 'number of likes',
	  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created date of photo\n',
	  `created_by` varchar(32) DEFAULT '' COMMENT 'created by who\n',
	  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'modified date of photo',
	  `updated_by` varchar(32) DEFAULT '' COMMENT 'modified by who',
		PRIMARY KEY (`id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
