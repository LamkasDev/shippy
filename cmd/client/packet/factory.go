package packet

import "github.com/LamkasDev/shippy/cmd/pb/auth"

func NewGatewayServerConnectPacket() *Packet {
	return NewPacket(10800, 0, &auth.GatewayServerConnectMessage{
		State:    58,
		Platform: "0",
	})
}
