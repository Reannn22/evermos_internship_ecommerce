# Evermos Internship E-Commerce
![E-Commerce API Diagram](https://github.com/user-attachments/assets/f3f5d81f-065c-4256-89e6-935076b630bf)

## 📌 Project Description
Evermos Internship E-Commerce is a project developed as part of the internship program at Evermos. This project aims to build an efficient e-commerce system with a **Golang**-based backend. This repository contains various core components of the system, including database models, API endpoints, middleware, and other essential services.

## 🚀 Technologies Used
- **Golang** - Primary programming language
- **MySQL** - Database for data storage
- **Fiber** - Framework for building REST API
- **JWT** - Authentication
- **Docker** - Containerization
- **GORM** - ORM for Golang

## 📂 Directory Structure
```
📦 evermos_internship_excommerce
├── 📂 config          # Application configuration
├── 📂 configs         # Additional settings
├── 📂 database        # Database schema and connections
├── 📂 docs            # API documentation
├── 📂 exceptions      # Error handlers
├── 📂 handlers        # API controllers
├── 📂 middleware      # Middleware such as authentication
├── 📂 migrations      # Database migration files
├── 📂 models          # Data structures and entities
├── 📂 repositories    # Database access layer
├── 📂 services        # Business logic
├── 📂 uploads         # Directory for storing files
├── 📂 utils           # Helper functions
├── .env              # Environment configuration file
├── LICENSE           # Project license (GPL-3.0)
├── go.mod            # Golang dependencies
├── go.sum            # Checksum dependencies
└── main.go           # Application entry point
```

## 🔧 Installation & Setup
1. **Install Golang**
   [Installation Guide](https://go.dev/doc/install)
2. **Install Fiber Framework**
   [Installation Guide](https://docs.gofiber.io/)
3. **Install GORM**
   [Installation Guide](https://gorm.io/)
4. **Install MySQL**
   [Installation Guide](https://dev.mysql.com/downloads/installer/)
5. **Install Postman** (For API testing)
   [Download Postman](https://www.postman.com/downloads/)
6. **Clone the repository**
   ```sh
   git clone https://github.com/Reannn22/evermos_internship_ecommerce.git
   cd evermos_internship_ecommerce
   ```
7. **Create a `.env` file** based on the required configuration:
   ```sh
   cp .env.example .env
   ```
8. **Run the application with Docker**
   ```sh
   docker-compose up --build
   ```
9. **Run locally** (without Docker):
   ```sh
   go run main.go
   ```

## 📌 API Endpoints
Complete API documentation is available in the `/docs` folder and can be accessed via the following link:
[API Documentation](https://drive.google.com/drive/folders/1qqLcsVxjqKUAaTr1hsoVoj9xxU47IZiZ?usp=sharing)

### API Reference Files:
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

## 📜 License

This project is licensed under the **GNU General Public License v3.0 (GPL-3.0)**.

You may obtain a copy of the license at:

[Evermos Internship E-Commerce License](https://github.com/Reannn22/evermos_internship_ecommerce/blob/main/LICENSE)

### Summary
- You are free to use, modify, and distribute this project under the terms of the GPL-3.0 license.
- Any modifications or derivative works must also be licensed under GPL-3.0.
- This project is provided "as-is" without any warranties.

For the full license text, please refer to the link above.
