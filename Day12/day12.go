// http://adventofcode.com/2017/day/12

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

	if err != nil {
		log.Fatal("Error ", err)
	}

	// split input into slices (removing line break character)
	s := strings.Split(string(input[:len(input)-1]), "\n")

	data := make(map[int][]int)
	// convert input slice into map
	for i := range s {
		arr := strings.Split(s[i], " <-> ")
		key, _ := strconv.Atoi(arr[0])
		v := strings.Split(arr[1], ", ")
		val := make([]int, len(v))
		for j := range val {
			val[j], _ = strconv.Atoi(v[j])
		}
		data[key] = val
	}

	initialKey := 0
	count := 0

	for len(data) > 0 {
		// get list of map keys
		i := 0
		keys := make([]int, len(data))
		for k := range data {
			keys[i] = k
			i++
		}

		key := keys[0]
		nodes := []int{keys[0]}

		if count == 0 {
			key = initialKey
			nodes = []int{0}
		}

		nodes = checkNodes(data, key, nodes)

		if count == 0 {
			fmt.Println("Programs in group 0:\t", len(nodes))
		}

		// remove initial key
		delete(data, key)
		// remove linked keys
		for _, k := range nodes {
			delete(data, k)
		}

		count++
	}

	fmt.Println("Total group count:\t", count)

}

func checkNodes(data map[int][]int, key int, nodes []int) []int {
	input := data[key]
	if len(input) == 0 {
		return nodes
	}
	for i := range input {
		if !hasVal(nodes, input[i]) {
			nodes = append(nodes, input[i])
			nodes = checkNodes(data, input[i], nodes)
		}
	}
	return nodes
}

func hasVal(s []int, v int) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}
	return false
}
