package repository

import "github.com/portfolio/internal/domain/model"

type SignupRepository interface {
	Create(model.Signup) error
}
