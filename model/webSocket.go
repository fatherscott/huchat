package model

import (
	"context"
	"errors"
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

//ReadWS 데이터 읽기
func ReadWS(ctx context.Context, c *websocket.Conn, counter *int64) (interface{}, error) {
	if counter != nil {
		defer atomic.AddInt64(counter, 2)
	}

	header := Packet.GetHeader()
	defer Packet.Release(header)

	err := wspb.Read(ctx, c, header)
	if err != nil {
		return nil, err
	}

	switch header.Type {
	case Packet.Header_CS_Enter:
		csEnter := Packet.GetCSEnter()
		err := wspb.Read(ctx, c, csEnter)
		return csEnter, err

	case Packet.Header_CS_Broadcast:
		csBroadcast := Packet.GetCSBroadcast()
		err := wspb.Read(ctx, c, csBroadcast)
		return csBroadcast, err
	}

	return nil, errors.New("invalid packet")
}
