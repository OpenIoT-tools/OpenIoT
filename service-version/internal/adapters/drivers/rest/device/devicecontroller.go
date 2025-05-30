package device

import (
	"fmt"
	"net/http"

	"github.com/OpenIoT-tools/OpenIoT/internal/core/ports"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type deviceController struct {
	deviceUseCase ports.DeviceUseCase
}

func newDeviceController(deviceUseCase ports.DeviceUseCase) *deviceController {
	return &deviceController{
		deviceUseCase: deviceUseCase,
	}
}

func (d *deviceController) findDevice(w http.ResponseWriter, r *http.Request) {
	err := validator.New().Var(r.URL.Query().Get("deviceId"), "required,numeric,min=0")
	if err != nil {
		fmt.Println(err)
	}

	device, err := d.deviceUseCase.FindDeviceById(r.URL.Query().Get("deviceId"))
	if err != nil {

	}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, device)
}

func (d *deviceController) createDevice(w http.ResponseWriter, r *http.Request) {

}

func (d *deviceController) removeDevice(w http.ResponseWriter, r *http.Request) {

}

func (d *deviceController) listDevices(w http.ResponseWriter, r *http.Request) {

}

func (d *deviceController) updateVersion(w http.ResponseWriter, r *http.Request) {

}

func (d *deviceController) updateVersionByCategory(w http.ResponseWriter, r *http.Request) {

}

func (d *deviceController) updateHardware(w http.ResponseWriter, r *http.Request) {

}

func (d *deviceController) syncDeviceVersion(w http.ResponseWriter, r *http.Request) {

}
