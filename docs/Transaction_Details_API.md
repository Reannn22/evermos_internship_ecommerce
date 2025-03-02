# Transaction Details API Documentation

## Overview

The Transaction Details API provides comprehensive functionality to manage detailed information about transactions within the system. This API allows clients to create, retrieve, update, and delete transaction details, enabling thorough tracking and management of transaction data.

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

### 1. Create Transaction Detail

Creates a new transaction detail record.

- **URL**: `/detail-trx`
- **Method**: `POST`
- **Authentication**: Required
- **Content-Type**: `application/json`

**Request Body**:

```json
{
    "trx_id": integer,
    "log_product_id": integer,
    "store_id": integer,
    "kuantitas": integer,
    "harga_total": integer
}
```

### 2. Get Specific Transaction Detail

Retrieves specific transaction detail information.

- **URL**: `/transaction-details/{id}`
- **Method**: `GET`
- **Authentication**: Required

### 3. Get All Transaction Details

Retrieves a list of all transaction details.

- **URL**: `/transaction-details`
- **Method**: `GET`
- **Authentication**: Required

### 4. Update Transaction Detail

Updates transaction detail information.

- **URL**: `/transaction-details/{id}`
- **Method**: `PUT`
- **Authentication**: Required
- **Content-Type**: `application/json`

**Request Body**:

```json
{
    "TrxID": integer,
    "LogProductID": integer,
    "StoreID": integer,
    "Kuantitas": integer,
    "HargaTotal": integer
}
```

### 5. Delete Transaction Detail

Removes a transaction detail from the system.

- **URL**: `/detail-trx/{id}`
- **Method**: `DELETE`
- **Authentication**: Required

## Response Codes

- `200 OK`: Request successful
- `201 Created`: Transaction detail created successfully
- `400 Bad Request`: Invalid request parameters
- `401 Unauthorized`: Authentication required
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: Transaction detail not found
- `500 Internal Server Error`: Server error

## Notes

- All monetary values are in Indonesian Rupiah (IDR)
- Transaction detail IDs are unique and auto-generated
- Each transaction detail is linked to a main transaction record
- All timestamps are in ISO 8601 format
- Deleted transaction details cannot be recovered
- The system maintains referential integrity with related transactions
- Updates to transaction details may affect overall transaction totals
