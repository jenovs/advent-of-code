// http://adventofcode.com/2017/day/16

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
	dances := 1000000000

	if err != nil {
		log.Fatal("Error ", err)
	}

	group := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p'}

	// split input into slice (removing line break character)
	s := strings.Split(string(input[:len(input)-1]), ",")

	afterFirst := ""
	for i := 0; i < dances; i++ {
		for _, d := range s {
			switch d[0] {
			case 's':
				n, _ := strconv.Atoi(d[1:])
				group = append(group[len(group)-n:], group[:len(group)-n]...)
			case 'x':
				n := strings.Split(d[1:], "/")
				n1, _ := strconv.Atoi(n[0])
				n2, _ := strconv.Atoi(n[1])
				group[n1], group[n2] = group[n2], group[n1]
			case 'p':
				n := strings.Split(d[1:], "/")
				i1 := strings.Index(string(group), n[0])
				i2 := strings.Index(string(group), n[1])
				group[i1], group[i2] = group[i2], group[i1]
			}
		}
		// Print the result for the first part
		if i == 0 {
			afterFirst = string(group)
			fmt.Println("After the first dance:", string(group))
		}
		// Find repeat period and calculate remaining operations
		if afterFirst == string(group) && i > 0 {
			rep := i
			rem := dances % rep
			i = dances - rem
		}
	}

	fmt.Println("After the last dance:", string(group))
}
