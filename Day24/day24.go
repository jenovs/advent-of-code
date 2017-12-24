// http://adventofcode.com/2017/day/24

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

	c := strings.Split(string(input[:len(input)-1]), "\n")
	s := make([][]int, len(c))
	for i, v := range c {
		ns := make([]int, 2)
		for j, n := range strings.Split(v, "/") {
			ns[j], _ = strconv.Atoi(n)
		}
		s[i] = ns
	}

	start := 0
	findMax(s, start, 0)
	fmt.Println("Maximum strength bridge:", max)
	findLongest(s, start, 0, 0)
	fmt.Println("Strength of the longest bridge:", max2)

}

var max int

func findMax(s [][]int, p int, sum int) {
	nn := findPieces(s, p)

	if len(nn) == 0 {
		return
	}

	for _, v := range nn {
		sum2 := sum + v[0] + v[1]

		if sum2 > max {
			max = sum2
		}
		next := v[0]
		if next == p {
			next = v[1]
		}
		findMax(discard(s, v), next, sum2)
	}
}

var max2 int
var longest int

func findLongest(s [][]int, p int, sum int, c int) {
	nn := findPieces(s, p)

	if len(nn) == 0 {
		if c == longest && sum > max2 {
			max2 = sum
		}
		if c > longest {
			longest = c
			max2 = sum
		}
		return
	}

	for _, v := range nn {
		sum2 := sum + v[0] + v[1]

		if sum2 > max {
			max2 = sum2
		}
		next := v[0]
		if next == p {
			next = v[1]
		}
		findLongest(discard(s, v), next, sum2, c+1)
	}
}

func discard(s [][]int, p []int) [][]int {
	var sn [][]int
	for _, v := range s {
		if (v[0] == p[0] && v[1] == p[1]) || (v[1] == p[0] && v[0] == p[1]) {
			continue
		}
		sn = append(sn, v)
	}
	return sn
}

func findPieces(s [][]int, n int) [][]int {
	nn := [][]int{}
	for _, a := range s {
		if a[0] == n || a[1] == n {
			nn = append(nn, a)
		}
	}
	return nn
}
