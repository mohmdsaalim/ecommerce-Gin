# Testing OTP Functionality - Step by Step Guide

This guide will help you test the OTP functionality with Redis.

## Prerequisites

1. **Start Redis Server**
   ```bash
   # On Mac with Homebrew
   brew services start redis
   
   # Or run Redis directly
   redis-server
   
   # Check if Redis is running
   redis-cli ping
   # Should respond: PONG
   ```

2. **Verify config.yaml**
   Make sure Redis configuration is correct in `config/config.yaml`:

3. **Start the Application**
   ```bash
   cd /Users/saalim/Desktop/ecommerce-Gin
   go run cmd/api/main.go
   ```

---

## Test 1: Register User and Generate OTP

### Request
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123"
  }'
```

### Expected Response
```json
{
  "message": "user registered successfully"
}
```

### What Happened Behind the Scenes
1. ✅ User created in database with `EmailVerified = false`
2. ✅ 6-digit OTP generated (e.g., "123456")
3. ✅ OTP stored in Redis with key: `otp:test@example.com`
4. ✅ OTP will expire in 5 minutes
5. ✅ OTP sent to email

### Verify in Redis
```bash
# Connect to Redis CLI
redis-cli

# Check if OTP exists
GET otp:test@example.com
# Should return the 6-digit OTP

# Check TTL (time remaining)
TTL otp:test@example.com
# Should return seconds remaining (up to 300 = 5 minutes)

# Exit Redis CLI
exit
```

---

## Test 2: Verify OTP (Correct OTP)

### Get User ID First
You need the user ID. Check your email or database for the user ID.

Let's assume the user ID is **1** for this example.

### Request
```bash
curl -X POST http://localhost:8080/verify-email-otp/1 \
  -H "Content-Type: application/json" \
  -d '{
    "otp": "123456"
  }'
```
*(Replace "123456" with the actual OTP from Redis or email)*

### Expected Response
```json
{
  "message": "Email verified successfully"
}
```

### What Happened Behind the Scenes
1. ✅ OTP retrieved from Redis: `otp:test@example.com`
2. ✅ Redis OTP compared with user input
3. ✅ OTPs matched!
4. ✅ User.EmailVerified set to true in database
5. ✅ OTP deleted from Redis

### Verify in Database
```sql
SELECT id, username, email, email_verified FROM users WHERE email = 'test@example.com';
-- email_verified should be TRUE
```

### Verify in Redis
```bash
redis-cli GET otp:test@example.com
# Should return: (nil) - OTP was deleted after verification
```

---

## Test 3: Verify OTP (Wrong OTP)

### Request
```bash
curl -X POST http://localhost:8080/verify-email-otp/1 \
  -H "Content-Type: application/json" \
  -d '{
    "otp": "999999"
  }'
```

### Expected Response
```json
{
  "error": "invalid OTP"
}
```

### What Happened
1. ✅ OTP retrieved from Redis
2. ❌ OTPs didn't match
3. ❌ User.EmailVerified remains false
4. ✅ OTP remains in Redis (not deleted)

---

## Test 4: Verify OTP (Expired OTP - After 5 Minutes)

### Request
Wait 5 minutes after registering, then:
```bash
curl -X POST http://localhost:8080/verify-email-otp/1 \
  -H "Content-Type: application/json" \
  -d '{
    "otp": "123456"
  }'
```

### Expected Response
```json
{
  "error": "OTP expired or invalid"
}
```

### What Happened
1. ✅ Tried to get OTP from Redis
2. ❌ Redis returned error (key not found - automatically deleted after 5 minutes)
3. ❌ User.EmailVerified remains false

---

## Test 5: Resend OTP

### Request
```bash
curl -X POST http://localhost:8080/request-email-otp/1
```

### Expected Response
```json
{
  "message": "OTP sent successfully"
}
```

### What Happened
1. ✅ New 6-digit OTP generated (e.g., "654321")
2. ✅ Old OTP in Redis overwritten with new OTP
3. ✅ New OTP expires in 5 minutes
4. ✅ New OTP sent to email

### Verify in Redis
```bash
redis-cli GET otp:test@example.com
# Should return the NEW OTP

redis-cli TTL otp:test@example.com
# Should return ~300 seconds (5 minutes)
```

---

## Test 6: Monitor Redis in Real-Time

### Watch OTP Expiration
```bash
# In one terminal, monitor Redis
redis-cli

# Watch TTL decrease
WATCH otp:test@example.com

# In another terminal, check TTL every few seconds
while true; do redis-cli TTL otp:test@example.com; sleep 1; done

# After 5 minutes (300 seconds), TTL will become -2 (key deleted)
```

---

## Debugging Tips

### 1. Check if Redis is Running
```bash
redis-cli ping
# Should return: PONG
```

### 2. List All OTP Keys in Redis
```bash
redis-cli KEYS "otp:*"
# Shows all OTP keys
```

### 3. Check OTP Value and TTL
```bash
redis-cli
GET otp:test@example.com    # Get OTP value
TTL otp:test@example.com    # Get time remaining in seconds
```

### 4. Manually Delete OTP (for testing)
```bash
redis-cli DEL otp:test@example.com
```

### 5. Check Application Logs
Look for these messages in your application logs:
- `✅ Redis connected successfully` - Redis connection successful
- `✅ Config loaded from config.yaml` - Config loaded

---

## Common Issues and Solutions

### Issue 1: "Error: failed to connect to Redis"
**Solution:**
```bash
# Make sure Redis is running
brew services start redis

# Or start manually
redis-server
```

### Issue 2: OTP Not Expiring
**Solution:**
- Check that you're using `5*time.Minute` in SetOTP call
- Verify TTL in Redis: `redis-cli TTL otp:test@example.com`

### Issue 3: "OTP expired or invalid" Immediately
**Solution:**
- OTP might not be stored in Redis
- Check Redis connection in application logs
- Verify `database.SetOTP()` is being called

---

## Complete Test Sequence

```bash
# 1. Start Redis
brew services start redis

# 2. Start Application
go run cmd/api/main.go

# 3. Register User
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username":"test","email":"test@example.com","password":"pass123"}'

# 4. Check OTP in Redis
redis-cli GET otp:test@example.com

# 5. Verify OTP (use the OTP from step 4)
curl -X POST http://localhost:8080/verify-email-otp/1 \
  -H "Content-Type: application/json" \
  -d '{"otp":"YOUR_OTP_HERE"}'

# 6. Verify OTP is deleted
redis-cli GET otp:test@example.com
# Should return: (nil)
```

---

## Success Checklist

- ✅ Redis is running and responding to PING
- ✅ Application connects to Redis successfully
- ✅ OTP is stored in Redis when user registers
- ✅ OTP has 5-minute TTL in Redis
- ✅ OTP is sent to user's email
- ✅ Correct OTP verification sets EmailVerified to true
- ✅ OTP is deleted from Redis after successful verification
- ✅ Wrong OTP returns error
- ✅ Expired OTP returns error
- ✅ Resend OTP generates new code

🎉 **All tests passing? Your OTP implementation is working perfectly!**
