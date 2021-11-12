// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"github.com/Ras96/gcg/internal/handler"
	"github.com/Ras96/gcg/internal/repository"
	"github.com/google/wire"
)

// Injectors from wire.go:

func Handlers() *handler.Handlers {
	parserRepository := repository.NewParserRepository()
	repositories := repository.NewRepositories(parserRepository)
	handlers := handler.NewHandlers(repositories)
	return handlers
}

// wire.go:

var (
	mainSet = wire.NewSet(repository.NewParserRepository, repository.NewRepositories, handler.NewHandlers)
)