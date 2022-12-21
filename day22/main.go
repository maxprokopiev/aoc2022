package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type coords struct {
	x int
	y int
}

func mod(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}

func BuildEdge(from, to coords) []coords {
	e := make([]coords, 0)
	if from.x == to.x {
		// go through y
		if from.y > to.y {
			// decrease y
			for i := from.y; i >= to.y; i-- {
				e = append(e, coords{from.x, i})
			}
		} else {
			// increase y
			for i := from.y; i <= to.y; i++ {
				e = append(e, coords{from.x, i})
			}
		}
	} else {
		// go through x
		if from.x > to.x {
			// decrease x
			for i := from.x; i >= to.x; i-- {
				e = append(e, coords{i, from.y})
			}
		} else {
			// increase x
			for i := from.x; i <= to.x; i++ {
				e = append(e, coords{i, from.y})
			}
		}
	}
	return e
}

type Edge struct {
	m1 map[coords]coords
	m2 map[coords]coords
	d1 int
	d2 int
}

var edges1 = [][]coords{
	[]coords{coords{9, 1}, coords{12, 1}, coords{4, 5}, coords{1, 5}},
	[]coords{coords{12, 1}, coords{12, 4}, coords{16, 9}, coords{16, 12}},
	[]coords{coords{12, 5}, coords{12, 8}, coords{16, 9}, coords{13, 9}},
	[]coords{coords{13, 12}, coords{16, 12}, coords{1, 8}, coords{1, 5}},
	[]coords{coords{9, 12}, coords{12, 12}, coords{4, 8}, coords{1, 8}},
	[]coords{coords{9, 9}, coords{9, 12}, coords{8, 8}, coords{5, 8}},
	[]coords{coords{5, 5}, coords{8, 5}, coords{9, 1}, coords{9, 4}},
}

var dirs1 = [][]int{
	[]int{1, 1},
	[]int{2, 2},
	[]int{1, 2},
	[]int{0, 3},
	[]int{3, 3},
	[]int{3, 0},
	[]int{0, 1},
}

var edges2 = [][]coords{
	[]coords{coords{51, 1}, coords{100, 1}, coords{1, 151}, coords{1, 200}},
	[]coords{coords{101, 1}, coords{150, 1}, coords{1, 200}, coords{50, 200}},
	[]coords{coords{101, 50}, coords{150, 50}, coords{100, 51}, coords{100, 100}},
	[]coords{coords{150, 1}, coords{150, 50}, coords{100, 150}, coords{100, 101}},
	[]coords{coords{51, 150}, coords{100, 150}, coords{50, 151}, coords{50, 200}},
	[]coords{coords{51, 51}, coords{51, 100}, coords{1, 101}, coords{50, 101}},
	[]coords{coords{51, 1}, coords{51, 50}, coords{1, 150}, coords{1, 101}},
}

var dirs2 = [][]int{
	[]int{0, 1},
	[]int{3, 1},
	[]int{2, 3},
	[]int{2, 2},
	[]int{2, 3},
	[]int{1, 0},
	[]int{0, 0},
}

func BuildMapping(edges [][]coords, dirs [][]int) []Edge {
	r := make([]Edge, 0)
	for j, e := range edges {
		e1 := BuildEdge(e[0], e[1])
		e2 := BuildEdge(e[2], e[3])

		m1 := make(map[coords]coords)
		for i, _ := range e1 {
			m1[e1[i]] = e2[i]
		}

		m2 := make(map[coords]coords)
		for i, _ := range e2 {
			m2[e2[i]] = e1[i]
		}

		d1 := dirs[j][0]
		d2 := dirs[j][1]

		r = append(r, Edge{m1, m2, d1, d2})
	}

	return r
}

