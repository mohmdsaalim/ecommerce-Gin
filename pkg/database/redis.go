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

var RedisClient *redis.Client
var Ctx = context.Background()

// ConnectRedis initializes Redis client using config.yaml values
func ConnectRedis() {
	// Parse port if it's string
	port, err := strconv.Atoi(config.AppConfig.Redis.Port)
	if err != nil {
		log.Fatalf("[error] invalid Redis port in config.yaml: %v", err)
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", config.AppConfig.Redis.Host, port),
		Password:     config.AppConfig.Redis.Password, // leave empty if no password
		DB:           config.AppConfig.Redis.DB,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		PoolSize:     10,
		MinIdleConns: 2,
	})

	if _, err := RedisClient.Ping(Ctx).Result(); err != nil {
		log.Fatalf("[error] failed to connect to Redis: %v", err)
	} else {
		log.Println("✅ Redis connected successfully")
	}
}

// SetOTP sets an OTP in Redis with TTL
func SetOTP(key string, otp string, ttl time.Duration) error {
	return RedisClient.Set(Ctx, key, otp, ttl).Err()
}

// GetOTP retrieves an OTP from Redis
func GetOTP(key string) (string, error) {
	return RedisClient.Get(Ctx, key).Result()
}

// DeleteOTP deletes an OTP from Redis
func DeleteOTP(key string) error {
	return RedisClient.Del(Ctx, key).Err()
}