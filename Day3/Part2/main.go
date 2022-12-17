package main

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//opens the input file
	f, err := os.Open("/Users/sinah/Code/AdventOfCode/Day3/input.txt")
	check(err)

	scanner := bufio.NewScanner(f)
	score := make([]int, 0)

}
