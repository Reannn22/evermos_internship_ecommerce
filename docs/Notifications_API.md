# Notifications API Documentation

## Overview

The Notifications API provides endpoints to manage user notifications within the application. This API allows for creating, retrieving, updating, and deleting notification messages sent to users. Notifications keep users informed about order status changes, promotions, system announcements, and other important events.

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

### 1. Create Notification

Creates a new notification message.

- **URL**: `/notifications`
- **Method**: `POST`
- **Authentication**: Required
- **Content-Type**: `application/json`

**Request Body**:

```json
{
  "pesan": "string"
}
```

### 2. Get Specific Notification

Retrieves details of a specific notification.

- **URL**: `/notifications/{id}`
- **Method**: `GET`
- **Authentication**: Required

### 3. Get All Notifications

Retrieves all notifications for the authenticated user.

- **URL**: `/notifications`
- **Method**: `GET`
- **Authentication**: Required
- **Query Parameters**:
  - unread: boolean (optional)
  - page: integer (optional)
  - limit: integer (optional)

### 4. Update Notification

Updates an existing notification.

- **URL**: `/notifications/{id}`
- **Method**: `PUT`
- **Authentication**: Required
- **Content-Type**: `application/json`

### 5. Delete Notification

Removes a notification from the system.

- **URL**: `/notifications/{id}`
- **Method**: `DELETE`
- **Authentication**: Required

## Response Codes

- `200 OK`: Request successful
- `201 Created`: Notification created successfully
- `400 Bad Request`: Invalid request parameters
- `401 Unauthorized`: Authentication required
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: Notification not found
- `500 Internal Server Error`: Server error

## Notes

- Notification IDs are unique and auto-generated
- Read/unread status is tracked per notification
- Notifications can be targeted to specific users
- Supports both system and administrative messages
- All timestamps are in ISO 8601 format
- Notifications can include action links
- Deletion removes from user's view only
- System retains notification history
- Push notifications may be triggered
- Supports notification categories and priorities
