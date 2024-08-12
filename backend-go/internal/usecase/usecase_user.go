package usecase

import (
	"strconv"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/sebastianaldi17/sample-app-go-sql/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

func (u *Usecase) CreateAccount(req entity.Login) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return err
	}

	return u.repo.CreateAccount(entity.Login{
		Username: req.Username,
		Password: string(hashedPassword),
	})
}

func (u *Usecase) ValidateLogin(req entity.Login) error {
	passwordHash, err := u.repo.GetPasswordHash(req.Username)
	if err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(req.Password))
}

func (u *Usecase) CreateJWT(req entity.Login) (string, error) {
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
