package model

import (
	"context"
	"errors"
	Packet "huchat/packet"
	"net"
	"net/http"
	"sync/atomic"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wspb"
)

//WSListener Wait for message
func (e *EndPoint) WSListener() {
	defer func() {
		e.WaitListener.Done()
	}()

	l, err := net.Listen("tcp", "0.0.0.0:"+e.WSBind)
	if err != nil {
		e.INFO.Panicln("failed to listen:", err)
	}

	e.INFO.Println("tcp Listen:0.0.0.0", e.WSBind)
	defer l.Close()

	e.Server = &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			//e.INFO.Println("serving", r.RemoteAddr)

			c, err := websocket.Accept(w, r, nil)
			if err != nil {
				e.INFO.Println(r.RemoteAddr, err)
				return
			}
			defer c.Close(websocket.StatusInternalError, "the sky is falling")

			//client full close check
			e.WaitClient.Add(1)
		}),
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}
	defer e.Server.Close()

	e.WSListenerMaked <- true

	for {
		select {
		case <-e.Ctx.Done():
			return
		}
	}

}

//ReadWS ReadWS
func (e *EndPoint) ReadWS(ctx context.Context, c *websocket.Conn) (interface{}, error) {
	defer atomic.AddInt64(&e.RecvCount, 1)

	header := GetHeader()
	defer SetHeader(header)

	err := wspb.Read(ctx, c, header)
	if err != nil {
		return nil, err
	}

	switch header.Type {
	case Packet.Header_Type_LoginRequest:
	case Packet.Header_Type_LogOutRequest:
	case Packet.Header_Type_MessageRequest:
	}

	return nil, errors.New("invalid packet")
}

//WriteWS WriteWS
func (e *EndPoint) WriteWS(ctx context.Context, c *websocket.Conn, v interface{}, accountId string) error {
	defer atomic.AddInt64(&e.SendCount, 1)

	defer Packet.Release(v)

	header := Packet.GetHeader()
	defer Packet.Release(header)

	switch packet := v.(type) {

	case *Packet.CS_Enter_Ack:
		header.Type = Packet.Header_CS_Enter_Ack
		err := wspb.Write(ctx, c, header)
		if err != nil {
			return err
		}
		Packet.Log(gameUID, "CS_Enter_Ack", v)

		return wspb.Write(ctx, c, packet)

	case *Packet.SC_Enter_Second:
		header.Type = Packet.Header_SC_Enter_Second
		err := wspb.Write(ctx, c, header)
		if err != nil {
			return err
		}
		Packet.Log(gameUID, "SC_Enter_Second", v)

		return wspb.Write(ctx, c, packet)

	case *Packet.CS_Broadcast_Ack:
		header.Type = Packet.Header_CS_Broadcast_Ack
		err := wspb.Write(ctx, c, header)
		if err != nil {
			return err
		}
		Packet.Log(gameUID, "CS_Broadcast_Ack", v)

		return wspb.Write(ctx, c, packet)

	case *Packet.SC_Broadcast:
		header.Type = Packet.Header_SC_Broadcast
		err := wspb.Write(ctx, c, header)
		if err != nil {
			return err
		}
		Packet.Log(gameUID, "SC_Broadcast", v)

		return wspb.Write(ctx, c, packet)
	}

	return errors.New("not support type")
}
