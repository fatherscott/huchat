package model

import (
	"fmt"
	"huchat/Protocol"
)

type Listener struct {
	Clients map[string]*Client
	Rooms   map[string]map[string]*Client
}

//CreateListener
func (e *EndPoint) CreateListener() {
	defer func() {
		e.WaitListener.Done()
	}()

	listener := &Listener{
		Clients: make(map[string]*Client),
		Rooms:   make(map[string]map[string]*Client),
	}

	e.ListenerMaked <- true

	for {
		if listener.Parse(e) {
			return
		}
	}
}

func (l *Listener) Leave(accountId string) bool {
	val, exists := l.Clients[accountId]
	var roomId string
	if exists {
		roomId = val.RoomId
		val.Cancel()
		delete(l.Clients, accountId)
	}

	if len(roomId) > 0 {
		room, exists := l.Rooms[roomId]
		if exists {
			delete(room, accountId)
		}
		if len(room) == 0 {
			delete(l.Rooms, roomId)
		}
	}
	return exists
}

func (l *Listener) Enter(c *Client) map[string]*Client {

	l.Clients[c.AccountId] = c
	_, exists := l.Rooms[c.RoomId]
	if !exists {
		l.Rooms[c.RoomId] = make(map[string]*Client)
	}

	l.Rooms[c.RoomId][c.AccountId] = c

	return l.Rooms[c.RoomId]
}

func (l *Listener) GetClient(accountId string) *Client {
	val, exists := l.Clients[accountId]

	if exists {
		return val
	}
	return nil
}

func (l *Listener) GetRoom(accountId string) map[string]*Client {
	val, exists := l.Clients[accountId]
	var roomId string
	if exists {
		roomId = val.RoomId
		val.Cancel()
		delete(l.Clients, accountId)
	}

	if len(roomId) > 0 {
		room, exists := l.Rooms[roomId]
		if exists {
			return room
		}
	}
	return nil
}

//Parse
func (l *Listener) Parse(e *EndPoint) (exit bool) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Parse", "RunTime Panic", string(Stack()), err)
		}
	}()

	exit = false

	select {
	case input := <-e.ListenerChannel:
		switch obj := input.(type) {

		case *Client:
			if l.Leave(obj.AccountId) {
				obj.Cancel()

				sendLogin := Protocol.GetSendLogin()
				sendLogin.Packet.Result = 0
				sendLogin.Connection = obj.Conn
				e.SenderChannel <- sendLogin

			} else {
				l.Enter(obj)

				sendLogin := Protocol.GetSendLogin()
				sendLogin.Packet.Result = 1
				sendLogin.Connection = obj.Conn
				e.SenderChannel <- sendLogin
			}
		case *Protocol.LeaveUser:
			l.Leave(obj.AccountId)

		case *Protocol.MessageUser:
			client := l.GetClient(obj.AccountId)
			if client != nil {
				sendMessage := Protocol.GetSendMessage()
				room := l.GetRoom(obj.AccountId)
				if room != nil {
					for _, v := range room {
						sendMessage.Connections = append(sendMessage.Connections, v.Conn)
					}
				}
				sendMessage.Packet.AccountId = client.AccountId
				sendMessage.Packet.Message = obj.Message
				sendMessage.Packet.RoomId = client.RoomId
				sendMessage.Packet.Level = client.Level
				sendMessage.Packet.NickName = client.NickName

				e.SenderChannel <- sendMessage
			}
		}

	case <-e.Ctx.Done():
		exit = true
	}

	return exit
}
