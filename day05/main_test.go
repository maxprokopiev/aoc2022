package main

import (
	"log"
	"os"
	"testing"
)

func TestSolve1TestInput(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

	ss, ms := ReadInput(input)
	result := Solve1(ss, ms, 9000)
	if result != "CMZ" {
		t.Errorf("got %v, want %v", result, "CMZ")
	}
}

func TestSolve2TestInput(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

	ss, ms := ReadInput(input)
	result := Solve1(ss, ms, 9001)
	if result != "MCD" {
		t.Errorf("got %v, want %v", result, "MCD")
	}
}

func TestSolve1Input(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	ss, ms := ReadInput(string(input))
	result := Solve1(ss, ms, 9000)
	if result != "CVCWCRTVQ" {
		t.Errorf("got %v, want %v", result, "CVCWCRTVQ")
	}
}

func TestSolve2Input(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	ss, ms := ReadInput(string(input))
	result := Solve1(ss, ms, 9001)
	if result != "CNSCZWLVT" {
		t.Errorf("got %v, want %v", result, "CNSCZWLVT")
	}
}
