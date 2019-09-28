package usecase

import (
	"github.com/portfolio/internal/domain/model"
	"github.com/portfolio/internal/domain/repository"
	"github.com/portfolio/internal/infra/persistence"
)

type LoginUsecase struct{}

func (loginUseCase LoginUsecase) Select(login model.Login) (model.Login, error) {
	result, err := repository.LoginRepository(persistence.LoginPersistence{}).Select(login)
	return result, err
}
