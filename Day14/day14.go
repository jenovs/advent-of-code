// http://adventofcode.com/2017/day/14

package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := "ljoxqyyw"

	sum := 0

	for i := 0; i < 128; i++ {
		ind := strconv.Itoa(i)
		hash := hasher(input + "-" + ind)
		for _, v := range hash {
			dec, _ := strconv.ParseInt(string(v), 16, 0)
			bin := fmt.Sprintf("%b", dec)
			for _, b := range bin {
				if b == '1' {
					sum++
				}
			}
		}
	}
	fmt.Println(sum)
}
