# Shopping Carts API Documentation

## Overview

This API enables comprehensive management of shopping carts within the e-commerce system. It provides functionality to create, retrieve, update, and delete shopping carts, allowing customers to manage their product selections before proceeding to checkout.

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

### 1. Create Shopping Cart

Creates a new shopping cart for a user.

- **URL**: `/keranjang-belanja`
- **Method**: `POST`
- **Authentication**: Required
- **Content-Type**: `application/json`

**Request Body**:

```json
{
    "id_toko": integer,
    "id_produk": integer,
    "jumlah_produk": integer
}
```

### 2. Get Specific Shopping Cart

Retrieves detailed information about a specific cart.

- **URL**: `/keranjang-belanja/{id}`
- **Method**: `GET`
- **Authentication**: Required

### 3. Get All Shopping Carts

Retrieves all shopping carts in the system.

- **URL**: `/keranjang-belanja`
- **Method**: `GET`
- **Authentication**: Required

### 4. Update Shopping Cart

Updates an existing shopping cart.

- **URL**: `/keranjang-belanja/{id}`
- **Method**: `PUT`
- **Authentication**: Required
- **Content-Type**: `application/json`

**Request Body**:

```json
{
    "id_toko": integer,
    "id_produk": integer,
    "jumlah_produk": integer
}
```

### 5. Delete Shopping Cart

Removes a specific shopping cart.

- **URL**: `/keranjang-belanja/{id}`
- **Method**: `DELETE`
- **Authentication**: Required

### 6. Clear Shopping Cart

Removes all items from a shopping cart while maintaining the cart structure.

- **URL**: `/keranjang-belanja/clear`
- **Method**: `DELETE`
- **Authentication**: Required

## Response Codes

- `200 OK`: Request successful
- `201 Created`: Cart created successfully
- `400 Bad Request`: Invalid request parameters
- `401 Unauthorized`: Authentication required
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: Cart not found
- `500 Internal Server Error`: Server error

## Notes

- Shopping cart IDs are unique and auto-generated
- Products in cart are validated for availability
- Prices are automatically updated to current rates
- Cart contents expire after 24 hours of inactivity
- Maximum items per cart: 50
- Quantities are subject to stock availability
- All monetary values are in Indonesian Rupiah (IDR)
- All timestamps are in ISO 8601 format
- Cart totals include item prices, taxes, and discounts
