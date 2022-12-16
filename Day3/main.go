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
	f, err := os.Open("/Users/sinah/Code/AdventOfCode/Day3/input.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	score := make([]int, 0)

	for scanner.Scan() {

		line := scanner.Text()
		backpack := strings.Split(line, "/n")

		for i := 0; i < len(backpack); i++ {
			currentBackpack := backpack[i]
			department1 := currentBackpack[0 : len(currentBackpack)/2]
			department2 := currentBackpack[len(currentBackpack)/2:]
			//fmt.Print("String 1: ", department1, " String 2: ", department2)
			matchingString := getSimilarString(department1, department2)
			//fmt.Print(" Match: ", matchingString, " ")
			//fmt.Print(getPriority(matchingString), " ")
			score = append(score, int(getPriority(matchingString)))
		}

	}
	fmt.Print(sum(score))
}
