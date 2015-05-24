package main

import (
	"golang.org/x/net/websocket"
)

type Client struct {
	server *Server

	ws     *websocket.Conn
	id     int
	name   string
	status int

	startCh chan bool
}

func NewClient(server *Server, ws *websocket.Conn) *Client {
	return &Client{server: server, ws: ws, startCh: make(chan bool)}
}

func (c *Client) Start() {
	<-c.startCh
	close(c.startCh)

	println("Client", c.id, "Start")
	for {
		var m Message
		err := websocket.JSON.Receive(c.ws, &m)
		if err != nil {
			println("Client", c.id, "Error", err.Error())
			c.server.delCh <- c
			return
		} else {
			println("Client", c.id, "Receive: m.Name =", m.Name)
			m.Id = c.id
			c.server.msgCh <- &m
		}
	}
}
