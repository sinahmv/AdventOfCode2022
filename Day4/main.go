package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	f, err := os.Open("/Users/sinah/Code/AdventOfCode2022/Day4/input.txt")
	check(err)
	scanner := bufio.NewScanner(f)

	threeElves := make([]string, 0)
	score := 0

	for scanner.Scan() {

		input := strings.Split(scanner.Text(), ",")
		threeElves = append(threeElves, input...)
		fmt.Println(threeElves)

		if len(threeElves) == 2 {

			range1 := strings.Split(threeElves[0], "-")
			range2 := strings.Split(threeElves[1], "-")

			int1, _ := strconv.Atoi(range1[0])
			int2, _ := strconv.Atoi(range1[1])
			int3, _ := strconv.Atoi(range2[0])
			int4, _ := strconv.Atoi(range2[1])

			fmt.Println(int1, int2, int3, int4)

			if int1 <= int3 && int2 >= int4 {
				fmt.Println("Hurra1")
				score++
			} else if int3 <= int1 && int4 >= int2 {
				fmt.Println("Hurra2")
				score++
			} else if int2 >= int3 && int4 >= int1 { //Abfrage für Part2
				fmt.Println("Hurra3")
				score++
			}
			threeElves = make([]string, 0)
		}

	}
	fmt.Print(score)
}
