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

**Endpoint:** `POST /signup`  

**Request Body:**
```json
{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "securePassword123"
}