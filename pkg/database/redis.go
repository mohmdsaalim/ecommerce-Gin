package database

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/mohmdsaalim/ecommerce-Gin/config"
	"github.com/redis/go-redis/v9"
)

// RedisClient is the global Redis connection that we use throughout the application
var RedisClient *redis.Client

// Ctx is the context used for Redis operations (required by Redis library)
var Ctx = context.Background()

// ConnectRedis initializes Redis client using values from config.yaml file
// This function reads Redis host, port, password, and database number from config.yaml
// and establishes a connection to the Redis server
func ConnectRedis() {
	// Step 1: Convert port from string to integer (config.yaml stores it as string)
	port, err := strconv.Atoi(config.AppConfig.Redis.Port)
	if err != nil {
		log.Fatalf("[error] invalid Redis port in config.yaml: %v", err)
	}

	// Step 2: Create a new Redis client with configuration from config.yaml
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", config.AppConfig.Redis.Host, port), // Redis server address (localhost:6379)
		Password:     config.AppConfig.Redis.Password,                         // Password (empty string if no password)
		DB:           config.AppConfig.Redis.DB,                               // Database number (0 by default)
		DialTimeout:  5 * time.Second,                                         // Maximum time to wait for connection
		ReadTimeout:  5 * time.Second,                                         // Maximum time to wait for read operations
		WriteTimeout: 5 * time.Second,                                         // Maximum time to wait for write operations
		PoolSize:     10,                                                      // Number of connections to maintain
		MinIdleConns: 2,                                                       // Minimum idle connections to keep open
	})

	// Step 3: Test the connection by sending a PING command to Redis
	if _, err := RedisClient.Ping(Ctx).Result(); err != nil {
		log.Fatalf("[error] failed to connect to Redis: %v", err)
	} else {
		log.Println("✅ Redis connected successfully")
	}
}

// SetOTP stores an OTP in Redis with automatic expiration (TTL = Time To Live)
// Parameters:
//   - key: unique identifier for the OTP (example: "otp:user@email.com")
//   - otp: the 6-digit OTP code to store
//   - ttl: how long to keep the OTP in Redis (example: 5 minutes)
//
// After TTL expires (5 minutes), Redis will automatically delete the OTP
// Returns: error if the operation fails
func SetOTP(key string, otp string, ttl time.Duration) error {
	// Store OTP in Redis with automatic expiration after 5 minutes
	return RedisClient.Set(Ctx, key, otp, ttl).Err()
}

// GetOTP retrieves an OTP from Redis using the key
// Parameters:
//   - key: unique identifier for the OTP (example: "otp:user@email.com")
//
// Returns:
//   - string: the OTP code if found
//   - error: if OTP not found or expired (Redis returns "redis: nil" error)
func GetOTP(key string) (string, error) {
	// Get OTP from Redis using the key
	return RedisClient.Get(Ctx, key).Result()
}

// DeleteOTP removes an OTP from Redis immediately
// This is called after successful verification to prevent OTP reuse
// Parameters:
//   - key: unique identifier for the OTP (example: "otp:user@email.com")
//
// Returns: error if the deletion fails
func DeleteOTP(key string) error {
	// Delete OTP from Redis immediately
	return RedisClient.Del(Ctx, key).Err()
}
