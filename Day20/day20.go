// http://adventofcode.com/2017/day/20

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type particles struct {
	p coords
	v coords
	a coords
}

type coords struct {
	x int
	y int
	z int
}

func main() {
	input, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal("Error ", err)
	}

	comm := strings.Split(string(input[:len(input)-1]), "\n")

	ps := make([]particles, len(comm))

	for i, c := range comm {
		particle := strings.Split(c, ", ")
		p := particle[0][3 : len(particle[0])-1]
		ps[i].p = toIntXYZ(p)

		v := particle[1][3 : len(particle[1])-1]
		ps[i].v = toIntXYZ(v)

		a := particle[2][3 : len(particle[2])-1]
		ps[i].a = toIntXYZ(a)
	}

	fmt.Println("Closest particle in the long run: ", calcMinDelta(ps))
	fmt.Println("Particles left after collisions:  ", afterCollisions(ps))
}

func calcMinDelta(ps []particles) int {
	minD := delta(ps[0].a)
	ind := 0
	for i, prt := range ps {
		d := delta(prt.a)
		if d < minD {
			minD, ind = d, i
		}
	}
	return ind
}

func delta(a coords) int {
	return int(math.Abs(float64(a.x)) + math.Abs(float64(a.y)) + math.Abs(float64(a.z)))
}

// TODO: Refactor this brute force guessing approach
func afterCollisions(ps []particles) int {
	for i := 0; i < 100; i++ {
		ps = removeCollisions(update(ps))
	}
	return len(ps)
}

func removeCollisions(ps []particles) []particles {
	res := []particles{}
	for _, prt := range ps {
		if !collide(ps, prt.p) {
			res = append(res, prt)
		}
	}
	return res
}

func collide(ps []particles, c coords) bool {
	count := 0
	for _, prt := range ps {
		if prt.p.x == c.x && prt.p.y == c.y && prt.p.z == c.z {
			count++
			if count > 1 {
				return true
			}
		}
	}
	return false
}

func update(ps []particles) []particles {
	for i := range ps {
		ps[i].v.x += ps[i].a.x
		ps[i].v.y += ps[i].a.y
		ps[i].v.z += ps[i].a.z
		ps[i].p.x += ps[i].v.x
		ps[i].p.y += ps[i].v.y
		ps[i].p.z += ps[i].v.z
	}
	return ps
}

func toIntXYZ(s string) coords {
	sa := strings.Split(s, ",")
	nums := [3]int{}
	for i, st := range sa {
		n, err := strconv.Atoi(st)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		nums[i] = n
	}
	cs := coords{}
	cs.x = nums[0]
	cs.y = nums[1]
	cs.z = nums[2]
	return cs
}
