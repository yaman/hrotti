package main

import (    "os"
	"os/signal"
	"syscall"
	"fmt"
	. "github.com/abdulkadiryaman/hrotti/broker"
)

func main() {
	listener := NewListenerConfig("tcp://0.0.0.0:1883")

	h := NewHrotti(100)

	fmt.Println("starting app")

	h.AddListener("self", listener)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	h.Stop()
}
