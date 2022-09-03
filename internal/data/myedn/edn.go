package myedn

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"osync/internal/biz/bizedn"
	"osync/internal/conf"
	"osync/internal/db"
	"osync/internal/model"
	"time"
)

type ednRepo struct {
	data *db.Data
	conf *conf.Server
	log  *log.Helper
}

func NewEdnRepo(data *db.Data, c *conf.Server, logger log.Logger) bizedn.EdnRepo {
	return &ednRepo{
		data: data,
		conf: c,
		log:  log.NewHelper(logger),
	}
}

func (r *ednRepo) Search(ctx context.Context) ([]*model.OsyncModel, error) {
	c, span := otel.Tracer("OSync.Search").Start(ctx, "ednRepo.Search")
	defer span.End()
	//获取同步期限内最早时间
	currentTime := time.Now()
	oldTime := currentTime.AddDate(0, 0, -int(r.conf.SyncDuration))
	unix := time.Date(oldTime.Year(), oldTime.Month(), oldTime.Day(), 0, 0, 0, 0, oldTime.Location()).Unix()

	var osyncs []*model.OsyncModel
	result := r.data.Mysql.WithContext(c).Raw(`SELECT
	"Agg" AS app_name,
	t.channel_id,
	t.package_id,
	c.channel_name,
	p.update_version 
FROM
	package_channel t
	LEFT JOIN package_info p ON ( t.package_id = p.id )
	LEFT JOIN channel c ON ( t.channel_id = c.id ) 
WHERE
	NOT EXISTS ( SELECT id FROM osync_models o WHERE t.channel_id = o.channel_id AND t.package_id = o.package_id ) 
	AND t.release_time > ?
	AND t.is_handled > ?`, unix, 0).Scan(&osyncs)

	return osyncs, result.Error
}

func (r *ednRepo) Create(ctx context.Context, osync []*model.OsyncModel) error {
	c, span := otel.Tracer("OSync.Create").Start(ctx, "ednRepo.Create")
	defer span.End()

	return r.data.Mysql.WithContext(c).Create(osync).Error
}

func (r *ednRepo) Update(ctx context.Context, o *model.OsyncModel) error {
	c, span := otel.Tracer("OSync.Update").Start(ctx, "ednRepo.Update")
	defer span.End()

	return r.data.Mysql.WithContext(c).Updates(o).Error
}

func (r *ednRepo) Filter(ctx context.Context, o *model.OsyncModel) ([]*model.OsyncModel, error) {
	c, span := otel.Tracer("OSync.Filter").Start(ctx, "ednRepo.Filter")
	defer span.End()

	var osync []*model.OsyncModel
	result := r.data.Mysql.WithContext(c).Where("status = ?", o.Status).Find(&osync)
	return osync, result.Error
}
