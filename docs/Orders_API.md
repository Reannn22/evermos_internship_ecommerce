# Orders API Documentation

## Overview

The Orders API provides endpoints to manage customer orders within the application. This API enables the creation, retrieval, updating, and deletion of order information, which is essential for e-commerce functionality. The Orders API tracks the entire lifecycle of a customer purchase from initial checkout to fulfillment and delivery.

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

### 1. Create Order

Creates a new order in the system.

- **URL**: `/orders`
- **Method**: `POST`
- **Authentication**: Required
- **Content-Type**: `application/json`

**Request Body**:

```json
{
    "transaction_detail_id": integer,
    "product_status": "string"
}
```

### 2. Get Specific Order

Retrieves details of a specific order.

- **URL**: `/orders/{id}`
- **Method**: `GET`
- **Authentication**: Required

### 3. Get All Orders

Retrieves all orders for the authenticated user.

- **URL**: `/orders`
- **Method**: `GET`
- **Authentication**: Required
- **Notes**:
  - Regular users see only their orders
  - Administrators can view all orders

### 4. Update Order

Updates an existing order.

- **URL**: `/orders/{id}`
- **Method**: `PUT`
- **Authentication**: Required

### 5. Delete Order

Removes an order from the system.

- **URL**: `/orders/{id}`
- **Method**: `DELETE`
- **Authentication**: Required

## Response Codes

- `200 OK`: Request successful
- `201 Created`: Order created successfully
- `400 Bad Request`: Invalid request parameters
- `401 Unauthorized`: Authentication required
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: Order not found
- `500 Internal Server Error`: Server error

## Notes

- Order IDs are unique and auto-generated
- All monetary values are in Indonesian Rupiah (IDR)
- Orders maintain complete audit trails
- Status changes are tracked with timestamps
- Deletion is typically replaced by cancellation
- Order history is preserved for reporting
- Updates may be restricted after processing begins
- All timestamps are in ISO 8601 format
- Payment status is tracked separately
- Order modifications trigger notifications
