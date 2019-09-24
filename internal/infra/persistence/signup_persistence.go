package persistence

import (
	"github.com/portfolio/internal/domain/model"
)

type SignupPersistence struct{}

/* Create Signup */
func (signupPersistence SignupPersistence) Create(signup model.Signup) error {
	ins, err := DbConnection.Prepare("INSERT INTO login(login_id, password) VALUES (?, ?)")
	if err != nil {
		return err
	}
	_, err = ins.Exec(signup.LoginID, signup.Password)
	if err != nil {
		return err
	}
	return err
}
