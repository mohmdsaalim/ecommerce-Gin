# OTP Redis Implementation - Complete Guide

## Overview
This document explains how OTP (One-Time Password) is stored in Redis, verified, and how it automatically expires after 5 minutes.

## Configuration (config.yaml)

```yaml
redis:
  host: localhost      # Redis server address
  port: 6379          # Redis server port
  password: ""        # Empty if no password is set
  db: 0               # Database number (0 is default)
```

The application reads these values from `config/config.yaml` to connect to Redis.

---

## Complete OTP Flow

### 1️⃣ **User Registration (OTP Generation)**

**File:** `internal/services/auth.service.go` → `Register()` function

**What happens:**
1. User registers with username, email, and password
2. System generates a random 6-digit OTP (example: "123456")
3. Creates a unique Redis key: `otp:user@email.com`
4. **Stores OTP in Redis with 5-minute expiration**
5. Sends OTP to user's email
6. **After 5 minutes, Redis automatically deletes the OTP**

**Code Flow:**
```go
// Generate 6-digit OTP
otp, _ := utils.GenerateOTP(6)

// Create Redis key
key := fmt.Sprintf("otp:%s", user.Email) // Example: "otp:john@example.com"

// Store in Redis for 5 minutes
database.SetOTP(key, otp, 5*time.Minute)

// Send OTP via email
workers.OTPChannel <- workers.OTPJob{
    Email: user.Email,
    Code:  otp,
}
```

---

### 2️⃣ **Resend OTP (If User Didn't Receive)**

**Endpoint:** `POST /request-email-otp/:userId`
**File:** `internal/controllers/auth.controller.go` → `RequestEmailOTP()`

**What happens:**
1. User requests a new OTP
2. System generates a NEW 6-digit OTP
3. **Overwrites the old OTP in Redis** with the new one
4. New OTP expires after 5 minutes
5. Sends new OTP to user's email

**Code Flow:**
```go
// Generate new OTP
otp, err := utils.GenerateOTP(6)

// Create same Redis key format
key := fmt.Sprintf("otp:%s", user.Email)

// Store new OTP (overwrites old one)
database.SetOTP(key, otp, 5*time.Minute)

// Send new OTP via email
workers.OTPChannel <- workers.OTPJob{
    Email: user.Email,
    Code:  otp,
}
```

---

### 3️⃣ **OTP Verification**

**Endpoint:** `POST /verify-email-otp/:userId`
**Request Body:** `{"otp": "123456"}`
**File:** `internal/services/auth.service.go` → `VerifyEmailOTP()`

**What happens:**
1. User submits the OTP code
2. System retrieves OTP from Redis using user's email
3. **Compares Redis OTP with user's input**
4. If they match:
   - ✅ Sets `User.EmailVerified = true` in database
   - ✅ Deletes OTP from Redis (prevents reuse)
5. If they don't match:
   - ❌ Returns "invalid OTP" error
6. If OTP expired (5 minutes passed):
   - ❌ Returns "OTP expired or invalid" error

**Code Flow:**
```go
// Create Redis key
key := fmt.Sprintf("otp:%s", user.Email)

// Get OTP from Redis
storedOTP, err := database.GetOTP(key)
if err != nil {
    // OTP not found or expired
    return errors.New("OTP expired or invalid")
}

// Compare OTPs
if storedOTP != code {
    // Wrong OTP
    return errors.New("invalid OTP")
}

// OTP is correct! Update user
s.repo.UpdateFields(&user, userID, map[string]interface{}{
    "email_verified": true,  // Set EmailVerified to true
})

// Delete OTP from Redis (security)
database.DeleteOTP(key)
```

---

## Redis Helper Functions

**File:** `pkg/database/redis.go`

### 1. `SetOTP(key, otp, ttl)`
**Purpose:** Store OTP in Redis with automatic expiration

```go
// Store OTP for 5 minutes
database.SetOTP("otp:john@example.com", "123456", 5*time.Minute)
```

**What it does:**
- Saves OTP in Redis
- Sets TTL (Time To Live) = 5 minutes
- **After 5 minutes, Redis automatically deletes the OTP**

---

### 2. `GetOTP(key)`
**Purpose:** Retrieve OTP from Redis

```go
// Get OTP from Redis
otp, err := database.GetOTP("otp:john@example.com")
if err != nil {
    // OTP not found or expired
}
```

**Returns:**
- OTP code if found
- Error if OTP expired or doesn't exist

---

### 3. `DeleteOTP(key)`
**Purpose:** Immediately delete OTP from Redis

```go
// Delete OTP after successful verification
database.DeleteOTP("otp:john@example.com")
```

**Why we delete:**
- Prevents OTP from being used again (security)
- User can't verify the same OTP twice

---

## Security Features

### ✅ **Automatic Expiration (5 Minutes)**
- OTP is stored in Redis with `TTL = 5 minutes`
- Redis automatically deletes OTP after 5 minutes
- User cannot use expired OTP

### ✅ **One-Time Use**
- OTP is deleted from Redis after successful verification
- User cannot reuse the same OTP

### ✅ **Unique OTP per Email**
- Each user email has a unique Redis key
- Format: `otp:user@email.com`
- No conflicts between different users

---

## Testing the Flow

### 1. Register a User
```bash
POST /register
{
    "username": "john",
    "email": "john@example.com",
    "password": "password123"
}
```
- OTP is generated and stored in Redis
- OTP sent to email
- OTP expires after 5 minutes

### 2. Verify Email
```bash
POST /verify-email-otp/123
{
    "otp": "123456"
}
```
- If OTP matches → `EmailVerified = true`
- If OTP wrong → Error
- If OTP expired → Error

### 3. Resend OTP
```bash
POST /request-email-otp/123
```
- New OTP generated
- Old OTP is overwritten
- New OTP sent to email

---

## Common Errors

| Error Message | Reason | Solution |
|--------------|--------|----------|
| "OTP expired or invalid" | 5 minutes passed OR OTP not found | Request new OTP |
| "invalid OTP" | User entered wrong code | Check email and try again |
| "user not found" | Invalid userId in URL | Check userId parameter |

---

## Summary

✅ **OTP is generated** when user registers or requests resend  
✅ **OTP is stored in Redis** with 5-minute auto-expiration  
✅ **Redis config** is read from `config.yaml`  
✅ **OTP is verified** by comparing Redis OTP with user input  
✅ **User.EmailVerified is set to true** when OTP matches  
✅ **OTP is deleted** after successful verification  
✅ **Redis automatically deletes OTP** after 5 minutes  

All code includes beginner-friendly comments explaining each step! 🎉
