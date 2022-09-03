package db

import (
	"osync/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"
)

func NewRedisConf(data *conf.Data) *redis.Options {
	return &redis.Options{
		Addr:     data.Redis.Addr,
		Password: data.Redis.Pwd,     // no password set
		DB:       int(data.Redis.Db), // use default DB
	}
}

func NewRedis(ro *redis.Options, logger log.Logger) *redis.Client {
	log.NewHelper(logger).Info("Redis连接中...")
	defer log.NewHelper(logger).Info("Redis连接成功")
	rdb := redis.NewClient(ro)
	rdb.AddHook(redisotel.NewTracingHook())
	return rdb
}
