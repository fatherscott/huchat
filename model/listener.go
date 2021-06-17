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

func (l *Listener) Leave(c *Client) {
	val, exists := l.Clients[c.AccountId]
	var roomId string
	if exists {
		roomId = val.RoomId
		val.Cancel()
		delete(l.Clients, c.AccountId)
	}

	if len(roomId) > 0 {
		room, exists := l.Rooms[roomId]
		if exists {
			delete(room, c.AccountId)
		}
	}
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
			l.Leave(obj)
			room := l.Enter(obj)

			sendLogin := Protocol.GetSendLogin()
			for _, v := range room {
				sendLogin.Connections = append(sendLogin.Connections, v.Conn)
			}

			sendLogin.Packet.AccountId = obj.AccountId
			sendLogin.Packet.RoomId = obj.RoomId
			sendLogin.Packet.Level = obj.Level
			sendLogin.Packet.NickName = obj.NickName

			e.SenderChannel <- sendLogin
		}

	case <-e.Ctx.Done():
		exit = true
	}

	return exit
}
