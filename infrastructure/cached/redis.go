package cached

import (
	"fmt"
	"github.com/go-redis/cache/v8"
	"github.com/rs/zerolog/log"
	"toko/config"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var once sync.Once

type RedisImpl struct {
	db    *redis.Client
	cache *cache.Cache
}

func NewRedisClient(cfg config.Config) *RedisImpl {
	log.Info().Msg("Initialize Redis connection")
	host := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)
	rdb := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: cfg.RedisPassword, // no password set
		DB:       0,                 // use default DB
	})

	ctx := rdb.Context()
	ping, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Error().Msgf("Redis Connection: %v", err)
		return nil
	}
	log.Info().Msgf("Redis Connection: %v", ping)

	cache := cache.New(&cache.Options{
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
		Redis:      rdb,
	})

	return &RedisImpl{
		db:    rdb,
		cache: cache,
	}
}

func (c RedisImpl) DB() *redis.Client {
	return c.db
}

func (c RedisImpl) Cache() *cache.Cache {
	return c.cache
}

func (c RedisImpl) Close() error {
	return c.db.Close()
}
