package main

import (
	"golang.org/x/net/websocket"
)

type Server struct {
	nextId  int
	clients map[int]*Client

	addCh chan *Client
	delCh chan *Client
	msgCh chan *Message
}

func NewServer() *Server {
	return &Server{
		nextId:  0,
		clients: map[int]*Client{},

		addCh: make(chan *Client, 1),
		delCh: make(chan *Client),
		msgCh: make(chan *Message),
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
			if c.game != nil {
				o := c.game.Opponent(c)
				o.SetDone("Opponent has been disconnected.")
			}
			delete(s.clients, c.id)
			s.Broadcast()

		case m := <-s.msgCh:
			c, ok := s.clients[m.Id]
			if ok {
				s.ProcessMessage(c, m)
			}
		}
	}
}

func (s *Server) ProcessMessage(c *Client, m *Message) {
	switch {
	case c.status == statusLogin && m.Command == "login" && m.Name != "":
		c.status = statusIdle
		c.name = m.Name
		s.Broadcast()

	case c.status == statusIdle && m.Command == "opponent":
		o, ok := s.clients[m.Opponent]
		if ok {
			c.opponent = o.id
			c.oppName = o.name
			if o.opponent == c.id {
				c.status = statusNumber1
				o.status = statusNumber1
				g := NewGame(c, o)
				c.game = g
				o.game = g
			}
			s.Broadcast()
		}

	case c.status == statusNumber1 && m.Command == "number":
		if GoodNumber(m.Number) {
			c.SetNumber(m.Number)
			s.Broadcast()
		}

	case c.status == statusPlay && m.Command == "answer" && c.game.Turn(c):
		if GoodNumber(m.Number) {
			c.game.AddAnswer(c, m.Number)
			s.Broadcast()
		}

	case c.status == statusDone && m.Command == "done":
		c.status = statusIdle
		c.message = ""
		s.Broadcast()
	}
}

func (s *Server) Broadcast() {
	ps := NewPlayers(s.clients)
	for _, c := range s.clients {
		m := NewMessage(c, ps)
		go c.ws.Send(m)
	}
}

func (s *Server) WebSocketHandler() websocket.Handler {
	return func(ws *websocket.Conn) {
		c := NewClient(s, WSConn{ws})
		s.addCh <- c
		c.Start()
	}
}
