package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type monkey struct {
	n   int
	num bool
	op1 string
	op2 string
	op  string
}

var rop = regexp.MustCompile(`(\w+): (\w+) (.) (\w+)`)
var rnum = regexp.MustCompile(`(\w+): (\d+)`)

func DoOp(x, y int, op string) int {
	switch op {
	case "+":
		return x + y
	case "-":
		return x - y
	case "*":
		return x * y
	case "/":
		return x / y
	default:
		return -1
	}
}

func GetNum(ms map[string]monkey, m string) int {
	if ms[m].op == "" {
		return ms[m].n
	} else {
		n1 := GetNum(ms, ms[m].op1)
		n2 := GetNum(ms, ms[m].op2)
		ms[m] = monkey{
			DoOp(n1, n2, ms[m].op),
			true,
			ms[m].op1,
			ms[m].op2,
			ms[m].op,
		}
		return ms[m].n
	}
}

func HasHumn(ms map[string]monkey, m string) bool {
	if m == "humn" {
		return true
	}

	if ms[m].num {
		return false
	} else {
		if HasHumn(ms, ms[m].op1) {
			return true
		}
		if HasHumn(ms, ms[m].op2) {
			return true
		}
	}

	return false
}

func Solve2(ms map[string]monkey) int {
	if HasHumn(ms, ms["root"].op1) {
		n := GetNum(ms, ms["root"].op2)

		for name, _ := range ms {
			if !HasHumn(ms, name) {
				num := GetNum(ms, name)
				ms[name] = monkey{num, true, "", "", ""}
			}
		}

		i := 3378273350000
		for {
			ms["humn"] = monkey{i, true, "", "", ""}
			n2 := GetNum(ms, ms["root"].op1)
			if n2 == n {
				return i
			} else {
				i++
			}
		}
	}

	return -1
}

func Solve1(ms map[string]monkey) int {
	r := GetNum(ms, "root")
	fmt.Println(ms)
	return r
}

func ParseInput(input string) map[string]monkey {
	monkeys := make(map[string]monkey)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case rop.MatchString(line):
			ms := rop.FindStringSubmatch(line)[1:]
			monkeys[ms[0]] = monkey{0, false, ms[1], ms[3], ms[2]}
		case rnum.MatchString(line):
			ms := rnum.FindStringSubmatch(line)[1:]
			n, _ := strconv.Atoi(ms[1])
			monkeys[ms[0]] = monkey{n, true, "", "", ""}
		}
	}

	return monkeys
}

func main() {
}
