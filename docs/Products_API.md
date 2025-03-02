# Products API Documentation

## Overview

The Products API provides comprehensive functionality to manage product information within the e-commerce system. This API enables creation, retrieval, updating, and deletion of products, along with advanced search and categorization features.

## Base URL

```
http://localhost:3000/api/v1
```

## Authentication

All endpoints require Bearer token authentication. Include the token in the Authorization header:

```
Authorization: Bearer <your_token>
```

## Endpoints

### 1. Create Product

Creates a new product in the system.

- **URL**: `/product`
- **Method**: `POST`
- **Authentication**: Required
- **Content-Type**: `multipart/form-data`

**Request Body**:

```
FormData:
- nama_produk: string
- category_id: string
- harga_reseller: string
- harga_konsumen: string
- stok: string
- deskripsi: string
- photo_url: string
```

### 2. Get Specific Product

Retrieves detailed information about a specific product.

- **URL**: `/product/{id}`
- **Method**: `GET`
- **Authentication**: Required

### 3. Get All Products

Retrieves a paginated list of all products.

- **URL**: `/product`
- **Method**: `GET`
- **Authentication**: Required

### 4. Get Products by Category

Retrieves products within a specific category.

- **URL**: `/produk/category/{id}`
- **Method**: `GET`
- **Authentication**: Required

### 5. Search Products

Searches products using keywords.

- **URL**: `/product/search`
- **Method**: `GET`
- **Authentication**: Required
- **Query Parameters**:
  - q: string (search query)
  - category: string (optional)
  - min_price: number (optional)
  - max_price: number (optional)

### 6. Get Related Products

Retrieves products related to a specific product.

- **URL**: `/product/{id}/related`
- **Method**: `GET`
- **Authentication**: Required

### 7. Update Product

Updates an existing product.

- **URL**: `/product/{id}`
- **Method**: `PUT`
- **Authentication**: Required
- **Content-Type**: `multipart/form-data`

**Request Body**:

```
FormData:
- nama_produk: string
- category_id: string
- harga_reseller: string
- harga_konsumen: string
- stok: string
- deskripsi: string
- photo_url: string
```

### 8. Delete Product

Removes a product from the system.

- **URL**: `/product/{id}`
- **Method**: `DELETE`
- **Authentication**: Required

## Response Codes

- `200 OK`: Request successful
- `201 Created`: Product created successfully
- `400 Bad Request`: Invalid request parameters
- `401 Unauthorized`: Authentication required
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: Product not found
- `500 Internal Server Error`: Server error

## Notes

- Product IDs are unique and auto-generated
- All monetary values are in Indonesian Rupiah (IDR)
- Product photos must be valid URLs or uploaded files
- Stock quantities must be non-negative integers
- Category assignment is required
- Product names must be unique within a category
- Prices must include taxes and other charges
- All timestamps are in ISO 8601 format
- Related products are determined by category and tags
