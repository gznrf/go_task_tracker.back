package h_middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gznrf/go_task_tracker.back.git/utils"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

type HMiddleWare struct {
}

func NewMiddleWare() *HMiddleWare {
	return &HMiddleWare{}
}

func (h *HMiddleWare) UserIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authorizationHeader)
		if header == "" {
			utils.WriteError(w, 401, fmt.Errorf("empty auth header"))
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			utils.WriteError(w, 401, fmt.Errorf("invalid auth header"))
			return
		}

		userId, err := utils.ParseToken(headerParts[1])
		if err != nil {
			utils.WriteError(w, 401, fmt.Errorf("invalid auth header"))
			return
		}

		ctx := context.WithValue(r.Context(), userCtx, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
