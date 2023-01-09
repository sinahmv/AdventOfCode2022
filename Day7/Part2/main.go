package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type directory struct {
	name     string
	size     int
	isFile   bool
	children map[string]*directory
	parent   *directory
}

func getMyInput(fileName string) [][]string {
	var input [][]string
	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, strings.Split(scanner.Text(), " "))
	}
	return input
}

func parseInput(input [][]string, currentDirectory *directory, directories []*directory) {

	for _, line := range input {

		if len(line) > 2 {

			if line[2] == ".." {
				currentDirectory = currentDirectory.parent
			} else if line[2] == "/" {

				currentDirectory = &directory{"/", 0, false, make(map[string]*directory), nil}
				directories = append(directories, currentDirectory)
			} else {

				currentDirectory = currentDirectory.children[line[2]]
			}
		} else if line[0] == "dir" {

			currentDirectory.children[line[1]] = &directory{line[1], 0, false, make(map[string]*directory), currentDirectory}
			directories = append(directories, currentDirectory.children[line[1]])
		} else if line[0] != "$" {

			size, _ := strconv.Atoi(line[0])
			currentDirectory.children[line[1]] = &directory{line[1], size, true, nil, currentDirectory}
		}
	}

	toFree := 30000000 - (70000000 - calcSize(*directories[0]))
	fmt.Println("toFree: ", toFree)
	var smallestEnaugthSize int = calcSize(*directories[0])
	fmt.Println(smallestEnaugthSize)

	for _, dir := range directories {
		size := calcSize(*dir)
		if size > toFree && size-toFree < smallestEnaugthSize-toFree {
			smallestEnaugthSize = size
			fmt.Println(size)
		}
	}

	fmt.Println(smallestEnaugthSize)
}

func main() {
	var InputFile = getMyInput("/Users/sinah/Code/AdventOfCode2022/Day7/input.txt")
	//zeigt auf mein akutelles Directory
	var currentDirectory *directory
	//Sammlung aller Directories an denen man vorbei geht
	directories := []*directory{}

	parseInput(InputFile, currentDirectory, directories)
}

func calcSize(root directory) (size int) {
	if root.isFile {
		return root.size
	}
	for _, d := range root.children {
		size += calcSize(*d)
	}
	return
}
