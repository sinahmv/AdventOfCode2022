package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part2(string1 string, string2 string, string3 string) string {
	var matchingString string
	for i := 0; i < len(string1); i++ {
		letter := string1[i : i+1]
		if strings.Contains(string2, letter) && strings.Contains(string3, letter) {
			matchingString = letter
		}
	}
	return matchingString
}

func getPriority(matchingString string) rune {
	var myPrio rune
	byteArray := []byte(matchingString)
	rs := bytes.Runes(byteArray)
	myByte := rs[0]
	if myByte >= 97 {
		myPrio = myByte - 96
	} else {
		myPrio = myByte - 38
	}
	return myPrio
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func main() {

	f, err := os.Open("/Users/sinah/Code/AdventOfCode2022/Day3/input.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	threeElves := make([]string, 0)
	score := make([]int, 0)

	for scanner.Scan() {

		input := strings.Split(scanner.Text(), "/n")
		threeElves = append(threeElves, input...)

		if len(threeElves) == 3 {
			elf1 := threeElves[0]
			elf2 := threeElves[1]
			elf3 := threeElves[2]
			matchingString := part2(elf1, elf2, elf3)
			score = append(score, int(getPriority(matchingString)))
			threeElves = make([]string, 0)
		}

	}
	fmt.Print("Score: ", sum(score))

}
