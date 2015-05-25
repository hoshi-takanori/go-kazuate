// +build message_test

package main

import (
	"encoding/json"
	"testing"
)

func TestPlayer(t *testing.T) {
	p := Player{Id: 123, Name: "abc", Status: "idle", Opponent: 456}
	var q Player
	MarshalUnmarshal("p", p, &q)
	PrintPlayer("q", q)
}

func TestMessage(t *testing.T) {
	m := Message{
		Command: "test",
		Player:  Player{Id: 123, Name: "abc", Status: "idle"},
		Players: []Player{},
	}
	m.Players = append(m.Players, Player{Id: 456, Name: "def", Status: "login"})
	m.Players = append(m.Players, Player{Id: 789, Name: "xyz", Status: "unknown"})
	var n Message
	MarshalUnmarshal("m", m, &n)
	PrintPlayer("n", n.Player)
	println("len(n.Players) =", len(n.Players))
	PrintPlayer("n.Players[0]", n.Players[0])
	PrintPlayer("n.Players[1]", n.Players[1])
}

func MarshalUnmarshal(name string, p interface{}, q interface{}) {
	buf, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	println(name, "=", string(buf))

	err = json.Unmarshal(buf, q)
	if err != nil {
		panic(err)
	}
}

func PrintPlayer(name string, p Player) {
	println(name+".Id =", p.Id)
	println(name+".Name =", p.Name)
	println(name+".Status =", p.Status)
	println(name+".Opponent =", p.Opponent)
}
