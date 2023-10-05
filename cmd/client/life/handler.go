package life

import (
	"time"

	"github.com/LamkasDev/shippy/cmd/client/connection"
	"github.com/LamkasDev/shippy/cmd/client/handler"
	"github.com/LamkasDev/shippy/cmd/client/packet"
	"github.com/LamkasDev/shippy/cmd/pb/auth"
)

func NewPacketHandlerContainer(instance *ShippyInstance) *handler.PacketHandlerContainer {
	container := &handler.PacketHandlerContainer{
		Handlers: map[uint16]*handler.PacketHandler{},
	}
	handler.AddPacketHandler(container, NewGatewayServerConnectReplyPacketHandler(instance))
	handler.AddPacketHandler(container, NewGatewayServerLoginReplyPacketHandler(instance))

	return container
}

func NewGatewayServerConnectReplyPacketHandler(instance *ShippyInstance) *handler.PacketHandler {
	return &handler.PacketHandler{ProtocolId: packet.GATEWAY_SERVER_CONNECT_REPLY_PROTOCOL_ID, Process: func(cpacket *packet.Packet) *error {
		return connection.WriteConnection(instance.ConnectionManager.Gateway, packet.NewGatewayServerLoginPacket())
	}}
}

func NewGatewayServerLoginReplyPacketHandler(instance *ShippyInstance) *handler.PacketHandler {
	return &handler.PacketHandler{ProtocolId: packet.GATEWAY_SERVER_LOGIN_REPLY_PROTOCOL_ID, Process: func(cpacket *packet.Packet) *error {
		message := cpacket.Data.(*auth.GatewayServerLoginReplyMessage)
		if err := connection.CloseConnection(instance.ConnectionManager.Gateway, "done"); err != nil {
			return err
		}

		timer := time.NewTimer(60 * time.Second)
		<-timer.C
		connection.StartConnection(instance.ConnectionManager.Proxy)
		return connection.WriteConnection(instance.ConnectionManager.Proxy, packet.NewProxyServerLoginPacket(message.AccountId, message.ServerTicket))
	}}
}
