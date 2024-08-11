package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/beetleman/go-chat/internal/client"
)

func main() {
	fmt.Println("Client started")
	if len(os.Args) != 2 {
		log.Fatalln("help: client userName")
	}
	sigCh := make(chan os.Signal, 1)
	doneCh := make(chan struct{}, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	c := client.New(os.Args[1])
	go func() {
		<-sigCh
		c.Stop()
		doneCh <- struct{}{}
	}()
	c.Connect()
	<-doneCh
}
