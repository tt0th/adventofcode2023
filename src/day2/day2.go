package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Outcome struct {
	red   int
	green int
	blue  int
}

type Game struct {
	gameId   int
	outcomes []Outcome
}

func main() {
	//path := "src/day2/test-input-1.txt"
	path := "src/day2/input.txt"
	file, _ := os.Open(path)
	defer file.Close()

	var sum = 0
	var sum2 = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		game := parseGame(scanner.Text())
		if isGameValid(game) {
			sum += game.gameId
		}
		sum2 += powerOfGame(game)
	}

	fmt.Printf("Sum: %d, sum2: %d\n", sum, sum2)
}

func powerOfGame(game Game) int {
	var reds []int
	var greens []int
	var blues []int
	for _, outcome := range game.outcomes {
		reds = append(reds, outcome.red)
		greens = append(greens, outcome.green)
		blues = append(blues, outcome.blue)
	}
	return slices.Max(reds) * slices.Max(greens) * slices.Max(blues)
}

// Game 100: 2 red, 9 green, 11 blue; 13 blue, 4 red, 16 green; 8 green, 13 blue; 10 green, 1 red, 12 blue
func parseGame(text string) Game {
	gameIdAndOutcomes := strings.Split(text, ":")
	gameId, _ := strconv.Atoi(strings.Split(gameIdAndOutcomes[0], " ")[1])
	outcomesParts := strings.Split(gameIdAndOutcomes[1], ";")
	var outcomes []Outcome
	for _, outcomePart := range outcomesParts {
		outcome := parseOutcome(outcomePart)
		outcomes = append(outcomes, outcome)
	}

	return Game{gameId: gameId, outcomes: outcomes}
}

func parseOutcome(outcomePart string) Outcome {
	outcome := Outcome{red: 0, green: 0, blue: 0}
	for _, outcomePerColorPart := range strings.Split(outcomePart, ",") {
		countPart := regexp.MustCompile(`[^0-9]+`).ReplaceAllString(outcomePerColorPart, "")
		count, _ := strconv.Atoi(countPart)
		if strings.Contains(outcomePerColorPart, "red") {
			outcome.red = count
		}
		if strings.Contains(outcomePerColorPart, "green") {
			outcome.green = count
		}
		if strings.Contains(outcomePerColorPart, "blue") {
			outcome.blue = count
		}
	}
	return outcome
}

func isGameValid(game Game) bool {
	for _, outcome := range game.outcomes {
		if outcome.red > 12 || outcome.green > 13 || outcome.blue > 14 {
			return false
		}
	}

	return true
}
