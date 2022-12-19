package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func main() {
	score := make([]int, 0)
	f, err := os.Open("/Users/sinah/Code/AdventOfCode2022/Day2/input.txt")
	check(err)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		line := scanner.Text()
		plays := strings.Split(line, "/n")

		for i := 0; i < len(plays); i++ {
			innerPlay := strings.Split(plays[i], " ")
			for j := 0; j <= 1; j++ {
				if innerPlay[j] == "A" {
					if innerPlay[j+1] == "X" {
						score = append(score, 4)
					} else if innerPlay[j+1] == "Y" {
						score = append(score, 8)
					} else if innerPlay[j+1] == "Z" {
						score = append(score, 3)
					}
				} else if innerPlay[j] == "B" {
					if innerPlay[j+1] == "X" {
						score = append(score, 1)
					} else if innerPlay[j+1] == "Y" {
						score = append(score, 5)
					} else if innerPlay[j+1] == "Z" {
						score = append(score, 9)
					}
				} else if innerPlay[j] == "C" {
					if innerPlay[j+1] == "X" {
						score = append(score, 7)
					} else if innerPlay[j+1] == "Y" {
						score = append(score, 2)
					} else if innerPlay[j+1] == "Z" {
						score = append(score, 6)
					}
				}
			}
		}
	}
	fmt.Print("Mein Score: ", sum(score))
}
