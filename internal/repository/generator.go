package repository

import "github.com/Ras96/gcg/internal/model"

type GeneratorRepository interface {
	GenerateConstructors(file *model.File) (string, error)
}