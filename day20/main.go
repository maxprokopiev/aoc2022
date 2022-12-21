package main

func mod(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}

func Solve1(input []int, dkey int, iterations int) int {
	xs := make([][]int, 0)

	for i := 0; i < len(input); i++ {
		xs = append(xs, []int{i, input[i] * dkey})
	}

	var e, val []int
	for t := 0; t < iterations; t++ {
		for i := 0; i < len(input); i++ {
			for {
				if xs[0][0] == i {
					break
				}
				e, xs = xs[0], xs[1:]
				xs = append(xs, e)
			}

			val, xs = xs[0], xs[1:]

			toPop := val[1]
			toPop = mod(toPop, len(xs))

			for i := 0; i < toPop; i++ {
				e, xs = xs[0], xs[1:]
				xs = append(xs, e)
			}

			xs = append(xs, val)
		}
	}

	var zero int
	for i := 0; i < len(xs); i++ {
		if xs[i][1] == 0 {
			zero = i
		}
	}

	return xs[(zero+1000)%len(xs)][1] + xs[(zero+2000)%len(xs)][1] + xs[(zero+3000)%len(xs)][1]
}

func main() {
}
