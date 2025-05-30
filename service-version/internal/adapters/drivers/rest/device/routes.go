package device

import (
	"github.com/OpenIoT-tools/OpenIoT/internal/adapters/drivers/rest/middlewares"
	"github.com/OpenIoT-tools/OpenIoT/internal/core/ports"
	"github.com/OpenIoT-tools/OpenIoT/internal/core/usecases"
	"github.com/OpenIoT-tools/OpenIoT/internal/security"
	"github.com/go-chi/chi/v5"
)

func StartDeviceApi(r *chi.Mux, deviceRepository ports.DeviceRespository, versionRepository ports.VersionRespository, securityToken security.SecurityToken) {
	deployService := usecases.NewDeploy(nil, securityToken)
	constroller := newDeviceController(usecases.NewDeviceService(deviceRepository, versionRepository, deployService))

	r.Route("/device", func(r chi.Router) {
		r.Use(middlewares.SetAuthMiddleware(securityToken))
		r.Get("/{id}", constroller.findDevice)
		r.Post("/", constroller.createDevice)
		r.Delete("/{id}", constroller.removeDevice)
		r.Get("/", constroller.listDevices)
		r.Patch("/version", constroller.updateVersionByCategory)
		r.Patch("/category/{id}/version", constroller.updateVersionByCategory)
		r.Patch("/{id}/hardware", constroller.updateHardware)
	})
}
