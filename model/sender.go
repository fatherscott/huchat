package model

import (
	"context"
	"fmt"
	"huchat/Protocol"
	"time"
)

type Sender struct {
}

//CreateSender
func (e *EndPoint) CreateSender() {
	defer func() {
		e.WaitListener.Done()
	}()

	sender := &Sender{}

	e.ListenerMaked <- true

	for {
		if sender.Parse(e) {
			return
		}
	}
}

//Parse
func (s *Sender) Parse(e *EndPoint) (exit bool) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Parse", "RunTime Panic", string(Stack()), err)
		}
	}()

	exit = false

	select {
	case input := <-e.SenderChannel:
		switch obj := input.(type) {
		case *Protocol.SendLogin:
			for _, con := range obj.Connections {

				ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
				defer cancel()

				e.WriteWS(ctx, con, &obj.Packet)
			}
			obj.Release()
		}

	case <-e.Ctx.Done():
		exit = true
	}

	return exit
}
