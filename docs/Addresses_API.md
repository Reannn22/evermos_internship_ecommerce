# Addresses API Documentation

## Overview

The Addresses API provides endpoints to manage user delivery addresses within the application. This API enables users to create, retrieve, update, and delete address information, which is essential for order fulfillment and delivery processes.

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

### 1. Create Address

Creates a new address entry.

- **URL**: `/alamat`
- **Method**: `POST`
- **Authentication**: Required
- **Content-Type**: `application/json`

**Request Body**:

```json
{
  "judul_alamat": "string",
  "nama_penerima": "string",
  "no_telp": "string",
  "detail_alamat": "string",
  "id_provinsi": "string",
  "id_kota": "string"
}
```

### 2. Get Specific Address

Retrieves details of a specific address.

- **URL**: `/address/{id}`
- **Method**: `GET`
- **Authentication**: Required

### 3. Get All Addresses

Retrieves all addresses for the authenticated user.

- **URL**: `/address`
- **Method**: `GET`
- **Authentication**: Required

### 4. Update Address

Updates an existing address.

- **URL**: `/address/{id}`
- **Method**: `PUT`
- **Authentication**: Required
- **Content-Type**: `application/json`

**Request Body**:

```json
{
  "judul_alamat": "string",
  "nama_penerima": "string",
  "no_telp": "string",
  "detail_alamat": "string",
  "id_provinsi": "string",
  "id_kota": "string"
}
```

### 5. Delete Address

Removes an address from the system.

- **URL**: `/address/{id}`
- **Method**: `DELETE`
- **Authentication**: Required

## Response Codes

- `200 OK`: Request successful
- `201 Created`: Address created successfully
- `400 Bad Request`: Invalid request parameters
- `401 Unauthorized`: Authentication required
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: Address not found
- `500 Internal Server Error`: Server error

## Notes

- Address IDs are unique and auto-generated
- Users can have multiple saved addresses
- Phone numbers must be in valid format
- Province and city IDs must be valid
- Addresses are user-specific
- All timestamps are in ISO 8601 format
- Default address can be marked
- Address validation is performed
- Postal codes are validated
- Address changes are logged for auditing
