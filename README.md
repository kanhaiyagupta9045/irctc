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


### 3. ADD Train

**Endpoint:** `POST{{baseurl}}/user/signup`  

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




