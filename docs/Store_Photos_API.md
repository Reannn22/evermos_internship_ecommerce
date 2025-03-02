# Store Photos API Documentation

## Overview

The Store Photos API provides functionality to manage and access store photos in the system. This API allows clients to upload, retrieve, update, and delete photos associated with stores.

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

### 1. Create Store Photo

Uploads a new store photo.

- **URL**: `/toko-foto`
- **Method**: `POST`
- **Authentication**: Required
- **Content-Type**: `multipart/form-data`

### 2. Get Specific Store Photo

Retrieves a specific store photo.

- **URL**: `/toko-foto/{id}`
- **Method**: `GET`
- **Authentication**: Required

### 3. Get All Store Photos

Retrieves all store photos.

- **URL**: `/toko-foto`
- **Method**: `GET`
- **Authentication**: Required

### 4. Update Store Photo

Updates an existing store photo.

- **URL**: `/toko-foto/{id}`
- **Method**: `PUT`
- **Authentication**: Required
- **Content-Type**: `application/json`

**Request Body**:

```json
{
    "id_toko": integer,
    "id_foto": string,
    "url": string
}
```

### 5. Delete Store Photo

Removes a store photo from the system.

- **URL**: `/toko-foto/{id}`
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
- Photos are stored securely with backup
- URLs must be valid and accessible
- Each photo is associated with exactly one store
- Photo IDs are unique and auto-generated
- Deletion is permanent and cannot be undone
- All timestamps are in ISO 8601 format
