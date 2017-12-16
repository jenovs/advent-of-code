// Code from Day 10-2

package main

import (
	"fmt"
)

func hasher(input string) string {
  listSize :=  256
  suffix := []int{17, 31, 73, 47, 23}

  lengths := make([]int, len(input))
  for i, v := range input {
    lengths[i] = int(v)
  }

  lengths = append(lengths, suffix...)

  list := make([]int, listSize)
  for i := 0; i < listSize; i++ {
    list[i] = i
  }

  pos := 0
  skip := 0

  for i := 0; i < 64; i++ {
    for _, l := range lengths {
      newRange := readRange(list, pos, l)
      substr := reverse(newRange)

      list = writeRange(list, pos, substr)
      pos += l + skip
      pos = pos % len(list)
      skip++
    }
  }
  denseHash := denseHash(list)
  hexHash := ""
  for i := range denseHash {
    hex := fmt.Sprintf("%x", denseHash[i])
    if len(hex) == 1 {
      hex = "0" + hex
    }
    hexHash += hex
  }
  return hexHash
}

func denseHash(hash []int) []int {
	dense := make([]int, 16)
	for i := 0; i < len(hash); i += 16 {
		temp := 0
		for j := 0; j < 16; j++ {
			if j == 0 {
				temp = hash[i+j]
			} else {
				temp = temp ^ hash[i+j]
			}
		}
		dense[i/16] = temp
	}
	return dense
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
