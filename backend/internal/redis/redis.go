package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"scarlet/internal/http/models"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func client() *redis.Client {

	opt, err := redis.ParseURL(fmt.Sprintf("redis://%s:%s@localhost:%s", os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("PORT")))

	if err != nil {
		log.Println("[Redis - Connect] -> ", err)
		return redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	}

	return redis.NewClient(opt)
}

func Set(key string, data any) {
	log.Printf("[CACHE - SET] -> %s \n", key)

	parser, _ := json.Marshal(data)
	duration := time.Minute * time.Duration(30)

	client().Set(ctx, key, parser, duration)
}

func Get[T any](key string) (models.BSResponse[T], error) {
	log.Printf("[CACHE - GET] -> %s \n", key)

	result, err := client().Get(ctx, key).Result()

	if err != nil {
		return models.BSResponse[T]{}, errors.New("value don't exists")
	}

	var data models.BSResponse[T]
	err = json.Unmarshal([]byte(result), &data)

	if err != nil {
		log.Println("[Redis - Unmarshal] -> ", err.Error())
		return models.BSResponse[T]{}, errors.New("internal.server.error")
	}

	return data, nil
}
