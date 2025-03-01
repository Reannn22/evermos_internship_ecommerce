-- Drop the old table if it exists
DROP TABLE IF EXISTS users;

-- Create the table with the correct schema
CREATE TABLE users (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    nama VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    kata_sandi VARCHAR(255) NOT NULL,
    notelp VARCHAR(15) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
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
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
