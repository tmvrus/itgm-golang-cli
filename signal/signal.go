package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	stopApplication := make(chan struct{})
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	go func() {
		s := <-sigChan
		fmt.Printf("got terminating signal %d(%s)\n", s, s)
		close(stopApplication)
	}()

	<-stopApplication
	fmt.Printf("application stopped\n")
}
