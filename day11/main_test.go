package main

import (
	"testing"
)

var testInput = []*monkey{
	&monkey{
		items: []int{79, 98},
		op:    operation{cmd: "mult", n: 19},
		test:  test{n: 23, b: map[bool]int{true: 2, false: 3}},
	},
	&monkey{
		items: []int{54, 65, 75, 74},
		op:    operation{cmd: "plus", n: 6},
		test:  test{n: 19, b: map[bool]int{true: 2, false: 0}},
	},
	&monkey{
		items: []int{79, 60, 97},
		op:    operation{cmd: "square"},
		test:  test{n: 13, b: map[bool]int{true: 1, false: 3}},
	},
	&monkey{
		items: []int{74},
		op:    operation{cmd: "plus", n: 3},
		test:  test{n: 17, b: map[bool]int{true: 0, false: 1}},
	},
}

func TestSolve2TestInput(t *testing.T) {
	result := Solve1(testInput, 10000)
	if result != 2713310158 {
		t.Errorf("got %v, want %v", result, 2713310158)
	}
}

func TestSolve1TestInput(t *testing.T) {
	result := Solve1(testInput, 20)
	if result != 10605 {
		t.Errorf("got %v, want %v", result, 10605)
	}
}

var input = []*monkey{
	&monkey{
		items: []int{92, 73, 86, 83, 65, 51, 55, 93},
		op:    operation{cmd: "mult", n: 5},
		test:  test{n: 11, b: map[bool]int{true: 3, false: 4}},
	},
	&monkey{
		items: []int{99, 67, 62, 61, 59, 98},
		op:    operation{cmd: "square"},
		test:  test{n: 2, b: map[bool]int{true: 6, false: 7}},
	},
	&monkey{
		items: []int{81, 89, 56, 61, 99},
		op:    operation{cmd: "mult", n: 7},
		test:  test{n: 5, b: map[bool]int{true: 1, false: 5}},
	},
	&monkey{
		items: []int{97, 74, 68},
		op:    operation{cmd: "plus", n: 1},
		test:  test{n: 17, b: map[bool]int{true: 2, false: 5}},
	},
	&monkey{
		items: []int{78, 73},
		op:    operation{cmd: "plus", n: 3},
		test:  test{n: 19, b: map[bool]int{true: 2, false: 3}},
	},
	&monkey{
		items: []int{50},
		op:    operation{cmd: "plus", n: 5},
		test:  test{n: 7, b: map[bool]int{true: 1, false: 6}},
	},
	&monkey{
		items: []int{95, 88, 53, 75},
		op:    operation{cmd: "plus", n: 8},
		test:  test{n: 3, b: map[bool]int{true: 0, false: 7}},
	},
	&monkey{
		items: []int{50, 77, 98, 85, 94, 56, 89},
		op:    operation{cmd: "plus", n: 2},
		test:  test{n: 13, b: map[bool]int{true: 4, false: 0}},
	},
}

func TestSolve2Input(t *testing.T) {
	result := Solve1(input, 10000)
	if result != 39109444654 {
		t.Errorf("got %v, want %v", result, 39109444654)
	}
}

func TestSolve1Input(t *testing.T) {
	result := Solve1(input, 20)
	if result != 120756 {
		t.Errorf("got %v, want %v", result, 120756)
	}
}
