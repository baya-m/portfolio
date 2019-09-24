package usecase

import (
	"github.com/portfolio/internal/domain/model"
	"github.com/portfolio/internal/domain/repository"
	"github.com/portfolio/internal/infra/persistence"
)

type SignupUsecase struct{}

func (signUsecase SignupUsecase) Create(signup model.Signup) error {
	err := repository.SignupRepository(persistence.SignupPersistence{}).Create(signup)
	return err
}
