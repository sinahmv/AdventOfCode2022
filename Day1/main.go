package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	top3 := []int{0, 0, 0}
	cals := 0
	f, err := os.Open("/Users/sinah/Code/AdventOfCode/Day1/input")
	check(err)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() || cals != 0 {
		line := scanner.Text()
		if line != "" {
			calsInLine, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalf("bad item %q: %v", line, err)
			}
			cals += calsInLine
			continue
		}

		for i := range top3 {
			if cals <= top3[i] {
				continue
			}

			if i < len(top3)-1 {
				copy(top3[i+1:], top3[i:])
			}

			top3[i] = cals
			break
		}

		cals = 0
	}

	fmt.Printf("Die Lösung ist: %d, Die Top3 sind: %v\n", top3[0]+top3[1]+top3[2], top3)
}
