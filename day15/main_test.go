package main

import (
	"log"
	"os"
	"testing"
)

var testInput = `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`

func TestSolve1TestInput(t *testing.T) {
	result := Solve1(ParseInput(testInput), 10)
	if result != 26 {
		t.Errorf("got %v, want %v", result, 26)
	}
}

func TestSolve2TestInput(t *testing.T) {
	result := Solve2(ParseInput(testInput), 20)
	if result != 56000011 {
		t.Errorf("got %v, want %v", result, 56000011)
	}
}

func TestSolve2Input(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := Solve2(ParseInput(string(input)), 4000000)
	if result != 10382630753392 {
		t.Errorf("got %v, want %v", result, 10382630753392)
	}
}

func TestSolve1Input(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := Solve1(ParseInput(string(input)), 2000000)
	if result != 4424278 {
		t.Errorf("got %v, want %v", result, 4424278)
	}
}

func TestHelloWorld(t *testing.T) {
	input := "Sensor at x=2, y=18: closest beacon is at x=-2, y=15"

	c1, c2 := ParseSensor(input)
	expected := coords{x: 2, y: 18}
	if c1 != expected {
		t.Errorf("Expected %v, got %v", expected, c1)
	}

	expected = coords{x: -2, y: 15}
	if c2 != expected {
		t.Errorf("Expected %v, got %v", expected, c2)
	}
}
