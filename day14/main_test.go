package main

import (
	"fmt"
	"os"
	"testing"
)

var testInput = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`

func TestSolve1TestInput(t *testing.T) {
	fmt.Println(Solve1(testInput))
}

func TestSolve2Input(t *testing.T) {
	input, _ := os.ReadFile("input.txt")

	result := Solve1(string(input))
	if result != 26375 {
		t.Errorf("got %v, want %v", result, 26375)
	}
}
