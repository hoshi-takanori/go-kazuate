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

func (server *Server) Start() {
	for {
		select {
		case client := <-server.addCh:
			server.nextId++
			server.clients[server.nextId] = client
			client.idCh <- server.nextId

		case client := <-server.delCh:
			delete(server.clients, client.id)
		}
	}
}

func (server *Server) WebSocketHandler() websocket.Handler {
	return func(ws *websocket.Conn) {
		client := NewClient(server, ws)
		server.addCh <- client
		client.Start()
	}
}
