// http://adventofcode.com/2017/day/6

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
	inputStrings := strings.Split(inputString, " ")

	// remove the last (empty) value from the input array
	inputStrings = inputStrings[:len(inputStrings)-1]

	input := make([]int, len(inputStrings))

	// map over input array and convert strings to ints
	for i, v := range inputStrings {
		intV, _ := strconv.ParseInt(v, 10, 0)
		x := int(intV)
		input[i] = x
	}

	var results []string

	results = append(results, toString(input))

	counter := 0
	duplicateInd := 0

	for duplicateInd == 0 {
		maxInd := findMaxInd(input)

		toRedistribute := input[maxInd]
		input[maxInd] = 0
		var startInd int
		if maxInd+1 > 15 {
			startInd = 0
		} else {
			startInd = maxInd + 1
		}
		input = redistribute(input, toRedistribute, startInd)
		results = append(results, toString(input))
		counter++
		duplicateInd = hasDuplicate(results)
	}

	fmt.Println("Cycles to duplicate:\t", counter)
	fmt.Println("Size of the loop:\t", counter-duplicateInd)
}

func hasDuplicate(input []string) int {
	str := input[len(input)-1]

	for i := 0; i < len(input)-2; i++ {
		if input[i] == str {
			return i
		}
	}
	return 0
}

func toString(input []int) string {
	output := make([]string, len(input))

	for i, n := range input {
		output[i] = strconv.Itoa(n)
	}

	return strings.Join(output, ",")
}

func findMaxInd(s []int) int {
	var max int
	var result int
	for i, v := range s {
		if v > max {
			result = i
			max = v
		}
	}
	return result
}

func redistribute(input []int, r int, ind int) []int {
	for r > 0 {
		input[ind]++
		r--
		if ind >= 15 {
			ind = 0
		} else {
			ind++
		}
	}
	return input
}
