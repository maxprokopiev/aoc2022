package main

import (
	"sort"
)

type operation struct {
	cmd string
	n   int
}

type test struct {
	n int
	b map[bool]int
}

type monkey struct {
	items []int
	op    operation
	test  test
}

func NewWorryLevel(lcm int, old int, op operation) int {
	var l int
	switch op.cmd {
	case "mult":
		l = old * op.n
	case "plus":
		l = old + op.n
	case "square":
		l = old * old
	}
	//return (l % lcm) / 3
	return (l % lcm)
}

func Shift(a []int) []int {
	result := make([]int, 0)
	for i := 1; i < len(a); i++ {
		result = append(result, a[i])
	}

	return result
}

func Solve1(ms []*monkey, rounds int) int {
	inspects := make(map[int]int)

	//lcm := 23 * 19 * 13 * 17
	lcm := 11 * 2 * 5 * 17 * 19 * 7 * 3 * 13

	for t := 1; t <= rounds; t++ {
		for i, m := range ms {
			for {
				if len(m.items) == 0 {
					break
				}
				inspects[i]++
				oldL := m.items[0]
				m.items = Shift(m.items)

				newL := NewWorryLevel(lcm, oldL, m.op)
				passTo := m.test.b[newL%m.test.n == 0]

				ms[passTo].items = append(ms[passTo].items, newL)
			}
		}
	}

	counts := make([]int, 0)
	for _, c := range inspects {
		counts = append(counts, c)
	}
	sort.Ints(counts)

	return counts[len(counts)-1] * counts[len(counts)-2]
}

func main() {
}
