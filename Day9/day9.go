// http://adventofcode.com/2017/day/8

package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal("Error ", err)
	}

	var filtered []byte

	// filter out `!` and the character after it
	// ! = 33 (10 is a newline character)
	for i := 0; i < len(input); i++ {
		if input[i] == 33 || input[i] == 10 {
			i++
		} else {
			filtered = append(filtered, input[i])
		}
	}

	input = filtered
	filtered = nil

	// filter out and count garbage
	// < - 60
	// > - 62
	isGarbageOpen := false
	garbageCounter := 0

	for i := 0; i < len(input); i++ {
		if isGarbageOpen && input[i] != 62 {
			garbageCounter++
			continue
		} else if isGarbageOpen && input[i] == 62 {
			isGarbageOpen = false
			continue
		}
		if input[i] == 60 && !isGarbageOpen {
			isGarbageOpen = true
			continue
		}
		filtered = append(filtered, input[i])
	}

	fmt.Println("Total score of groups:\t", counter(filtered))
	fmt.Println("Characters in garbage:\t", garbageCounter)
}

// Count groups of {}
// { - 123
// } - 125
func counter(s []byte) int {
	sum := 0
	count := 0
	for _, v := range s {
		if v == 123 {
			count++
		} else if v == 125 {
			sum += count
			count--
		}
	}
	return sum
}
