package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, syscall.SIGTERM)

	fmt.Println("Starting timer application..")
	for {
		select {
		case t := <-ticker.C:
			fmt.Printf("Current timestamp is %s\n", t.String())
		case <-signalChannel:
			fmt.Println("Shutting down timer application")
			return
		}
	}
}
