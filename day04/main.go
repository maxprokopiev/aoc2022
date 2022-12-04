package main

import (
	"regexp"
	"strconv"
	"strings"
)

func IsIncluded(s1 int, s2 int, s3 int, s4 int) bool {
	return ((s1 <= s3) && (s2 >= s4)) || ((s3 <= s1) && (s4 >= s2))
}

func DoesOverlap(s1 int, s2 int, s3 int, s4 int) bool {
	return ((s1 <= s3) && (s3 <= s2)) || ((s3 <= s1) && (s1 <= s4))
}

func Solve2(input string) int {
	total := 0
	for _, s := range strings.Fields(input) {
		if DoesOverlap(GetSections(s)) {
			total += 1
		}
	}

	return total
}

func Solve1(input string) int {
	total := 0
	for _, s := range strings.Fields(input) {
		if IsIncluded(GetSections(s)) {
			total += 1
		}
	}

	return total
}

var r = regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)

func GetSections(s string) (int, int, int, int) {
	ss := r.FindStringSubmatch(s)

	s1, _ := strconv.Atoi(ss[1])
	s2, _ := strconv.Atoi(ss[2])
	s3, _ := strconv.Atoi(ss[3])
	s4, _ := strconv.Atoi(ss[4])

	return s1, s2, s3, s4
}

func main() {
}
