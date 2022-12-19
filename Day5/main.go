package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type stack struct {
	id  int
	content []string
}

func moveItems() {

}

func main() {

	/*var ship []stack = []stack{
		{id: 1, content: []string{"T", "R", "G", "W", "Q", "M", "F", "P"}},
		{id: 2, content: []string{"R", "F", "H"}},
		{id: 3, content: []string{"D", "S", "H", "G", "V", "R", "Z", "P"}},
		{id: 4, content: []string{"G", "W", "F", "B", "P", "H", "Q"}},
		{id: 5, content: []string{"H", "J", "M", "S", "P"}},
		{id: 6, content: []string{"L", "P", "R", "S", "H", "T", "Z", "M"}},
		{id: 7, content: []string{"L", "M", "N", "H", "T", "P"}},
		{id: 8, content: []string{"R", "Q", "D", "F"}},
		{id: 9, content: []string{"H", "P", "L", "N", "C", "S", "D"}},
	}*/

	f, _ := os.Open("/Users/sinah/Code/AdventOfCode2022/Day5/input.txt")
	scanner := bufio.NewScanner(f)
	result := make([]string, 0)

	for scanner.Scan() {

		input := strings.Split(scanner.Text(), "/n")
		
		

	}
	fmt.Print(result)
}
