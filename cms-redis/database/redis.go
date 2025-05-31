// database/redis.go - Kết nối Redis và tạo index Redisearch

package database

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     getEnv("REDIS_ADDR", "10.73.48.243:6379"),
		Password: getEnv("REDIS_PASSWORD", "Tuan123"),
		DB:       0,
	})
	if err := Rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Không thể kết nối Redis: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// Tạo index Redisearch cho Location (các model khác tương tự)
func EnsureLocationIndex() error {
	_, err := Rdb.Do(context.Background(), "FT.CREATE", "idx:location",
		"ON", "JSON",
		"PREFIX", "1", "location:",
		"SCHEMA",
		"$.name", "AS", "name", "TEXT",
		"$.address", "AS", "address", "TEXT",
		"$.status", "AS", "status", "NUMERIC",
	).Result()
	if err != nil && !isIndexExistsError(err) {
		return err
	}
	return nil
}

func isIndexExistsError(err error) bool {
	return err != nil && // Redisearch trả lỗi nếu index đã tồn tại
		// go-redis v9 trả lỗi dạng string
		(err.Error() == "Index already exists" || err.Error() == "ERR index already exists")
}
