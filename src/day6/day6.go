package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Race struct {
	duration          int
	recordDistance    int
	holdingTimesToWin []int
}

type RaceOutcome struct {
	holdingTime int
	distance    int
}

func main() {
	//path := "src/day6/test-input-1.txt"
	path := "src/day6/input.txt"
	var raceRecords = parseInput(path)

	raceRecords = findWaysToWin(raceRecords)
	printObject(raceRecords)
	product := lo.Reduce(raceRecords, func(agg int, race Race, index int) int {
		return agg * len(race.holdingTimesToWin)
	}, 1)

	fmt.Printf("product: %d\n", product)

	var race = Race{duration: 55999793, recordDistance: 401148522741405}
	var races = findWaysToWin([]Race{race})
	fmt.Printf("ways to win: %d\n", len(races[0].holdingTimesToWin))
}

func printObject(object interface{}) {
	fmt.Printf("%v\n", object)
}

func findWaysToWin(races []Race) []Race {
	for i := 0; i < len(races); i++ {
		holdingTimes := lo.Map(make([]int, races[i].duration), func(_ int, index int) int {
			return index
		})
		distances := lo.Map(holdingTimes, distanceCalculator(races[i].duration))
		outcomes := lo.Map(lo.Zip2(holdingTimes, distances), func(item lo.Tuple2[int, int], index int) RaceOutcome {
			return RaceOutcome{holdingTime: item.A, distance: item.B}
		})
		winningOutcomes := lo.Filter(outcomes, func(outcome RaceOutcome, _ int) bool {
			return outcome.distance > races[i].recordDistance
		})
		races[i].holdingTimesToWin = lo.Map(winningOutcomes, func(outcome RaceOutcome, _ int) int {
			return outcome.holdingTime
		})
	}
	return races
}

func distanceCalculator(duration int) func(holdingTime int, _ int) int {
	return func(holdingTime int, _ int) int {
		return holdingTime * (duration - holdingTime)
	}
}

func parseInput(path string) []Race {
	file, _ := os.Open(path)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var times []int
	var distances []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Time:") {
			splitByColon := strings.Split(line, ": ")
			times = stringToIntArray(splitByColon[1])
		}
		if strings.Contains(line, "Distance:") {
			splitByColon := strings.Split(line, ": ")
			distances = stringToIntArray(splitByColon[1])
		}
	}

	return lo.Map(lo.Zip2(times, distances), func(item lo.Tuple2[int, int], index int) Race {
		return Race{duration: item.A, recordDistance: item.B}
	})
}

func stringToIntArray(input string) []int {
	numbersAsString := regexp.MustCompile(`\s+`).Split(strings.Trim(input, " "), -1)

	var numbers []int
	for _, numberAsString := range numbersAsString {
		numbers = append(numbers, unsafeParseInt(numberAsString))
	}
	return numbers
}

func unsafeParseInt(input string) int {
	number, _ := strconv.Atoi(input)
	return number
}
