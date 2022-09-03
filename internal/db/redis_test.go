package db

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"os"
	"osync/internal/conf"
	"testing"
	"time"
)

var RedisConf *conf.Data_Redis

func TestMain(m *testing.M) {
	viper.AutomaticEnv()
	viper.SetDefault("REDIS_DSN", "redis://redis-a53eecff.myyun.org:6379/1")
	dsn := viper.GetString("REDIS_DSN")
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		panic(err)
	}
	RedisConf = &conf.Data_Redis{
		Addr: opt.Addr,
		Pwd:  opt.Password,
		Db:   int32(opt.DB),
	}
	m.Run()
}

func TestNewRedis(t *testing.T) {
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
	t.Run("redis", func(t *testing.T) {
		opt := NewRedisConf(&conf.Data{
			Redis: RedisConf,
		})
		cli := NewRedis(opt, logger)
		err := cli.Ping(context.Background()).Err()
		assert.NoError(t, err)
		err = cli.Set(context.Background(), "test", 123, 20*time.Second).Err()
		assert.NoError(t, err)
	})
}
