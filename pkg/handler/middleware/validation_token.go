package h_middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gznrf/go_task_tracker.back.git/utils"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

type MiddleWareHandler struct {
}

func NewMiddleWareHandler() *MiddleWareHandler {
	return &MiddleWareHandler{}
}

func (h *MiddleWareHandler) UserIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			next.ServeHTTP(w, r)

			return
		}
		var userId int64
		header := r.Header.Get(authorizationHeader)
		if header == "" {
			cookie, err := r.Cookie("auth_token")
			if err != nil {
				if errors.Is(err, http.ErrNoCookie) {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}
			userId, err = findInCookie(cookie)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusBadRequest)
				return
			}
			ctx := context.WithValue(r.Context(), userCtx, userId)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		userId, err := findInHeaders(header)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), userCtx, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func findInHeaders(header string) (int64, error) {
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return 0, fmt.Errorf("invalid header")
	}

	userId, err := utils.ParseToken(headerParts[1])
	if err != nil {
		return 0, err
	}
	return userId, nil
}
func findInCookie(cookie *http.Cookie) (int64, error) {
	token := cookie.Value
	userId, err := utils.ParseToken(token)
	if err != nil {
		return 0, err
	}
	return userId, nil
}
