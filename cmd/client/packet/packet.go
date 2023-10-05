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

	// Encode packet header
	data = binary.BigEndian.AppendUint16(data, 0)
	data = append(data, 0)
	data = binary.BigEndian.AppendUint16(data, packet.ProtocolId)
	data = binary.BigEndian.AppendUint16(data, packet.PacketId)

	// Encode packet message
	messageData, err := proto.Marshal(packet.Data.(protoreflect.ProtoMessage))
	if err != nil {
		return data, &err
	}
	data = append(data, messageData...)

	// Assert size
	binary.BigEndian.PutUint16(data, uint16(len(data))-2)

	return data, nil
}

func DecodePacket(data []byte) (*Packet, *error, uint16) {
	if len(data) == 0 {
		return nil, nil, 0
	}

	// Check if there's enough data
	size := binary.BigEndian.Uint16(data[0:2])
	if uint16(len(data)) < size+2 {
		return nil, nil, 0
	}

	// Decode packet header
	protocolId := binary.BigEndian.Uint16(data[3:5])
	packetId := binary.BigEndian.Uint16(data[5:7])

	// Decode packet mesage
	message := MessagesByProtocolId[protocolId]
	proto.Unmarshal(data[7:size+2], message.(protoreflect.ProtoMessage))

	return NewPacket(protocolId, packetId, message), nil, size
}
