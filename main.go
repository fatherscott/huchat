package main

import "huchat/model"

func main() {
	e := model.NewServer()

	//Verify socket operation
	<-e.WSListenerMaked

	e.Cancel()
	e.WaitListener.Wait()
}
