package model

import (
	"net"
	"net/http"
	"time"

	"nhooyr.io/websocket"
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
