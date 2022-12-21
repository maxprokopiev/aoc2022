package main

import (
	"bufio"
	"fmt"
	"strings"
)

type Value struct {
	v   coords
	set bool
}

type coords struct {
	x int
	y int
}

func Print(m map[coords]Value) int {
	var maxX, maxY, minX, minY int
	for _, v := range m {
		if v.set {
			if v.v.x > maxX {
				maxX = v.v.x
			}
			if v.v.x < minX {
				minX = v.v.x
			}
			if v.v.y > maxY {
				maxY = v.v.y
			}
			if v.v.y < minY {
				minY = v.v.y
			}
		}
	}

	fmt.Println(m)
	fmt.Println(minY, maxY, minX, maxX)

	empty := 0
	for i := maxY; i >= minY; i-- {
		for j := minX; j <= maxX; j++ {
			if m[coords{j, i}].set {
				fmt.Print("#")
			} else {
				empty += 1
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	return empty
}

func Alone(m map[coords]Value, e coords) bool {
	dirs := []coords{
		coords{0, 1},
		coords{1, 0},
		coords{1, 1},
		coords{0, -1},
		coords{-1, 0},
		coords{-1, -1},
		coords{-1, 1},
		coords{1, -1},
	}

	for _, d := range dirs {
		if m[coords{e.x + d.x, e.y + d.y}].set {
			return false
		}
	}

	return true
}

var moves = [][]coords{
	[]coords{coords{0, 1}, coords{1, 1}, coords{-1, 1}},    // north
	[]coords{coords{0, -1}, coords{1, -1}, coords{-1, -1}}, // south
	[]coords{coords{-1, 0}, coords{-1, 1}, coords{-1, -1}}, // west
	[]coords{coords{1, 0}, coords{1, 1}, coords{1, -1}},    // east
}

func CanMove(m map[coords]Value, from coords, move []coords) bool {
	for _, c := range move {
		if m[coords{from.x + c.x, from.y + c.y}].set {
			return false
		}
	}

	return true
}

func Round(m map[coords]Value, rn int) bool {
	fmt.Println("round", rn)
	newPos := make(map[coords]int)
	for current, v := range m {
		mi := rn % len(moves)
		move := moves[rn%len(moves)]
		if (v.set && Alone(m, current)) || !v.set {
			//fmt.Println("alone", current)
			continue
		} else {
			j := 0
			canMove := true
			for {
				//fmt.Println(mi, j, canMove, move[0])
				if j == 4 {
					canMove = false
					break
				}

				if CanMove(m, current, move) {
					//fmt.Println(v.v, "can move to", move[0])
					break
				} else {
					//fmt.Println(v.v, "cannot move to", move[0])
					move = moves[(mi+1)%4]
				}
				j++
				mi++
			}
			if canMove {
				//fmt.Println(current, "moving to", move[0])
				newCoords := coords{current.x + move[0].x, current.y + move[0].y}
				newPos[newCoords] += 1
				m[current] = Value{newCoords, true}
			} else {
				//fmt.Println(current, "is not moving")
			}
		}
	}

	// second half of the round
	for c, count := range newPos {
		if count > 1 {
			for k, _ := range m {
				if m[k].set && (m[k].v == c) {
					//fmt.Println(k, "cannot move because multiple moves")
					m[k] = Value{k, true}
				}
			}
		}
	}

	moved := false
	for currentPos, v := range m {
		if v.set && (currentPos != v.v) {
			moved = true
			m[v.v] = Value{v.v, true}
			m[currentPos] = Value{currentPos, false}
		}
	}

	//Print(m)
	return moved
}

func BuildMap(input string) map[coords]Value {
	m := make(map[coords]Value)

	y := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		for x, e := range line {
			if e == '#' {
				m[coords{x, y}] = Value{coords{x, y}, true}
			}
		}
		y--
	}

	return m
}

func main() {
}
