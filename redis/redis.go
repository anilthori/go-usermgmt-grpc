package database

import "github.com/go-redis/redis"

type redisDatabase struct {
	client *redis.Client
}

type Database interface {
	Set(key string, value string) (string, error)
	Get(key string) (string, error)
}

func CreateRedisDatabase() (Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return &redisDatabase{client: client}, nil
}

func (r *redisDatabase) Set(key string, value string) (string, error) {
	_, err := r.client.Set(key, value, 0).Result()
	if err != nil {
		return "", nil
	}
	return key, nil
}

func (r *redisDatabase) Get(key string) (string, error) {
	value, err := r.client.Get(key).Result()
	if err != nil {
		return "", nil
	}
	return value, nil
}
