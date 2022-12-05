package main

import (
	"regexp"
	"strconv"
	"strings"
)

func Reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func Solve1(ss [][]string, ms [][]int, v int) string {
	for _, m := range ms {
		n, from, to := m[0], m[1]-1, m[2]-1

		newTo := make([]string, 0)
		if v == 9000 {
			newTo = append(newTo, Reverse(ss[from][:n])...)
		} else if v == 9001 {
			newTo = append(newTo, ss[from][:n]...)

		}
		newTo = append(newTo, ss[to]...)
		ss[to] = newTo

		newFrom := make([]string, len(ss[from][n:]))
		copy(newFrom, ss[from][n:])
		ss[from] = newFrom

	}

	result := ""
	for _, s := range ss {
		result += s[0]
	}

	return result
}

func ReadInput(input string) ([][]string, [][]int) {
	var r = regexp.MustCompile(`(?:(\[[A-Z]\])|\s{3})(?:\s|$)`)
	crates := strings.Split(input, "\n\n")[0]
	ss := make([][]string, 0)
	for _, tier := range strings.Split(crates, "\n") {
		if string(tier[1]) == "1" {
			break
		}

		stacks := r.FindAll([]byte(tier), -1)
		for j, crate := range stacks {
			if len(ss) <= j {
				ss = append(ss, make([]string, 0))
			}

			if string(crate[0]) != " " {
				ss[j] = append(ss[j], string(crate[1]))
			}
		}
	}

	var mr = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	moves := strings.Split(input, "\n\n")[1]
	ms := make([][]int, 0)
	for _, move := range strings.Split(moves, "\n") {
		if len(mr.FindStringSubmatch(move)) == 0 {
			break
		}

		ns := mr.FindStringSubmatch(move)[1:]
		c, _ := strconv.Atoi(ns[0])
		i, _ := strconv.Atoi(ns[1])
		j, _ := strconv.Atoi(ns[2])

		ms = append(ms, []int{c, i, j})
	}

	return ss, ms
}

func main() {
}
