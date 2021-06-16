package model

import (
	"context"
	"sync"

	"nhooyr.io/websocket"
)

//Client Information
type Client struct {
	EndPoint  *EndPoint
	AccountId string
	Conn      *websocket.Conn

	Ctx    context.Context
	Cancel context.CancelFunc
}

// ClientPool sync.Pool
var ClientPool = sync.Pool{
	New: func() interface{} {
		client := new(Client)
		return client
	},
}

//Make 처음 생성시
func (c *Client) Make() {
	c.Ctx, c.Cancel = context.WithCancel(context.Background())
}

//Reset Reset
func (c *Client) Reset() {

	c.EndPoint = nil
	c.AccountId = ""
	c.Conn = nil

	c.Ctx = nil
	c.Cancel = nil
}

// NewClient 개인 유저 처리
func (e *EndPoint) NewClient(conn *websocket.Conn) {

	// var (
	// 	accountId string
	// 	seq       uint32
	// )

	// client := ClientPool.Get().(*Client)
	// client.EndPoint = e
	// client.Conn = conn

	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("NewClient", "RunTime Panic", string(Stack()), err)
	// 	}
	// 	client.Reset()
	// 	ClientPool.Put(client)
	// 	e.WaitClient.Done()
	// }()

	//타임 아웃 설정
	// firstCtx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	// defer cancel()

	// v, err := ReadWS(firstCtx, client.Conn, &e.RecvCount)
	// if err != nil {
	// 	close(client.SyncSender)
	// 	return
	// }

	// switch message := v.(type) {
	// case *Packet.CS_Enter:
	// 	gameUID = uint64(message.GameUID)
	// 	seq = message.SEQ
	// }
	// Packet.Release(v)

	// if gameUID == 0 {
	// 	close(client.SyncSender)
	// 	return
	// }

	// Packet.Log(gameUID, "CS_Enter", v)

	// dbSeq, result := e.Redis.GetSEQ(gameUID)
	// if result == false {
	// 	close(client.SyncSender)
	// 	return
	// }

	// if dbSeq != seq {
	// 	close(client.SyncSender)
	// 	return
	// }

	// client.GameUID = gameUID

	// go client.SendWorker(e)
	// <-client.SyncSender

	// //세션에 등록
	// sessionEnter := GetSessionEnter()
	// sessionEnter.Client = client
	// e.SessionChannel <- sessionEnter
	// <-client.SyncSession

	// for {
	// 	select {
	// 	case <-client.Context.Done():
	// 		return
	// 	default:
	// 		//5분동안 반응 없으면 에러처리
	// 		ctx, mainCancel := context.WithTimeout(context.Background(), time.Second*600)
	// 		defer mainCancel()

	// 		v, err := ReadWS(ctx, client.Conn, &e.RecvCount)
	// 		if err != nil {
	// 			return
	// 		}

	// 		switch message := v.(type) {
	// 		case *Packet.CS_Broadcast:

	// 			Packet.Log(gameUID, "CS_Broadcast", v)

	// 			broadcast := GetBroadcast()
	// 			broadcast.Client = client
	// 			broadcast.Message = message.Message
	// 			e.SessionChannel <- broadcast
	// 		}
	// 		Packet.Release(v)
	// 	}
	// }
}
