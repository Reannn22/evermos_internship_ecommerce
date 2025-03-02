# Product Logs API Documentation

## Overview

The Product Logs API provides endpoints to track and manage the history of product-related activities within the application. This API enables the recording, retrieval, updating, and deletion of log entries that document changes to products, including inventory updates, price modifications, attribute changes, and other significant events in a product's lifecycle.

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

### 1. Create Product Log

Creates a new log entry for product changes.

- **URL**: `/product-logs`
- **Method**: `POST`
- **Authentication**: Required
- **Content-Type**: `application/json`

**Request Body**:

```json
{
    "product_id": integer,
    "nama_produk": "string",
    "slug": "string",
    "harga_reseller": "string",
    "harga_konsumen": "string",
    "deskripsi": "string",
    "store_id": integer,
    "category_id": integer
}
```

### 2. Get Specific Log

Retrieves details of a specific log entry.

- **URL**: `/product-logs/{id}`
- **Method**: `GET`
- **Authentication**: Required

### 3. Get All Product Logs

Retrieves all product log entries.

- **URL**: `/product-logs`
- **Method**: `GET`
- **Authentication**: Required

### 4. Update Log

Updates an existing log entry.

- **URL**: `/product-logs/{id}`
- **Method**: `PUT`
- **Authentication**: Required
- **Content-Type**: `application/json`

**Request Body**:

```json
{
    "product_id": integer,
    "nama_produk": "string",
    "slug": "string",
    "harga_reseller": "string",
    "harga_konsumen": "string",
    "deskripsi": "string",
    "store_id": integer,
    "category_id": integer
}
```

### 5. Delete Log

Removes a log entry from the system.

- **URL**: `/product-logs/{id}`
- **Method**: `DELETE`
- **Authentication**: Required

## Response Codes

- `200 OK`: Request successful
- `201 Created`: Log entry created successfully
- `400 Bad Request`: Invalid request parameters
- `401 Unauthorized`: Authentication required
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: Log entry not found
- `500 Internal Server Error`: Server error

## Notes

- Log entries are immutable by design
- All monetary values are in Indonesian Rupiah (IDR)
- Timestamps are automatically recorded in ISO 8601 format
- Log deletion requires highest administrative privileges
- Each log entry maintains complete change history
- Automatic logging for product modifications
- Supports audit trail requirements
- Log retention policies may apply
- Changes are tracked with user attribution
- Bulk operations are not supported for data integrity
