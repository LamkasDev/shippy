package handler

import "github.com/LamkasDev/shippy/cmd/client/packet"

type PacketHandler struct {
	ProtocolId uint16
	Process    func(cpacket *packet.Packet) *error
}
