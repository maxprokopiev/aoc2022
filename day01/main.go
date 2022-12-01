package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Solve1(elfs [][]int) int {
	var sums []int
	for _, elf := range elfs {
		sum := 0
		for _, calories := range elf {
			sum += calories
		}
		sums = append(sums, sum)
	}

	max := sums[0]
	for _, sum := range sums {
		if max < sum {
			max = sum
		}
	}

	return max
}

func Solve2(elfs [][]int) int {
	var sums []int
	for _, elf := range elfs {
		sum := 0
		for _, calories := range elf {
			sum += calories
		}
		sums = append(sums, sum)
	}

	for i := 0; i < len(sums); i++ {
		for j := i + 1; j < len(sums); j++ {
			if sums[i] < sums[j] {
				tmp := sums[i]
				sums[i] = sums[j]
				sums[j] = tmp
			}
		}
	}

	return sums[0] + sums[1] + sums[2]
}

func ReadInput() ([][]int, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return make([][]int, 0), err
	}
	defer f.Close()

	var elfs [][]int
	elf := 0
	elfs = append(elfs, make([]int, 0))

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			calories, _ := strconv.Atoi(scanner.Text())
			elfs[elf] = append(elfs[elf], calories)
		} else {
			elf++
			elfs = append(elfs, make([]int, 0))
		}
	}

	return elfs, nil
}

func main() {
	elfs, _ := ReadInput()

	result := Solve1(elfs)
	fmt.Println(result)

	result = Solve2(elfs)
	fmt.Println(result)
}
