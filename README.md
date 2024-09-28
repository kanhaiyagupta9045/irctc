# Railway Management System API

## Overview

This API provides a platform for users to check train availability, book seats, and manage their bookings in a railway management system similar to IRCTC. It supports two types of users: Admins and Regular Users.

## Table of Contents

- [Base URL](#base-url)
- [Authentication](#authentication)
- [Endpoints](#endpoints)
  - [User Sign-Up](#user-sign-up)
  - [User Sign-In](#user-sign-in)
  - [Add Train](#add-train)
  - [Get Seat Availability](#get-seat-availability)
  - [Book Seat](#book-seat)
  - [Get Booking Details](#get-booking-details)
- [Error Handling](#error-handling)
- [Running the API](#running-the-api)

## Base URL
http://localhost:3000

### 1. User Sign-Up

**Endpoint:** `POST{{baseurl}}/user/signup`  

**Request Body:**
```json
{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "securePassword123"
}
```

**Response Body:**
```json
{
  "message": "User registered successfully"
}
```

### 2. User Sign-In

**Endpoint:** `POST {{baseurl}}/user/signin`  

**Request Body:**
```json
{
  "email":"kanhaiya@gmail.com",
  "password":"********",
  "username":"kanhaiyagupta9045",
  "usertype" :"ADMIN"
}
```

**Response Body:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImthbmhhaXlhcmF1bml5YXJAZ21haWwuY29tIiwiZXhwIjoxNzI3NTg4MTE2fQ.uBoGPnJyZfOhdvf1irqyG6G9GXRxaJaSBH7L3V07YEc"
}
```


### 3. ADD Train (User should be ADMIN)

**Endpoint:** `POST{{baseurl}}/add/train`  

**Request Body:**
```json

  {
  "train_number": "12346",
  "source": "Mumbai",
  "destination": "Delhi",
  "total_seats": 100,
  "available_seats": 100
}
```

**Response Body:**
```json
{
  "message": "Train added successfully"
}
```


### 4. Book Train Ticket 

**Endpoint:** `POST{{baseurl}}/book/seat`  

**Request Body:**
```json

{
  "train_id" : "12346"
}
```

**Response Body:**
```json
{
  "booking": {
    "train_number": "12346",
    "status": "booked",
    "SeatNumber": 5
  },
  "message": "Train booked successfully"
}
```

### 5. GET Train Ticket Availabilty

**Endpoint:** `GET{{baseurl}}seat-availability?src=Mumbai&dst=Delhi`  


**Response Body:**
```json
[
  {
    "ID": 1,
    "CreatedAt": "2024-09-28T10:21:42.462942+05:30",
    "UpdatedAt": "2024-09-28T10:25:30.940654+05:30",
    "DeletedAt": null,
    "train_number": "12346",
    "source": "Mumbai",
    "destination": "Delhi",
    "total_seats": 100,
    "available_seats": 96
  }
]
```


### 6. Booking Details 

**Endpoint:** `GET{{baseurl}}/add/train`  

**Request Body:**
```json

  {
  "Authorization" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImthbmhhaXlhcmF1bml5YXI5MDQ1QGdtYWlsLmNvbSIsImV4cCI6MTcyNzU3OTIzNn0.x1ntuBZd7xqPEfuoWJcWghA36hJvlqOfH2yjelRTHfk"
}
```

**Response Body:**
```json
{
  "bookings": [
    {
      "ID": 4,
      "CreatedAt": "2024-09-28T10:25:30.940654+05:30",
      "UpdatedAt": "2024-09-28T10:25:30.941439+05:30",
      "DeletedAt": null,
      "user_id": 1,
      "train_id": "12346",
      "status": "booked",
      "seat_number": 4
    }
  ]
}
```




