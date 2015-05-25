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
			println("c.id =", c.id, ", o.id =", o.id)
			if o.opponent == c.id {
				c.status = statusNumber1
				o.status = statusNumber1
			}
			s.Broadcast()
		}

	case c.status == statusNumber1 && m.Command == "number":
		x := m.Number / 100
		y := m.Number / 10 % 10
		z := m.Number % 10
		println("x =", x, ", y =", y, ", z =", z)
		if x != y && y != z && z != x {
			c.status = statusNumber2
			c.number = m.Number
			s.Broadcast()
		}
	}
}

func (s *Server) Broadcast() {
	players := []Player{}
	for _, c := range s.clients {
		players = append(players, c.ToPlayer())
	}
	for _, c := range s.clients {
		m := Message{Player: c.ToPlayer(), OppName: c.oppName, Players: players}
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
