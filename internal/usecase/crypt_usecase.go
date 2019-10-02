package usecase

import (
	"time"

	"github.com/davecgh/go-spew/spew"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/portfolio/internal/domain/model"
	"golang.org/x/crypto/bcrypt"
)

type CryptUsecase struct {
}

func (cryptUsecase CryptUsecase) CreateToken(login model.Login) (t string, err error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  login.LoginID,
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	spew.Dump(token)

	tokenString, err := token.SignedString([]byte("secret"))

	return tokenString, err
}

func (cryptUsecase CryptUsecase) HashPassword(signup model.Signup) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(signup.Password), 10)
	return hash, err
}
