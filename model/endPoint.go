package model

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"
)

type EndPoint struct {
	WSBind string
	INFO   *log.Logger

	Server *http.Server

	WaitClient   sync.WaitGroup
	WaitListener sync.WaitGroup

	WSListenerMaked   chan bool
	MainListenerMaked chan bool

	MainChannel chan interface{}

	Ctx context.Context
	//Related Function Exit Request
	Cancel context.CancelFunc
}

//NewServer NewServer
func NewServer() *EndPoint {

	e := &EndPoint{
		WSListenerMaked:   make(chan bool),
		MainListenerMaked: make(chan bool),
		MainChannel:       make(chan interface{}, 2048),
	}
	e.INFO = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	e.LoadConfig("conf.json")

	e.Ctx, e.Cancel = context.WithCancel(context.Background())

	e.WaitListener.Add(2)

	go e.WSListener()
	//Verify socket operation
	<-e.WSListenerMaked

	go e.MainListener()
	//Verify parse operation
	<-e.MainListenerMaked
	return e
}

//LoadConfig Read the settings from the file.
func (e *EndPoint) LoadConfig(filePath string) bool {
	file, _ := os.Open(filePath)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(e)
	if err != nil {
		e.INFO.Panicln(err)
		return false
	}
	return true
}
