package model

import (
	"fmt"
)

//MainListener 작업
func (e *EndPoint) MainListener() {
	defer func() {
		e.WaitListener.Done()
	}()

	clients := make(map[string]*Client)
	rooms := make(map[string]*Client)

	e.MainListenerMaked <- true

	for {
		if e.MainParser(clients, rooms) {
			return
		}
	}
}

//MainParser
func (e *EndPoint) MainParser(clients map[string]*Client, rooms map[string]*Client) (exit bool) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("MainParser", "RunTime Panic", string(Stack()), err)
		}
	}()

	exit = false

	select {
	case messageInterface := <-e.MainChannel:
		switch message := messageInterface.(type) {

		default:
			e.INFO.Println(message)
		}

	case <-e.Ctx.Done():
		exit = true
	}

	return exit
}
