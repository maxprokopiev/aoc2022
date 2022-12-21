package main

import (
	"os"
	"testing"
)

var testInput = `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5`

func TestSolve2TestInput(t *testing.T) {
	result := Solve2(ParseInput(testInput), 1000)
	if result != 58 {
		t.Errorf("got %v, want %v", result, 58)
	}
}

func TestSolve2Input(t *testing.T) {
	input, _ := os.ReadFile("input.txt")
	result := Solve2(ParseInput(string(input)), 3000)
	if result != 2546 {
		t.Errorf("got %v, want %v", result, 2546)
	}
}

func TestSolve1TestInput(t *testing.T) {
	result := Solve2(ParseInput(testInput), 0)
	if result != 64 {
		t.Errorf("got %v, want %v", result, 64)
	}
}

func TestSolve1Input(t *testing.T) {
	input, _ := os.ReadFile("input.txt")
	result := Solve2(ParseInput(string(input)), 0)
	if result != 4348 {
		t.Errorf("got %v, want %v", result, 4348)
	}
}
