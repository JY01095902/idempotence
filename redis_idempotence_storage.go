package idempotence

import (
	"context"
	"errors"
	"log"
	"strings"
	"sync"

	"github.com/go-redis/redis/v8"
)

var redisClientOnce sync.Once
var redisClient *redis.Client

func newRedisClient(ip, port string, dbNum int) *redis.Client {
	redisClientOnce.Do(func() {
		address := ip + ":" + port
		if len(address) < 3 {
			log.Printf("create redis client faild, error: address(%s) is invalid \n", address)

			return
		}
		client := redis.NewClient(&redis.Options{
			Addr:     address,
			Password: "",
			DB:       dbNum,
		})

		_, err := client.Ping(context.Background()).Result()
		if err != nil {
			log.Printf("create redis client faild, error: %s \n", err.Error())

			return
		}

		redisClient = client
	})

	return redisClient
}

type RedisIdempotenceStorageConfig struct {
	IP       string
	Port     string
	DBNumber int
}

type RedisIdempotenceStorage struct {
	client *redis.Client
}

func NewRedisIdempotenceStorage(config RedisIdempotenceStorageConfig) (RedisIdempotenceStorage, error) {
	client := newRedisClient(config.IP, config.Port, config.DBNumber)
	if client == nil {
		return RedisIdempotenceStorage{}, errors.New("create redis client failed")
	}

	return RedisIdempotenceStorage{
		client: client,
	}, nil
}

func (storage RedisIdempotenceStorage) SaveIfAbsent(key, group string) error {
	if storage.client == nil {
		return errors.New("invalid redis client")
	}

	groupName := strings.ToLower(strings.TrimSpace(group)) + "_request_keys"
	cmd := storage.client.HSetNX(context.Background(), groupName, key, GetTimestamp())
	if cmd.Val() {
		return nil
	}

	return errors.New("idempotence key already exists")
}

func (storage RedisIdempotenceStorage) Remove(key, group string) error {
	if storage.client == nil {
		return errors.New("invalid redis client")
	}

	groupName := strings.ToLower(strings.TrimSpace(group)) + "_request_keys"
	cmd := storage.client.HDel(context.Background(), groupName, key)

	return cmd.Err()
}
