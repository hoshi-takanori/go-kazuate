package main

import (
	"math/rand"
	"time"
)

type Answer struct {
	Number int `json:"number"`
	Exact  int `json:"exact"`
	Near   int `json:"near"`
}

type Game struct {
	player1 *Client
	player2 *Client
	turn    bool

	answers1 []Answer
	answers2 []Answer
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewGame(p1, p2 *Client) *Game {
	return &Game{
		player1:  p1,
		player2:  p2,
		turn:     rand.Intn(2) == 0,
		answers1: []Answer{},
		answers2: []Answer{},
	}
}

func (g *Game) Opponent(c *Client) *Client {
	if c == g.player1 {
		return g.player2
	} else {
		return g.player1
	}
}

func (g *Game) Turn(c *Client) bool {
	if g.turn {
		return c == g.player1
	} else {
		return c == g.player2
	}
}

func (g *Game) Answers(c *Client) []Answer {
	if c == g.player1 {
		return g.answers1
	} else {
		return g.answers2
	}
}
