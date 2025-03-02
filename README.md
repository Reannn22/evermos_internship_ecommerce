# Evermos Internship E-Commerce
![E-Commerce API Diagram](https://github.com/user-attachments/assets/f3f5d81f-065c-4256-89e6-935076b630bf)

## 📌 Deskripsi Proyek
Evermos Internship E-Commerce adalah proyek yang dikembangkan sebagai bagian dari program internship di Evermos. Proyek ini bertujuan untuk membangun sistem e-commerce yang efisien dengan backend berbasis **Golang**. Repository ini berisi berbagai komponen utama dari sistem, termasuk model database, API endpoints, middleware, dan layanan lainnya.

## 🚀 Teknologi yang Digunakan
- **Golang** - Bahasa pemrograman utama
- **MySQL** - Database untuk penyimpanan data
- **Fiber** - Framework untuk membangun REST API
- **JWT** - Authentication
- **Docker** - Containerization
- **GORM** - ORM untuk Golang

## 📂 Struktur Direktori
```
📦 evermos_internship_excommerce
├── 📂 config          # Konfigurasi aplikasi
├── 📂 configs         # Pengaturan tambahan
├── 📂 database        # Skema database dan koneksi
├── 📂 docs            # Dokumentasi API
├── 📂 exceptions      # Handler untuk error
├── 📂 handlers        # Controller API
├── 📂 middleware      # Middleware seperti autentikasi
├── 📂 migrations      # File migrasi database
├── 📂 models          # Struktur data dan entity
├── 📂 repositories    # Layer akses database
├── 📂 services        # Logika bisnis aplikasi
├── 📂 uploads         # Direktori untuk menyimpan file
├── 📂 utils           # Helper functions
├── .env              # File konfigurasi environment
├── LICENSE           # Lisensi proyek (GPL-3.0)
├── go.mod            # Dependencies Golang
├── go.sum            # Checksum dependencies
└── main.go           # Entry point aplikasi
```

## 🔧 Instalasi & Setup
1. **Install Bahasa Golang**
   [Panduan Instalasi](https://go.dev/doc/install)
2. **Install Framework Fiber**
   [Panduan Instalasi](https://docs.gofiber.io/)
3. **Install GORM**
   [Panduan Instalasi](https://gorm.io/)
4. **Install MySQL/MariaDB**
   [Panduan Instalasi](https://dev.mysql.com/downloads/installer/)
5. **Install Postman** (Untuk uji API)
   [Download Postman](https://www.postman.com/downloads/)
6. **Clone repository**
   ```sh
   git clone https://github.com/Reannn22/evermos_internship_ecommerce.git
   cd evermos_internship_ecommerce
   ```
7. **Buat file .env** berdasarkan konfigurasi yang diperlukan:
   ```sh
   cp .env.example .env
   ```
8. **Jalankan aplikasi dengan Docker**
   ```sh
   docker-compose up --build
   ```
9. **Jalankan secara lokal** (tanpa Docker):
   ```sh
   go run main.go
   ```

## 📌 API Endpoints
Dokumentasi lengkap tersedia di folder `/docs`. Berikut adalah daftar dokumentasi API:

- [Addresses API](https://github.com/Reannn22/evermos_internship_ecommerce/blob/main/docs/Addresses_API.md)
- [Categories API](https://github.com/Reannn22/evermos_internship_ecommerce/blob/main/docs/Categories_API.md)
- [Notifications API](https://github.com/Reannn22/evermos_internship_ecommerce/blob/main/docs/Notifications_API.md)
- [Orders API](https://github.com/Reannn22/evermos_internship_ecommerce/blob/main/docs/Orders_API.md)
- [Product Coupons API](https://github.com/Reannn22/evermos_internship_ecommerce/blob/main/docs/Product_Coupons_API.md)
- [Product Discounts API](https://github.com/Reannn22/evermos_internship_ecommerce/blob/main/docs/Product_Discounts_API.md)
- [Product Logs API](https://github.com/Reannn22/evermos_internship_ecommerce/blob/main/docs/Product_Logs_API.md)
- [Product Photos API](https://github.com/Reannn22/evermos_internship_ecommerce/blob/main/docs/Product_Photos_API.md)
- [Product Promos API](https://github.com/Reannn22/evermos_internship_ecommerce/blob/main/docs/Product_Promos_API.md)
- [Product Reviews API](https://github.com/Reannn22/evermos_internship_ecommerce/blob/main/docs/Product_Reviews_API.md)
- [Products API](https://github.com/Reannn22/evermos_internship_ecommerce/blob/main/docs/Products_API.md)
- [Shopping Carts API](https://github.com/Reannn22/evermos_internship_ecommerce/blob/main/docs/Shopping_Carts_API.md)
- [Shopping Wishlists API](https://github.com/Reannn22/evermos_internship_ecommerce/blob/main/docs/Shopping_Wishlists_API.md)
- [Store Photos API](https://github.com/Reannn22/evermos_internship_ecommerce/blob/main/docs/Store_Photos_API.md)
- [Stores API](https://github.com/Reannn22/evermos_internship_ecommerce/blob/main/docs/Stores_API.md)
- [Transaction Details API](https://github.com/Reannn22/evermos_internship_ecommerce/blob/main/docs/Transaction_Details_API.md)
- [Transactions API](https://github.com/Reannn22/evermos_internship_ecommerce/blob/main/docs/Transactions_API.md)
- [Users API](https://github.com/Reannn22/evermos_internship_ecommerce/blob/main/docs/Users_API.md)

## 📜 Lisensi
Proyek ini dilisensikan di bawah **GPL-3.0**. Silakan lihat file `LICENSE` untuk detailnya.
