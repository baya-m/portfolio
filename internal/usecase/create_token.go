package usecase

import (
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/portfolio/internal/domain/model"
)

type CreateToken struct {
}

func (createToken CreateToken) Create(login model.Login) (t string, err error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  login.LoginID,
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	spew.Dump(token)

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		log.Println(err)
		return "", err
	}
	return tokenString, nil
}
