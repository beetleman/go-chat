package client

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/beetleman/go-chat/internal/encoding"
	"github.com/beetleman/go-chat/internal/server"
)

type Client struct {
	port     int
	host     string
	conn     *net.Conn
	UserName string
}

func (client *Client) Handle() {
	scanner := bufio.NewScanner(*client.conn)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		message := encoding.Decode(scanner.Bytes())
		fmt.Println(message)
	}
}

func (client *Client) UserInput() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Printf("\033[1A\033[K")
		msg := encoding.Message{
			Text: text,
			User: client.UserName,
		}
		conn := *client.conn
		conn.Write(msg.Encode())
	}
}

func (client *Client) Connect() {
	address := fmt.Sprintf("%s:%d", client.host, client.port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}
	client.conn = &conn
	go client.Handle()
	client.UserInput()
}

func New(userName string) *Client {
	return &Client{
		host:     server.Host,
		port:     server.Port,
		UserName: userName,
	}
}
