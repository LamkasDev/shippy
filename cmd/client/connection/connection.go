package connection

import (
	"fmt"
	"net"
	"sync"

	"github.com/LamkasDev/shippy/cmd/client/handler"
	"github.com/LamkasDev/shippy/cmd/client/packet"
	"github.com/jwalton/gchalk"
)

type Connection struct {
	Descriptor       ConnectionDescriptor
	TCP              *net.TCPConn
	Open             bool
	Closed           bool
	HandlerContainer *handler.PacketHandlerContainer

	Buffer    packet.PacketBuffer
	WriteLock sync.Mutex
	PacketId  uint16
}

type ConnectionDescriptor struct {
	Id   string
	Host string
	Port uint32
}

func NewConnection(descriptor ConnectionDescriptor) *Connection {
	return &Connection{
		Descriptor:       descriptor,
		TCP:              nil,
		Open:             false,
		Closed:           false,
		HandlerContainer: nil,

		Buffer:   packet.NewPacketBuffer(),
		PacketId: 0,
	}
}

func StartConnection(connection *Connection) *error {
	address, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", connection.Descriptor.Host, connection.Descriptor.Port))
	if err != nil {
		return &err
	}

	connection.TCP, err = net.DialTCP("tcp", nil, address)
	if err != nil {
		return &err
	}

	connection.Open = true
	fmt.Printf("🟢 connected to %s!\n", gchalk.Red(fmt.Sprintf("%s:%d", connection.Descriptor.Host, connection.Descriptor.Port)))

	return nil
}

func ProcessConnection(connection *Connection) *error {
	return ReadConnection(connection)
}

func ReadConnection(connection *Connection) *error {
	if _, err := connection.TCP.Read(connection.Buffer.Data); err != nil {
		return &err
	}
	cpacket, err, size := packet.DecodePacket(connection.Buffer.Data)
	if err != nil {
		return err
	}
	copy(connection.Buffer.Data, connection.Buffer.Data[size:])

	// Process packet handler
	fmt.Printf("⬅️  received packet from %s: %+v\n", gchalk.Red(connection.Descriptor.Id), cpacket)
	if handler := handler.GetPacketHandler(connection.HandlerContainer, cpacket.ProtocolId); handler != nil {
		if err := handler.Process(cpacket); err != nil {
			return err
		}
	}

	return nil
}

func WriteConnection(connection *Connection, cpacket *packet.Packet) *error {
	// Encode packet
	connection.WriteLock.Lock()
	cpacket.PacketId = connection.PacketId
	data, err := packet.EncodePacket(cpacket)
	if err != nil {
		fmt.Printf("❌ failed to encode packet: %+v\n", cpacket)
		return err
	}
	connection.PacketId++

	// Send data
	if _, err := connection.TCP.Write(data); err != nil {
		return &err
	}
	connection.WriteLock.Unlock()
	fmt.Printf("➡️  sent packet to %s: %+v\n", gchalk.Red(connection.Descriptor.Id), cpacket)

	return nil
}

func CloseConnection(connection *Connection, reason string) *error {
	if !connection.Open || connection.Closed {
		return nil
	}
	if err := connection.TCP.Close(); err != nil {
		return &err
	}
	connection.Closed = true
	fmt.Printf("🔴 disconnected from %s (%s)!\n", gchalk.Red(fmt.Sprintf("%s:%d", connection.Descriptor.Host, connection.Descriptor.Port)), reason)

	return nil
}
