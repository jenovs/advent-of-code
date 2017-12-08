// http://adventofcode.com/2017/day/8

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Register struct {
	name  string
	value int
}

func main() {
	inputBytes, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal("Error ", err)
	}

	inputStr := string(inputBytes)
	input := strings.Split(inputStr, "\n")

	// remove the last (empty) value from the input array
	input = input[:len(input)-1]

	// create a data map
	m := make(map[string]Register)

  highestEver := 0

	for _, v := range input {
		instr := strings.Split(v, " ")
		name := instr[0]
		command := instr[1]
		amount, _ := strconv.Atoi(instr[2])
		nameToCheck := instr[4]
		condition := instr[5]
		valueToCheck, _ := strconv.Atoi(instr[6])
		registerToCheck := m[nameToCheck]
    register := m[name]
		check := checkCondition(registerToCheck.value, valueToCheck, condition)
    if check {
      if command == "inc" {
        register.value += amount
        m[name] = register
      } else {
        register.value -= amount
        m[name] = register
      }
    }
    maxValue := findMaxValue(m)
    if maxValue > highestEver {
      highestEver = maxValue
    }
	}

  fmt.Println("Max value:\t", findMaxValue(m))
  fmt.Println("Max value ever:\t", highestEver)
}

func findMaxValue(register map[string]Register) int {
  max := 0
  for _, v := range register {
    if v.value > max {
      max = v.value
    }
  }
  return max
}

func checkCondition(regValue int, number int, condition string) bool {
	switch condition {
	case "==":
		return regValue == number
	case "!=":
		return regValue != number
	case "<":
		return regValue < number
	case "<=":
		return regValue <= number
	case ">":
		return regValue > number
	case ">=":
		return regValue >= number
	}
	return false
}
