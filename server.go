package main

import (
	"golang.org/x/net/websocket"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (server *Server) Start() {
}

func (server *Server) WebSocketHandler() websocket.Handler {
	return func(ws *websocket.Conn) {
	}
}
