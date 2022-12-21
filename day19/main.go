package main

import (
	"bufio"
	"fmt"
	"strings"
)

type bp struct {
	id  int
	oo  int
	co  int
	bo  int
	bc  int
	gro int
	grb int
}

type st struct {
	t  int
	o  int
	or int
	c  int
	cr int
	b  int
	br int
	g  int
	gr int
}

func ParseBlueprint(line string) bp {
	pattern := "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian."
	b := bp{}
	fmt.Sscanf(line, pattern, &b.id, &b.oo, &b.co, &b.bo, &b.bc, &b.gro, &b.grb)
	return b
}

func ParseInput(input string) []bp {
	bps := make([]bp, 0)

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		b := ParseBlueprint(line)
		bps = append(bps, b)
	}

	return bps
}

func search(b bp, s st, max int) int {
	if s.t == 0 {
		return s.g
	}

	if s.g+s.t*s.gr+s.t*(s.t-1)/2 <= max {
		return 0
	}

	mx := b.bo
	if mx < b.co {
		mx = b.co
	}
	if mx < b.gro {
		mx = b.gro
	}
	tryOre := s.o >= b.oo && s.or < mx
	tryClay := s.o >= b.co
	tryObsidian := s.o >= b.bo && s.c >= b.bc
	tryGeode := s.o >= b.gro && s.b >= b.grb

	s.t -= 1
	s.o += s.or
	s.c += s.cr
	s.b += s.br
	s.g += s.gr

	if tryGeode {
		ns := s
		ns.o -= b.gro
		ns.b -= b.grb
		ns.gr += 1
		r := search(b, ns, max)
		if r > max {
			max = r
		}
	}

	if tryObsidian {
		ns := s
		ns.o -= b.bo
		ns.c -= b.bc
		ns.br += 1
		r := search(b, ns, max)
		if r > max {
			max = r
		}
	}

	if tryClay {
		ns := s
		ns.o -= b.co
		ns.cr += 1
		r := search(b, ns, max)
		if r > max {
			max = r
		}
	}

	if tryOre {
		ns := s
		ns.o -= b.oo
		ns.or += 1
		r := search(b, ns, max)
		if r > max {
			max = r
		}
	}

	ns := s
	r := search(b, ns, max)
	if r > max {
		max = r
	}
	return max
}

func Solve2(bps []bp) int {
	total := 1
	dst := st{32, 0, 1, 0, 0, 0, 0, 0, 0}
	for _, b := range bps[0:3] {
		s := dst
		total *= search(b, s, 0)
	}
	return total
}

func Solve1(bps []bp) int {
	total := 0
	dst := st{24, 0, 1, 0, 0, 0, 0, 0, 0}
	for _, b := range bps {
		s := dst
		total += search(b, s, 0) * b.id
	}
	return total
}

func main() {
}
