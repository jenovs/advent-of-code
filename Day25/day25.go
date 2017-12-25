// http://adventofcode.com/2017/day/25

package main

import "fmt"

func main() {
	tape := []int{0}
	pos := 0
	state := 'A'
	steps := 12399302

	for ; steps > 0; steps-- {
		v := tape[pos]
		switch state {
		case 'A':
			if v == 0 {
				tape[pos] = 1
				state = 'B'
			} else {
				tape[pos] = 0
				state = 'C'
			}
			pos++
		case 'B':
			tape[pos] = 0
			if v == 0 {
				pos--
				state = 'A'
			} else {
				pos++
				state = 'D'
			}
		case 'C':
			tape[pos] = 1
			pos++
			if v == 0 {
				state = 'D'
			} else {
				state = 'A'
			}
		case 'D':
			if v == 0 {
				tape[pos] = 1
				state = 'E'
			} else {
				tape[pos] = 0
			}
			pos--
		case 'E':
			tape[pos] = 1
			if v == 0 {
				pos++
				state = 'F'
			} else {
				pos--
				state = 'B'
			}
		case 'F':
			tape[pos] = 1
			pos++
			if v == 0 {
				state = 'A'
			} else {
				state = 'E'
			}
		}
		pos, tape = adjust(pos, tape)
	}

	c := 0
	for _, t := range tape {
		c += t
	}
	fmt.Println("Checksum:", c)
}

func adjust(p int, s []int) (int, []int) {
	if p >= len(s) {
		s = append(s, 0)
	} else if p < 0 {
		s = append([]int{0}, s...)
		p++
	}
	return p, s
}
