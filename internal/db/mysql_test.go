package db

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/stretchr/testify/assert"
	"os"
	"osync/internal/conf"
	"osync/internal/model"
	"testing"
)

func TestNewMysql(t *testing.T) {
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)

	t.Run("mysql", func(t *testing.T) {
		conn, err := NewMysql(&conf.Data{
			Mysql: &conf.Data_Mysql{
				Dsn: "root:antiy@2018@tcp(10.251.21.20:31963)/enginedn?charset=utf8&parseTime=True&loc=Local",
			},
		}, logger)
		assert.Nil(t, err, "database connect error")

		var ready *model.OsyncModel
		err = conn.Where("app_name = ?", "agg").Find(&ready).Error
		assert.Nil(t, err, "database select error")
	})

	t.Run("tidb", func(t *testing.T) {
		conn, err := NewTidb(&conf.Data{
			Tidb: &conf.Data_Tidb{
				Dsn: "l_engine_admin:H7BzcYQEVJk7@tcp(192.168.200.130:3390)/L_direct_engine_monitor?charset=utf8mb4&parseTime=True&loc=Local",
			},
		}, logger)
		assert.Nil(t, err, "database connect error")

		var ready *model.ChannelPackageModel
		err = conn.Where("channel_name = ?", "yw_test").Find(&ready).Error
		assert.Nil(t, err, "database select error")
	})
}
