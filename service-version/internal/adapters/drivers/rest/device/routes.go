package device

import "github.com/go-chi/chi/v5"

func StartDeviceApi(r *chi.Mux) {
	constroller := newDeviceController()

	r.Route("/device", func(r chi.Router) {
		r.Get("/{id}", constroller.findDevice)
		r.Post("", constroller.createDevice)
		r.Delete("/{id}", constroller.removeDevice)
		r.Get("/", constroller.listDevices)
		r.Patch("/category/{id}", constroller.updateVersionByCategory)
		r.Patch("/{id}/hardware", constroller.updateHardware)

	})
}

func setMiddlewares() {

}
