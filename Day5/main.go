package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack struct {
	id        int
	resources []string
}

func getMyInput(fileName string) []string {
	var input []string
	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

func getUpperResource(currentStack stack) (string, stack) {
	resource := currentStack.resources[len(currentStack.resources)-1]
	stackContent := currentStack.resources[:len(currentStack.resources)-1]
	return resource, stack{
		id:        currentStack.id,
		resources: stackContent,
	}
}

func addResourceToStack(currentstack stack, movingResource string) stack {
	var content []string = currentstack.resources
	content = append(content, movingResource)
	currentstack.resources = content
	return currentstack
}

func moveItems(currentstack []stack, amount int, startStackID int, endStackID int) []stack {
	for i := 0; i < amount; i++ {
		resource, newStack := getUpperResource(currentstack[startStackID-1])
		currentstack[startStackID-1] = newStack

		newStackAdded := addResourceToStack(currentstack[endStackID-1], resource)
		currentstack[endStackID-1] = newStackAdded
	}
	return currentstack
}

func getResult(currentstack []stack) string {
	result := ""
	for i := 0; i < len(currentstack); i++ {
		upperResource := currentstack[i].resources[len(currentstack[i].resources)-1]
		result += upperResource
	}
	return result
}

func main() {
	var currentStacks []stack = []stack{
		{id: 1, resources: []string{"P", "F", "M", "Q", "W", "G", "R", "T"}},
		{id: 2, resources: []string{"H", "F", "R"}},
		{id: 3, resources: []string{"P", "Z", "R", "V", "G", "H", "S", "D"}},
		{id: 4, resources: []string{"Q", "H", "P", "B", "F", "W", "G"}},
		{id: 5, resources: []string{"P", "S", "M", "J", "H"}},
		{id: 6, resources: []string{"M", "Z", "T", "H", "S", "R", "P", "L"}},
		{id: 7, resources: []string{"P", "T", "H", "N", "M", "L"}},
		{id: 8, resources: []string{"F", "D", "Q", "R"}},
		{id: 9, resources: []string{"D", "S", "C", "N", "L", "P", "H"}},
	}
	var input = getMyInput("/Users/sinah/Code/AdventOfCode2022/Day5/input.txt")

	for i := 0; i < len(input); i++ {

		for i := 0; i < len(input); i++ {
			rearrangements := strings.Split(input[i], " ")
			amount, _ := strconv.Atoi(rearrangements[1])
			startStackID, _ := strconv.Atoi(rearrangements[3])
			endStackID, _ := strconv.Atoi(rearrangements[5])
			currentStacks = moveItems(currentStacks, amount, startStackID, endStackID)
		}
	}
	fmt.Print(getResult(currentStacks))
}
