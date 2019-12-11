package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func intcode(program []int, position int) []int {
	var operation int = program[0+position]
	if operation == 99 {
		return program
	}

	var input1Position int = program[1+position]
	var input2Position int = program[2+position]
	var outputPosition int = program[3+position]

	if operation == 1 {
		program[outputPosition] = program[input1Position] + program[input2Position]
	}

	if operation == 2 {
		program[outputPosition] = program[input1Position] * program[input2Position]
	}

	return intcode(program, position+4)
}

func main() {
	data, loadErr := ioutil.ReadFile("./part1.txt")
	if loadErr != nil {
		fmt.Println("Load data error")
	}
	tmp := strings.Split(string(data), ",")
	intCodeValues := make([]int, 0, len(tmp))
	for _, raw := range tmp {
		intVal, parsingErr := strconv.Atoi(raw)
		if parsingErr != nil {
			fmt.Println("Parsing error", parsingErr)
			continue
		}
		intCodeValues = append(intCodeValues, intVal)
	}

	intCodeValues = intcode(intCodeValues, 0)

	fmt.Println(intCodeValues[0])
}
