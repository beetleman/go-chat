package server

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
)

const (
	Port int    = 8888
	Host string = "localhost"
)

type Connection struct {
	conn net.Conn
}

func (c *Connection) Write(data []byte) error {
	_, err := c.conn.Write(append(data, []byte("\n")...))
	return err
}

func (c *Connection) Close() {
	c.conn.Close()
}

type Server struct {
	connections []*Connection
	Address     string
	listener    net.Listener
	wg          sync.WaitGroup
}

// Handle run handlers for messages
func (server *Server) Handle(conn net.Conn) {
	server.connections = append(
		server.connections,
		&Connection{conn: conn},
	)
	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		data := scanner.Bytes()
		log.Println("data: ", string(data))
		for idx, c := range server.connections {
			server.wg.Add(1)
			go func() {
				err := c.Write(data)
				if err != nil {
					c.Close()
					server.connections[idx] = server.connections[len(server.connections)-1]
					server.connections = server.connections[:len(server.connections)-1]
					log.Println("Handle data", err)
				}
				server.wg.Done()
			}()
		}
	}
	server.wg.Done()
}

// Stop server`
func (server *Server) Stop() {
	server.listener.Close()
	for _, c := range server.connections {
		c.conn.Close()
	}
	server.connections = make([]*Connection, 0)
	server.wg.Wait()
	log.Println("Server stooped")
}

// Listen for messages
func (server *Server) Listen() {
	listener, err := net.Listen("tcp", server.Address)
	if err != nil {
		panic(err)
	}
	server.listener = listener
	log.Println("Server started")
	for {
		conn, err := listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return
			}
			log.Println(err, net.ErrClosed)
			continue
		}
		server.wg.Add(1)
		go server.Handle(conn)
	}
}

// New create new `Server`
func New() *Server {
	server := &Server{
		Address: fmt.Sprintf("%s:%d", Host, Port),
	}
	return server
}
