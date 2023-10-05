package packet

type PacketBuffer struct {
	Data []byte
}

func NewPacketBuffer() PacketBuffer {
	return PacketBuffer{
		Data: make([]byte, 16384),
	}
}
