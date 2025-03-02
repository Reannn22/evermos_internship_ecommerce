# Product Reviews API Documentation

## Overview

This API provides comprehensive functionality for managing product reviews within the system. It enables users and administrators to create, access, modify, and remove customer reviews and ratings for products.

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

### 1. Create Product Review

Creates a new product review.

- **URL**: `/reviews`
- **Method**: `POST`
- **Authentication**: Required
- **Content-Type**: `application/json`

**Request Body**:

```json
{
    "ulasan": "string",
    "rating": integer,
    "id_toko": integer,
    "id_produk": integer
}
```

### 2. Get Specific Product Review

Retrieves details of a specific review.

- **URL**: `/reviews/{id}`
- **Method**: `GET`
- **Authentication**: Required

### 3. Get All Product Reviews

Retrieves all product reviews with filtering options.

- **URL**: `/reviews`
- **Method**: `GET`
- **Authentication**: Required
- **Query Parameters**:
  - product_id: integer (optional)
  - rating: integer (optional)
  - sort: string (optional)
  - page: integer (optional)
  - limit: integer (optional)

### 4. Update Product Review

Updates an existing review.

- **URL**: `/reviews/{id}`
- **Method**: `PUT`
- **Authentication**: Required
- **Content-Type**: `application/json`

**Request Body**:

```json
{
    "ulasan": "string",
    "rating": integer,
    "id_toko": integer,
    "id_produk": integer
}
```

### 5. Delete Product Review

Removes a review from the system.

- **URL**: `/reviews/{id}`
- **Method**: `DELETE`
- **Authentication**: Required

## Response Codes

- `200 OK`: Request successful
- `201 Created`: Review created successfully
- `400 Bad Request`: Invalid request parameters
- `401 Unauthorized`: Authentication required
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: Review not found
- `500 Internal Server Error`: Server error

## Notes

- Rating values must be between 1 and 5
- Reviews can only be modified by their authors or administrators
- Review IDs are unique and auto-generated
- Users can only review products they have purchased
- All timestamps are in ISO 8601 format
- Review content is subject to moderation
- Updates to reviews maintain edit history
- Deleted reviews may be soft-deleted
- Review metrics are updated in real-time
