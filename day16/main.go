package main

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

type valve struct {
	rate   int
	valves []int
}

func Solve1(m []valve) int {
	return Search(make(map[int]been), m, 0, 0, 30, 0)
}

func Solve2(m []valve) int {
	return Search(make(map[int]been), m, 0, 0, 26, 1)
}

type been struct {
	n int
	f bool
}

func Search(mm map[int]been, m []valve, current int, opened int, time int, pp int) int {
	if time == 0 {
		if pp > 0 {
			return Search(mm, m, 0, opened, 26, 0)
		} else {
			return 0
		}
	}

	key := opened*len(m)*31*2 + current*31*2 + time*2 + pp
	if mm[key].f && (mm[key].n >= 0) {
		return mm[key].n
	}

	mx := 0
	notOpen := (opened & (1 << current)) == 0
	if (m[current].rate > 0) && notOpen {
		s := (time-1)*m[current].rate + Search(mm, m, current, opened|(1<<current), time-1, pp)
		if s > mx {
			mx = s
		}
	}

	for _, v := range m[current].valves {
		s := Search(mm, m, v, opened, time-1, pp)
		if s > mx {
			mx = s
		}
	}

	mm[key] = been{n: mx, f: true}

	return mx
}

func ParseInput(input string) []valve {
	m := make(map[string]int)

	r := regexp.MustCompile(`Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? (.+)`)

	i := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		ms := r.FindStringSubmatch(line)[1:]
		m[ms[0]] = i
		i++
	}

	nm := make([]valve, 0)
	scanner = bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		ms := r.FindStringSubmatch(line)[1:]

		rate, _ := strconv.Atoi(ms[1])

		valves := make([]int, 0)
		for _, v := range strings.Split(ms[2], ", ") {
			valves = append(valves, m[v])
		}

		nm = append(nm, valve{rate, valves})
	}

	return nm
}

func main() {
}
