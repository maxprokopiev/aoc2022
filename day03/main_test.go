package main

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func TestSolve1TestInput(t *testing.T) {
	bags := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}

	result := Solve1(bags)
	if result != 157 {
		t.Errorf("got %v, want %v", result, 157)
	}
}

func TestSolve1Input(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer f.Close()

	var bags []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		bags = append(bags, line)
	}

	result := Solve1(bags)
	if result != 7824 {
		t.Errorf("got %v, want %v", result, 7824)
	}
}

func TestIntersect1(t *testing.T) {
	bags := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
	}

	result := Intersect(bags)
	if result != 18 {
		t.Errorf("got %v, want %v", result, 18)
	}
}

func TestIntersect2(t *testing.T) {
	bags := []string{
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}

	result := Intersect(bags)
	if result != 52 {
		t.Errorf("got %v, want %v", result, 52)
	}
}

func TestSolve2TestInput(t *testing.T) {
	bagGroups := [][]string{
		{
			"vJrwpWtwJgWrhcsFMMfFFhFp",
			"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			"PmmdzqPrVvPwwTWBwg",
		},
		{
			"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			"ttgJtRGJQctTZtZT",
			"CrZsJsPPZsGzwwsLwLmpwMDw",
		},
	}

	result := Solve2(bagGroups)
	if result != 70 {
		t.Errorf("got %v, want %v", result, 70)
	}
}

func TestSolve2Input(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer f.Close()

	var bags []string
	var bagGroups [][]string
	i := 1
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		bags = append(bags, line)
		if i%3 == 0 {
			bagGroups = append(bagGroups, bags)
			bags = []string{}
		}
		i += 1
	}

	result := Solve2(bagGroups)
	if result != 2798 {
		t.Errorf("got %v, want %v", result, 2798)
	}
}
