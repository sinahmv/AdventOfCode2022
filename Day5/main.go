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

func moveItems(ship []ship, amount int, startship int, endship int) []ship {

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

		for i := 0; i < len(input); i++ {
			rearrangements := strings.Split(input[i], " ")
			amount, _ := strconv.Atoi(rearrangements[1])
			startship, _ := strconv.Atoi(rearrangements[3])
			endship, _ := strconv.Atoi(rearrangements[5])
			moveItems(currentShip, amount, startship, endship)
		}

	}
	fmt.Print(getResult(currentShip))
}
