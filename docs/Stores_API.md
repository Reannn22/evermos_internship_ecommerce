# Stores API Documentation

## Overview

The Stores API provides comprehensive functionality to manage store entities within the system. This API allows clients to create, retrieve, update, and delete store information, enabling seamless store management operations.

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

### 1. Create Store

Creates a new store record.

- **URL**: `/toko`
- **Method**: `POST`
- **Authentication**: Required
- **Content-Type**: `multipart/form-data`

**Request Body**:

```
FormData:
- nama_toko: string
- url_foto: string
- id_foto: string
- photo: file
- deskripsi_toko: string
```

### 2. Get Specific Store

Retrieves details of a specific store.

- **URL**: `/toko/{id}`
- **Method**: `GET`
- **Authentication**: Required

### 3. Get All Stores

Retrieves a list of all stores with pagination.

- **URL**: `/toko`
- **Method**: `GET`
- **Authentication**: Required
- **Query Parameters**:
  - limit: integer (items per page)
  - page: integer (page number)

### 4. Get Store by Registered User

Retrieves the authenticated user's store.

- **URL**: `/toko/my`
- **Method**: `GET`
- **Authentication**: Required

### 5. Update Store

Updates store information.

- **URL**: `/toko/{id}`
- **Method**: `PUT`
- **Authentication**: Required
- **Content-Type**: `multipart/form-data`

**Request Body**:

```
FormData:
- nama_toko: string
- url_foto: string
```

### 6. Delete Store

Removes a store from the system.

- **URL**: `/toko/{id}`
- **Method**: `DELETE`
- **Authentication**: Required

## Response Codes

- `200 OK`: Request successful
- `201 Created`: Store created successfully
- `400 Bad Request`: Invalid request parameters
- `401 Unauthorized`: Authentication required
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: Store not found
- `500 Internal Server Error`: Server error

## Notes

- Store IDs are unique and auto-generated
- Photo uploads support common image formats (JPEG, PNG)
- Store deletion may affect related products and transactions
- Each user can have one associated store
- All timestamps are in ISO 8601 format
- URLs must be valid and accessible
- Store names must be unique in the system
