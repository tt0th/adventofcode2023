package main

import (
	"bufio"
	"fmt"
	"os"
)

const ROCK = 'O'
const EMPTY = '.'

func main() {
	//path := "src/day14/test-input-1.txt"
	path := "src/day14/input.txt"
	var platform = parseInput(path)

	sum := summarizePlatformWeightOnNorth(platform)

	const spinCount = 1000000000
	platform = parseInput(path)
	loop := spin(platform, spinCount)
	platform = parseInput(path)
	spin(platform, loop.start+(spinCount-loop.start)%loop.length)

	sum2 := summarizePlatformWeightOnNorth(platform)

	fmt.Printf("sum: %d, sum2: %d\n", sum, sum2)
}

func summarizePlatformWeightOnNorth(platform [][]rune) int {
	var sum int
	for i, row := range platform {
		for _, column := range row {
			if column == ROCK {
				sum += len(platform) - i
			}
		}
	}
	return sum
}

type Loop struct {
	start  int
	length int
}

func spin(platform [][]rune, maxTimes int) Loop {
	var platformStateHistory [][][]rune
	for i := 0; i < maxTimes; i++ {
		rollToNorth(platform)
		rollToWest(platform)
		rollToSouth(platform)
		rollToEast(platform)
		for historyIndex, prevPlatform := range platformStateHistory {
			if arePlatformsEqual(platform, prevPlatform) {
				return Loop{
					start:  historyIndex,
					length: i - historyIndex,
				}
			}
		}
		platformStateHistory = append(platformStateHistory, copyPlatform(platform))
	}
	return Loop{}
}

func arePlatformsEqual(platform [][]rune, otherPlatform [][]rune) bool {
	for i := 0; i < len(otherPlatform); i++ {
		for j := 0; j < len(otherPlatform[0]); j++ {
			if platform[i][j] != otherPlatform[i][j] {
				return false
			}
		}
	}
	return true
}
func copyPlatform(platform [][]rune) [][]rune {
	platformCopy := make([][]rune, len(platform))
	for i := 0; i < len(platform); i++ {
		platformCopy[i] = make([]rune, len(platform[0]))
		copy(platformCopy[i], platform[i])
	}
	return platformCopy
}

func rollToNorth(platform [][]rune) {
	wasChanged := true
	for wasChanged {
		wasChanged = false
		for i := 0; i < len(platform)-1; i++ {
			for j := 0; j < len(platform[0]); j++ {
				if platform[i][j] == EMPTY && platform[i+1][j] == ROCK {
					wasChanged = true
					platform[i+1][j] = EMPTY
					platform[i][j] = ROCK
				}
			}
		}
	}
}

func rollToSouth(platform [][]rune) {
	wasChanged := true
	for wasChanged {
		wasChanged = false
		for i := len(platform) - 1; i > 0; i-- {
			for j := 0; j < len(platform[0]); j++ {
				if platform[i][j] == EMPTY && platform[i-1][j] == ROCK {
					wasChanged = true
					platform[i-1][j] = EMPTY
					platform[i][j] = ROCK
				}
			}
		}
	}
}

func rollToEast(platform [][]rune) {
	wasChanged := true
	for wasChanged {
		wasChanged = false
		for j := len(platform[0]) - 1; j > 0; j-- {
			for i := 0; i < len(platform); i++ {
				if platform[i][j] == EMPTY && platform[i][j-1] == ROCK {
					wasChanged = true
					platform[i][j-1] = EMPTY
					platform[i][j] = ROCK
				}
			}
		}
	}
}

func rollToWest(platform [][]rune) {
	wasChanged := true
	for wasChanged {
		wasChanged = false
		for j := 0; j < len(platform[0])-1; j++ {
			for i := 0; i < len(platform); i++ {
				if platform[i][j] == EMPTY && platform[i][j+1] == ROCK {
					wasChanged = true
					platform[i][j+1] = EMPTY
					platform[i][j] = ROCK
				}
			}
		}
	}
}

func parseInput(path string) [][]rune {
	file, _ := os.Open(path)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var inputs [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputs = append(inputs, []rune(line))
	}

	return inputs
}
