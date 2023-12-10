package main

import (
	"bufio"
	"fmt"
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
	path := "src/day10/input.txt"
	var maze = parseInput(path)

	var stepsToFarthest int
	loop := getLoop(maze)
	stepsToFarthest = len(loop) / 2

	fmt.Printf("stepsToFarthest: %d\n", stepsToFarthest)
}

func getLoop(maze Maze) []C {
	var route []C
	position := getStartPosition(maze)
	direction := guessStartDirection(maze, position)
	for maze.get(position) != 'S' || len(route) == 0 {
		route = append(route, position)
		position = position.add(direction)
		direction = getNewDirection(maze.get(position), direction)
	}
	return route
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
