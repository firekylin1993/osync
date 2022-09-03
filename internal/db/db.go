package db

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(
	NewData,
	NewMysql,
	NewTidb,

	NewRedisConf,
	NewRedis,
)

// Data .
type Data struct {
	Mysql *gorm.DB
	Tidb  Tidb
	Redis *redis.Client
}

func NewData(db *gorm.DB, tidb Tidb, rdb *redis.Client) (*Data, error) {
	return &Data{
		Mysql: db,
		Tidb:  tidb,
		Redis: rdb,
	}, nil
}
