package repository

import "github.com/portfolio/internal/domain/model"

type LoginRepository interface {
	Select(model.Login) (model.Login, error)
}
