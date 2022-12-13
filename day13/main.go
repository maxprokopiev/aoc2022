package main

import (
	"sort"
	"strconv"
	"strings"
)

func Solve2(input string) int {
	packets := make([]interface{}, 0)

	packets = append(packets, ParsePacket("[[2]]"))
	packets = append(packets, ParsePacket("[[6]]"))

	for _, p := range strings.Split(input, "\n\n") {
		pairs := strings.Split(p, "\n")
		l := ParsePacket(pairs[0])
		r := ParsePacket(pairs[1])
		packets = append(packets, l)
		packets = append(packets, r)
	}

	sort.SliceStable(packets, func(i, j int) bool {
		pi, _ := packets[i].([]interface{})
		pj, _ := packets[j].([]interface{})

		r := IsRightOrder(pi, pj)

		return r > 0
	})

	total := 1
	for i, p := range packets {
		a, _ := p.([]interface{})
		if len(a) == 1 {
			a0, ok := a[0].([]interface{})
			if ok && (len(a0) == 1) {
				if (a0[0] == 2) || (a0[0] == 6) {
					total *= i + 1
				}
			}
		}
	}

	return total
}

func Solve1(input string) int {
	total := 0
	for i, p := range strings.Split(input, "\n\n") {
		pairs := strings.Split(p, "\n")
		l := ParsePacket(pairs[0])
		r := ParsePacket(pairs[1])
		res := IsRightOrder(l, r)
		if res == 1 {
			total += i + 1
		}
	}
	return total
}

func IsRightOrder(l, r []interface{}) int {
	i := 0
	for {
		if len(l) == i {
			if len(l) == len(r) {
				return 0
			}
			return 1
		} else if len(r) == i {
			return -1
		}
		lei, leInt := l[i].(int)
		rei, reInt := r[i].(int)
		lel, leList := l[i].([]interface{})
		rel, reList := r[i].([]interface{})
		i++

		if leInt && reInt {
			if lei == rei {
				continue
			} else if lei < rei {
				return 1
			} else if lei > rei {
				return -1
			}
		} else if leList && reList {
			return IsRightOrder(lel, rel)
		} else if leList && reInt {
			l := make([]interface{}, 0)
			l = append(l, rei)
			return IsRightOrder(lel, l)
		} else if leInt && reList {
			l := make([]interface{}, 0)
			l = append(l, lei)
			return IsRightOrder(l, rel)
		} else {
			panic("boom")
		}
	}
	panic("boom")
	return 0
}

func Pop(s [][]interface{}) [][]interface{} {
	r := make([][]interface{}, 0)

	for _, i := range s[:len(s)-1] {
		r = append(r, i)
	}

	return r
}

func ParsePacket(input string) []interface{} {
	s := make([][]interface{}, 0)
	var n string

	for _, i := range input {
		switch string(i) {
		case "[":
			newL := make([]interface{}, 0)
			s = append(s, newL)
		case "]":
			if len(n) > 0 {
				num, _ := strconv.Atoi(n)
				n = ""
				s[len(s)-1] = append(s[len(s)-1], num)
			}
			if len(s) == 1 {
				return s[0]
			} else {
				curL := s[len(s)-1]
				s = Pop(s)
				s[len(s)-1] = append(s[len(s)-1], curL)
			}
		case ",":
			if len(n) > 0 {
				num, _ := strconv.Atoi(n)
				n = ""
				s[len(s)-1] = append(s[len(s)-1], num)
			}
		default:
			n += string(i)
		}
	}

	return s[0]
}

func main() {
}
