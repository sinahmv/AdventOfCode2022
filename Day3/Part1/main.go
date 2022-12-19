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

func getSimilarString(substring1 string, substring2 string) string {
	var matchingString string
	for i := 0; i < len(substring1); i++ {
		string2 := substring1[i : i+1]
		if strings.Contains(substring2, string2) {
			matchingString = string2
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
	score := make([]int, 0)

	for scanner.Scan() {

		input := strings.Split(scanner.Text(), "/n")

		currentinput := input[0]
		department1 := currentinput[0 : len(currentinput)/2]
		department2 := currentinput[len(currentinput)/2:]
		fmt.Print("String 1: ", department1, " String 2: ", department2, " ")
		matchingString := getSimilarString(department1, department2)
		fmt.Print("Match: ", matchingString, " ")
		fmt.Println("Priority: ", getPriority(matchingString), " ")
		score = append(score, int(getPriority(matchingString)))

	}
	fmt.Print("Score: ", sum(score))
}
