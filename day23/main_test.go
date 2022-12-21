package main

import (
	"fmt"
	"os"
	"testing"
)

var testInput = `.....
..##.
..#..
.....
..##.
.....`

func TestSmallInput(t *testing.T) {
	m := BuildMap(testInput)
	Print(m)
	i := 0
	for {
		moved := Round(m, i)
		i++
		if !moved {
			break
		}
	}
	fmt.Println(i)
}

var testInput2 = `....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..`

func TestBiggerInput(t *testing.T) {
	m := BuildMap(testInput2)
	Print(m)
	i := 0
	for {
		moved := Round(m, i)
		i++
		if !moved {
			break
		}
	}
	fmt.Println(i)
}

func TestRealInput(t *testing.T) {
	input, _ := os.ReadFile("input.txt")
	m := BuildMap(string(input))
	Print(m)
	//for i := 0; i < 10; i++ {
	//	Round(m, i)
	//}
	//fmt.Println(Print(m)) // 4254

	i := 0
	for {
		moved := Round(m, i)
		i++
		if !moved {
			break
		}
	}
	fmt.Println(i) // 992
}
