-- First drop related tables
SET FOREIGN_KEY_CHECKS = 0;
DROP TABLE IF EXISTS stores;
DROP TABLE IF EXISTS users;
SET FOREIGN_KEY_CHECKS = 1;

-- Then create users table
CREATE TABLE users (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    nama VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    kata_sandi VARCHAR(255) NOT NULL,
    notelp VARCHAR(15) NOT NULL,
    email VARCHAR(255) NOT NULL,
    tanggal_lahir DATETIME,
    jenis_kelamin VARCHAR(255),
    tentang TEXT,
    pekerjaan VARCHAR(255),
    id_provinsi VARCHAR(255),
    id_kota VARCHAR(255),
    is_admin BOOLEAN DEFAULT false,
    created_at DATETIME(3),
    updated_at DATETIME(3),
    deleted_at DATETIME(3),
    PRIMARY KEY (id),
    UNIQUE KEY idx_username (username),
    UNIQUE KEY idx_email (email),
    UNIQUE KEY idx_notelp (notelp)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Recreate stores table
CREATE TABLE stores (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    nama_toko VARCHAR(255) NOT NULL,
    url_foto TEXT,
    user_id BIGINT UNSIGNED NOT NULL,
    created_at DATETIME(3),
    updated_at DATETIME(3),
    deleted_at DATETIME(3),
    PRIMARY KEY (id),
    CONSTRAINT stores_ibfk_1 FOREIGN KEY (user_id) REFERENCES users (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
