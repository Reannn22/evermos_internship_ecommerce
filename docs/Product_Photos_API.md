# Product Photos API Documentation

## Overview

This API provides access to manage product photos within the system. You can use these endpoints to upload, retrieve, update, and delete product photos.

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

### 1. Create Product Photo

Creates a new product photo entry.

- **URL**: `/product-photos`
- **Method**: `POST`
- **Authentication**: Required
- **Content-Type**: `application/json`

**Request Body**:

```json
{
    "product_id": integer,
    "photo_id": integer,
    "url": "string"
}
```

### 2. Get Specific Product Photo

Retrieves a specific product photo.

- **URL**: `/product-photos/{id}`
- **Method**: `GET`
- **Authentication**: Required

### 3. Get All Product Photos

Retrieves all product photos.

- **URL**: `/product-photos`
- **Method**: `GET`
- **Authentication**: Required

### 4. Update Product Photo

Updates an existing product photo.

- **URL**: `/product-photos/{id}`
- **Method**: `PUT`
- **Authentication**: Required
- **Content-Type**: `application/json`

**Request Body**:

```json
{
    "product_id": integer,
    "photo_id": integer,
    "url": "string"
}
```

### 5. Delete Product Photo

Removes a product photo from the system.

- **URL**: `/product-photos/{id}`
- **Method**: `DELETE`
- **Authentication**: Required

## Response Codes

- `200 OK`: Request successful
- `201 Created`: Photo uploaded successfully
- `400 Bad Request`: Invalid request parameters
- `401 Unauthorized`: Authentication required
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: Photo not found
- `500 Internal Server Error`: Server error

## Notes

- Supported image formats: JPEG, PNG
- Maximum file size: 5MB
- Photo IDs are unique and auto-generated
- URLs must be valid and accessible
- Each photo must be associated with a product
- Photos are stored securely with backup
- Photo metadata includes upload timestamp
- Deletion is permanent and cannot be undone
- High-resolution originals are preserved
- Thumbnails are automatically generated
