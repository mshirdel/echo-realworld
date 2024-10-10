CREATE TABLE IF NOT EXISTS `users` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(64) DEFAULT NULL,
    `email` varchar(256) DEFAULT NULL,
    `password` varchar(512) DEFAULT NULL,
    `last_login` timestamp NULL DEFAULT NULL,
    `is_active` tinyint(1) DEFAULT 1,
    `created_at` timestamp NOT NULL,
    `updated_at` timestamp NOT NULL,
    `deleted_at` datetime(3) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uni_users_username` (`username`),
    UNIQUE KEY `uni_users_email` (`email`),
    KEY `idx_users_username` (`username`),
    KEY `idx_users_email` (`email`),
    KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci