package Protocol

import "nhooyr.io/websocket"

type LeaveUser struct {
	AccountId string
}

type MessageUser struct {
	Msaage    string
	AccountId string
}

type SendLogin struct {
	Connections []*websocket.Conn
	Packet      LoginResponse
}
