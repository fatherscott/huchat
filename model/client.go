package model

import (
	"context"
	"huchat/Protocol"
	"strings"
	"sync"
	"time"

	"nhooyr.io/websocket"
)

//Client Information
type Client struct {
	EndPoint  *EndPoint
	AccountId string

	Channel chan interface{}

	Conn *websocket.Conn

	Ctx    context.Context
	Cancel context.CancelFunc
}

// ClientPool sync.Pool
var ClientPool = sync.Pool{
	New: func() interface{} {
		client := new(Client)
		client.Channel = make(chan interface{}, 8)
		client.Ctx, client.Cancel = context.WithCancel(context.Background())
		return client
	},
}

//Reset Reset
func (c *Client) Reset() {

	c.EndPoint = nil
	c.AccountId = ""
	c.Conn = nil
	close(c.Channel)

	c.Channel = make(chan interface{}, 16)
	c.Ctx, c.Cancel = context.WithCancel(context.Background())
}

// NewClient
func (e *EndPoint) NewClient(conn *websocket.Conn) {

	var (
		accountId string
		level     int32
		nickName  string
	)

	client := ClientPool.Get().(*Client)
	client.EndPoint = e
	client.Conn = conn

	defer func() {
		if err := recover(); err != nil {
			e.INFO.Println("NewClient", "RunTime Panic", string(Stack()), err)
		}
		client.Reset()
		ClientPool.Put(client)
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
		//deepCopy
		var sb1 strings.Builder
		sb1.WriteString(login.AccountId)
		accountId = sb1.String()

		var sb2 strings.Builder
		sb2.WriteString(login.NickName)
		nickName = sb2.String()

		level = login.Level

		e.INFO.Panicln(accountId, nickName, level)

		login.Release()

	default:
		return
	}

	for {
		select {
		case <-client.Ctx.Done():
			return
		case <-client.Channel:
			return
		default:
			//If you don't react for five minutes, you'll fail.
			ctx, mainCancel := context.WithTimeout(context.Background(), time.Second*600)
			defer mainCancel()

			v, err := e.ReadWS(ctx, client.Conn)
			if err != nil {
				return
			}
			switch message := v.(type) {
			case *Protocol.MessageRequest:
				message.Release()
			}
		}
	}
}
