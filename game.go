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

func GoodNumber(num int) bool {
	n := [...]int{num / 100, num / 10 % 10, num % 10}
	return n[0] != n[1] && n[1] != n[2] && n[2] != n[0]
}

func NewAnswer(ans, num int) Answer {
	a := Answer{Number: ans}
	m := [...]int{ans / 100, ans / 10 % 10, ans % 10}
	n := [...]int{num / 100, num / 10 % 10, num % 10}
	for i, x := range m {
		if x == n[i] {
			a.Exact++
		} else {
			for _, y := range n {
				if x == y {
					a.Near++
				}
			}
		}
	}
	return a
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

func (g *Game) AddAnswer(c *Client, ans int) {
	a := NewAnswer(ans, g.Opponent(c).number)
	if c == g.player1 {
		g.answers1 = append(g.answers1, a)
	} else {
		g.answers2 = append(g.answers2, a)
	}
	g.turn = !g.turn
}
