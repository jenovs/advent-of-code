// http://adventofcode.com/2017/day/13

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
	// input, err := ioutil.ReadFile("input_test.txt")

	if err != nil {
		log.Fatal("Error ", err)
	}

	// split input into slice (removing line break character)
	s := strings.Split(string(input[:len(input)-1]), "\n")

	// convert input slice into map
	data := make(map[int]int)
	for i := range s {
		kv := strings.Split(s[i], ": ")
		k, _ := strconv.Atoi(kv[0])
		v, _ := strconv.Atoi(kv[1])
		data[k] = v
	}

	caught := true
	delay := -1

	for caught == true {
		caught = false
		severity := 0

		delay++

		for k, p := range data {
			scannerPos := (k + delay) % ((p - 1) * 2)
			if scannerPos == 0 {
				caught = true
				if delay == 0 {
					severity += (k * p)
				}
			}
		}

		if delay == 0 {
			fmt.Println("Severity:\t", severity)
		}

	}
	fmt.Println("Min. delay:\t", delay)
}
