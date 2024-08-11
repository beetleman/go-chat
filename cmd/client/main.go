package main

import (
	"fmt"
	"log"
	"os"

	"github.com/beetleman/go-chat/internal/client"
)

func main() {
	fmt.Println("Client started")
	if len(os.Args) != 2 {
		log.Fatalln("help: client userName")
	}
	c := client.New(os.Args[1])
	c.Connect()
}
