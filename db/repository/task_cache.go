package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
	"web-svc/db/cache"
)

var redisClient = cache.NewClient()

type TaskCache struct{}

func (t TaskCache) List() ([]Task, error) {
	result, err := redisClient.Get(context.Background(), "tasks").Result()

	if err == redis.Nil {
		return nil, fmt.Errorf("cache not found %v", err)
	}

	var tasks []Task

	if err := json.Unmarshal([]byte(result), &tasks); err != nil {
		return nil, fmt.Errorf("cannot unmarshal json: %v", err)
	}

	return tasks, nil
}

func (t TaskCache) Find(id string) (Task, error) {
	log.Printf("finding key: %s", fmt.Sprintf("task:%s", id))

	result, err := redisClient.Get(context.Background(), fmt.Sprintf("task:%s", id)).Result()

	if err == redis.Nil {
		return Task{}, fmt.Errorf("cache not found")
	}

	var task Task

	if err := json.Unmarshal([]byte(result), &task); err != nil {
		return Task{}, fmt.Errorf("cannot unmarshal json: %v", err)
	}

	return task, nil
}

func (t TaskCache) Set(task Task) error {
	err := redisClient.Set(context.Background(), "task:"+task.Id, task, 5*time.Minute).Err()

	if err != nil {
		return fmt.Errorf("cannot set task cache: %v", err)
	}

	return nil
}
