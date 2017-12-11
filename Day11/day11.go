// http://adventofcode.com/2017/day/11

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	buffer, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal("Error ", err)
	}
	buffer = buffer[:len(buffer)-1]
	input := strings.Split(string(buffer), ",")

	position := make([]int, 2)
	maxDistance := 0

	for _, v := range input {
		position = move(v, position)
		if d := calcDistance(position); maxDistance < d {
			maxDistance = d
		}
	}
	
	fmt.Println("Distance to him:\t", calcDistance(position))
	fmt.Println("Furthest he got:\t", maxDistance)
}

func calcDistance(p []int) int {
	p0 := p[0]
	p1 := p[1]

	// Math.abs()
	if p0 < 0 {
		p0 *= -1
	}
	if p1 < 0 {
		p1 *= -1
	}
	return ((p0 + p1) /2)
}

func move(s string, pos []int) []int {
	switch s {
	case "n":
		pos[0] += 2
	case "s":
		pos[0] -= 2
	case "ne":
		pos[0] += 1
		pos[1] += 1
	case "se":
		pos[0] -= 1
		pos[1] += 1
	case "sw":
		pos[0] -= 1
		pos[1] -= 1
	case "nw":
		pos[0] += 1
		pos[1] -= 1
	}
	return pos
}
