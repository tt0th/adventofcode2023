package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"math"
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
	//path := "src/day8/test-input-3.txt"
	path := "src/day8/input.txt"
	var nodes, directions = parseInput(path)

	requiredSteps := getStepsRequiredToZZZ(directions, nodes)
	fmt.Printf("Steps to ZZZ: %d\n", requiredSteps)

	currentNodes := lo.Filter(lo.Keys(nodes), func(nodeId string, _ int) bool {
		return []rune(nodeId)[2] == 'A'
	})
	loops := lop.Map(currentNodes, func(nodeId string, _ int) Loop {
		return getLoop(directions, nodes, nodeId)
	})
	printObject(loops)
	printObject(lo.Reduce(loops, func(agg int, loop Loop, _ int) int {
		return lcm(agg, loop.length)
	}, 1))
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

type Step struct {
	index          int
	directionIndex int
	node           string
}
type Loop struct {
	start           int
	length          int
	solutionIndexes []int
}

func getLoop(directions []rune, nodes map[string]Node, startNode string) Loop {
	currentNode := startNode
	var solutionIndexes []int
	var steps []Step
	for i := 0; true; i++ {
		step, _, found := lo.FindIndexOf(steps, func(step Step) bool {
			return step.directionIndex == i%len(directions) && step.node == currentNode
		})
		if found {
			return Loop{
				start:           step.index,
				length:          i - step.index,
				solutionIndexes: solutionIndexes,
			}
		}
		if []rune(currentNode)[2] == 'Z' {
			solutionIndexes = append(solutionIndexes, i)
		}
		steps = append(steps, Step{index: i, directionIndex: i % len(directions), node: currentNode})
		if directions[i%len(directions)] == 'L' {
			currentNode = nodes[currentNode].left
		} else {
			currentNode = nodes[currentNode].right
		}
	}
	return Loop{}
}

func printObject(object interface{}) {
	fmt.Printf("%v\n", object)
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

// lcm and gcd from https://github.com/TheAlgorithms/Go
func lcm(a, b int) int {
	return int(math.Abs(float64(a*b)) / float64(gcd(a, b)))
}
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
