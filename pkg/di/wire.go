//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/athunlal/Note-Taking-Application/pkg/api"
	"github.com/athunlal/Note-Taking-Application/pkg/api/handler"
	"github.com/athunlal/Note-Taking-Application/pkg/db"
	"github.com/athunlal/Note-Taking-Application/pkg/repository"
	"github.com/athunlal/Note-Taking-Application/pkg/usecase"
	"github.com/google/wire"

	"github.com/athunlal/Note-Taking-Application/pkg/config"
)

func InitApi(cfg config.Config) (*http.ServerHttp, error) {
	wire.Build(
		db.ConnectToDb,
		repository.NewUserRepo,
		usecase.NewUserUseCase,
		handler.NewUserHandler,
		http.NewServerHttp,
	)
	return &http.ServerHttp{}, nil
}
