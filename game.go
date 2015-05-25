package main

import (
	"math/rand"
)

type Game struct {
	player1 *Client
	player2 *Client
	turn    bool
}

func NewGame(p1, p2 *Client) *Game {
	return &Game{player1: p1, player2: p2, turn: rand.Intn(2) == 0}
}

func (g *Game) Opponent(c *Client) *Client {
	o := g.player1
	if o == c {
		o = g.player2
	}
	return o
}
