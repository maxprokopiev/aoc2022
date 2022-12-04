package main

import (
	"log"
	"os"
	"testing"
)

func TestSolve2TestInput(t *testing.T) {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`

	result := Solve2(input)
	if result != 4 {
		t.Errorf("got %v, want %v", result, 4)
	}
}

func TestDoesOverlap(t *testing.T) {
	result := DoesOverlap(2, 4, 6, 8)
	if result != false {
		t.Errorf("got %v, want %v", result, false)
	}
}

func TestIsIncluded(t *testing.T) {
	result := IsIncluded(6, 6, 4, 6)
	if result != true {
		t.Errorf("got %v, want %v", result, true)
	}
}

func TestSolve1TestInput(t *testing.T) {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`

	result := Solve1(input)
	if result != 2 {
		t.Errorf("got %v, want %v", result, 2)
	}
}

func TestSolve2Input(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := Solve2(string(input))
	if result != 861 {
		t.Errorf("got %v, want %v", result, 861)
	}
}

func TestSolve1Input(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := Solve1(string(input))
	if result != 441 {
		t.Errorf("got %v, want %v", result, 441)
	}
}
