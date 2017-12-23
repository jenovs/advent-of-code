// http://adventofcode.com/2017/day/23

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal("Error ", err)
	}

	comm := strings.Split(string(input[:len(input)-1]), "\n")
	commands := [][]string{}

	for _, c := range comm {
		t := strings.Split(c, " ")
		commands = append(commands, t)
	}

	fmt.Printf("Multiplied %v times.\n", countMul(commands))
	fmt.Printf("Final value of h: %v.\n", loop(commands))
}

func countMul(commands [][]string) int {
	register := make(map[string]int)
	i := 0
	mulCount := 0
	for {
		if i < 0 || i >= len(commands) {
			break
		}
		c := commands[i]

		switch c[0] {
		case "set", "sub":
			executeMath(c, register)
			i++
		case "mul":
			executeMath(c, register)
			mulCount++
			i++
		case "jnz":
			x, err := strconv.Atoi(c[1])
			if err != nil {
				x = register[c[1]]
			}
			if x != 0 {
				y, _ := strconv.Atoi(c[2])
				i += y
			} else {
				i++
			}
		}
	}
	return mulCount
}

func loop(commands [][]string) int {
	r := make(map[string]int)
	// use first 8 commands to set variables
	for i := 0; i < 8; i++ {
		cm := commands[i]
		switch cm[0] {
		case "set", "sub", "mul":
			executeMath(cm, r)
		}
	}

	b := r["b"]
	c := r["c"]
	h := 0
	// get subtracting value
	s, _ := strconv.Atoi(commands[30][2])

	// simplified loop
	for c-b >= 0 {
		if !isPrime(b) {
			h++
		}
		b -= s
	}
	return h
}

func isPrime(n int) bool {
	if n%2 == 0 && n != 2 {
		return false
	}
	for i := 3; i*i < n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func executeMath(c []string, r map[string]int) {
	y, err := strconv.Atoi(c[2])
	switch c[0] {
	case "set":
		if err != nil {
			r[c[1]] = r[c[2]]
		} else {
			r[c[1]] = y
		}
	case "sub":
		if err != nil {
			r[c[1]] -= r[c[2]]
		} else {
			r[c[1]] -= y
		}
	case "mul":
		if err != nil {
			r[c[1]] *= r[c[2]]
		} else {
			r[c[1]] *= y
		}
	}
}
