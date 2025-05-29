package ws

import (
	"fmt"
	"io"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

// constructor to init conns map
func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) HandleWS(ws *websocket.Conn) {
	fmt.Println("new connection from client: ", ws.RemoteAddr())
	s.conns[ws] = true
	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("error while reading", err)
			continue
		}
		msg := buf[:n]
		fmt.Println(string(msg))

		s.broadcast(msg)
	}
}

// write message to all clients connected to websocket (should later become "room")
// spawn a go routine for each client so it happens concurrently
func (s *Server) broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			_, err := ws.Write(b)
			if err != nil {
				fmt.Println("error while writing", err)
			}
		}(ws)
	}
}
