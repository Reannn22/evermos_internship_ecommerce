CREATE TABLE IF NOT EXISTS `notifikasi` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `pesan` text NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
