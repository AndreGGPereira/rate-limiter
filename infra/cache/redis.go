package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func newRepository(rdb *redis.Client) Repository {
	return &Cache{
		client: rdb,
	}
}

type Cache struct {
	client *redis.Client
}

func NewConnection(ctx context.Context, host, port, pwd string, db int) (Repository, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: pwd,
		DB:       db,
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	if pong != "PONG" {
		return nil, nil
	}
	return newRepository(rdb), nil
}

func (c *Cache) Incr(ctx context.Context, key string) (int64, error) {
	return c.client.Incr(ctx, key).Result()
}

func (c *Cache) Expire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return c.client.Expire(ctx, key, expiration).Result()
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.client.Set(ctx, key, value, expiration).Err()
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

func (c *Cache) Close() error {
	return c.client.Close()
}
