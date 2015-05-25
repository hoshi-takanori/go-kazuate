// +build server_test

package main

import (
	"encoding/json"
	"runtime"
	"testing"
)

type TestConn struct {
}

func (c TestConn) Receive(v interface{}) error {
	return nil
}

func (c TestConn) Send(v interface{}) error {
	buf, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	println(string(buf))
	return nil
}

func TestServer(t *testing.T) {
	s := NewServer()
	go s.Start()

	c1 := NewClient(s, TestConn{})
	s.addCh <- c1
	<-c1.startCh
	close(c1.startCh)

	println("player1 login")
	s.msgCh <- &Message{Command: "login", Player: Player{Id: c1.id, Name: "player1"}}
	runtime.Gosched()
	println()

	c2 := NewClient(s, TestConn{})
	s.addCh <- c2
	<-c2.startCh
	close(c2.startCh)

	println("player2 login")
	s.msgCh <- &Message{Command: "login", Player: Player{Id: c2.id, Name: "player2"}}
	runtime.Gosched()
	println()

	println("player1 choose opponent")
	s.msgCh <- &Message{Command: "opponent", Player: Player{Id: c1.id, Opponent: c2.id}}
	runtime.Gosched()
	println()

	println("player2 choose opponent")
	s.msgCh <- &Message{Command: "opponent", Player: Player{Id: c2.id, Opponent: c1.id}}
	runtime.Gosched()
	println()

	println("player1 decide number")
	s.msgCh <- &Message{Command: "number", Number: 123, Player: Player{Id: c1.id}}
	runtime.Gosched()
	println()

	println("player2 decide number")
	s.msgCh <- &Message{Command: "number", Number: 456, Player: Player{Id: c2.id}}
	runtime.Gosched()
	println()

	g := c1.game
	if g == nil {
		panic("no game")
	}

	var c *Client
	if g.Turn(c1) {
		c = c1
	} else {
		c = c2
	}

	println("player", c.id, "answer")
	s.msgCh <- &Message{Command: "answer", Number: 789, Player: Player{Id: c.id}}
	runtime.Gosched()
	println()

	println("player1 disconnect")
	s.delCh <- c1
	runtime.Gosched()
	println()

	println("player2 disconnect")
	s.delCh <- c2
	runtime.Gosched()
	println()
}
