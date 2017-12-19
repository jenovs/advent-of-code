// http://adventofcode.com/2017/day/19

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	// input, err := ioutil.ReadFile("input_test.txt")

	if err != nil {
		log.Fatal("Error ", err)
	}

	comm := strings.Split(string(input[:len(input)-1]), "\n")

	start := strings.Index(comm[0], "|")

	pos := []int{0, start}
	dir := "D"
	letters := ""
	steps := 0

	for {
		// current positon is outside the map
		if pos[0] < 0 || pos[0] > len(comm) || pos[1] < 0 || pos[1] > len(comm[0]) {
			break
		}

		// value of the current position
		p := comm[pos[0]][pos[1]]

		// current position is an empty field
		if p == 32 {
			break
		}

		// Turn
		if p == '+' {
			dir, err = calcTurn(comm, dir, pos)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else
		// Letter
		if p >= 65 && p <= 90 {
			letters += string(p)
		}

		// Calculate next position
		switch dir {
		case "U":
			pos[0]--
		case "D":
			pos[0]++
		case "R":
			pos[1]++
		case "L":
			pos[1]--
		}

		steps++
	}

	fmt.Println("Letters:", letters)
	fmt.Println("Steps:\t", steps)
}

// calcTurn takes array of commands c, current direction d, current position p
// and returns new direction nd and error e
// `32` is an ASCII for space
func calcTurn(c []string, d string, p []int) (nd string, e error) {
	if d == "R" || d == "L" {
		np0 := p[0] - 1
		if np0 >= 0 && c[np0][p[1]] != 32 {
			return "U", nil
		}
		np0 = p[0] + 1
		if np0 < len(c) && c[np0][p[1]] != 32 {
			return "D", nil
		}
	} else {
		np1 := p[1] - 1
		if np1 >= 0 && c[p[0]][np1] != 32 {
			return "L", nil
		}
		np1 = p[1] + 1
		if np1 < len(c[0]) && c[p[0]][np1] != 32 {
			return "R", nil
		}
	}
	return "", errors.New("Error while calculating turn")
}
