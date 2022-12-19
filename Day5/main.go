package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type ship struct {
	id  int
	resources []string
}

func removeResourceFromStack(currentShip []ship, amount int, startStackID int) []string {
	stack := currentShip[startStackID-1].resources
	movingResources := make([]string, 0)
	for i := len(stack)-1; i >= amount; i-- {
		movingResources = append(movingResources, stack[i])
		if len(stack) > 0 {
			stack = stack[0:len(stack)-1]
		} else {
			continue
		}
	}
	currentShip[startStackID-1].resources = stack
	return movingResources
}

func addResourcesToStack(currentShip []ship, endStackID int, movingResources []string) []ship {
	stack := currentShip[endStackID-1].resources 
	stack = append(stack, movingResources...)
	currentShip[endStackID-1].resources = stack
	return currentShip
}

func moveItems(currentShip []ship, amount int, startStackID int, endStackID int) []ship {
	movingResources := removeResourceFromStack(currentShip, amount, startStackID)
	currentShip = addResourcesToStack(currentShip, endStackID, movingResources)
	return currentShip
}

func getResult(currentShip []ship) []string {
	result := make([]string, 0)
	upperResource := ""
	for i := 0; i < len(currentShip); i++ {
		upperResource = currentShip[i].resources[len(currentShip[i].resources)-1]
		result = append(result, upperResource)
	}
	return result
}

func main() {

	f, _ := os.Open("/Users/sinah/Code/AdventOfCode2022/Day5/input.txt")
	scanner := bufio.NewScanner(f)
	var currentShip []ship = []ship{
		{id: 1, resources: []string{"T", "R", "G", "W", "Q", "M", "F", "P"}},
		{id: 2, resources: []string{"R", "F", "H"}},
		{id: 3, resources: []string{"D", "S", "H", "G", "V", "R", "Z", "P"}},
		{id: 4, resources: []string{"G", "W", "F", "B", "P", "H", "Q"}},
		{id: 5, resources: []string{"H", "J", "M", "S", "P"}},
		{id: 6, resources: []string{"L", "P", "R", "S", "H", "T", "Z", "M"}},
		{id: 7, resources: []string{"L", "M", "N", "H", "T", "P"}},
		{id: 8, resources: []string{"R", "Q", "D", "F"}},
		{id: 9, resources: []string{"H", "P", "L", "N", "C", "S", "D"}},
	}

	for scanner.Scan() {

		input := strings.Split(scanner.Text(), "/n")
		//fmt.Print(input)

		for i := 0; i < len(input); i++ {
			rearrangements := strings.Split(input[i], " ")
			amount, _ := strconv.Atoi(rearrangements[1])
			startStackID, _ := strconv.Atoi(rearrangements[3])
			endStackID, _ := strconv.Atoi(rearrangements[5])
			moveItems(currentShip, amount, startStackID, endStackID)
		}

	}
	fmt.Print(getResult(currentShip))
}
