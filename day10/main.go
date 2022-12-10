package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Draw(i int, x int) {
	p := (i - 1) % 40
	if p == 0 {
		fmt.Println()
	}
	if ((x - 1) == p) || (x == p) || ((x + 1) == p) {
		fmt.Print("🟩")
	} else {
		fmt.Print("⬛")
	}
}

func AddUp(i int, x int) int {
	if (i == 20) || (i == 60) || (i == 100) || (i == 140) || (i == 180) || (i == 220) {
		return i * x
	} else {
		return 0
	}
}

func Solve1(input string) int {
	x := 1
	i := 0

	total := 0
	r := regexp.MustCompile(`(addx|noop)\s?(-?\d+)?`)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		ms := r.FindStringSubmatch(scanner.Text())
		n, _ := strconv.Atoi(ms[2])
		switch ms[1] {
		case "addx":
			for j := 0; j < 2; j++ {
				i++
				total += AddUp(i, x)
				Draw(i, x)
			}
			x += n
		case "noop":
			i++
			total += AddUp(i, x)
			Draw(i, x)
		}
	}

	return total
}

func main() {
}
