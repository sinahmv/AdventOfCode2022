package main

import (
	"bufio"
	"fmt"
	"os"
)

func getMyInput(fileName string) string {
	var input string
	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = scanner.Text()
	}
	return input
}

func differentLetters(letters string) bool {
	for i := 0; i < len(letters); i++ {
		for j := i + 1; j < len(letters); j++ {
			if letters[i] == letters[j] {
				return false
			}
		}
	}
	return true
}

func part1(input string) int {
	for i := 0; i+4 <= len(input); i++ {
		if differentLetters(input[i : i+4]) {
			return i + 4
		}
	}

	return 0
}

func part2(input string) int {
	for i := 0; i+14 <= len(input); i++ {
		if differentLetters(input[i : i+14]) {
			return i + 14
		}
	}

	return 0
}

func main() {
	var input = getMyInput("/Users/sinah/Code/AdventOfCode2022/Day6/input.txt")
	count1 := part1(input)
	count2 := part2(input)
	fmt.Println("Count1:", count1)
	fmt.Println("Count2:", count2)
}
