package main

import (
	"os"
	"testing"
)

var testInput = `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32`

func TestSolve2TestInput(t *testing.T) {
	result := Solve2(ParseInput(testInput))
	if result != 301 {
		t.Errorf("got %v, want %v", result, 301)
	}
}

func TestSolve1TestInput(t *testing.T) {
	result := Solve1(ParseInput(testInput))
	if result != 152 {
		t.Errorf("got %v, want %v", result, 152)
	}
}

func TestSolve2Input(t *testing.T) {
	input, _ := os.ReadFile("input.txt")
	result := Solve2(ParseInput(string(input)))
	if result != 3378273370680 {
		t.Errorf("got %v, want %v", result, 3378273370680)
	}
}

func TestSolve1Input(t *testing.T) {
	input, _ := os.ReadFile("input.txt")
	result := Solve1(ParseInput(string(input)))
	if result != 331120084396440 {
		t.Errorf("got %v, want %v", result, 331120084396440)
	}
}
