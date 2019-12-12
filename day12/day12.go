package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
	z int
}

const (
	rounds int = 10
)

var (
	posititions []Coordinate = make([]Coordinate, 4)
	velocitys   []Coordinate = make([]Coordinate, 4)
)

func readInput(path string) error {
	file, fileError := os.Open(path)
	if fileError != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	index := 0
	for scanner.Scan() {
		coor := parseCoordinate(scanner.Text())
		posititions[index] = coor
		velocitys[index] = Coordinate{x: 0, y: 0, z: 0}
		index += 1
	}
	return fileError
}

func splitString(str string, index int, sep rune) string {
	ss := strings.FieldsFunc(str, func(r rune) bool {
		if r == sep {
			return true
		} else {
			return false
		}
	})
	return ss[index]
}

func trimLastCharFromString(str string) string {
	if len(str) > 0 {
		return str[:len(str)-1]
	} else {
		return str
	}
}

func parseCoordinate(coor string) Coordinate {
	xCor, _ := strconv.Atoi(splitString(splitString(coor, 0, ','), 1, '='))
	yCor, _ := strconv.Atoi(splitString(splitString(coor, 1, ','), 1, '='))
	zCor, _ := strconv.Atoi(trimLastCharFromString(splitString(splitString(coor, 2, ','), 1, '=')))
	return Coordinate{x: xCor, y: yCor, z: zCor}
}

func applyVelocity(position Coordinate, velocity Coordinate) Coordinate {
	position.x += velocity.x
	position.y += velocity.y
	position.z += velocity.z
	return position
}

func part1(r int) {
	readError := readInput("./example.txt")
	if readError != nil {
		fmt.Println("Error reading input file and parsing")
	}

	for i := 0; i < r; i++ {
		fmt.Println(i)
	}
}

func main() {
	part1(rounds)
}
