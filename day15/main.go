package main

import (
	"bufio"
	"fmt"
	"strings"
)

type coords struct {
	x int
	y int
}

func Abs(x int) int {
	if x < 0 {
		return x * (-1)
	} else {
		return x
	}
}

func md(c1, c2 coords) int {
	return Abs(c1.x-c2.x) + Abs(c1.y-c2.y)
}

func Solve2(m map[coords]coords, max int) int {
	dirs := []coords{coords{-1, 1}, coords{-1, -1}, coords{1, 1}, coords{1, -1}}
	for s, b := range m {
		d := md(s, b)

		for dx := 0; dx <= d; dx++ {
			dy := (d + 1) - dx
			for _, sign := range dirs {
				x := s.x + (dx * sign.x)
				y := s.y + (dy * sign.y)
				if (x < 0) || (x > max) || (y < 0) || (y > max) {
					continue
				}

				valid := true
				for s, b := range m {
					if md(s, coords{x, y}) <= md(s, b) {
						valid = false
						break
					}
				}
				if valid {
					return x*4000000 + y
				}
			}
		}
	}

	return -1
}

func Solve1(m map[coords]coords, y int) int {
	cs := make(map[coords]int)

	for x := -1000000; x <= 7000000; x++ {
		for s, b := range m {
			c := coords{x, y}
			if (c != b) && (md(s, c) <= md(s, b)) {
				cs[c] += 1
			}
		}
	}

	return len(cs)
}

func ParseInput(input string) map[coords]coords {
	m := make(map[coords]coords)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		c1, c2 := ParseSensor(line)
		m[c1] = c2
	}

	return m
}

func ParseSensor(input string) (coords, coords) {
	c1 := coords{x: 0, y: 0}
	c2 := coords{x: 0, y: 0}

	fmt.Sscanf(input,
		"Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
		&c1.x,
		&c1.y,
		&c2.x,
		&c2.y,
	)

	return c1, c2
}

func main() {
}
