package main

import (
	"log"
	"os"
	"testing"
)

var testInput = `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II`

func TestSolve1TestInput(t *testing.T) {
	m := ParseInput(testInput)
	result := Solve1(m)
	if result != 1651 {
		t.Errorf("got %v, want %v", result, 1651)
	}
}

func TestSolve2TestInput(t *testing.T) {
	m := ParseInput(testInput)
	result := Solve2(m)
	if result != 1707 {
		t.Errorf("got %v, want %v", result, 1707)
	}
}

// 1209 is too low
func TestSolve2Input(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := Solve2(ParseInput(string(input)))
	if result != 2261 {
		t.Errorf("got %v, want %v", result, 2261)
	}
}

func TestSolve1Input(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := Solve1(ParseInput(string(input)))
	if result != 1641 {
		t.Errorf("got %v, want %v", result, 1641)
	}
}
