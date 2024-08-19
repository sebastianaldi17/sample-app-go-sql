package handler

import (
	"context"
	"strconv"

	"github.com/go-chi/jwtauth/v5"
)

func getUserIDFromContext(ctx context.Context) (int64, error) {
	_, claims, err := jwtauth.FromContext(ctx)
	if err != nil {
		return 0, err
	}

	userIDInterface, ok := claims["id"]
	if !ok {
		return 0, err
	}

	userIDStr, ok := userIDInterface.(string)
	if !ok {
		return 0, err
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
