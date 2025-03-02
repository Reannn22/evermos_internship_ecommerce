# Users API Documentation

## Overview

The Users API provides comprehensive functionality to manage user accounts within the system. This API allows clients to register, authenticate, update, and manage user profiles and credentials across the platform.

## Base URL

```
http://localhost:3000/api/v1
```

## Authentication

Most endpoints require Bearer token authentication. Include the token in the Authorization header:

```
Authorization: Bearer <your_token>
```

## Endpoints

### 1. Register User

Creates a new user account.

- **URL**: `/auth/register`
- **Method**: `POST`
- **Content-Type**: `application/json`

**Request Body**:

```json
{
    "id": 1,
    "nama": "string",
    "kata_sandi": "string",
    "no_telp": "string",
    "tanggal_lahir": "string",
    "jenis_kelamin": "string",
    "tentang": "string",
    "pekerjaan": "string",
    "email": "string",
    "id_provinsi": "string",
    "id_kota": "string",
    "is_admin": boolean
}
```

### 2. Login User

Authenticates a user and returns a session token.

- **URL**: `/auth/login`
- **Method**: `POST`
- **Content-Type**: `application/json`

**Request Body**:

```json
{
  "email": "string",
  "kata_sandi": "string"
}
```

### 3. Logout User

Ends the current user session.

- **URL**: `/auth/logout`
- **Method**: `POST`

### 4. Change User Password

Updates the user's password.

- **URL**: `/user/{id}/password`
- **Method**: `PATCH`
- **Authentication**: Required

**Request Body**:

```json
{
  "old_password": "string",
  "new_password": "string"
}
```

### 5. Forgot Password

Initiates password recovery process.

- **URL**: `/user/forgot-password`
- **Method**: `POST`
- **Content-Type**: `application/json`

**Request Body**:

```json
{
  "email": "string"
}
```

### 6. Reset Password

Completes the password reset process.

- **URL**: `/user/reset-password`
- **Method**: `POST`
- **Content-Type**: `application/json`

**Request Body**:

```json
{
  "reset_token": "string",
  "new_password": "string"
}
```

### 7. Get Specific User

Retrieves details of a specific user.

- **URL**: `/user`
- **Method**: `GET`
- **Authentication**: Required

### 8. Get All Users

Retrieves a list of all users.

- **URL**: `/user`
- **Method**: `GET`
- **Authentication**: Required

### 9. Update User

Updates user information.

- **URL**: `/user`
- **Method**: `PUT`
- **Authentication**: Required
- **Content-Type**: `application/json`

**Request Body**:

```json
{
    "id": 1,
    "nama": "string",
    "kata_sandi": "string",
    "no_telp": "string",
    "tanggal_lahir": "string",
    "jenis_kelamin": "string",
    "tentang": "string",
    "pekerjaan": "string",
    "email": "string",
    "id_provinsi": "string",
    "id_kota": "string",
    "is_admin": boolean
}
```

### 10. Delete User

Removes a user from the system.

- **URL**: `/user`
- **Method**: `DELETE`
- **Authentication**: Required

## Response Codes

- `200 OK`: Request successful
- `201 Created`: Resource created successfully
- `400 Bad Request`: Invalid request parameters
- `401 Unauthorized`: Authentication required
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: Resource not found
- `500 Internal Server Error`: Server error

## Notes

- All timestamps are in ISO 8601 format
- Authentication tokens should be kept secure
- Passwords must meet minimum security requirements
