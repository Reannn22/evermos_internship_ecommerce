# Product Discounts API Documentation

## Overview

The Product Discounts API provides endpoints to manage product-specific discounts within the application. Unlike coupons that users apply at checkout, product discounts are automatically applied to specific products or product categories. This API enables the creation, retrieval, updating, and deletion of discount configurations that determine special pricing, sales events, and promotional offers directly at the product level.

## Base URL

```
http://localhost:3000/api/v1
```

## Authentication

All endpoints require administrative authentication via JWT token:

```
Authorization: Bearer <your_token>
```

## Endpoints

### 1. Create Product Discount

Creates a new product discount configuration.

- **URL**: `/diskon-produk`
- **Method**: `POST`
- **Authentication**: Required
- **Content-Type**: `application/json`

**Request Body**:

```json
{
    "product_id": integer,
    "harga_konsumen": "string"
}
```

### 2. Get Specific Product Discount

Retrieves details of a specific discount.

- **URL**: `/diskon-produk/{id}`
- **Method**: `GET`
- **Authentication**: Required

### 3. Get All Product Discounts

Retrieves all product discounts.

- **URL**: `/diskon-produk`
- **Method**: `GET`
- **Authentication**: Required

### 4. Update Product Discount

Updates an existing discount configuration.

- **URL**: `/diskon-produk/{id}`
- **Method**: `PUT`
- **Authentication**: Required
- **Content-Type**: `application/json`

**Request Body**:

```json
{
    "product_id": integer,
    "harga_konsumen": "string"
}
```

### 5. Delete Product Discount

Removes a discount configuration.

- **URL**: `/diskon-produk/{id}`
- **Method**: `DELETE`
- **Authentication**: Required

## Response Codes

- `200 OK`: Request successful
- `201 Created`: Discount created successfully
- `400 Bad Request`: Invalid request parameters
- `401 Unauthorized`: Authentication required
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: Discount not found
- `500 Internal Server Error`: Server error

## Notes

- Discount IDs are unique and auto-generated
- All monetary values are in Indonesian Rupiah (IDR)
- Discounts are applied automatically at product level
- Multiple discounts can be stacked based on rules
- Discount history is maintained for auditing
- All timestamps are in ISO 8601 format
- Changes require administrative privileges
- Price calculations are performed in real-time
- Discounts can be time-bound or permanent
- System validates discount against base price
