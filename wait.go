package wait2die

import (
	"os"
	"os/signal"
	"fmt"
)

func WaitToDie(beforeDie func())  {
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			fmt.Printf("\nReceived an interrupt...\n\n")
			beforeDie()
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}