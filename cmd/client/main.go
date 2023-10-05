package main

import (
	"os"
	"os/signal"

	"github.com/LamkasDev/shippy/cmd/client/life"
)

func main() {
	// Start the game instance
	instance := life.NewShippyInstance()
	if err := life.StartShippyInstance(instance); err != nil {
		panic(err)
	}

	// Block until program termination
	terminateSignal := make(chan os.Signal, 1)
	signal.Notify(terminateSignal, os.Interrupt)
	<-terminateSignal

	// End the game instance
	if err := life.EndShippyInstance(instance); err != nil {
		panic(err)
	}
	os.Exit(0)
}
