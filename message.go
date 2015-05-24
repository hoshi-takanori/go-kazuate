package main

const (
	statusLogin = iota
	statusIdle
	statusPlay
)

var statusNames = [...]string{"login", "idle", "play"}

type Player struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	Opponent int    `json:"opponent"`
}

type Message struct {
	Command string `json:"command"`
	Player
	Players []Player `json:"players"`
}

func (c *Client) ToPlayer() Player {
	return Player{
		Id:       c.id,
		Name:     c.name,
		Status:   statusNames[c.status],
		Opponent: c.opponent,
	}
}
