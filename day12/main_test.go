package main

import (
	"log"
	"os"
	"testing"
)

var testInput = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

func TestSolve1TestInput(t *testing.T) {
	result := Solve1(ReadInput(testInput))
	if result != 31 {
		t.Errorf("got %v, want %v", result, 31)
	}
}

func TestSolve2TestInput(t *testing.T) {
	result := Solve2(ReadInput(testInput))
	if result != 29 {
		t.Errorf("got %v, want %v", result, 29)
	}
}

func TestSolve2Input(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := Solve2(ReadInput(string(input)))
	if result != 399 {
		t.Errorf("got %v, want %v", result, 399)
	}
}

func TestSolve1Input(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := Solve1(ReadInput(string(input)))
	if result != 408 {
		t.Errorf("got %v, want %v", result, 408)
	}
}
