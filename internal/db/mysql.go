package db

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"osync/internal/conf"
)

type Tidb struct {
	*gorm.DB
}

// NewMysql .
func NewMysql(c *conf.Data, logger log.Logger) (*gorm.DB, error) {
	log.NewHelper(logger).Info("数据库连接中...")
	defer log.NewHelper(logger).Info("数据库连接成功")

	db, err := gorm.Open(mysql.Open(c.Mysql.Dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.WithMessage(err, "mysql连接失败")
	}
	if err := db.Use(otelgorm.NewPlugin(otelgorm.WithDBName("edn-mysql"))); err != nil {
		return nil, errors.WithMessage(err, "otelgorm插件部署失败")
	}
	return db, nil
}

func NewTidb(c *conf.Data, logger log.Logger) (Tidb, error) {
	log.NewHelper(logger).Info("Tidb数据库连接中...")
	defer log.NewHelper(logger).Info("Tidb数据库连接成功")

	db, err := gorm.Open(mysql.Open(c.Tidb.Dsn), &gorm.Config{})
	if err != nil {
		return Tidb{}, errors.WithMessage(err, "Tidb连接失败")
	}
	if err := db.Use(otelgorm.NewPlugin(otelgorm.WithDBName("edn-mysql"))); err != nil {
		return Tidb{}, errors.WithMessage(err, "tidb的otelgorm插件部署失败")
	}
	return Tidb{db}, nil
}
