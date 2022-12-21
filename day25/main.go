package main

import (
	"bufio"
	"math"
	"strconv"
	"strings"
)

func Convert(input string) int {
	sum := 0
	for i := len(input) - 1; i >= 0; i-- {
		pow := math.Pow(5, float64(i))
		switch input[len(input)-i-1] {
		case '2':
			pow = pow * 2
		case '1':
			pow = pow * 1
		case '0':
			pow = pow * 0
		case '-':
			pow = pow * (-1)
		case '=':
			pow = pow * (-2)
		}
		sum += int(pow)
	}
	return sum
}

func ConvertBack(x int) string {
	m := false
	r := ""
	for {
		div := x / 5
		mod := x % 5
		x = div

		if m {
			mod += 1
			m = false
		}

		switch mod {
		case 3:
			r = "=" + r
			m = true
		case 4:
			r = "-" + r
			m = true
		case 5:
			r = "0" + r
		default:
			r = strconv.Itoa(mod) + r
		}

		if div == 0 {
			return r
			break
		}
	}

	return ""
}

func Solve1(input string) string {
	total := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		total += Convert(line)
	}
	return ConvertBack(total)
}

func main() {
}
