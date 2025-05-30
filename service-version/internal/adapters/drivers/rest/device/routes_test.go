package device

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/OpenIoT-tools/OpenIoT/internal/security"
	"github.com/OpenIoT-tools/OpenIoT/tests/mocks"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestDeviceRoutes_ProtectedWithoutToken(t *testing.T) {
	r := chi.NewRouter()

	deviceRepo := mocks.NewDeviceRespository(t)
	versionRepo := mocks.NewVersionRespository(t)
	securityToken, _ := security.NewJWT("ASY", "SYM")

	StartDeviceApi(r, deviceRepo, versionRepo, securityToken)

	req := httptest.NewRequest(http.MethodGet, "/device", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
}
