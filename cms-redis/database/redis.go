// database/redis.go - Kết nối Redis và tạo index Redisearch

package database

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func InitRedis() {
	// Tối ưu cấu hình Redis để tăng tốc độ đọc key
	Rdb = redis.NewClient(&redis.Options{
		Addr:     getEnv("REDIS_ADDR", "10.73.48.243:6379"),
		Password: getEnv("REDIS_PASSWORD", "Tuan123"),
		DB:       0,

		// Connection Pool Configuration
		PoolSize:        getEnvInt("REDIS_POOL_SIZE", 100),              // Số connection tối đa trong pool
		MinIdleConns:    getEnvInt("REDIS_MIN_IDLE_CONNS", 20),          // Số connection idle tối thiểu
		MaxIdleConns:    getEnvInt("REDIS_MAX_IDLE_CONNS", 50),          // Số connection idle tối đa
		ConnMaxLifetime: getDuration("REDIS_CONN_MAX_LIFETIME", "1h"),   // Thời gian sống của connection
		ConnMaxIdleTime: getDuration("REDIS_CONN_MAX_IDLE_TIME", "30m"), // Thời gian idle tối đa

		// Timeout Configuration cho tối ưu read operations
		DialTimeout:  getDuration("REDIS_DIAL_TIMEOUT", "5s"),  // Timeout khi connect
		ReadTimeout:  getDuration("REDIS_READ_TIMEOUT", "3s"),  // Timeout khi đọc
		WriteTimeout: getDuration("REDIS_WRITE_TIMEOUT", "3s"), // Timeout khi ghi
		PoolTimeout:  getDuration("REDIS_POOL_TIMEOUT", "4s"),  // Timeout khi chờ connection từ pool

		// Retry Configuration
		MaxRetries:      getEnvInt("REDIS_MAX_RETRIES", 3), // Số lần retry
		MinRetryBackoff: getDuration("REDIS_MIN_RETRY_BACKOFF", "8ms"),
		MaxRetryBackoff: getDuration("REDIS_MAX_RETRY_BACKOFF", "512ms"),
	})

	if err := Rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Không thể kết nối Redis: %v", err)
	}

	log.Printf("Redis connected successfully!")
	log.Printf("Connection pool configured: PoolSize=%d, MinIdle=%d, MaxIdle=%d",
		Rdb.Options().PoolSize, Rdb.Options().MinIdleConns, Rdb.Options().MaxIdleConns)
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// Helper function để đọc environment variable kiểu int
func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return fallback
}

// Helper function để đọc environment variable kiểu duration
func getDuration(key, fallback string) time.Duration {
	if v := os.Getenv(key); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
	}
	if d, err := time.ParseDuration(fallback); err == nil {
		return d
	}
	return 5 * time.Second // fallback an toàn
}

// Tạo context với timeout để tối ưu Redis operations
func GetRedisContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

// Hàm helper để đọc nhiều keys cùng lúc với pipeline (tối ưu performance)
func GetMultipleKeys(keys []string) ([]interface{}, error) {
	if len(keys) == 0 {
		return nil, nil
	}

	ctx, cancel := GetRedisContext()
	defer cancel()

	// Sử dụng pipeline để giảm round-trips
	pipe := Rdb.Pipeline()
	cmds := make([]*redis.Cmd, len(keys))

	for i, key := range keys {
		cmds[i] = pipe.Do(ctx, "JSON.GET", key)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}

	results := make([]interface{}, len(keys))
	for i, cmd := range cmds {
		results[i], _ = cmd.Result()
	}

	return results, nil
}

// Tạo index Redisearch cho Location (các model khác tương tự)
func EnsureLocationIndex() error {
	ctx, cancel := GetRedisContext()
	defer cancel()

	_, err := Rdb.Do(ctx, "FT.CREATE", "idx:location",
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
