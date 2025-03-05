package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"redis-sample/infra/redis"
	"redis-sample/internal/model"
	"redis-sample/internal/repository"
)

func main() {

	personID := uuid.NewString()

	jsonString, err := json.Marshal(model.Person{
		ID:   personID, // uuid ...
		Name: "go",
		Age:  22,
	})

	if err != nil {
		fmt.Printf("Failed to marshal JSON: %s", err.Error())
		return
	}

	ctx := context.Background()

	if err := redis.NewRedisClient(ctx, "localhost:6379", "", 0); err != nil {
		log.Fatalf("Redis 연결 실패: %v", err)
	}
	client, err := redis.GetRedisClient()
	if err != nil {
		log.Fatalf("Redis 클라이언트 가져오기 실패: %v", err)
	}

	// set
	if err := repository.SetKeyRedis(ctx, client, personID, jsonString, 0); err != nil {
		log.Printf("데이터 저장(set)에 실패하였습니다. %v\n", err)
	}

	// get
	val, err := repository.GetKeyRedis(ctx, client, personID)

	if err != nil {
		log.Printf("데이터 조회(get)에 실패하였습니다. %v\n\n", err)
	}
	log.Printf("조회된 값은: %s", val)
}
