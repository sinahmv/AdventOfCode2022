package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMyInput(fileName string) [][]string {
	var input [][]string
	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, strings.Split(scanner.Text(), " "))
	}
	return input
}

func CalculateDirSize(DirName string, InputFile [][]string) int {
	var DirSize int = 0
	var ReachedCurrentDir bool = false
	for _, line := range InputFile {
		// End of dir ls is reached
		if ReachedCurrentDir && line[1] == "cd" {
			ReachedCurrentDir = false
			if DirSize <= 100000 {
				result += DirSize
			}
			return DirSize
		}
		// Add File size to dir size
		if ReachedCurrentDir && line[0] != "dir" && line[0] != "ls" {
			size, _ := strconv.Atoi(line[0])
			DirSize += size
		}
		// Calculate subfolder size
		if ReachedCurrentDir && line[0] == "dir" {
			DirSize += CalculateDirSize(line[1], InputFile)
		}
		// Directory is reached
		if line[0] == "$" && line[1] == "cd" && line[2] == DirName {
			ReachedCurrentDir = true
		}
	}
	return 1
}

var result int = 0

func main() {
	var InputFile = getMyInput("/Users/sinah/Code/AdventOfCode2022/Day7/input.txt")
	fmt.Print(InputFile)
}
