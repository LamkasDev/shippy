package handler

type PacketHandlerContainer struct {
	Handlers map[uint16]*PacketHandler
}

func NewPacketHandlerContainer() *PacketHandlerContainer {
	container := &PacketHandlerContainer{
		Handlers: map[uint16]*PacketHandler{},
	}

	return container
}

func AddPacketHandler(container *PacketHandlerContainer, handler *PacketHandler) {
	container.Handlers[handler.ProtocolId] = handler
}

func GetPacketHandler(container *PacketHandlerContainer, protocolId uint16) *PacketHandler {
	return container.Handlers[protocolId]
}
