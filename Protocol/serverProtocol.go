package Protocol

import "nhooyr.io/websocket"

type LeaveUser struct {
	AccountId string
}

type MessageUser struct {
	Message   string
	AccountId string
}

type SendLogin struct {
	Connection *websocket.Conn
	Packet     LoginResponse
}

type SendMessage struct {
	Connections []*websocket.Conn
	Packet      MessageResponse
}
