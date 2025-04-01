package ports

import "github.com/OpenIoT-tools/OpenIoT/internal/core/models/entity"

type Broker interface {
	sendMessage(message any, device ...*entity.Device)
}
