package main

import (
	"golang.org/x/net/websocket"
)

type Client struct {
	server *Server

	ws *websocket.Conn
	id int

	idCh chan int
}

func NewClient(server *Server, ws *websocket.Conn) *Client {
	return &Client{server: server, ws: ws, idCh: make(chan int)}
}

func (client *Client) Start() {
	client.id = <-client.idCh
	close(client.idCh)

	println("client", client.id, "Start")
	for {
		var message string
		err := websocket.Message.Receive(client.ws, &message)
		if err != nil {
			println("client", client.id, "Error", err.Error())
			client.server.delCh <- client
			return
		} else {
			println("client", client.id, "Receive", message)
		}
	}
}
