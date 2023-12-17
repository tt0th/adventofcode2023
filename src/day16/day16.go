package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	. "github.com/tt0th/adventofcode2023/src/utils"
	"os"
)

const EMPTY = '.'
const FORWARD_SLASH_MIRROR = '/'
const BACK_SLASH_MIRROR = '\\'
const HORIZONTAL_SPLITTER = '-'
const VERTICAL_SPLITTER = '|'

func main() {
	//path := "src/day16/test-input-1.txt"
	path := "src/day16/input.txt"
	var contraption = parseInput(path)

	initialBeam := Beam{position: C{I: 0, J: -1}, direction: Right}
	simulateBeamTravel(initialBeam, contraption)
	printContraption(contraption)

	sum := countEnergizedTiles(contraption)
	fmt.Printf("sum: %d\n", sum)
}

func printContraption(contraption [][]Tile) {
	for i := 0; i < len(contraption); i++ {
		for j := 0; j < len(contraption[0]); j++ {
			if contraption[i][j].energyLevel > 0 {
				print("#")
			} else {
				print(".")
			}
		}
		println("")
	}
}

func countEnergizedTiles(contraption [][]Tile) int {
	var count int
	for _, tiles := range contraption {
		for _, tile := range tiles {
			if tile.energyLevel > 0 {
				count++
			}
		}
	}
	return count
}

type Tile struct {
	content             rune
	energyLevel         int
	beamEnterDirections []Direction
}
type Beam struct {
	position  C
	direction Direction
}

func simulateBeamTravel(beam Beam, contraption [][]Tile) {
	position := beam.position.Add(beam.direction)
	if position.I < 0 || position.J < 0 || position.I >= len(contraption) || position.J >= len(contraption[0]) {
		return
	}
	tile := &contraption[position.I][position.J]
	if lo.Contains(tile.beamEnterDirections, beam.direction) {
		return
	}
	tile.beamEnterDirections = append(tile.beamEnterDirections, beam.direction)
	tile.energyLevel++

	switch tile.content {
	case EMPTY:
		simulateBeamTravel(Beam{position: position, direction: beam.direction}, contraption)
	case FORWARD_SLASH_MIRROR:
		if beam.direction == Left {
			simulateBeamTravel(Beam{position: position, direction: Down}, contraption)
		} else if beam.direction == Right {
			simulateBeamTravel(Beam{position: position, direction: Up}, contraption)
		} else if beam.direction == Up {
			simulateBeamTravel(Beam{position: position, direction: Right}, contraption)
		} else if beam.direction == Down {
			simulateBeamTravel(Beam{position: position, direction: Left}, contraption)
		}
	case BACK_SLASH_MIRROR:
		if beam.direction == Left {
			simulateBeamTravel(Beam{position: position, direction: Up}, contraption)
		} else if beam.direction == Right {
			simulateBeamTravel(Beam{position: position, direction: Down}, contraption)
		} else if beam.direction == Up {
			simulateBeamTravel(Beam{position: position, direction: Left}, contraption)
		} else if beam.direction == Down {
			simulateBeamTravel(Beam{position: position, direction: Right}, contraption)
		}
	case HORIZONTAL_SPLITTER:
		if beam.direction == Left || beam.direction == Right {
			simulateBeamTravel(Beam{position: position, direction: beam.direction}, contraption)
		} else if beam.direction == Up || beam.direction == Down {
			simulateBeamTravel(Beam{position: position, direction: Left}, contraption)
			simulateBeamTravel(Beam{position: position, direction: Right}, contraption)
		}
	case VERTICAL_SPLITTER:
		if beam.direction == Up || beam.direction == Down {
			simulateBeamTravel(Beam{position: position, direction: beam.direction}, contraption)
		} else if beam.direction == Left || beam.direction == Right {
			simulateBeamTravel(Beam{position: position, direction: Up}, contraption)
			simulateBeamTravel(Beam{position: position, direction: Down}, contraption)
		}
	}
}

func parseInput(path string) [][]Tile {
	file, _ := os.Open(path)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var inputs [][]Tile

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		tiles := lo.Map(runes, func(item rune, _ int) Tile {
			return Tile{content: item}
		})
		inputs = append(inputs, tiles)
	}

	return inputs
}
