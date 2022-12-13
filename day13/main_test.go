package main

import (
	"log"
	"os"
	"testing"
)

var testInput = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`

func TestSolve1TestInput(t *testing.T) {
	result := Solve1(testInput)
	if result != 13 {
		t.Errorf("got %v, want %v", result, 13)
	}
}

func TestSolve2TestInput(t *testing.T) {
	result := Solve2(testInput)
	if result != 140 {
		t.Errorf("got %v, want %v", result, 140)
	}
}

func TestSolve2Input(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := Solve2(string(input))
	if result != 19570 {
		t.Errorf("got %v, want %v", result, 19570)
	}
}

func TestSolve1Input(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := Solve1(string(input))
	if result != 5350 {
		t.Errorf("got %v, want %v", result, 5350)
	}
}
