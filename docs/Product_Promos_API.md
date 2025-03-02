# Product Promos API Documentation

## Overview

This API provides comprehensive functionality for managing product promotions within the system. It enables creation, retrieval, modification, and removal of promotional offers that can be applied to products or product categories.

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

### 1. Create Product Promo

Creates a new product promotion.

- **URL**: `/promos`
- **Method**: `POST`
- **Authentication**: Required
- **Content-Type**: `application/json`

**Request Body**:

```json
{
    "id_toko": integer,
    "id_produk": integer,
    "promo": "string"
}
```

### 2. Get Specific Product Promo

Retrieves details of a specific promotion.

- **URL**: `/promos/{id}`
- **Method**: `GET`
- **Authentication**: Required

### 3. Get All Product Promos

Retrieves all product promotions.

- **URL**: `/promos`
- **Method**: `GET`
- **Authentication**: Required

### 4. Update Product Promo

Updates an existing promotion.

- **URL**: `/promos/{id}`
- **Method**: `PUT`
- **Authentication**: Required
- **Content-Type**: `application/json`

### 5. Delete Product Promo

Removes a specific promotion.

- **URL**: `/promos/{id}`
- **Method**: `DELETE`
- **Authentication**: Required

### 6. Clear Product Promos

Deactivates all active promotions.

- **URL**: `/promos/clear`
- **Method**: `DELETE`
- **Authentication**: Required

## Response Codes

- `200 OK`: Request successful
- `201 Created`: Promotion created successfully
- `400 Bad Request`: Invalid request parameters
- `401 Unauthorized`: Authentication required
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: Promotion not found
- `500 Internal Server Error`: Server error

## Notes

- Promotion IDs are unique and auto-generated
- Active promotions can be scheduled for future dates
- Promotions can be applied to specific products or categories
- Multiple active promotions are allowed per product
- Promotion history is maintained for auditing
- All timestamps are in ISO 8601 format
- Deletion of active promotions requires admin privileges
- Promotional discounts are calculated in real-time
- Usage statistics are tracked for each promotion
