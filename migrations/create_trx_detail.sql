CREATE TABLE `trx_detail` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `id_trx` bigint unsigned NOT NULL,
  `id_log_produk` bigint unsigned NOT NULL,
  `id_toko` bigint unsigned NOT NULL,
  `kuantitas` int NOT NULL,
  `harga_total` double NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_trx_detail_deleted_at` (`deleted_at`),
  KEY `fk_trx_detail_transaction` (`id_trx`),
  KEY `fk_trx_detail_product_log` (`id_log_produk`),
  KEY `fk_trx_detail_store` (`id_toko`),
  CONSTRAINT `fk_trx_detail_transaction` FOREIGN KEY (`id_trx`) REFERENCES `trx` (`id`),
  CONSTRAINT `fk_trx_detail_product_log` FOREIGN KEY (`id_log_produk`) REFERENCES `product_log` (`id`),
  CONSTRAINT `fk_trx_detail_store` FOREIGN KEY (`id_toko`) REFERENCES `toko` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
