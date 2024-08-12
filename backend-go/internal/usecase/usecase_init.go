package usecase

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/sebastianaldi17/sample-app-go-sql/internal/repo"
)

type Usecase struct {
	jwtTokenAuth *jwtauth.JWTAuth
	repo         repo.Repo
}

func New(repo repo.Repo, tokenAuth *jwtauth.JWTAuth) *Usecase {
	return &Usecase{
		jwtTokenAuth: tokenAuth,
		repo:         repo,
	}
}
