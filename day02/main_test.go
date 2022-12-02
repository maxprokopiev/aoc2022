package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"
)

func TestSolve1TestInput(t *testing.T) {
	guide := []string{"AY", "BX", "CZ"}

	result := Solve1(guide)
	if result != 15 {
		t.Errorf("got %v, want %v", result, 15)
	}
}

func TestSolve1Input(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer f.Close()

	guide := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		guide = append(guide, strings.Replace(line, " ", "", 1))
	}

	result := Solve1(guide)
	if result != 13005 {
		t.Errorf("got %v, want %v", result, 13005)
	}
}

func TestSolve2TestInput(t *testing.T) {
	guide := []string{"AY", "BX", "CZ"}

	result := Solve2(guide)
	if result != 12 {
		t.Errorf("got %v, want %v", result, 12)
	}
}

func TestSolve2Input(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer f.Close()

	guide := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		guide = append(guide, strings.Replace(line, " ", "", 1))
	}

	result := Solve2(guide)
	if result != 11373 {
		t.Errorf("got %v, want %v", result, 11373)
	}
}
