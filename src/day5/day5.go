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

type Range struct {
	destinationStart int
	sourceStart      int
	length           int
}

type Mapping struct {
	ranges []Range
}

type Seed struct {
	seedId  int
	idChain []int
}

func main() {
	//path := "src/day5/test-input-1.txt"
	path := "src/day5/input.txt"
	var seedIds, mappings = parseInput(path)

	seeds := collectSeedProperties(seedIds, mappings)

	minLocation := lo.Min(lo.Map(seeds, func(seed Seed, index int) int {
		return seed.idChain[len(seed.idChain)-1]
	}))
	fmt.Printf("Min location: %d\n", minLocation)
}

func collectSeedProperties(seedIds []int, mappings []Mapping) []Seed {
	var seeds []Seed
	for _, seedId := range seedIds {
		seed := Seed{seedId: seedId, idChain: []int{seedId}}
		for _, mapping := range mappings {
			nextId := getMappedId(seed.idChain[len(seed.idChain)-1], mapping)
			seed.idChain = append(seed.idChain, nextId)
		}
		seeds = append(seeds, seed)
	}
	return seeds
}

func getMappedId(id int, mapping Mapping) int {
	for _, mappingRange := range mapping.ranges {
		if id >= mappingRange.sourceStart && id < mappingRange.sourceStart+mappingRange.length {
			return mappingRange.destinationStart + id - mappingRange.sourceStart
		}
	}
	return id
}

func parseInput(path string) ([]int, []Mapping) {
	var seedIds []int
	var mappings []Mapping
	file, _ := os.Open(path)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.Contains(line, "seeds:") {
			seedIds = parseSeedIds(line)
		}
		if strings.Contains(line, "-to-") {
			mappings = append(mappings, parseMapping(scanner))
		}
	}

	return seedIds, mappings
}

func parseMapping(scanner *bufio.Scanner) Mapping {
	mapping := Mapping{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		numbers := stringToIntArray(line)
		mappingRange := Range{destinationStart: numbers[0], sourceStart: numbers[1], length: numbers[2]}
		mapping.ranges = append(mapping.ranges, mappingRange)
	}
	return mapping
}

func parseSeedIds(line string) []int {
	seedsHeaderAndValues := strings.Split(line, ": ")
	return stringToIntArray(seedsHeaderAndValues[1])
}

func stringToIntArray(input string) []int {
	numbersAsString := regexp.MustCompile(`\s+`).Split(input, -1)
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
