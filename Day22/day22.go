// http://adventofcode.com/2017/day/22

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	// input, err := ioutil.ReadFile("input_test.txt")

	if err != nil {
		log.Fatal("Error ", err)
	}

	grid := strings.Split(string(input[:len(input)-1]), "\n")
	fmt.Println("Nodes infected (regular mode):", phaseOne(grid))

	grid = strings.Split(string(input[:len(input)-1]), "\n")
	fmt.Println("Nodes infected (angry mode):", phaseTwo(grid))
}

func phaseTwo(grid []string) int {
	w := len(grid[0])
	h := len(grid)
	c := 0

	pos := []int{h / 2, w / 2}
	// Direction: 0 - up, 1 - right, 2 - down, 3 - left
	dir := 0
	for i := 0; i < 10000000; i++ {
		currCell := grid[pos[0]][pos[1]]
		if currCell == 35 { // #
			dir = turn(dir, "R")
		} else if currCell == 46 { // .
			dir = turn(dir, "L")
		} else if currCell == 70 { // F
			if dir > 1 {
				dir -= 2
			} else {
				dir += 2
			}
		}
		flipX(grid, pos)
		if grid[pos[0]][pos[1]] == 35 {
			c++
		}
		pos = move(pos, dir)
		if pos[1] >= w {
			growRight(grid)
			w++
		}
		if pos[0] >= h {
			grid = append(grid, strings.Repeat(".", w))
			h++
		}
		if pos[1] < 0 {
			growLeft(grid)
			w++
			pos[1]++
		}
		if pos[0] < 0 {
			grid = append([]string{strings.Repeat(".", w)}, grid...)
			h++
			pos[0]++
		}
	}
	return c
}

func phaseOne(grid []string) int {
	w := len(grid[0])
	h := len(grid)
	c := 0

	pos := []int{h / 2, w / 2}
	// Direction: 0 - up, 1 - right, 2 - down, 3 - left
	dir := 0
	for i := 0; i < 10000; i++ {
		currCell := grid[pos[0]][pos[1]]
		if currCell == 35 {
			dir = turn(dir, "R")
		} else {
			dir = turn(dir, "L")
		}
		flip(grid, pos)
		if grid[pos[0]][pos[1]] == 35 {
			c++
		}
		pos = move(pos, dir)
		if pos[1] >= w {
			growRight(grid)
			w++
		}
		if pos[0] >= h {
			grid = append(grid, strings.Repeat(".", w))
			h++
		}
		if pos[1] < 0 {
			growLeft(grid)
			w++
			pos[1]++
		}
		if pos[0] < 0 {
			grid = append([]string{strings.Repeat(".", w)}, grid...)
			h++
			pos[0]++
		}
	}
	return c
}

func growRight(g []string) {
	for i := range g {
		g[i] = g[i] + "."
	}
}

func growLeft(g []string) {
	for i := range g {
		g[i] = "." + g[i]
	}
}

func move(p []int, d int) []int {
	switch d {
	case 0:
		p[0]--
	case 1:
		p[1]++
	case 2:
		p[0]++
	case 3:
		p[1]--
	}
	return p
}

func flip(g []string, pos []int) {
	row := strings.Split(g[pos[0]], "")
	if row[pos[1]] == "#" {
		row[pos[1]] = "."
	} else {
		row[pos[1]] = "#"
	}
	g[pos[0]] = strings.Join(row, "")
}

func flipX(g []string, pos []int) {
	row := strings.Split(g[pos[0]], "")
	switch row[pos[1]] {
	case "#":
		row[pos[1]] = "F"
	case ".":
		row[pos[1]] = "W"
	case "W":
		row[pos[1]] = "#"
	case "F":
		row[pos[1]] = "."
	}
	g[pos[0]] = strings.Join(row, "")
}

func turn(d int, t string) int {
	tn := 1
	if t == "L" {
		tn = -1
	}
	d = d + tn
	if d > 3 {
		d = 0
	}
	if d < 0 {
		d = 3
	}
	return d
}
