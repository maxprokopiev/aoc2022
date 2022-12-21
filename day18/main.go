package main

import (
	"bufio"
	"strconv"
	"strings"
)

type coords struct {
	x int
	y int
	z int
}

func outside(cc []coords, c coords, threshold int) bool {
	q := make([]coords, 0)
	q = append(q, c)

	cm := make(map[coords]bool)
	for _, e := range cc {
		cm[e] = true
	}

	seen := make(map[coords]bool)
	var e coords
	for {
		if len(q) == 0 {
			break
		}

		e, q = q[0], q[1:]

		if cm[e] {
			continue
		}
		if seen[e] {
			continue
		}
		seen[e] = true
		if len(seen) > threshold {
			return true
		}

		q = append(q, coords{e.x + 1, e.y, e.z})
		q = append(q, coords{e.x - 1, e.y, e.z})
		q = append(q, coords{e.x, e.y + 1, e.z})
		q = append(q, coords{e.x, e.y - 1, e.z})
		q = append(q, coords{e.x, e.y, e.z + 1})
		q = append(q, coords{e.x, e.y, e.z - 1})
	}

	return false
}

func Solve2(cs []coords, threshold int) int {
	total := 0
	for _, e := range cs {
		if outside(cs, coords{e.x + 1, e.y, e.z}, threshold) {
			total += 1
		}
		if outside(cs, coords{e.x - 1, e.y, e.z}, threshold) {
			total += 1
		}
		if outside(cs, coords{e.x, e.y + 1, e.z}, threshold) {
			total += 1
		}
		if outside(cs, coords{e.x, e.y - 1, e.z}, threshold) {
			total += 1
		}
		if outside(cs, coords{e.x, e.y, e.z + 1}, threshold) {
			total += 1
		}
		if outside(cs, coords{e.x, e.y, e.z - 1}, threshold) {
			total += 1
		}
	}

	return total
}

func ParseInput(input string) []coords {
	cs := make([]coords, 0)

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		r := make([]int, 0)
		for _, e := range strings.Split(line, ",") {
			n, _ := strconv.Atoi(e)
			r = append(r, n)
		}
		cs = append(cs, coords{r[0], r[1], r[2]})
	}

	return cs
}

func main() {
}
