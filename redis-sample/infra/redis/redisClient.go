package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
)

// 전역 변수
var (
	rdb  *redis.Client
	once sync.Once
)

// context 를 호출하는 쪽에서 파라미터로 넣어주는 걸로 변경
func NewRedisClient(ctx context.Context, addr string, password string, db int) error {

	var connErr error

	// Redis Server Connection
	once.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		})

		// 연결 테스트
		if err := rdb.Ping(ctx).Err(); err != nil {
			connErr = fmt.Errorf("rdeis 연결이 실패 하였습니다. %v", err)
			return
		}
		log.Println("Redis 연결이 성공하였습니다.")
	})

	return connErr
}

func GetRedisClient() (*redis.Client, error) {
	if rdb == nil {
		log.Fatal("Redis 클라이언트가 초기화 되지 않았습니다.")
	}
	return rdb, nil
}
