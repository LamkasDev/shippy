package life

import (
	"github.com/LamkasDev/shippy/cmd/client/connection"
	"github.com/LamkasDev/shippy/cmd/client/handler"
)

type ShippyInstance struct {
	ConnectionManager *connection.ConnectionManager
	HandlerContainer  *handler.PacketHandlerContainer
}

func NewShippyInstance() *ShippyInstance {
	instance := &ShippyInstance{}
	var err *error

	instance.ConnectionManager, err = connection.NewConnectionManager()
	if err != nil {
		panic(err)
	}

	instance.HandlerContainer = NewPacketHandlerContainer(instance)
	connection.SetupConnectionManager(instance.ConnectionManager, instance.HandlerContainer)

	return instance
}

func StartShippyInstance(instance *ShippyInstance) *error {
	connection.StartConnectionManager(instance.ConnectionManager)
	return nil
}

func EndShippyInstance(instance *ShippyInstance) *error {
	return connection.CloseConnectionManager(instance.ConnectionManager)
}
