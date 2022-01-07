package application

import (
	"github.com/J0eppp/api.j0eppp.dev/pkg/config"
	"github.com/J0eppp/api.j0eppp.dev/pkg/db"
)

type Application struct {
	DB     *db.DB
	Config *config.Config
}

func Get() (*Application, error) {
	config := config.Get()
	db, err := db.Get(config.GetDBConnectionString())

	if err != nil {
		return nil, err
	}

	return &Application{
		DB:     db,
		Config: config,
	}, nil
}
