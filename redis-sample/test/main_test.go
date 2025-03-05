package test

import (
	"context"
	"fmt"
	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func init() {
	fmt.Println("(1) init 함수 (테스트 시작)")
}

func setup() {
	fmt.Println("(2) Setup 함수 (Redis Client Setting)")
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	rdb = redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
}

func teardown() {
	fmt.Println("(3) teardown 함수")
	rdb.FlushDB(ctx)
	err := rdb.Close()
	if err != nil {
		panic(err)
	}
	return
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run() //
	teardown()
	os.Exit(code)
}

// Set, Get
func Test_Set_Get(t *testing.T) {
	fmt.Println("exec1 테스트")

	err := rdb.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		t.Error("저장 실패")
	}

	val, err := rdb.Get(ctx, "foo").Result()
	if err != nil || val != "bar" {
		t.Error("잘못된 값입니다.")
	}

	assert.Equal(t, "bar", val)

	rdb.FlushDB(ctx)
}

// RPush, RPop foo, [a,b,c,d,c]
func Test_RPush(t *testing.T) {
	fmt.Println("exec2 테스트")
	err := rdb.RPush(ctx, "foo", "bar1", "bar2", "bar3").Err()

	if err != nil {
		t.Error("저장 실패")
	}
	//TODO: assert error 체크

	val1, err := rdb.RPop(ctx, "foo").Result()

	if err != nil || val1 != "bar3" {
		t.Error("잘못된 값입니다.")
	}

	assert.Equal(t, "bar3", val1)

	val2, err := rdb.RPop(ctx, "foo").Result()
	if err != nil || val2 != "bar2" {
		t.Error("잘못된 값입니다.")
	}

	assert.Equal(t, "bar2", val2)

	rdb.FlushDB(ctx)
}
