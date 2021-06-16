package model

import (
	"fmt"
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
		if listener.MainParser(e) {
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

func (l *Listener) Enter(c *Client) {

	l.Clients[c.AccountId] = c
	_, exists := l.Rooms[c.RoomId]
	if !exists {
		l.Rooms[c.RoomId] = make(map[string]*Client)
	}

	l.Rooms[c.RoomId][c.AccountId] = c
}

//MainParser
func (l *Listener) MainParser(e *EndPoint) (exit bool) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("MainParser", "RunTime Panic", string(Stack()), err)
		}
	}()

	exit = false

	select {
	case input := <-e.MainChannel:
		switch obj := input.(type) {

		case *EnterRequest:
			l.Leave(obj.Client)
			l.Enter(obj.Client)

		case *SendMessageRequest:
			e.INFO.Println(obj)

		case *LeaveRequest:
			l.Leave(obj.Client)
		}

	case <-e.Ctx.Done():
		exit = true
	}

	return exit
}
