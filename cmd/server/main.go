package main

import (
	"fmt"

	"github.com/beetleman/go-chat/internal/server"
)

func main() {
	fmt.Println("Server started")
	s := server.New()
	s.Listen()
}
