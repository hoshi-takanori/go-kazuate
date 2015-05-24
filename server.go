package main

import (
	"golang.org/x/net/websocket"
)

type Server struct {
	nextId  int
	clients map[int]*Client

	addCh chan *Client
	delCh chan *Client
}

func NewServer() *Server {
	return &Server{
		nextId:  0,
		clients: map[int]*Client{},

		addCh: make(chan *Client, 1),
		delCh: make(chan *Client),
	}
}

func (s *Server) Start() {
	for {
		select {
		case c := <-s.addCh:
			s.nextId++
			c.id = s.nextId
			s.clients[c.id] = c
			c.startCh <- true

		case c := <-s.delCh:
			delete(s.clients, c.id)
		}
	}
}

func (s *Server) WebSocketHandler() websocket.Handler {
	return func(ws *websocket.Conn) {
		c := NewClient(s, ws)
		s.addCh <- c
		c.Start()
	}
}
