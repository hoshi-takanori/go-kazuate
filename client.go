package main

import (
	"golang.org/x/net/websocket"
)

type Conn interface {
	Receive(interface{}) error
	Send(interface{}) error
}

type WSConn struct {
	ws *websocket.Conn
}

func (c WSConn) Receive(v interface{}) error {
	return websocket.JSON.Receive(c.ws, v)
}

func (c WSConn) Send(v interface{}) error {
	return websocket.JSON.Send(c.ws, v)
}

type Client struct {
	server *Server

	ws       Conn
	id       int
	name     string
	status   int
	opponent int
	oppName  string
	number   int
	game     *Game

	startCh chan bool
}

func NewClient(server *Server, ws Conn) *Client {
	return &Client{server: server, ws: ws, startCh: make(chan bool)}
}

func (c *Client) Start() {
	<-c.startCh
	close(c.startCh)

	println("Client", c.id, "Start")
	for {
		var m Message
		err := c.ws.Receive(&m)
		if err != nil {
			println("Client", c.id, "Error", err.Error())
			c.server.delCh <- c
			return
		} else {
			println("Client", c.id, "Receive: m.Command =", m.Command)
			m.Id = c.id
			c.server.msgCh <- &m
		}
	}
}
