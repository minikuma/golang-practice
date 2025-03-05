package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// SetKeyRedis /** Set 함수: 외부에서 접근 가능하도록 대문자로 시작 */

func SetKeyRedis(cxt context.Context, client *redis.Client, id string, v []byte, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(cxt, 2*time.Second)
	defer cancel()

	err := client.Set(ctx, id, v, expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set key: %w", err)
	}
	return nil
}

// GetKeyRedis /** Get  함수: 외부에서 접근 가능하도록 대문자로 시작 */
func GetKeyRedis(cxt context.Context, client *redis.Client, id string) (string, error) {
	ctx, cancel := context.WithTimeout(cxt, 1*time.Second)
	defer cancel()

	val, err := client.Get(ctx, id).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get key: %w", err)
	}
	return val, nil
}
