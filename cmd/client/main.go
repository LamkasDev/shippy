package main

import (
	"net"

	"github.com/LamkasDev/shippy/cmd/client/packet"
)

func main() {
	data, derr := packet.EncodePacket(packet.NewGatewayServerConnectPacket())
	if derr != nil {
		panic(*derr)
	}

	tcpServer, err := net.ResolveTCPAddr("tcp", "blhxusgate.yo-star.com:80")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpServer)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	_, err = conn.Write(data)
	if err != nil {
		panic(err)
	}

	// buffer to get data
	buffer := packet.NewPacketBuffer()
	for {
		read, err := conn.Read(buffer.Buffer)
		if err != nil {
			panic(err)
		}
		buffer.Read += uint32(read)
		println("Received message:", string(buffer.Buffer))
		break
	}
}
