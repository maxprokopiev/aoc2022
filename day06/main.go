package main

import (
	"strings"
)

func Solve1(input string, size int) int {
	stream, marker, i := input, "", 0

	for {
		if strings.Contains(marker, string(stream[i])) {
			_, stream, _ = strings.Cut(stream, string(stream[i]))
			i, marker = 0, ""
		} else {
			marker += string(stream[i])
			if len(marker) == size {
				return strings.Index(input, marker) + size
			}
			i++
		}
	}

	return -1
}

func main() {
}
