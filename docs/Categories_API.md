# Categories API Documentation

## Overview

The Categories API provides endpoints to manage product categories within the application. This API enables administrators and authorized users to create, retrieve, update, and delete product category information, which is essential for organizing products, filtering, and navigation throughout the e-commerce platform.

## Base URL

```
http://localhost:3000/api/v1
```

## Authentication

All endpoints require authentication via JWT token:

```
Authorization: Bearer <your_token>
```

## Endpoints

### 1. Create Category

Creates a new product category.

- **URL**: `/category`
- **Method**: `POST`
- **Authentication**: Required
- **Content-Type**: `application/json`

**Request Body**:

```json
{
  "nama_category": "string"
}
```

### 2. Get Specific Category

Retrieves details of a specific category.

- **URL**: `/category/{id}`
- **Method**: `GET`
- **Authentication**: Required

### 3. Get All Categories

Retrieves all available categories.

- **URL**: `/category`
- **Method**: `GET`
- **Authentication**: Required

### 4. Update Category

Updates an existing category.

- **URL**: `/category/{id}`
- **Method**: `PUT`
- **Authentication**: Required
- **Content-Type**: `application/json`

**Request Body**:

```json
{
  "nama_category": "string"
}
```

### 5. Delete Category

Removes a category from the system.

- **URL**: `/category/{id}`
- **Method**: `DELETE`
- **Authentication**: Required

## Response Codes

- `200 OK`: Request successful
- `201 Created`: Category created successfully
- `400 Bad Request`: Invalid request parameters
- `401 Unauthorized`: Authentication required
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: Category not found
- `500 Internal Server Error`: Server error

## Notes

- Category IDs are unique and auto-generated
- Categories can't be deleted if products are assigned
- Category names must be unique
- All timestamps are in ISO 8601 format
- Category changes are logged for auditing
- Supports hierarchical category structure
- Category operations require admin privileges
- Categories are used for product navigation
- Changes affect product categorization
- Bulk operations are not supported
