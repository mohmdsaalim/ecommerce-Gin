# Quick Reference - OTP Files and Flow

## 📁 File Structure

```
ecommerce-Gin/
│
├── config/
│   └── config.yaml                    # Redis config (host, port, password, db)
│
├── pkg/database/
│   └── redis.go                       # ✅ UPDATED: Redis connection + OTP functions
│                                      # - ConnectRedis() → Connects using config.yaml
│                                      # - SetOTP() → Store OTP with 5-min expiration
│                                      # - GetOTP() → Retrieve OTP from Redis
│                                      # - DeleteOTP() → Delete OTP after verification
│
├── internal/
│   ├── controllers/
│   │   └── auth.controller.go         # ✅ UPDATED: HTTP request handlers
│   │                                  # - RequestEmailOTP() → Handle resend OTP
│   │                                  # - VerifyEmailOTP() → Handle OTP verification
│   │
│   └── services/
│       └── auth.service.go            # ✅ UPDATED: Business logic
│                                      # - Register() → Create user + generate OTP
│                                      # - SendEmailOTP() → Resend new OTP
│                                      # - VerifyEmailOTP() → Verify + set EmailVerified
│
└── Documentation (NEW):
    ├── OTP_REDIS_IMPLEMENTATION.md    # 📄 Complete implementation guide
    ├── OTP_FLOW_DIAGRAM.md            # 📊 Visual flow diagram
    ├── OTP_TESTING_GUIDE.md           # 🧪 Testing instructions
    └── SUMMARY_OF_CHANGES.md          # 📝 What changed and why
```

---

## 🔄 OTP Flow (Quick Version)

```
1. USER REGISTERS
   ↓
2. GENERATE 6-DIGIT OTP
   ↓
3. STORE IN REDIS (5-min expiration)
   Key: "otp:user@email.com"
   Value: "123456"
   TTL: 300 seconds
   ↓
4. SEND OTP TO EMAIL
   ↓
5. USER ENTERS OTP
   ↓
6. GET OTP FROM REDIS
   ↓
7. COMPARE: Redis OTP vs User Input
   ↓
   ├─ ✅ MATCH
   │  ├─ Set EmailVerified = true
   │  ├─ Delete OTP from Redis
   │  └─ Return success
   │
   └─ ❌ NO MATCH
      └─ Return error
```

---

## 🔐 Redis Keys Format

| Purpose | Key Format | Example | TTL |
|---------|-----------|---------|-----|
| Store OTP | `otp:{email}` | `otp:john@example.com` | 5 minutes |

---

## 🚀 Quick Start

```bash
# 1. Start Redis
brew services start redis

# 2. Verify Redis is running
redis-cli ping
# Should return: PONG

# 3. Start your application
go run cmd/api/main.go

# 4. Check logs for:
# ✅ Config loaded from config.yaml
# ✅ Redis connected successfully
```

---

## 📝 Key Functions

### SetOTP(key, otp, ttl)
```go
database.SetOTP("otp:user@email.com", "123456", 5*time.Minute)
```
- Stores OTP in Redis
- Auto-deletes after 5 minutes
- Overwrites existing OTP if present

### GetOTP(key)
```go
otp, err := database.GetOTP("otp:user@email.com")
```
- Returns OTP if exists and not expired
- Returns error if expired or not found

### DeleteOTP(key)
```go
database.DeleteOTP("otp:user@email.com")
```
- Immediately deletes OTP
- Called after successful verification

---

## 🧪 Test Commands

```bash
# Check OTP in Redis
redis-cli GET otp:user@email.com

# Check time remaining
redis-cli TTL otp:user@email.com

# List all OTP keys
redis-cli KEYS "otp:*"

# Delete OTP manually (for testing)
redis-cli DEL otp:user@email.com
```

---

## ✅ What Works Now

| Feature | Status | Details |
|---------|--------|---------|
| OTP Generation | ✅ | 6-digit random OTP |
| Redis Storage | ✅ | Stored with 5-min TTL |
| Auto-Expiration | ✅ | Redis deletes after 5 min |
| Email Sending | ✅ | Via worker channel |
| OTP Verification | ✅ | Compares Redis vs input |
| EmailVerified Update | ✅ | Set to true on success |
| OTP Deletion | ✅ | Removed after verification |
| Config from YAML | ✅ | Redis settings from config.yaml |
| Beginner Comments | ✅ | All code explained |
| Documentation | ✅ | 4 guides created |

---

## 📚 Documentation Guide

| Document | What It Covers | When to Read |
|----------|---------------|--------------|
| **SUMMARY_OF_CHANGES.md** | What changed and why | Read first |
| **OTP_FLOW_DIAGRAM.md** | Visual flow with diagrams | Understanding flow |
| **OTP_REDIS_IMPLEMENTATION.md** | Complete technical guide | Deep dive |
| **OTP_TESTING_GUIDE.md** | How to test everything | Before testing |

---

## 🎯 Common Tasks

### Register User
```bash
POST /register
{
  "username": "john",
  "email": "john@example.com",
  "password": "password123"
}
```

### Verify OTP
```bash
POST /verify-email-otp/:userId
{
  "otp": "123456"
}
```

### Resend OTP
```bash
POST /request-email-otp/:userId
```

---

## 🔍 Debug Checklist

- [ ] Redis is running (`redis-cli ping`)
- [ ] Config.yaml has correct Redis settings
- [ ] Application shows "Redis connected successfully"
- [ ] OTP appears in Redis after registration
- [ ] TTL is set to ~300 seconds (5 minutes)
- [ ] OTP deleted after successful verification
- [ ] EmailVerified is true in database after verification

---

## 💡 Remember

1. **Redis auto-deletes OTP after 5 minutes** - No manual cleanup needed!
2. **OTP is deleted after successful verification** - Prevents reuse
3. **All settings come from config.yaml** - Easy to configure
4. **Simple code with lots of comments** - Perfect for beginners!

🎉 **Everything is ready to use!**
