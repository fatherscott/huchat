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

	ListenerMaked   chan bool

	ListenerChannel chan interface{}
	SenderChannel chan interface{}

	Ctx context.Context
	//Related Function Exit Request
	Cancel context.CancelFunc

	RecvCount int64
	SendCount int64
}

//NewServer NewServer
func NewServer() *EndPoint {

	e := &EndPoint{
		ListenerMaked: make(chan bool),
		ListenerChannel:       make(chan interface{}, 2048),
		SenderChannel:       make(chan interface{}, 256),
	}
	e.INFO = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	e.LoadConfig("conf.json")

	e.Ctx, e.Cancel = context.WithCancel(context.Background())

	e.WaitListener.Add(2)

	go e.WSListener()
	//Verify socket operation
	<-e.ListenerMaked

	go e.CreateListener()
	//Verify parse operation
	<-e.ListenerMaked

	go e.CreateSender()
	//Verify parse operation
	<-e.ListenerMaked

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
