package main

import (
	"log"
	"os"
	"testing"
)

var testInput = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func TestBuildFs(t *testing.T) {
	root := BuildFs(testInput)
	dExtSize := root.children["d"].children["d.ext"].size
	if dExtSize != 5626152 {
		t.Errorf("got %v, want %v", dExtSize, 5626152)
	}
}

func TestSolve2TestInput(t *testing.T) {
	result := Solve2(testInput)
	if result != 24933642 {
		t.Errorf("got %v, want %v", result, 24933642)
	}
}

func TestSolve2Input(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := Solve2(string(input))
	if result != 4370655 {
		t.Errorf("got %v, want %v", result, 4370655)
	}
}

func TestSolve1TestInput(t *testing.T) {
	result := Solve1(testInput)
	if result != 95437 {
		t.Errorf("got %v, want %v", result, 95437)
	}
}

func TestSolve1Input(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := Solve1(string(input))
	if result != 1783610 {
		t.Errorf("got %v, want %v", result, 1783610)
	}
}
