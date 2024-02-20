package log

import (
	"database/sql"
	"sync"

	"authen.agnoshealth.com/domain"
	"github.com/google/wire"
)

var (
	mdw     *middleware
	mdwOnce sync.Once

	svc     *service
	svcOnce sync.Once

	repo     *repository
	repoOnce sync.Once
)

func ProvideHandler(svc domain.LogService) *middleware {
	mdwOnce.Do(func() {
		mdw = &middleware{
			svc: svc,
		}
	})

	return mdw
}

func ProvideService(repo domain.LogRepository) *service {
	svcOnce.Do(func() {
		svc = &service{
			repo: repo,
		}
	})

	return svc
}

func ProvideRepository(db *sql.DB) *repository {
	repoOnce.Do(func() {
		repo = &repository{
			db: db,
		}
	})

	return repo
}


var ProviderSet wire.ProviderSet = wire.NewSet(
    ProvideHandler,
    ProvideService,
    ProvideRepository,

    wire.Bind(new(domain.LogMiddleware), new(*middleware)),
    wire.Bind(new(domain.LogService), new(*service)),
    wire.Bind(new(domain.LogRepository), new(*repository)),
)