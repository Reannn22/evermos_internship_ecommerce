-- First, make sure the table exists
CREATE TABLE IF NOT EXISTS `product_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `id_produk` bigint unsigned NOT NULL,
  `nama_produk` varchar(255) NOT NULL,
  `slug` varchar(255) NOT NULL,
  `harga_reseller` varchar(255) NOT NULL,
  `harga_konsumen` varchar(255) NOT NULL,
  `deskripsi` text,
  `id_toko` bigint unsigned NOT NULL,
  `id_category` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_product_log_deleted_at` (`deleted_at`),
  KEY `fk_product_log_product` (`id_produk`),
  KEY `fk_product_log_store` (`id_toko`),
  KEY `fk_product_log_category` (`id_category`),
  CONSTRAINT `fk_product_log_product` FOREIGN KEY (`id_produk`) REFERENCES `produk` (`id`),
  CONSTRAINT `fk_product_log_store` FOREIGN KEY (`id_toko`) REFERENCES `toko` (`id`),
  CONSTRAINT `fk_product_log_category` FOREIGN KEY (`id_category`) REFERENCES `category` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
