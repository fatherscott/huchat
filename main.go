package main

import (
	"fmt"
	"huchat/model"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("HuChat Start")

	e := model.NewServer()

	//Verify socket operation
	<-e.WSListenerMaked

	// Setup our Ctrl+C handler
	SetupCloseHandler(e)

	//All listen is waiting until closed.
	e.WaitListener.Wait()

	fmt.Println("HuChat Stop")
}

// SetupCloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS. We then handle this by calling
// our clean up procedure and exiting the program.
func SetupCloseHandler(e *model.EndPoint) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		e.Cancel()

	}()
}
