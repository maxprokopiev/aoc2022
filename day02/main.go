package main

func CalculateTotal(scores map[string]int, guide []string) int {
	total := 0
	for _, round := range guide {
		total += scores[round]
	}

	return total
}

func Solve2(guide []string) int {
	var scores = map[string]int{
		"AX": 3,
		"AY": 4,
		"AZ": 8,
		"BX": 1,
		"BY": 5,
		"BZ": 9,
		"CX": 2,
		"CY": 6,
		"CZ": 7,
	}

	return CalculateTotal(scores, guide)
}

func Solve1(guide []string) int {
	var scores = map[string]int{
		"AX": 4,
		"AY": 8,
		"AZ": 3,
		"BX": 1,
		"BY": 5,
		"BZ": 9,
		"CX": 7,
		"CY": 2,
		"CZ": 6,
	}

	return CalculateTotal(scores, guide)
}

func main() {
}