func Move2(s coords, m map[coords]int, mm []Edge, path string) int {
	n := ""
	cd := map[rune]int{'R': 1, 'L': -1}
	dirs := []coords{
		coords{1, 0},
		coords{0, 1},
		coords{-1, 0},
		coords{0, -1},
	}
	di := 0
	for i, e := range path {
		if (e == 'R') || (e == 'L') {
			num, _ := strconv.Atoi(n)

			fmt.Println("walking", num, "steps to", dirs[di], "from", s)
			n = ""
			for j := 0; j < num; j++ {
				nc := coords{s.x + dirs[di].x, s.y + dirs[di].y}
				fmt.Println("moving to", nc)
				//fmt.Println(m[nc])
				if m[nc] == 2 {
					s = nc
				} else if m[nc] == 0 {
					// find the next plane of the cube based on mapping
					nc = coords{nc.x - dirs[di].x, nc.y - dirs[di].y}
					oldNc := nc
					oldDi := di
					fmt.Println("at the edge, moving back to", nc)
					for _, e := range mm {
						if e.m1[nc].x > 0 {
							nc = e.m1[nc]
							di = e.d1
							break
						} else if e.m2[nc].x > 0 {
							nc = e.m2[nc]
							di = e.d2
							break
						}
					}
					fmt.Println("found a mapping at", nc, m[nc])
					if m[nc] == 2 {
						s = nc
					} else if m[nc] == 3 {
						s = oldNc
						di = oldDi
						fmt.Println("it's a wall, moving back to before the mapping", s)
					}
				}
			}
			di = mod(di+cd[e], len(dirs))
		} else {
			n += string(e)
			// TODO: last move

			if i == len(path)-1 {
				num, _ := strconv.Atoi(n)

				fmt.Println("walking", num, "steps to", dirs[di], "from", s)
				n = ""
				for j := 0; j < num; j++ {
					nc := coords{s.x + dirs[di].x, s.y + dirs[di].y}
					fmt.Println("moving to", nc)
					//fmt.Println(m[nc])
					if m[nc] == 2 {
						s = nc
					} else if m[nc] == 0 {
						// find the next plane of the cube based on mapping
						nc = coords{nc.x - dirs[di].x, nc.y - dirs[di].y}
						oldNc := nc
						oldDi := di
						fmt.Println("at the edge, moving back to", nc)
						for _, e := range mm {
							if e.m1[nc].x > 0 {
								nc = e.m1[nc]
								di = e.d1
								break
							} else if e.m2[nc].x > 0 {
								nc = e.m2[nc]
								di = e.d2
								break
							}
						}
						fmt.Println("found a mapping at", nc, m[nc])
						if m[nc] == 2 {
							s = nc
						} else if m[nc] == 3 {
							s = oldNc
							di = oldDi
							fmt.Println("it's a wall, moving back to before the mapping", s)
						}
					}
				}
				di = mod(di+cd[e], len(dirs))
			}
		}
	}

	fmt.Println(s, di)
	// 5 7 3
	fmt.Println(s.y, s.x, di)

	return s.y*1000 + s.x*4 + di
}

func Move(m map[coords]int, path string) int {
	s := coords{51, 1}
	n := ""
	cd := map[rune]int{'R': 1, 'L': -1}
	dirs := []coords{
		coords{1, 0},
		coords{0, 1},
		coords{-1, 0},
		coords{0, -1},
	}
	di := 0
	for i, e := range path {
		if (e == 'R') || (e == 'L') {
			num, _ := strconv.Atoi(n)
			n = ""
			for j := 0; j < num; j++ {
				nc := coords{s.x + dirs[di].x, s.y + dirs[di].y}
				if m[nc] == 2 {
					s = nc
				} else if m[nc] == 0 {
					for {
						nc = coords{nc.x - dirs[di].x, nc.y - dirs[di].y}
						if m[nc] == 0 {
							nc = coords{nc.x + dirs[di].x, nc.y + dirs[di].y}
							break
						}
					}
					if m[nc] == 2 {
						s = nc
					}
				}
			}
			di = mod(di+cd[e], len(dirs))
		} else {
			n += string(e)
			if i == len(path)-1 {
				num, _ := strconv.Atoi(n)
				for j := 0; j < num; j++ {
					nc := coords{s.x + dirs[di].x, s.y + dirs[di].y}
					if m[nc] == 2 {
						s = nc
					} else if m[nc] == 0 {
						for {
							nc = coords{nc.x - dirs[di].x, nc.y - dirs[di].y}
							if m[nc] == 0 {
								nc = coords{nc.x + dirs[di].x, nc.y + dirs[di].y}
								break
							}
						}
						if m[nc] == 2 {
							s = nc
						}
					}
				}
				di = mod(di+cd[e], len(dirs))
			}
		}
	}

	return s.y*1000 + s.x*4 + di
}

func BuildMap(input string) map[coords]int {
	m := make(map[coords]int)
	scanner := bufio.NewScanner(strings.NewReader(input))
	y := 0
	for scanner.Scan() {
		y++
		line := scanner.Text()
		x := 0
		for _, c := range line {
			x++
			switch c {
			case '.':
				m[coords{x, y}] = 2
			case '#':
				m[coords{x, y}] = 3
			}
		}
	}

	return m
}

func main() {
}
