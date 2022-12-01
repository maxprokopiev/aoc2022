package main

import "testing"

func TestSolve1(t *testing.T) {
	elfs, _ := ReadInput()
	result := Solve1(elfs)
	if result != 66487 {
		t.Errorf("got %v", result)
	}
}

func TestSolve2(t *testing.T) {
	elfs, _ := ReadInput()
	result := Solve2(elfs)
	if result != 197301 {
		t.Errorf("got %q", result)
	}
}
