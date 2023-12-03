package svc

import (
	"github.com/aheadIV/NightVoyager/appsvc/internal/config"
	"github.com/go-rel/postgres"
	"github.com/go-rel/rel"
)

type ServiceContext struct {
	Config config.Config
	DB     rel.Repository
}

func NewServiceContext(c config.Config) *ServiceContext {
	adapter, _ := postgres.Open(c.Postgres.DataSource)

	// initialize rel's repo.
	repo := rel.New(adapter)
	// db :=
	return &ServiceContext{
		Config: c,
		DB:     repo,
	}
}
