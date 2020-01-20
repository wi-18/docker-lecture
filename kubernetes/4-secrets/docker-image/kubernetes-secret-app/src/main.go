package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	username, err := ioutil.ReadFile("/etc/app/config/username")
	if err != nil {
		fmt.Println("could not find configuration username for username")
		os.Exit(1)
	}

	password, err := ioutil.ReadFile("/etc/app/config/password")
	if err != nil {
		fmt.Println("could not find configuration username for password")
		os.Exit(1)
	}

	fmt.Println("Found config, well done!")
	fmt.Println("username: " + string(username))
	fmt.Println("password: " + string(password))

	fmt.Println("Now waiting for container termination...")

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, syscall.SIGINT)
	<-signalChannel
}
