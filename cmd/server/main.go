package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/beetleman/go-chat/internal/server"
)

func main() {
	s := server.New()
	sigCh := make(chan os.Signal, 1)
	doneCh := make(chan struct{}, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigCh
		s.Stop()
		doneCh <- struct{}{}
	}()
	go s.Listen()
	<-doneCh
}
