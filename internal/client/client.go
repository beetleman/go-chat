package client

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"sync"

	"github.com/beetleman/go-chat/internal/encoding"
	"github.com/beetleman/go-chat/internal/server"
)

type Client struct {
	port     int
	host     string
	conn     net.Conn
	UserName string
	wg       sync.WaitGroup
	r        io.ReadCloser
	w        io.WriteCloser
}

func (client *Client) Handle() {
	scanner := bufio.NewScanner(client.conn)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		message := encoding.Decode(scanner.Bytes())
		fmt.Println(message)
	}
	client.wg.Done()
}

func (client *Client) UserInput() {
	client.r, client.w = io.Pipe()
	go func() {
		io.Copy(client.w, os.Stdin) // TODO: not stop, waiting for intput
	}()
	scanner := bufio.NewScanner(client.r)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Printf("\033[1A\033[K")
		msg := encoding.Message{
			Text: text,
			User: client.UserName,
		}
		client.conn.Write(msg.Encode())
	}
	client.wg.Done()
}

func (client *Client) Stop() {
	client.conn.Close()
	client.r.Close()
	client.w.Close()
	client.wg.Wait()
}

func (client *Client) Connect() {
	address := fmt.Sprintf("%s:%d", client.host, client.port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}
	client.wg = sync.WaitGroup{}
	client.conn = conn
	client.wg.Add(1)
	go client.Handle()
	client.wg.Add(1)
	go client.UserInput()
}

func New(userName string) *Client {
	return &Client{
		host:     server.Host,
		port:     server.Port,
		UserName: userName,
	}
}
