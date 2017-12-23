// http://adventofcode.com/2017/day/7

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Data struct {
	weight   int
	children []string
}

var m map[string]Data
var sum int

func main() {
	inputBytes, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal("Error ", err)
	}

	inputString := string(inputBytes)
	inputStrings := strings.Split(inputString, "\n")

	// remove the last (empty) value from the input array
	input := inputStrings[:len(inputStrings)-1]

	// create a data map
	m = make(map[string]Data)

	// populate the data map
	for _, v := range input {
		arr := strings.Split(v, " -> ")
		fullRoot := strings.Split(arr[0], " ")
		weight, _ := strconv.Atoi(fullRoot[1][1 : len(fullRoot[1])-1])
		root := fullRoot[0]
		if len(arr) > 1 {
			m[root] = Data{weight, strings.Split(arr[1], ", ")}
		} else {
			m[root] = Data{weight, []string{}}
		}
	}

	var children []string

	for _, v := range input {
		arr := strings.Split(v, " -> ")
		if len(arr) > 1 {
			children = append(children, strings.Split(arr[1], ", ")...)
		}
	}

	var rootNode string

	for _, v := range input {
		arr := strings.Split(v, " -> ")
		word := strings.Split(arr[0], " ")[0]

		if !includes(children, word) {
			rootNode = word
			fmt.Println("Root word:", rootNode, "\n")
		}
	}

	// TODO Refactor to navigate unbalanced node automatically
	// currently it works by copy pasting unbalanced node starting with rootNode
	// until unbalanced child node is met
	// The answer is that unbalanced node's weight +/- difference of nodes
	navigateNodes(m, "gynfwly")
	// And get rid of global sum variable
}

func navigateNodes(data map[string]Data, node string) {
	weights := make(map[string]int)
	for _, node := range m[node].children {
		sumChildrenWeights(m, node, 0)
		weights[node] = sum
		fmt.Println("node", node, "weight", data[node].weight)
		sum = 0
	}
		fmt.Println("\nweights", weights)
}

func sumChildrenWeights(data map[string]Data, node string, total int) {
	children := data[node].children
	sum += data[node].weight
	if len(children) == 0 {
		return
	}

	for _, child := range children {
		sumChildrenWeights(data, child, total)
	}
	return
}

func includes(arr []string, word string) bool {
	for _, v := range arr {
		if v == word {
			return true
		}
	}
	return false
}
