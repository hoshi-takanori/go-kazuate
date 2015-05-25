package main

const (
	statusLogin = iota
	statusIdle
	statusNumber1
	statusNumber2
	statusPlay
	statusDone
)

var statusNames = [...]string{"login", "idle", "num1", "num2", "play", "done"}

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
	Message string `json:"message"`
	Player
	Players []Player `json:"players"`
}

func NewPlayer(c *Client) Player {
	return Player{
		Id:       c.id,
		Name:     c.name,
		Status:   statusNames[c.status],
		Opponent: c.opponent,
	}
}

func NewPlayers(cs map[int]*Client) []Player {
	ps := []Player{}
	for _, c := range cs {
		if c.status == statusIdle {
			ps = append(ps, NewPlayer(c))
		}
	}
	return ps
}

func NewMessage(c *Client, ps []Player) Message {
	m := Message{Player: NewPlayer(c), OppName: c.oppName, Message: c.message}
	if c.status == statusIdle {
		m.Players = ps
	}
	return m
}
