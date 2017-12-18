// http://adventofcode.com/2017/day/18

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

	fmt.Println("Last frequency:", recoverFrequency(commands))
}

func recoverFrequency(commands [][]string) int {
	register := make(map[string]int)
	i := 0
	var lastSnd int
	for true {
		c := commands[i]

		switch c[0] {
		case "set", "add", "mul", "mod":
			executeMath(c, register)
		case "snd":
			lastSnd = register[c[1]]
		case "rcv":
			x := register[c[1]]
			if x != 0 {
				return lastSnd
			}
		case "jgz":
			x := register[c[1]]
			if x > 0 {
				y, _ := strconv.Atoi(c[2])
				i += y
				continue
			}
		}

		i++
	}
	return -1
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
	case "add":
		if err != nil {
			r[c[1]] += r[c[2]]
		} else {
			r[c[1]] += y
		}
	case "mul":
		if err != nil {
			r[c[1]] *= r[c[2]]
		} else {
			r[c[1]] *= y
		}
	case "mod":
		if err != nil {
			r[c[1]] %= r[c[2]]
		} else {
			r[c[1]] %= y
		}
	}
}
