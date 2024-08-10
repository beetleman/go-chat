package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	port int    = 8888
	host string = "localhost"
)

type Server struct {
	stopCh   chan struct{}
	handlers []func([]byte) error
	Address  string
}

// Handle run handlers for messages
func (server *Server) Handle(conn net.Conn) {
	server.handlers = append(server.handlers, func(data []byte) error {
		_, err := conn.Write(append(data, []byte("\n")...))
		return err
	})
	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		data := scanner.Bytes()
		println("data: ", string(data))
		for idx, h := range server.handlers {
			go func() {
				err := h(data)
				if err != nil {
					server.handlers[idx] = server.handlers[len(server.handlers)-1]
					server.handlers = server.handlers[:len(server.handlers)-1]
					log.Println("Handle data", err)
				}
			}()
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Scanner err", err)
		}
	}
}

// Listen for messages
func (server *Server) Listen() {
	listener, err := net.Listen("tcp", server.Address)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go server.Handle(conn)
	}
}

// New create new `Server`
func New() *Server {
	server := &Server{
		stopCh:  make(chan struct{}),
		Address: fmt.Sprintf("%s:%d", host, port),
	}
	return server
}
