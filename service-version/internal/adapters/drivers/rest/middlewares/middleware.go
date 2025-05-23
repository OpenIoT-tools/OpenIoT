package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/OpenIoT-tools/OpenIoT/internal/security"
)

type ctxKeyUser string

const userIDKey ctxKeyUser = "userID"

func SetAuthMiddleware(securityToken security.SecurityToken) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if !strings.HasPrefix(token, "Bearer ") {
				http.Error(w, "token unavailable", http.StatusUnauthorized)
				return
			}
			data, err := securityToken.ValidateSymmetricalToken(token)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			ctx, err := setContext(data, "sub", userIDKey, r.Context())
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func setContext[T comparable](data map[string]interface{}, field string, ctxField T, ctx context.Context) (context.Context, error) {
	if value, ok := data[field]; ok {
		return context.WithValue(ctx, ctxField, value), nil
	}
	return nil, fmt.Errorf("token sent does not have valid content")
}
