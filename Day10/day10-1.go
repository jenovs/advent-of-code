// http://adventofcode.com/2017/day/10

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
	listSize := 256
	// input, err := ioutil.ReadFile("input_test.txt")
	// listSize := 5

	if err != nil {
		log.Fatal("Error ", err)
	}

	s := strings.Split(string(input[:len(input)-1]), ",")
	lengths := make([]int, len(s))
	for i, v := range s {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal("Error ", err)
		}
		lengths[i] = n
	}

	list := make([]int, listSize)
	for i := 0; i < listSize; i++ {
		list[i] = i
	}

	pos := 0
	skip := 0

	for _, l := range lengths {
		substr := reverse(readRange(list, pos, l))
		list = writeRange(list, pos, substr)

		pos += l + skip
		pos = pos % len(list)
		skip++
	}
	fmt.Println(list[0] * list[1])
}

func readRange(list []int, start int, length int) []int {
	result := make([]int, length)

	for i := 0; length > 0; i++ {
		listInd := (start + i) % len(list)
		result[i] = list[listInd]
		length--
	}

	return result
}

func writeRange(list []int, start int, newList []int) []int {
	for i := 0; i < len(newList); i++ {
		listInd := (start + i) % len(list)
		list[listInd] = newList[i]
	}
	return list
}

func reverse(s []int) []int {
	r := make([]int, len(s))
	c := 0
	for i := len(s) - 1; i >= 0; i-- {
		r[c] = s[i]
		c++
	}
	return r
}
