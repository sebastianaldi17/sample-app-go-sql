package usecase

import (
	"strconv"
	"time"

	"github.com/go-chi/jwtauth/v5"
	userEntity "github.com/sebastianaldi17/sample-app-go-sql/internal/entity/user"
	"golang.org/x/crypto/bcrypt"
)

func (u *Usecase) CreateAccount(req userEntity.Login) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return err
	}

	return u.repo.CreateAccount(userEntity.Login{
		Username: req.Username,
		Password: string(hashedPassword),
	})
}

func (u *Usecase) ValidateLogin(req userEntity.Login) error {
	passwordHash, err := u.repo.GetPasswordHash(req.Username)
	if err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(req.Password))
}

func (u *Usecase) CreateJWT(req userEntity.Login) (string, error) {
	err := u.ValidateLogin(req)
	if err != nil {
		return "", err
	}

	userID, err := u.repo.GetUserIDFromUsername(req.Username)
	if err != nil {
		return "", err
	}

	claims := map[string]interface{}{
		"id":       strconv.FormatInt(userID, 10),
		"username": req.Username,
	}
	jwtauth.SetExpiryIn(claims, time.Hour)

	_, token, err := u.jwtTokenAuth.Encode(claims)
	return token, err
}
