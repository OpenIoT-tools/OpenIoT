package device

import (
	"net/http"

	"github.com/OpenIoT-tools/OpenIoT/internal/core/ports"
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
