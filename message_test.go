package main

import (
	"encoding/json"
	"testing"
)

func TestPlayer(t *testing.T) {
	p := Player{Id: 123, Name: "abc", Status: "idle"}
	buf, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	println("p =", string(buf))

	var q Player
	err = json.Unmarshal(buf, &q)
	if err != nil {
		panic(err)
	}
	println("q.Id =", q.Id)
	println("q.Name =", q.Name)
	println("q.Status =", q.Status)
}

func TestMessage(t *testing.T) {
	m := Message{Player{Id: 123, Name: "abc", Status: "idle"}, []Player{}}
	m.Players = append(m.Players, Player{Id: 456, Name: "def", Status: "login"})
	m.Players = append(m.Players, Player{Id: 789, Name: "xyz", Status: "unknown"})
	buf, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	println("m =", string(buf))

	var n Message
	err = json.Unmarshal(buf, &n)
	if err != nil {
		panic(err)
	}
	println("n.Id =", n.Id)
	println("n.Name =", n.Name)
	println("n.Status =", n.Status)
	println("len(n.Players) =", len(n.Players))
	println("n.Players[0].Id =", n.Players[0].Id)
	println("n.Players[0].Name =", n.Players[0].Name)
	println("n.Players[0].Status =", n.Players[0].Status)
	println("n.Players[1].Id =", n.Players[1].Id)
	println("n.Players[1].Name =", n.Players[1].Name)
	println("n.Players[1].Status =", n.Players[1].Status)
}
