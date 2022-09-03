package myagg

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"osync/internal/biz/bizagg"
	"osync/internal/db"
	"osync/internal/model"
)

type aggRepo struct {
	data *db.Data
	log  *log.Helper
}

func NewAggRepo(data *db.Data, logger log.Logger) bizagg.AggRepo {
	return &aggRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *aggRepo) Sync(ctx context.Context, o *model.ChannelPackageModel) (int64, error) {
	c, span := otel.Tracer("OSync.Sync").Start(ctx, "aggRepo.Sync")
	defer span.End()
	result := r.data.Tidb.WithContext(c).Create(o)
	return result.RowsAffected, result.Error
}
