package packet

import (
	"crypto/md5"
	"fmt"

	"github.com/LamkasDev/shippy/cmd/pb/auth"
	"github.com/LamkasDev/shippy/cmd/pb/heartbeat"
)

var MessagesByProtocolId map[uint16]interface{} = map[uint16]interface{}{
	10801: &auth.GatewayServerConnectReplyMessage{},
	10021: &auth.GatewayServerLoginReplyMessage{},
	10023: &auth.ProxyServerLoginReplyMessage{},
	10101: &heartbeat.HeartbeatReplyMessage{},
}

func NewGatewayServerConnectPacket() *Packet {
	return NewPacket(GATEWAY_SERVER_CONNECT_PROTOCOL_ID, 0, &auth.GatewayServerConnectMessage{
		State:    58,
		Platform: "0",
	})
}

func NewGatewayServerLoginPacket() *Packet {
	return NewPacket(GATEWAY_SERVER_LOGIN_PROTOCOL_ID, 0, &auth.GatewayServerLoginMessage{
		LoginType: 1,
		Arg1:      "yostarus",
		Arg2:      "23601570",
		Arg3:      "62388b36dc7c44b02360157096605494",
		Arg4:      "0",
		CheckKey:  "93d39b95b0071107950ceb98b5e44a05",
		Device:    11,
	})
}

func NewProxyServerLoginPacket(accountId uint32, serverTicket string) *Packet {
	checkKeyText := []byte(serverTicket)
	checkKeyText = append(checkKeyText, []byte("dettimrepsignihtyrevednaeurtsignihton")...)
	checkKey := md5.Sum(checkKeyText)

	return NewPacket(PROXY_SERVER_LOGIN_PROTOCOL_ID, 0, &auth.ProxyServerLoginMessage{
		AccountId:    accountId,
		ServerTicket: serverTicket,
		Platform:     "0",
		ServerId:     6,
		CheckKey:     fmt.Sprintf("%x", checkKey),
		DeviceId:     "21994e06-46b7-44f5-bb0b-31a0cf77603c1696173667449",
	})
}

func NewHeartbeatPacket() *Packet {
	return NewPacket(HEARTBEAT_PROTOCOL_ID, 0, &heartbeat.HeartbeatMessage{
		Request: 1,
	})
}
