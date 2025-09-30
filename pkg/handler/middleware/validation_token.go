package handler

import (
	"errors"
	"fmt"
	"net/http"
)

func getUserId(w http.ResponseWriter, r *http.Request) (int, error) {
	id := r.Context().Value(userCtx)
	if id == nil {
		WriteError(w, 500, fmt.Errorf("user id not found"))
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		WriteError(w, 500, fmt.Errorf("user id not found"))
		return 0, errors.New("invalid user id")
	}

	return idInt, nil
}
