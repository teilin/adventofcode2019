package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func loadData(path string) map[int]int {
	data, loadErr := ioutil.ReadFile(path)
	if loadErr != nil {
		fmt.Println("Load data error")
	}
	tmp := strings.Split(string(data), ",")
	var intCodeValues map[int]int = make(map[int]int)
	var index int = 0
	for _, raw := range tmp {
		intVal, parsingErr := strconv.Atoi(raw)
		if parsingErr != nil {
			fmt.Println("Parsing error", parsingErr)
			continue
		}
		intCodeValues[index] = intVal
		index += 1
	}
	return intCodeValues
}

func intcode(program map[int]int, position int) map[int]int {
	var operation int = program[0+position]

	var input1Position int = program[1+position]
	var input2Position int = program[2+position]
	var outputPosition int = program[3+position]

	if operation == 1 {
		program[outputPosition] = program[input1Position] + program[input2Position]
	}

	if operation == 2 {
		program[outputPosition] = program[input1Position] * program[input2Position]
	}
	if operation == 99 {
		return program
	}

	return intcode(program, position+4)
}

func part1() error {
	intCodeValues := loadData("./part1.txt")
	intCodeValues = intcode(intCodeValues, 0)
	fmt.Println(intCodeValues[0])
	return nil
}

func partTwo(noun int, verb int, result int) bool {
	intCodeValues := loadData("./part1.txt")
	intCodeValues[1] = noun
	intCodeValues[2] = verb
	intCodeValues = intcode(intCodeValues, 0)
	return intCodeValues[0] == result
}

func main() {
	var possibleValues [100]int
	for i := 0; i < 100; i++ {
		possibleValues[i] = i
	}

	for _, noun := range possibleValues {
		for _, verb := range possibleValues {
			isCorrect := partTwo(noun, verb, 19690720)
			resultCalculation := 100*noun + verb
			if isCorrect {
				fmt.Println("Correct answer is " + strconv.Itoa(resultCalculation))
			}
		}
	}
}
