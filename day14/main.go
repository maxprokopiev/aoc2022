package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func min(x, y int) int {
	if x >= y {
		return y
	} else {
		return x
	}
}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func PrintMap(m [][]string) {
	for _, r := range m {
		for _, e := range r {
			fmt.Print(e)
		}
		fmt.Println()
	}
}

func Solve1(input string) int {
	m := make([][]string, 0)
	// 168 x 75
	lenX := 355
	for i := 0; i < 168; i++ {
		m = append(m, make([]string, 0))
		for j := 0; j < lenX; j++ {
			m[i] = append(m[i], ".")
		}
	}

	m = append(m, make([]string, 0))
	for j := 0; j < lenX; j++ {
		m[len(m)-1] = append(m[len(m)-1], ".")
	}

	m = append(m, make([]string, 0))
	for j := 0; j < lenX; j++ {
		m[len(m)-1] = append(m[len(m)-1], "#")
	}
	// 470 x 544

	minX := 330
	maxX := 684

	mnx := 600
	mxx := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		pairs := strings.Split(line, " -> ")
		for i := 0; i < len(pairs)-1; i++ {
			cs := strings.Split(pairs[i], ",")
			x1, _ := strconv.Atoi(cs[0])
			y1, _ := strconv.Atoi(cs[1])
			if x1 < mnx {
				mnx = x1
			}
			if x1 > mxx {
				mxx = x1
			}
			x1 -= minX

			cs2 := strings.Split(pairs[i+1], ",")
			x2, _ := strconv.Atoi(cs2[0])
			y2, _ := strconv.Atoi(cs2[1])
			if x2 < mnx {
				mnx = x2
			}
			if x2 > mxx {
				mxx = x2
			}

			x2 -= minX

			for j := min(x1, x2); j <= max(x1, x2); j++ {
				for k := min(y1, y2); k <= max(y1, y2); k++ {
					m[k][j] = "#"
				}
			}
		}
	}
	fmt.Println(mnx, mxx)
	PrintMap(m)

	x0 := 500 - minX
	y0 := 0

	nx := x0
	ny := y0

	total := 0

	for {
		if ny+1 == 168 {
			//break

			//fmt.Println(nx)

			//PrintMap(m)
		}

		if m[ny+1][nx] == "." {
			ny += 1
		} else {
			if nx-1 < 0 {
				break
			}

			if m[ny+1][nx-1] == "." {
				nx -= 1
				ny += 1
				continue
			}

			if nx+1 > maxX-minX {
				break
			}

			if m[ny+1][nx+1] == "." {
				nx += 1
				ny += 1
			} else {
				m[ny][nx] = "o"
				total += 1
				if (nx == x0) && (ny == y0) {
					fmt.Println("here")

					break
				}
				nx = x0
				ny = y0
			}
		}
	}

	//PrintMap(m)

	return total
}

func main() {
}
