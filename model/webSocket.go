package model

import (
	"context"
	"errors"
	"huchat/Protocol"
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

	header := Protocol.GetRequestHeader()
	defer Protocol.SetHeader(header)

	err := wspb.Read(ctx, c, header)
	if err != nil {
		return nil, err
	}

	switch header.Type {
	case Protocol.HeaderType_T_LoginRequest:
		packet := Protocol.GetLoginRequest()
		err := wspb.Read(ctx, c, packet)
		return packet, err

	case Protocol.HeaderType_T_LogoutRequest:
		packet := Protocol.GetLogoutRequest()
		err := wspb.Read(ctx, c, packet)
		return packet, err

	case Protocol.HeaderType_T_MessageRequest:
		packet := Protocol.GetMessageRequest()
		err := wspb.Read(ctx, c, packet)
		return packet, err
	}

	return nil, errors.New("invalid packet")
}

//WriteWS WriteWS
func (e *EndPoint) WriteWS(ctx context.Context, c *websocket.Conn, v interface{}) error {
	defer atomic.AddInt64(&e.SendCount, 1)

	header := Protocol.GetResponseHeader(v)
	defer Protocol.SetHeader(header)

	err := wspb.Write(ctx, c, header)
	if err != nil {
		return err
	}

	switch packet := v.(type) {
	case *Protocol.LoginResponse:
		return wspb.Write(ctx, c, packet)

	case *Protocol.LogoutResponse:
		return wspb.Write(ctx, c, packet)

	case *Protocol.MessageResponse:
		return wspb.Write(ctx, c, packet)
	}

	return errors.New("not support type")
}
