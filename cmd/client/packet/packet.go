package packet

import (
	"encoding/binary"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Packet struct {
	ProtocolId uint16
	PacketId   uint16
	Data       interface{}
}

func NewPacket(protocolId uint16, packetId uint16, data interface{}) *Packet {
	return &Packet{
		ProtocolId: protocolId,
		PacketId:   packetId,
		Data:       data,
	}
}

func EncodePacket(packet *Packet) ([]byte, *error) {
	data := []byte{}

	// Placeholder size
	data = binary.BigEndian.AppendUint16(data, 0)

	// Zero
	data = append(data, 0)

	// Protocol ID
	data = binary.BigEndian.AppendUint16(data, packet.ProtocolId)

	// Packet ID
	data = binary.BigEndian.AppendUint16(data, packet.PacketId)

	// Data
	messageData, err := proto.Marshal(packet.Data.(protoreflect.ProtoMessage))
	if err != nil {
		return data, &err
	}
	data = append(data, messageData...)

	// Assert size
	binary.BigEndian.PutUint16(data, uint16(len(data))-2)

	return data, nil
}

func DecodePacket(data []byte) (*Packet, *error) {
	if len(data) == 0 {
		return nil, nil
	}

	packet := NewPacket(0, 0, nil)

	return packet, nil
}
