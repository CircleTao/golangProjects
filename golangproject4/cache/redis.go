package cache

import (
	"context"
	"golangproject4/config"

	"github.com/redis/go-redis/v9"
)

var (
	Rdb  *redis.Client
	Rtcx context.Context
)

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: config.RedisPassword,
		DB:       config.RedisDb,
	})
	Rtcx = context.Background()
}

func Zscore(id int, score int) redis.Z {
	return redis.Z{
		Score:  float64(score),
		Member: id,
	}
}
