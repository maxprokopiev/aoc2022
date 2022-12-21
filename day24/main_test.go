package main

import (
	"fmt"
	"os"
	"testing"
)

// -1
var testInput = `#.#####
#.....#
#>....#
#.....#
#...v.#
#.....#
#####.#`

var testInput2 = `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`

func TestBuildMap(t *testing.T) {
	mins := 0
	bs := BuildMap(testInput2)
	Print(coords{0, -1}, bs)

	cache := PreCompile(bs)

	mins = Round(cache, mins, make(map[coords]int), coords{maxX - 1, maxY}, coords{0, -1})
	fmt.Println(mins) // 18

	mins = Round(cache, mins, make(map[coords]int), coords{0, -1}, coords{maxX - 1, maxY})
	fmt.Println(mins) // 23 for the total of 41

	mins = Round(cache, mins, make(map[coords]int), coords{maxX - 1, maxY}, coords{0, -1})
	fmt.Println(mins) // 13 for the total of 54
}

func TestRealInput(t *testing.T) {
	input, _ := os.ReadFile("input.txt")
	mins := 0
	bs := BuildMap(string(input))
	Print(coords{0, -1}, bs)

	cache := PreCompile(bs)

	mins = Round(cache, mins, make(map[coords]int), coords{maxX - 1, maxY}, coords{0, -1})
	fmt.Println(mins) // 266

	mins = Round(cache, mins, make(map[coords]int), coords{0, -1}, coords{maxX - 1, maxY})
	fmt.Println(mins) // 279

	mins = Round(cache, mins, make(map[coords]int), coords{maxX - 1, maxY}, coords{0, -1})
	fmt.Println(mins) // 289

	// 834 is too low
	// -> 853
}
