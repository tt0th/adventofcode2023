package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	id    string
	left  string
	right string
}

func main() {
	//path := "src/day8/test-input-1.txt"
	//path := "src/day8/test-input-2.txt"
	path := "src/day8/input.txt"
	var nodes, directions = parseInput(path)

	requiredSteps := getStepsRequiredToZZZ(directions, nodes)
	fmt.Printf("Steps to ZZZ: %d\n", requiredSteps)
}

func getStepsRequiredToZZZ(directions []rune, nodes map[string]Node) int {
	currentNode := "AAA"
	var i int
	for i = 0; currentNode != "ZZZ"; i++ {
		if directions[i%len(directions)] == 'L' {
			currentNode = nodes[currentNode].left
		} else {
			currentNode = nodes[currentNode].right
		}
	}
	return i
}

func parseInput(path string) (map[string]Node, []rune) {
	file, _ := os.Open(path)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	nodes := make(map[string]Node)
	var directions []rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.Contains(line, " = ") {
			splitByEqualSign := strings.Split(line, " = ")
			currentNodeId := splitByEqualSign[0]
			targetsPart := strings.Trim(splitByEqualSign[1], "()")
			targets := strings.Split(targetsPart, ", ")
			nodes[currentNodeId] = Node{
				id:    currentNodeId,
				left:  targets[0],
				right: targets[1],
			}
			continue
		}
		directions = []rune(line)
	}

	return nodes, directions
}
