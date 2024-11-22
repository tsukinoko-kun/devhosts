package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func AwaitCancel() <-chan struct{} {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	cancelChan := make(chan struct{})

	go func() {
		<-signalChan
		close(cancelChan)
	}()

	fmt.Println("Your defined hosts have been added")
	fmt.Println("Press Ctrl+C to stop and reset hosts")
	return cancelChan
}
