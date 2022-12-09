package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Solve2(input string) int {
	m := BuildMap(input)
	maxScore := 0

	for i := 1; i < len(m)-1; i++ {
		for j := 1; j < len(m[i])-1; j++ {
			tree := m[i][j]
			score := 1

			dUp := 0
			for u := i - 1; u >= 0; u-- {
				dUp++
				if tree <= m[u][j] {
					break
				}
			}
			score *= dUp

			dDown := 0
			for d := i + 1; d <= len(m)-1; d++ {
				dDown++
				if tree <= m[d][j] {
					break
				}
			}
			score *= dDown

			dLeft := 0
			for l := j - 1; l >= 0; l-- {
				dLeft++
				if tree <= m[i][l] {
					break
				}
			}
			score *= dLeft

			dRight := 0
			for r := j + 1; r <= len(m[i])-1; r++ {
				dRight++
				if tree <= m[i][r] {
					break
				}
			}
			score *= dRight
			if score > maxScore {
				maxScore = score
			}
		}
	}

	return maxScore
}

func Solve1(input string) int {
	m := BuildMap(input)

	total := 0
	for i := 1; i < len(m)-1; i++ {
		for j := 1; j < len(m[i])-1; j++ {
			tree := m[i][j]

			vUp := true
			for u := i - 1; u >= 0; u-- {
				if tree <= m[u][j] {
					vUp = false
					break
				}
			}
			if vUp {
				total++
				continue
			}

			vDown := true
			for d := i + 1; d <= len(m)-1; d++ {
				if tree <= m[d][j] {
					vDown = false
					break
				}
			}
			if vDown {
				total++
				continue
			}

			vLeft := true
			for l := j - 1; l >= 0; l-- {
				if tree <= m[i][l] {
					vLeft = false
					break
				}
			}
			if vLeft {
				total++
				continue
			}

			vRight := true
			for r := j + 1; r <= len(m[i])-1; r++ {
				if tree <= m[i][r] {
					vRight = false
					break
				}
			}
			if vRight {
				total++
				continue
			}
		}
	}

	return total + (len(m)*2 - 4) + (len(m[0]) * 2)
}

func BuildMap(input string) [][]int {
	m := make([][]int, 0)

	i := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		m = append(m, make([]int, 0))
		for _, s := range line {
			t, _ := strconv.Atoi(string(s))
			m[i] = append(m[i], t)
		}
		i++
	}
	return m
}

func main() {
	fmt.Println("test")
}
