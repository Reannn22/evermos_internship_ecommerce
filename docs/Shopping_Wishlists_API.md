# Shopping Wishlists API Documentation

## Overview

This API provides comprehensive functionality for managing customer wishlists within the e-commerce system. It enables users to create and maintain lists of desired products for future consideration or purchase, enhancing the shopping experience by allowing customers to save and organize items of interest.

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

### 1. Create Shopping Wishlist

Creates a new wishlist for a user.

- **URL**: `/wishlist-shopping`
- **Method**: `POST`
- **Authentication**: Required
- **Content-Type**: `application/json`

**Request Body**:

```json
{
    "store_id": integer,
    "product_id": integer
}
```

### 2. Get Specific Shopping Wishlist

Retrieves detailed information about a specific wishlist.

- **URL**: `/wishlist-shopping/{id}`
- **Method**: `GET`
- **Authentication**: Required

### 3. Get All Shopping Wishlists

Retrieves all wishlists for the authenticated user.

- **URL**: `/wishlist-shopping`
- **Method**: `GET`
- **Authentication**: Required

### 4. Update Shopping Wishlist

Updates an existing wishlist.

- **URL**: `/wishlist-shopping/{id}`
- **Method**: `PUT`
- **Authentication**: Required
- **Content-Type**: `application/json`

**Request Body**:

```json
{
    "store_id": integer,
    "product_id": integer
}
```

### 5. Delete Shopping Wishlist

Removes a specific wishlist.

- **URL**: `/wishlist-shopping/{id}`
- **Method**: `DELETE`
- **Authentication**: Required

### 6. Clear Shopping Wishlist

Removes all items from a wishlist while maintaining the wishlist structure.

- **URL**: `/wishlist-shopping/clear`
- **Method**: `DELETE`
- **Authentication**: Required

## Response Codes

- `200 OK`: Request successful
- `201 Created`: Wishlist created successfully
- `400 Bad Request`: Invalid request parameters
- `401 Unauthorized`: Authentication required
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: Wishlist not found
- `500 Internal Server Error`: Server error

## Notes

- Users can create multiple wishlists
- Wishlist IDs are unique and auto-generated
- Products can be in multiple wishlists
- All timestamps are in ISO 8601 format
- Deleting a wishlist is permanent
- Clearing a wishlist preserves the wishlist structure
- Product availability is checked in real-time
- Price changes are reflected in wishlists
- Maximum items per wishlist: 100
