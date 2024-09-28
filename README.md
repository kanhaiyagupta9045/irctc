# Railway Management System API

## Overview

This API provides a platform for users to check train availability, book seats, and manage their bookings in a railway management system similar to IRCTC. It supports two types of users: Admins and Regular Users.

## Table of Contents

- [Base URL](#base-url)
- [Authentication](#authentication)
- [Endpoints](#endpoints)
  - [Add Train](#add-train)
  - [Get Seat Availability](#get-seat-availability)
  - [Book Seat](#book-seat)
  - [Get Booking Details](#get-booking-details)
- [Error Handling](#error-handling)
- [Running the API](#running-the-api)

## Base URL
http://localhost:3000


## Authentication

All endpoints require a valid JWT token for authentication. You can obtain the token by logging in as a user.

## Endpoints

### 1. Add Train

**Endpoint:** `POST /add/train`  
**Authorization:** Admin required  

**Request Body:**
```json
{
  "train_number": "12345",
  "source": "Mumbai",
  "destination": "Delhi",
  "total_seats": 100,
  "available_seats": 100
}
```
