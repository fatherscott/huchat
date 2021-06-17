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
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
			e.WriteWS(ctx, obj.Connection, &obj.Packet)

			cancel()
			obj.Release()
		}

	case <-e.Ctx.Done():
		exit = true
	}

	return exit
}
