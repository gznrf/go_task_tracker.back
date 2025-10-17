package utils

import (
	"errors"
	"net/http"
)

func GetUserIdFromCtx(r *http.Request) (int64, error) {
	userId := r.Context().Value("userId")
	if userId == 0 {
		return 0, errors.New("user id not found")
	}
	userIdInt, ok := userId.(int64)
	if !ok {
		return 0, errors.New("user id not found")
	}

	return int64(userIdInt), nil
}
