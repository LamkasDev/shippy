package connection

import (
	"fmt"

	"github.com/LamkasDev/shippy/cmd/client/handler"
	"github.com/LamkasDev/shippy/cmd/client/packet"
)

type ConnectionManager struct {
	Gateway *Connection
	Proxy   *Connection
}

func NewConnectionManager() (*ConnectionManager, *error) {
	manager := &ConnectionManager{}
	var err *error

	fmt.Printf("🕒 connecting to servers...\n")
	manager.Gateway = NewConnection(ConnectionDescriptor{Id: "gateway", Host: "blhxusgate.yo-star.com", Port: 80})
	if StartConnection(manager.Gateway); err != nil {
		return manager, err
	}
	manager.Proxy = NewConnection(ConnectionDescriptor{Id: "proxy", Host: "blhxusproxy.yo-star.com", Port: 20000})

	return manager, nil
}

func SetupConnectionManager(manager *ConnectionManager, handlerContainer *handler.PacketHandlerContainer) {
	manager.Gateway.HandlerContainer = handlerContainer
	manager.Proxy.HandlerContainer = handlerContainer
}

func StartConnectionManager(manager *ConnectionManager) {
	go func() {
		for !manager.Gateway.Open {
		}
		WriteConnection(manager.Gateway, packet.NewGatewayServerConnectPacket())
		for !manager.Gateway.Closed {
			if err := ProcessConnection(manager.Gateway); err != nil {
				CloseConnection(manager.Gateway, (*err).Error())
			}
		}
	}()
	go func() {
		for !manager.Proxy.Open {
		}
		for !manager.Proxy.Closed {
			if err := ProcessConnection(manager.Proxy); err != nil {
				CloseConnection(manager.Proxy, (*err).Error())
			}
		}
	}()
	/* ticker := time.NewTicker(60 * time.Second)
	go func() {
		WriteConnection(manager.Gateway, packet.NewHeartbeatPacket())
		WriteConnection(manager.Proxy, packet.NewHeartbeatPacket())

		for {
			select {
			case <-ticker.C:
				WriteConnection(manager.Gateway, packet.NewHeartbeatPacket())
				WriteConnection(manager.Proxy, packet.NewHeartbeatPacket())
			}
		}
	}() */
}

func CloseConnectionManager(manager *ConnectionManager) *error {
	if err := CloseConnection(manager.Gateway, "app termination"); err != nil {
		return err
	}
	if err := CloseConnection(manager.Proxy, "app termination"); err != nil {
		return err
	}

	return nil
}
