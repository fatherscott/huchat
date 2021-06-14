package main

import (
	"huchat/model"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	e := model.NewServer()
	// Setup our Ctrl+C handler
	SetupCloseHandler(e)

	e.INFO.Println("HuChat Start")

	//All listen is waiting until closed.
	e.WaitListener.Wait()

	e.INFO.Println("HuChat Stop")
}

// SetupCloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS. We then handle this by calling
// our clean up procedure and exiting the program.
func SetupCloseHandler(e *model.EndPoint) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		e.INFO.Println("\r- Ctrl+C pressed in Terminal")
		e.Cancel()

	}()
}
