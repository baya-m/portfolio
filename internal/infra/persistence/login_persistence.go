package persistence

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/portfolio/internal/domain/model"
)

type LoginPersistence struct{}

/* Select Login */
func (loginPersistence LoginPersistence) Select(login model.Login) (model.Login, error) {
	row := DbConnection.QueryRow("SELECT login_id, password FROM LOGIN WHERE login_id = ?", login.LoginID)
	err := row.Scan(&login.LoginID, &login.Password)
	fmt.Printf("login_id is %v", login.LoginID)
	fmt.Printf("login_id is %v", login.Password)
	if err != nil {
		//loginを変えさないようにしたい
		return login, err
	}
	return login, nil
}
