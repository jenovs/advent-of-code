// http://adventofcode.com/2017/day/5

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	inputBytes, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal("Error ", err)
	}

	inputString := string(inputBytes)
	input := strings.Split(inputString, "\n")
	// remove the last (empty) value from the input array
	input = input[:len(input)-1]
	inputNums := make([]int, len(input))

	// map over input array and convert strings to ints
	for i, v := range input {
		intV, _ := strconv.ParseInt(v, 10, 0)
		x := int(intV)
		inputNums[i] = x
	}

	// Part 1 solution
	fmt.Println(countJumps(inputNums, false))
	// Part 2 solution
	fmt.Println(countJumps(inputNums, true))
}

func countJumps(arr []int, withDecrease bool) int {
	// copy the input slice
	input := make([]int, len(arr))
	copy(input, arr)

	counter := 0
	position := 0

	for position >= 0 && position < len(input) {
		jump := input[position]

		if input[position] >= 3 && withDecrease {
			input[position] -= 1
		} else {
			input[position] += 1
		}

		counter += 1
		position += jump
	}
	return counter
}
