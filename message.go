package main

const (
	statusLogin = iota
	statusIdle
	statusNumber1
	statusNumber2
	statusPlay
)

var statusNames = [...]string{"login", "idle", "num1", "num2", "play"}

type Player struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	Opponent int    `json:"opponent"`
}

type Message struct {
	Command string `json:"command"`
	Number  int    `json:"number"`
	OppName string `json:"opp_name"`
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
