package main

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"osync/internal/biz/bizagg"
	"osync/internal/biz/bizedn"
	"osync/internal/kits"
	"osync/internal/model"
)

func newOsync(ctx context.Context, edn *bizedn.EdnUsecase, agg *bizagg.AggUsecase, logger log.Logger) error {
	log.NewHelper(logger).Info("初始化同步数据...")
	defer log.NewHelper(logger).Info("数据初始化同步完成")

	c, span := otel.Tracer("OSync.newOsync").Start(ctx, "newOsync")
	defer span.End()

	searchPack, err := edn.SearchPackage(c)
	if err != nil {
		return errors.WithMessage(err, "查询同步数据失败")
	}

	if len(searchPack) > 0 {
		err = edn.CreatePackage(c, searchPack)
		if err != nil {
			return errors.WithMessage(err, "插入同步数据失败")
		}
	}

	filterPackage, err := edn.FilterPackage(c, &model.OsyncModel{Status: 0})
	if err != nil {
		return errors.WithMessage(err, "过滤未同步agg数据失败")
	}
	if len(filterPackage) == 0 {
		return nil
	}

	pool := kits.NewGoPool(10)
	for _, v := range filterPackage {
		pool.Add(1)
		go func(o *model.OsyncModel) error {
			defer pool.Done()
			_, err := agg.SyncPackage(c, o)
			o.Status = 1
			if err != nil {
				log.NewHelper(logger).Errorf("渠道：%s，版本：%s 同步%s失败:%s", o.ChannelName, o.UpdateVersion, o.AppName, err.Error())
				o.Status = 2
			}
			return edn.UpdatePackage(c, o)
		}(v) //nolint:errcheck
	}
	pool.Wait()

	return nil
}
