package main

import (
	"bufio"
	"regexp"
	"strconv"
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

func IsAdjacent(h *coords, t *coords) bool {
	return (Abs(h.x-t.x) < 2) && (Abs(h.y-t.y) < 2)
}

var r = regexp.MustCompile(`(R|L|D|U) (\d+)`)

func Solve2(input string, size int) int {
	knots := make([]coords, 0)
	for i := 0; i <= size-1; i++ {
		knots = append(knots, coords{x: 0, y: 0})
	}
	m := map[coords]int{knots[size-1]: 1}

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		ms := r.FindStringSubmatch(scanner.Text())
		steps, _ := strconv.Atoi(ms[2])

		for i := 1; i <= steps; i++ {
			h := &knots[0]
			switch ms[1] {
			case "R":
				h.x += 1
			case "L":
				h.x -= 1
			case "U":
				h.y += 1
			case "D":
				h.y -= 1
			}

			for cur := 0; cur <= size-2; cur++ {
				h = &knots[cur]
				t := &knots[cur+1]

				if IsAdjacent(h, t) {
					continue
				} else {
					if h.x == t.x {
						if h.y > t.y {
							t.y += 1
						} else {
							t.y -= 1
						}
					} else if h.y == t.y {
						if h.x > t.x {
							t.x += 1
						} else {
							t.x -= 1
						}
					} else {
						if h.x > t.x {
							t.x += 1
						} else {
							t.x -= 1
						}

						if h.y > t.y {
							t.y += 1
						} else {
							t.y -= 1
						}
					}
				}

				if cur == size-2 {
					m[knots[size-1]] += 1
				}
			}
		}
	}

	return len(m)
}

func main() {
}
