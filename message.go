package main

const (
	statusLogin = iota
	statusIdle
)

var statusNames = [...]string{"login", "idle"}

type Player struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type Message struct {
	Player
	Players []Player `json:"players"`
}

func (c *Client) ToPlayer() Player {
	return Player{Id: c.id, Name: c.name, Status: statusNames[c.status]}
}
