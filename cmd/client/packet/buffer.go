package packet

type PacketBuffer struct {
	Read   uint32
	Buffer []byte
}

func NewPacketBuffer() *PacketBuffer {
	return &PacketBuffer{
		Read:   0,
		Buffer: make([]byte, 4096),
	}
}
