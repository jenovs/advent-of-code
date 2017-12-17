// http://adventofcode.com/2017/day/17

package main

import "fmt"

func main() {
	input := 355
	limit := int(5e+7)
	num := 2017

	fmt.Printf("Value after %v:\t%v\n", num, findAfterNum(input, num))
	fmt.Printf("Value after zero:\t%v\n", findAfter0(input, limit))
}

func findAfterNum(input int, limit int) int {
	data := []int{0}
	pos := 0
	n := 1

	for n < limit+1 {
		moves := input % len(data)
		pos = (pos + moves) % len(data)
		data = append(data, -1)
		copy(data[pos+1:], data[pos:])
		data[pos] = n

		n++
		pos++
	}

	return data[find(data, limit)+1]
}

func findAfter0(input int, limit int) int {
	pos := 0
	res := 0

	for n := 1; n < limit; n++ {
		pos = (pos + input) % n
		if pos == 0 {
			res = n
		}
		pos++
	}

	return res
}

func find(s []int, n int) int {
	for i, v := range s {
		if v == n {
			return i
		}
	}
	return -1
}
