package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	name     string
	size     int
	isFile   bool
	children map[string]*node
	parent   *node
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

func calculateSize(root node) (size int) {
	if root.isFile {
		return root.size
	}
	for _, d := range root.children {
		size += calculateSize(*d)
	}
	return
}

func parseInput(input [][]string, currentDirectory *node, directories []*node) {

	for _, line := range input {

		if len(line) > 2 {
			//geht zum Eltern Directory
			if line[2] == ".." {
				currentDirectory = currentDirectory.parent
				//geht zum initialen Directory
			} else if line[2] == "/" {
				currentDirectory = &node{"/", 0, false, make(map[string]*node), nil}
			} else {
				currentDirectory = currentDirectory.children[line[2]]
			}

		} else if line[0] == "dir" {
			//es wird ein neues Directory hinzugefügt (können unter anderem die gleichen Namen haben, sind aber trotzdem neu)
			currentDirectory.children[line[1]] = &node{line[1], 0, false, make(map[string]*node), currentDirectory}
			directories = append(directories, currentDirectory.children[line[1]])

		} else if line[0] != "$" {
			//wenn die Zeile nicht mit einem $ beginnt, gibt sie die Größe eines Files wieder
			size, _ := strconv.Atoi(line[0])
			currentDirectory.children[line[1]] = &node{line[1], size, true, nil, currentDirectory}
		}
	}

	var totalSize int

	for _, directories := range directories {
		size := calculateSize(*directories)
		if size <= 100000 {
			totalSize += size
		}
	}

	fmt.Println(totalSize)
}

func main() {
	var InputFile = getMyInput("/Users/sinah/Code/AdventOfCode2022/Day7/input.txt")
	//zeigt auf mein akutelles Directory
	var currentDirectory *node
	//Sammlung aller Directories an denen man vorbei geht
	directories := []*node{}

	parseInput(InputFile, currentDirectory, directories)
}
