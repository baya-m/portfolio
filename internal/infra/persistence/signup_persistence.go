package persistence

import (
	"github.com/portfolio/internal/domain/model"
	"github.com/portfolio/internal/infra"
)

type SignupPersistence struct{}

/* Create  signup */
func (signupPersistence SignupPersistence) Create(signup model.Signup) error {
	_, err := infra.DbConnection.Exec("INSERT INTO LOGIN VALUES ('?','?')", signup.Id, signup.Password)
	if err != nil {
		return err
	}
	return err
}
