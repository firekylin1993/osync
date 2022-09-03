package bizagg

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"osync/internal/conf"
	"osync/internal/model"
)

type Agg struct {
	ChannelId int
	PackageId int
	Status    int
}

type AggRepo interface {
	Sync(context.Context, *model.ChannelPackageModel) (int64, error)
}

type AggUsecase struct {
	repo AggRepo
	conf *conf.Server
	log  *log.Helper
}

func NewAggUsecase(repo AggRepo, c *conf.Server, logger log.Logger) *AggUsecase {
	return &AggUsecase{
		repo: repo,
		conf: c,
		log:  log.NewHelper(logger),
	}
}

func (us *AggUsecase) SyncPackage(ctx context.Context, edn *model.OsyncModel) (int64, error) {
	c, span := otel.Tracer("OSync.SyncPackage").Start(ctx, "AggUsecase.SyncPackage")
	defer span.End()

	return us.repo.Sync(c, &model.ChannelPackageModel{
		ChannelName:   edn.ChannelName,
		UpdateVersion: edn.UpdateVersion,
	})
}
