package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type coords struct {
	x int
	y int
}

func Print(t int, rocks [][]coords, r []coords) {
	//return
	tmp := make([][]coords, len(rocks))
	copy(tmp, rocks)
	tmp = append(tmp, r)

	//fmt.Println("---")
	//fmt.Println(len(rocks))
	for i := 0; i < t; i++ {
		for j := 0; j < 7; j++ {
			printed := false
			for _, r := range tmp {
				for _, b := range r {
					if (b.x == j) && (b.y == i) {
						printed = true
						fmt.Print("#")
					}
				}
			}
			if !printed {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	time.Sleep(1 * time.Second) //(500 * time.Millisecond)
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func Drop(ri, t int) []coords {
	switch ri {
	case 0:
		return []coords{
			coords{2, t},
			coords{3, t},
			coords{4, t},
			coords{5, t},
		}
	case 1:
		return []coords{
			coords{3, t + 2},
			coords{2, t + 1},
			coords{3, t + 1},
			coords{4, t + 1},
			coords{3, t},
		}
	case 2:
		return []coords{
			coords{2, t},
			coords{3, t},
			coords{4, t},
			coords{4, t + 1},
			coords{4, t + 2},
		}
	case 3:
		return []coords{
			coords{2, t},
			coords{2, t + 1},
			coords{2, t + 2},
			coords{2, t + 3},
		}
	case 4:
		return []coords{
			coords{2, t + 1},
			coords{2, t},
			coords{3, t + 1},
			coords{3, t},
		}
	default:
		panic("boom")
	}
}

func collide(rock []coords, rocks []coords) bool {
	for _, b1 := range rocks {
		for _, b2 := range rock {
			if (b1.x == b2.x) && (b1.y == b2.y) {
				return true
			}
		}
	}

	return false
}

func MoveUp(r []coords) []coords {
	for i := 0; i < len(r); i++ {
		r[i].y += 1
	}

	return r
}

func MoveDown(r []coords) []coords {
	for i := 0; i < len(r); i++ {
		r[i].y -= 1
	}

	return r
}

func MoveLeft(r []coords) []coords {
	for i := 0; i < len(r); i++ {
		if r[i].x == 0 {
			return r
		}
	}

	for i := 0; i < len(r); i++ {
		r[i].x -= 1
	}

	return r
}

func MoveRight(r []coords) []coords {
	for i := 0; i < len(r); i++ {
		if r[i].x == 6 {
			return r
		}
	}
	for i := 0; i < len(r); i++ {
		r[i].x += 1
	}

	return r
}

type Value struct {
	i    int
	t    int
	s    []coords
	oldT int
	oldY int
}

func sign(rocks []coords) []coords {
	maxY := 0

	for _, e := range rocks {
		if e.y > maxY {
			maxY = e.y
		}
	}

	r := make([]coords, 0)
	for _, rock := range rocks {
		if maxY-rock.y <= 20 {
			r = append(r, coords{rock.x, maxY - rock.y})
		}
	}

	return r
}

// 28 4
func InSeen(i int, v Value, seen []Value) int {
	for i, e := range seen {
		if (v.i == e.i) && (v.t == e.t) {
			//if (v.i == 28) && (v.t == 4) {
			//	fmt.Println(v.s)
			//	fmt.Println(e.s)
			//	fmt.Println("---")
			//}
			equal := true
			for i, c1 := range e.s {
				c2 := v.s[i]
				if c1 != c2 {
					equal = false
					break
				}
			}
			if equal {
				return i
			}
		}
	}

	return -1
}

func Solve1(jets string) int {
	rocks := make([]coords, 0)
	i := 0
	top := 0

	seen := make([]Value, 0)

	for x := 0; x < 7; x++ {
		rocks = append(rocks, coords{x, 0})
	}

	added := 0
	t := 0

	for {
		if t >= 1000000000000 {
			break
		}
		r := Drop(t%5, top+4)

		for {
			switch jets[i] {
			case '>':
				r = MoveRight(r)
				if collide(r, rocks) {
					r = MoveLeft(r)
				}
			case '<':
				r = MoveLeft(r)
				if collide(r, rocks) {
					r = MoveRight(r)
				}
			default:
				panic("boom")
			}

			i = (i + 1) % len(jets)

			r = MoveDown(r)
			if collide(r, rocks) {
				r = MoveUp(r)
				for _, e := range r {
					rocks = append(rocks, e)
				}
				for _, e := range rocks {
					if e.y > top {
						top = e.y
					}
				}

				nv := Value{i, t % 5, sign(rocks), t, top}
				si := InSeen(t, nv, seen)
				//fmt.Println(si)
				if (si != -1) && (t >= 2022) {
					fmt.Println("found signature", top, t)
					fmt.Println(seen[si])
					oldT := seen[si].oldT
					oldY := seen[si].oldY
					dy := top - oldY
					dt := t - oldT
					amt := (1000000000000 - t) / dt
					added += amt * dy
					t += amt * dt
				}
				if si != -1 {
					seen[si] = Value{seen[si].i, seen[si].t, seen[si].s, t, top}
				} else {
					//fmt.Println("adding", i, t, sign(rocks))
					seen = append(seen, nv)
				}

				break
			}
		}
		t++
		//fmt.Println(t)
		//if t == 2022 {
		//	return top
		//}
	}

	return top + added
}

func main() {
}
