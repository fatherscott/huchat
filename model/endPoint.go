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

	WSListenerMaked chan bool

	Ctx    context.Context
	Cancel context.CancelFunc
}

//NewServer NewServer
func NewServer() *EndPoint {

	e := &EndPoint{
		WSListenerMaked: make(chan bool),
	}
	e.INFO = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	e.LoadConfig("conf.json")

	e.Ctx, e.Cancel = context.WithCancel(context.Background())

	e.WaitListener.Add(1)

	go e.WSListener()
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
