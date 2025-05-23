package middlewares_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/OpenIoT-tools/OpenIoT/internal/adapters/drivers/rest/middlewares"
	"github.com/OpenIoT-tools/OpenIoT/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func getMiddlewareHandler(tokenValidator *mocks.SecurityToken) http.Handler {
	authMiddleware := middlewares.SetAuthMiddleware(tokenValidator)
	return authMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))
}

func TestSetAuthMiddleware(t *testing.T) {
	t.Run("when token is not sent, should return an error", func(t *testing.T) {
		authHandler := getMiddlewareHandler(mocks.NewSecurityToken(t))
		req := httptest.NewRequest("Get", "/", nil)
		w := httptest.NewRecorder()

		authHandler.ServeHTTP(w, req)

		assert.Equal(t, 401, w.Code)
	})

	t.Run("when an invalid token is sent, should return an error", func(t *testing.T) {
		tokenValidator := mocks.NewSecurityToken(t)
		authHandler := getMiddlewareHandler(tokenValidator)

		req := httptest.NewRequest("Get", "/", nil)
		w := httptest.NewRecorder()
		req.Header.Set("Authorization", "invalid")

		authHandler.ServeHTTP(w, req)
		assert.Equal(t, 401, w.Code)
	})

	t.Run("when cannot validate token, should return an error", func(t *testing.T) {
		tokenValidator := mocks.NewSecurityToken(t)
		tokenValidator.EXPECT().ValidateSymmetricalToken("Bearer invalid").Return(nil, fmt.Errorf("invalid token"))
		authHandler := getMiddlewareHandler(tokenValidator)

		req := httptest.NewRequest("Get", "/", nil)
		w := httptest.NewRecorder()
		req.Header.Set("Authorization", "Bearer invalid")

		authHandler.ServeHTTP(w, req)
		assert.Equal(t, 401, w.Code)
	})

	t.Run("when the token payload is invalid, should return an error", func(t *testing.T) {
		tokenValidator := mocks.NewSecurityToken(t)
		tokenValidator.EXPECT().ValidateSymmetricalToken("Bearer invalid").
			Return(map[string]interface{}{}, nil)
		authHandler := getMiddlewareHandler(tokenValidator)

		req := httptest.NewRequest("Get", "/", nil)
		w := httptest.NewRecorder()
		req.Header.Set("Authorization", "Bearer invalid")

		authHandler.ServeHTTP(w, req)
		assert.Equal(t, 401, w.Code)
	})

	t.Run("when valid token is sent, should return success", func(t *testing.T) {
		tokenValidator := mocks.NewSecurityToken(t)
		tokenValidator.EXPECT().ValidateSymmetricalToken("Bearer invalid").
			Return(map[string]interface{}{"sub": 456789123}, nil)
		authHandler := getMiddlewareHandler(tokenValidator)

		req := httptest.NewRequest("Get", "/", nil)
		w := httptest.NewRecorder()
		req.Header.Set("Authorization", "Bearer invalid")

		authHandler.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
	})
}
