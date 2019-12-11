package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func fuelForFuel(moduleMass int64, extra int64) int64 {
	if extra <= 0 {
		return 0
	}
	return extra + fuelForFuel(moduleMass+extra, extra/3-2)
}

func main() {
	file, fileError := os.Open("./input.txt")
	if fileError != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var totalFuelRequirement int64 = 0

	for scanner.Scan() {
		mass, parseError := strconv.ParseInt(scanner.Text(), 10, 64)
		if parseError != nil {
			fmt.Println("Parse string to int error")
		}
		//moduleFuelReq := mass/3 - 2
		moduleFuelReq := fuelForFuel(mass/3-2, mass/3-2)

		totalFuelRequirement += moduleFuelReq
	}

	fmt.Println(totalFuelRequirement)
}
