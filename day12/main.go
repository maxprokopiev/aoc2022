package main

import (
	"strings"
)

type coords struct {
	x int
	y int
}

func IsInsideMap(s coords, mx, my int) bool {
	return (s.x >= 0) && (s.x < mx) && (s.y >= 0) && (s.y < my)
}

func CanClimb(hm [][]int, currentPos coords, nextPos coords) bool {
	from := hm[currentPos.y][currentPos.x]
	to := hm[nextPos.y][nextPos.x]

	return ((from <= to) && (to-from <= 1)) || (from >= to)
}

func PrepareDirs(hm [][]int, currentPos coords) []coords {
	r := make([]coords, 0)

	cs := []coords{
		coords{currentPos.x + 1, currentPos.y},
		coords{currentPos.x, currentPos.y + 1},
		coords{currentPos.x - 1, currentPos.y},
		coords{currentPos.x, currentPos.y - 1},
	}

	for _, possiblePos := range cs {
		if IsInsideMap(possiblePos, len(hm[0]), len(hm)) &&
			CanClimb(hm, currentPos, possiblePos) {

			r = append(r, possiblePos)
		}
	}

	return r
}

func Shift(a [][]int) [][]int {
	r := make([][]int, 0)
	for i := 1; i < len(a); i++ {
		r = append(r, a[i])
	}

	return r
}

func bfs(q [][]int, hm [][]int, finalPos coords) int {
	var been = make(map[coords]int)

	for {
		if len(q) == 0 {
			break
		}

		v := q[0]
		x, y, d := v[0], v[1], v[2]
		q = Shift(q)

		c := coords{x, y}
		if been[c] == 1 {
			continue
		}
		been[c]++

		if c == finalPos {
			return d
		}

		dirs := PrepareDirs(hm, c)
		for _, ds := range dirs {
			q = append(q, []int{ds.x, ds.y, d + 1})
		}
	}
	return -1
}

func Solve2(hm [][]int, sy, sx int, fy, fx int) int {
	f := coords{fx, fy}

	var q = make([][]int, 0)
	for y, r := range hm {
		for x, h := range r {
			if h == int('a') {
				q = append(q, []int{x, y, 0})
			}
		}
	}

	return bfs(q, hm, f)
}

func Solve1(hm [][]int, sy, sx int, fy, fx int) int {
	s := coords{sx, sy}
	f := coords{fx, fy}

	var q = make([][]int, 0)
	q = append(q, []int{s.x, s.y, 0})

	return bfs(q, hm, f)
}

func ReadInput(input string) ([][]int, int, int, int, int) {
	hm := make([][]int, 0)
	var i, j, sx, sy, fx, fy int
	for _, s := range strings.Split(input, "\n") {
		if s == "" {
			break
		}
		hm = append(hm, make([]int, 0))
		j = 0
		for _, h := range s {
			if int(h) == int('S') {
				sx, sy = i, j
				hm[len(hm)-1] = append(hm[len(hm)-1], int('a'))
			} else if int(h) == int('E') {
				fx, fy = i, j
				hm[len(hm)-1] = append(hm[len(hm)-1], int('z'))
			} else {
				hm[len(hm)-1] = append(hm[len(hm)-1], int(h))
			}
			j++
		}
		i++
	}

	return hm, sx, sy, fx, fy
}

func main() {
}
