package middlewares

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func CreateRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis-17808.c252.ap-southeast-1-1.ec2.cloud.redislabs.com:17808", // Alamat server Redis
		Password: "JlbeF8p9TQuOjW0bSRQwykPzBCEZ0h1A",                                // Kata sandi (kosong jika tidak ada)
		DB:       0,                                                                 // Nomor database Redis (default adalah 0)
	})

	// Tes koneksi ke Redis
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return nil
	}
	fmt.Println("Connected to Redis:", pong)
	return client
}
