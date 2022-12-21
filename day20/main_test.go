package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestSolve1TestInput(t *testing.T) {
	input := []int{1, 2, -3, 3, -2, 0, 4}

	result := Solve1(input, 1, 1)
	if result != 3 {
		t.Errorf("got %v, want %v", result, 3)
	}
}

func TestSolve2TestInput(t *testing.T) {
	input := []int{1, 2, -3, 3, -2, 0, 4}

	result := Solve1(input, 811589153, 10)
	if result != 1623178306 {
		t.Errorf("got %v, want %v", result, 1623178306)
	}
}

func TestSolve2Input(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	a := make([]int, 0)
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)
		a = append(a, n)
	}

	result := Solve1(a, 811589153, 10)
	if result != 11893839037215 {
		t.Errorf("got %v, want %v", result, 11893839037215)
	}
}

func TestSolve1Input(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	a := make([]int, 0)
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)
		a = append(a, n)
	}

	result := Solve1(a, 1, 1)
	if result != 6640 {
		t.Errorf("got %v, want %v", result, 6640)
	}
}
