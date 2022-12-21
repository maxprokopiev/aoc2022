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

type Value struct {
	c     coords
	steps int
}

type blz struct {
	c coords
	d coords
	v string
}

//var maxX = 6
//var maxY = 4

var maxX = 120
var maxY = 25

func Print(c coords, bs []blz) {
	for y := -1; y < maxY+1; y++ {
		for x := -1; x < maxX+1; x++ {
			cs := coords{x, y}
			if c == cs {
				fmt.Print("E")
			}
			if (x == -1) || (x == maxX) || (y == -1) || (y == maxY) {
				if c != cs {
					fmt.Print("#")
				}
			} else {
				n := 0
				var v string
				for _, b := range bs {
					cs := coords{x, y}
					if b.c == cs {
						n += 1
						v = b.v
					}
				}
				if n == 0 {
					if c != cs {
						fmt.Print(".")
					}
				} else if n == 1 {
					fmt.Print(v)
				} else {
					fmt.Print(n)
				}
			}
		}
		fmt.Println()
	}
}

func mod(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}

func PreCompile(bs []blz) [][]blz {
	r := make([][]blz, 0)
	r = append(r, bs)

	for j := 0; j < 1000; j++ {
		n := make([]blz, 0)
		for _, b := range r[len(r)-1] {
			newPos := coords{mod(b.c.x+b.d.x, maxX), mod(b.c.y+b.d.y, maxY)}
			n = append(n, blz{newPos, b.d, b.v})
		}
		r = append(r, n)
	}

	return r
}

func Round(cache [][]blz, mins int, been map[coords]int, f coords, v0 coords) int {
	q := make([]Value, 0)
	q = append(q, Value{v0, mins})

	dirs := []coords{
		coords{0, 1},
		coords{0, -1},
		coords{1, 0},
		coords{-1, 0},
	}

	for {
		if len(q) == 0 {
			break
		}

		var value Value
		value, q = q[0], q[1:]

		if (value.steps > 0) && (been[value.c] == value.steps) {
			continue
		}
		been[value.c] = value.steps

		inBlz := false
		for _, b := range cache[value.steps+1] {
			if value.c == b.c {
				inBlz = true
				break
			}
		}
		if !inBlz {
			q = append(q, Value{value.c, value.steps + 1})
		}
		for _, d := range dirs {
			newPos := coords{value.c.x + d.x, value.c.y + d.y}

			if newPos == f {
				Print(newPos, cache[value.steps+2])
				return value.steps + 1
			}

			if (newPos.x <= -1) || (newPos.x >= maxX) ||
				(newPos.y <= -1) || (newPos.y >= maxY) {
				continue
			}

			inBlz := false
			for _, b := range cache[value.steps+1] {
				if newPos == b.c {
					inBlz = true
					break
				}
			}
			if inBlz {
				continue
			}

			q = append(q, Value{newPos, value.steps + 1})
		}
	}

	panic("booom")
	return -1
}

func BuildMap(input string) []blz {
	m := make([]blz, 0)
	y := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		for x, e := range line {
			if (e == '>') || (e == '<') || (e == '^') || (e == 'v') {
				var d coords
				switch e {
				case '>':
					d = coords{1, 0}
				case '<':
					d = coords{-1, 0}
				case 'v':
					d = coords{0, 1}
				case '^':
					d = coords{0, -1}
				}
				m = append(m, blz{coords{x - 1, y - 1}, d, string(e)})
			}
		}
		y++
	}

	return m
}

func main() {
}
