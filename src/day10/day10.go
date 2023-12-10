package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"github.com/tt0th/adventofcode2023/src/utils"
	"os"
)

type C struct {
	i int
	j int
}

func (coordinate C) add(other C) C {
	return C{
		i: coordinate.i + other.i,
		j: coordinate.j + other.j,
	}
}

var down = C{i: 1, j: 0}
var up = C{i: -1, j: 0}
var right = C{i: 0, j: 1}
var left = C{i: 0, j: -1}

type Maze struct {
	maze [][]rune
	j    int
}

func (maze Maze) get(coordinate C) rune {
	return maze.maze[coordinate.i][coordinate.j]
}
func (maze Maze) getByIndex(i int, j int) rune {
	return maze.maze[i][j]
}
func (maze Maze) getHeight() int {
	return len(maze.maze)
}
func (maze Maze) getWidth() int {
	return len(maze.maze[0])
}

func main() {
	//path := "src/day10/test-input-1.txt"
	//path := "src/day10/test-input-2.txt"
	//path := "src/day10/test-input-3.txt"
	//path := "src/day10/test-input-4.txt"
	path := "src/day10/input.txt"
	var maze = parseInput(path)

	loop, innerPoints := getLoop(maze)
	innerPoints = extendInnerPoints(innerPoints, loop)

	fmt.Printf("stepsToFarthest: %d, countOfInnerPoints: %d \n", len(loop)/2, len(innerPoints))
}

func getLoop(maze Maze) ([]C, []C) {
	var route []C
	var pointsToTheRight, pointsToTheLeft []C
	turns := make(map[string]int)
	position := getStartPosition(maze)
	direction := guessStartDirection(maze, position)
	for maze.get(position) != 'S' || len(route) == 0 {
		route = append(route, position)
		position = position.add(direction)
		pointsToTheRight = append(pointsToTheRight, getPointsToTheRight(maze.get(position), direction, position)...)
		pointsToTheLeft = append(pointsToTheLeft, getPointsToTheLeft(maze.get(position), direction, position)...)
		turns[getTurnDirection(maze.get(position), direction)]++
		direction = getNewDirection(maze.get(position), direction)
		//printMaze(maze, position, turnDirection)
	}
	var innerPoints []C
	if turns["L"] > turns["R"] {
		innerPoints = pointsToTheLeft
	} else {
		innerPoints = pointsToTheRight
	}
	innerPoints = lo.Filter(innerPoints, func(point C, _ int) bool {
		return !lo.Contains(route, point)
	})
	return route, innerPoints
}

func extendInnerPoints(points []C, excludedPoints []C) []C {
	var pointCollection []C
	pointsToVisit := points

	for len(pointsToVisit) > 0 {
		point := pointsToVisit[0]
		pointsToVisit = pointsToVisit[1:]
		if lo.Contains(excludedPoints, point) {
			continue
		}
		if lo.Contains(pointCollection, point) {
			continue
		}
		pointCollection = append(pointCollection, point)
		pointsToVisit = append(pointsToVisit, point.add(up), point.add(down), point.add(left), point.add(right))
	}

	return pointCollection
}

func getPointsToTheRight(pipe rune, direction C, position C) []C {
	if pipe == '|' {
		if direction == up {
			return []C{position.add(right)}
		}
		if direction == down {
			return []C{position.add(left)}
		}
	}
	if pipe == '-' {
		if direction == left {
			return []C{position.add(up)}
		}
		if direction == right {
			return []C{position.add(down)}
		}
	}
	if pipe == 'L' {
		if direction == left {
			return []C{}
		}
		if direction == down {
			return []C{position.add(down), position.add(left)}
		}
	}
	if pipe == 'J' {
		if direction == right {
			return []C{position.add(down), position.add(right)}
		}
		if direction == down {
			return []C{}
		}
	}
	if pipe == '7' {
		if direction == right {
			return []C{}
		}
		if direction == up {
			return []C{position.add(up), position.add(right)}
		}
	}
	if pipe == 'F' {
		if direction == left {
			return []C{position.add(up), position.add(left)}
		}
		if direction == up {
			return []C{}
		}
	}
	return []C{}
}

func getPointsToTheLeft(pipe rune, direction C, position C) []C {
	if pipe == '|' {
		if direction == up {
			return []C{position.add(left)}
		}
		if direction == down {
			return []C{position.add(right)}
		}
	}
	if pipe == '-' {
		if direction == left {
			return []C{position.add(down)}
		}
		if direction == right {
			return []C{position.add(up)}
		}
	}
	if pipe == 'L' {
		if direction == left {
			return []C{position.add(down), position.add(left)}
		}
		if direction == down {
			return []C{}
		}
	}
	if pipe == 'J' {
		if direction == right {
			return []C{}
		}
		if direction == down {
			return []C{position.add(down), position.add(right)}
		}
	}
	if pipe == '7' {
		if direction == right {
			return []C{position.add(up), position.add(right)}
		}
		if direction == up {
			return []C{}
		}
	}
	if pipe == 'F' {
		if direction == left {
			return []C{}
		}
		if direction == up {
			return []C{position.add(up), position.add(left)}
		}
	}
	return []C{}
}

func printMaze(maze Maze, position C, turnDirection string) {
	for i := 0; i < maze.getHeight(); i++ {
		for j := 0; j < maze.getWidth(); j++ {
			if position.i == i && position.j == j {
				fmt.Printf("\u2588")
			} else {
				fmt.Printf("%c", maze.getByIndex(i, j))
			}
		}
		fmt.Printf("    ")
		fmt.Printf(turnDirection)
		fmt.Printf("\n")
	}
	utils.PrintObject(turnDirection)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func getTurnDirection(pipe rune, lastDirection C) string {
	var directions = map[C]map[rune]string{
		down: {
			'L': "L",
			'J': "R",
		},
		up: {
			'7': "L",
			'F': "R",
		},
		left: {
			'L': "R",
			'F': "L",
		},
		right: {
			'J': "L",
			'7': "R",
		},
	}

	return directions[lastDirection][pipe]
}

func getNewDirection(pipe rune, lastDirection C) C {
	var directions = map[C]map[rune]C{
		down: {
			'|': down,
			'L': right,
			'J': left,
		},
		up: {
			'|': up,
			'7': left,
			'F': right,
		},
		left: {
			'-': left,
			'L': up,
			'F': down,
		},
		right: {
			'-': right,
			'J': up,
			'7': down,
		},
	}

	return directions[lastDirection][pipe]
}

func guessStartDirection(maze Maze, position C) C {
	below := maze.get(position.add(down))
	if below == '|' || below == 'L' || below == 'J' {
		return down
	}

	above := maze.get(position.add(up))
	if above == '|' || above == '7' || above == 'F' {
		return up
	}

	toTheLeft := maze.get(position.add(left))
	if toTheLeft == '-' || toTheLeft == 'L' || toTheLeft == 'F' {
		return left
	}

	toTheRight := maze.get(position.add(right))
	if toTheRight == '-' || toTheRight == 'J' || toTheRight == '7' {
		return right
	}

	panic("can not start at any direction :(")
}

func getStartPosition(maze Maze) C {
	for i := 0; i < maze.getHeight(); i++ {
		for j := 0; j < maze.getWidth(); j++ {
			if maze.getByIndex(i, j) == 'S' {
				return C{i, j}
			}
		}
	}
	panic("no starting point o.O")
}

func parseInput(path string) Maze {
	file, _ := os.Open(path)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var runes [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		runes = append(runes, []rune(line))
	}

	return Maze{maze: runes}
}
