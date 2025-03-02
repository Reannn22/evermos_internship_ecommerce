# Product Coupons API Documentation

## Overview

The Product Coupons API provides endpoints to manage promotional coupons within the application. This API enables the creation, retrieval, updating, and deletion of coupon codes that can be applied to products or orders for discounts. Coupons are a critical marketing tool that incentivizes purchases, rewards customer loyalty, and drives sales during promotional periods.

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

### 1. Create Product Coupon

Creates a new coupon code in the system.

- **URL**: `/coupons`
- **Method**: `POST`
- **Authentication**: Required
- **Content-Type**: `application/json`

### 2. Get Specific Coupon

Retrieves details of a specific coupon.

- **URL**: `/coupons/{id}`
- **Method**: `GET`
- **Authentication**: Required

### 3. Get All Coupons

Retrieves all coupons in the system.

- **URL**: `/coupons`
- **Method**: `GET`
- **Authentication**: Required

### 4. Update Coupon

Updates an existing coupon.

- **URL**: `/coupons/{id}`
- **Method**: `PUT`
- **Authentication**: Required
- **Content-Type**: `application/json`

**Request Body**:

```json
{
    "product_id": integer,
    "code": "string",
    "discount": integer,
    "valid_from": "string",
    "valid_to": "string"
}
```

### 5. Delete Coupon

Removes a coupon from the system.

- **URL**: `/coupons/{id}`
- **Method**: `DELETE`
- **Authentication**: Required

## Response Codes

- `200 OK`: Request successful
- `201 Created`: Coupon created successfully
- `400 Bad Request`: Invalid request parameters
- `401 Unauthorized`: Authentication required
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: Coupon not found
- `500 Internal Server Error`: Server error

## Notes

- Coupon codes must be unique
- Discounts can be percentage or fixed amount
- Validity periods are in ISO 8601 format
- Usage limits can be set per coupon
- Coupons can be restricted to specific products
- Active coupons are validated at checkout
- Usage statistics are tracked per coupon
- Expired coupons are automatically deactivated
- Multiple coupons may be stackable based on rules
- Admin privileges required for management
