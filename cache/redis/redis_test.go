package redis

import (
	"context"
	"fmt"
	"testing"
)

func TestInitGlobalRedisPool(t *testing.T) {

	dsn := "redis://localhost:63790/0"

	pool, err := InitGlobalRedisPool(dsn)
	if err != nil {
		t.Fail()
	}
	reply, err := pool.Get().Do("SET", "test", "test")
	if err != nil {
		t.Fail()
		fmt.Println(err)
	} else {
		fmt.Println(reply)
	}

	client, err := InitGlobalRedisClient(dsn)
	if err != nil {
		t.Fail()
	}
	if v, err := client.Get(context.Background(), "test").Result(); v != "test" || err != nil {
		t.Fail()
	}

}
