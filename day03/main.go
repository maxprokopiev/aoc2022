package main

import (
	"math"
)

func Solve1(bags []string) int {
	total := 0
	for _, bag := range bags {
		otherHalf := len(bag) / 2
		total += Intersect([]string{bag[:otherHalf], bag[otherHalf:]})
	}

	return total
}

func Solve2(bagGroups [][]string) int {
	total := 0
	for _, bags := range bagGroups {
		total += Intersect(bags)
	}

	return total
}

func Intersect(bags []string) int {
	items := make(map[rune]int)
	for i, bag := range bags {
		for _, item := range bag {
			if (items[item] == 0) && (i == 0) {
				items[item] = 1
			} else if items[item] == i {
				items[item] += 1
			}

			if items[item] == len(bags) {
				return Convert(item)
			}
		}
	}

	return -1
}

const lowerCaseOffset int = 96
const upperCaseOffset int = 38

func Convert(item rune) int {
	if int(item) > 90 {
		return int(math.Abs(float64(lowerCaseOffset - int(item))))
	} else {
		return int(math.Abs(float64(upperCaseOffset - int(item))))
	}
}

func main() {
}
