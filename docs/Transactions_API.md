# Transactions API Documentation

## Overview

The Transactions API provides comprehensive functionality to manage financial transactions within the system. This API allows clients to create, retrieve, update, and delete transaction records, enabling efficient transaction processing and management across the platform.

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

### 1. Create Transaction

Creates a new transaction record.

- **URL**: `/trx`
- **Method**: `POST`
- **Authentication**: Required
- **Content-Type**: `application/json`

**Request Body**:

```json
{
    "user_id": integer,
    "alamat_pengiriman": integer,
    "harga_total": integer,
    "method_bayar": "string"
}
```

### 2. Get Specific Transaction

Retrieves details of a specific transaction.

- **URL**: `/trx/{id}`
- **Method**: `GET`
- **Authentication**: Required

### 3. Get All Transactions

Retrieves a list of all transactions.

- **URL**: `/trx`
- **Method**: `GET`
- **Authentication**: Required

### 4. Update Transaction

Updates transaction information.

- **URL**: `/trx/{id}`
- **Method**: `PUT`
- **Authentication**: Required
- **Content-Type**: `application/json`

**Request Body**:

```json
{
    "method_bayar": "string",
    "alamat_kirim": integer,
    "detail_trx": [
        {
            "id_produk": integer,
            "quantity": integer
        }
    ]
}
```

### 5. Delete Transaction

Removes a transaction from the system.

- **URL**: `/trx/{id}`
- **Method**: `DELETE`
- **Authentication**: Required

## Response Codes

- `200 OK`: Request successful
- `201 Created`: Transaction created successfully
- `400 Bad Request`: Invalid request parameters
- `401 Unauthorized`: Authentication required
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: Transaction not found
- `500 Internal Server Error`: Server error

## Notes

- All monetary values are in Indonesian Rupiah (IDR)
- Transaction IDs are unique and auto-generated
- Method of payment options include "BANK_TRANSFER" and others
- Deleted transactions cannot be recovered
- Transactions are linked to user accounts and delivery addresses
- All timestamps are in ISO 8601 format
