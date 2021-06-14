package model

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type EndPoint struct {
	WSBind string
	INFO   *log.Logger
}

//Initialize Initialize object
func (e *EndPoint) Initialize() {
	e.INFO = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	e.LoadConfig("conf.json")

}

//Load Read the settings from the file.
func (e *EndPoint) LoadConfig(filePath string) bool {
	file, _ := os.Open(filePath)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(e)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
