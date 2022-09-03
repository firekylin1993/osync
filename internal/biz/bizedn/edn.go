package bizedn

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"osync/internal/conf"
	"osync/internal/model"
)

type Edn struct {
	ChannelName string
	AppName     string
	Status      int
}

type EdnRepo interface {
	Search(context.Context) ([]*model.OsyncModel, error)
	Create(context.Context, []*model.OsyncModel) error
	Update(context.Context, *model.OsyncModel) error
	Filter(context.Context, *model.OsyncModel) ([]*model.OsyncModel, error)
}

type EdnUsecase struct {
	repo EdnRepo
	log  *log.Helper
	conf *conf.Server
}

func NewEdnUsecase(repo EdnRepo, c *conf.Server, logger log.Logger) *EdnUsecase {
	return &EdnUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
		conf: c,
	}
}

func (uc *EdnUsecase) SearchPackage(ctx context.Context) ([]*model.OsyncModel, error) {
	c, span := otel.Tracer("OSync.EdnUsecase").Start(ctx, "EdnUsecase.SearchMap")
	defer span.End()

	return uc.repo.Search(c)
}

func (uc *EdnUsecase) CreatePackage(ctx context.Context, osync []*model.OsyncModel) error {
	c, span := otel.Tracer("OSync.EdnUsecase").Start(ctx, "EdnUsecase.CreatePackage")
	defer span.End()

	return uc.repo.Create(c, osync)
}

func (uc *EdnUsecase) UpdatePackage(ctx context.Context, edn *model.OsyncModel) error {
	c, span := otel.Tracer("OSync.UpdatePackage").Start(ctx, "AggUsecase.UpdatePackage")
	defer span.End()

	return uc.repo.Update(c, edn)
}

func (uc *EdnUsecase) FilterPackage(ctx context.Context, edn *model.OsyncModel) ([]*model.OsyncModel, error) {
	c, span := otel.Tracer("OSync.SearchPackage").Start(ctx, "AggUsecase.SearchPackage")
	defer span.End()

	return uc.repo.Filter(c, edn)
}
