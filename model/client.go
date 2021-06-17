package model

import (
	"context"
	"huchat/Protocol"
	"time"

	"nhooyr.io/websocket"
)

//Client Information
type Client struct {
	EndPoint  *EndPoint
	AccountId string
	Level     int32
	NickName  string
	RoomId    string

	Channel chan interface{}

	Conn *websocket.Conn

	Ctx    context.Context
	Cancel context.CancelFunc
}

// // ClientPool sync.Pool
// var ClientPool = sync.Pool{
// 	New: func() interface{} {
// 		client := new(Client)
// 		client.Channel = make(chan interface{}, 8)
// 		client.Ctx, client.Cancel = context.WithCancel(context.Background())
// 		return client
// 	},
// }

// //Reset Reset
// func (c *Client) Reset() {

// 	c.EndPoint = nil
// 	c.AccountId = ""
// 	c.Level = 0
// 	c.NickName = ""
// 	c.Conn = nil
// 	close(c.Channel)

// 	c.Channel = make(chan interface{}, 8)
// 	c.Ctx, c.Cancel = context.WithCancel(context.Background())
// }

// NewClient
func (e *EndPoint) NewClient(conn *websocket.Conn) {

	client := &Client{
		Channel:  make(chan interface{}, 8),
		EndPoint: e,
		Conn:     conn,
	}
	client.Ctx, client.Cancel = context.WithCancel(context.Background())

	defer func() {
		if err := recover(); err != nil {
			e.INFO.Println("NewClient", "RunTime Panic", string(Stack()), err)
		}
		close(client.Channel)
		e.WaitClient.Done()
	}()

	//Timeout Settings
	firstCtx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	v, err := e.ReadWS(firstCtx, client.Conn)
	if err != nil {
		return
	}

	switch login := v.(type) {
	case *Protocol.LoginRequest:
		client.AccountId = login.AccountId
		client.NickName = login.NickName
		client.Level = login.Level
		client.RoomId = login.RoomId
		e.INFO.Println(client.AccountId, client.Level, client.NickName, client.RoomId)

	default:
		return
	}

	e.ListenerChannel <- client

	for {
		select {
		case <-client.Ctx.Done():
			e.INFO.Println("hard close", client.NickName)
			return
		case <-client.Channel:
			return
		default:
			//If you don't react for five minutes, you'll fail.
			ctx, mainCancel := context.WithTimeout(context.Background(), time.Second*600)
			defer mainCancel()

			v, err := e.ReadWS(ctx, client.Conn)
			if err != nil {
				e.INFO.Println("close", client.NickName)
				leave := Protocol.GetLeaveUser()
				e.ListenerChannel <- leave
				return
			}
			switch obj := v.(type) {
			case *Protocol.MessageRequest:
				message := Protocol.GetMessageUser()
				message.AccountId = client.AccountId
				message.Msaage = obj.Message
				e.ListenerChannel <- message
			}
		}
	}
}
