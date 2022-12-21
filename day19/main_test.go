package main

import (
	"fmt"
	"os"
	"testing"
)

var testInput = `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.`

func TestSolve1TestInput(t *testing.T) {
	bps := ParseInput(testInput)
	fmt.Println(Solve1(bps))
}

func TestSolve2Input(t *testing.T) {
	input, _ := os.ReadFile("input.txt")
	bps := ParseInput(string(input))
	fmt.Println(Solve2(bps)) // 88160
}

func TestSolve1Input(t *testing.T) {
	input, _ := os.ReadFile("input.txt")
	bps := ParseInput(string(input))
	fmt.Println(Solve1(bps)) // 1962
}
