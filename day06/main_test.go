package main

import (
	"log"
	"os"
	"strings"
	"testing"
)

func TestSolve1TestInput(t *testing.T) {
	var result int
	input := map[string]int{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    7,
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      5,
		"nppdvjthqldpwncqszvftbrmjlhg":      6,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 10,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  11,
	}

	for s, i := range input {
		result = Solve1(s, 4)
		if result != i {
			t.Errorf("got %v, want %v", result, i)
		}
	}
}

func TestSolve1Input(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := Solve1(strings.TrimSpace(string(input)), 4)
	if result != 1876 {
		t.Errorf("got %v, want %v", result, 1876)
	}
}

func TestSolve2Input(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := Solve1(strings.TrimSpace(string(input)), 14)
	if result != 2202 {
		t.Errorf("got %v, want %v", result, 2202)
	}
}
