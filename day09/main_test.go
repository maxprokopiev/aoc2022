package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAdjacent(t *testing.T) {
	assert.Equal(
		t,
		true,
		IsAdjacent(&coords{x: 1, y: 1}, &coords{x: 2, y: 1}),
	)

	assert.Equal(
		t,
		true,
		IsAdjacent(&coords{x: 1, y: 2}, &coords{x: 2, y: 1}),
	)

	assert.Equal(
		t,
		true,
		IsAdjacent(&coords{x: 3, y: 3}, &coords{x: 3, y: 3}),
	)

	assert.Equal(
		t,
		false,
		IsAdjacent(&coords{x: 3, y: 3}, &coords{x: 3, y: 5}),
	)
}

var testInput = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

func TestSolve1TestInput(t *testing.T) {
	result := Solve2(testInput, 2)
	if result != 13 {
		t.Errorf("got %v, want %v", result, 13)
	}
}

var testInput2 = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

func TestSolve2TestInput(t *testing.T) {
	result := Solve2(testInput, 10)
	if result != 1 {
		t.Errorf("got %v, want %v", result, 1)
	}

	result = Solve2(testInput2, 10)
	if result != 36 {
		t.Errorf("got %v, want %v", result, 36)
	}
}

func TestSolve2Input(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := Solve2(string(input), 10)
	if result != 2467 {
		t.Errorf("got %v, want %v", result, 2467)
	}
}

func TestSolve1Input(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := Solve2(string(input), 2)
	if result != 5874 {
		t.Errorf("got %v, want %v", result, 5874)
	}
}
