package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	. "github.com/tt0th/adventofcode2023/src/utils"
	"os"
	"sort"
)

func main() {
	//path := "src/day17/test-input-1.txt"
	path := "src/day17/input.txt"
	var field = parseInput(path)

	destination := C{I: len(field) - 1, J: len(field[0]) - 1}
	source := C{I: 0, J: 0}

	//sum := calculateLowestSum(field, source, destination)
	//fmt.Printf("sum: %d\n", sum)

	//state := DFSState{minSolution: 750, costMemory: make(map[CostMemoryKey]int)}
	//for i := 0; i < math.MaxInt; i++ {
	//	println("start depth limit ", i)
	//	dfs(field, []C{source}, 0, []Direction{}, destination, &state, i)
	//}
	state := DFSState{minSolution: 750, costMemory: make(map[CostMemoryKey]int)}
	dfs(field, []C{source}, 0, []Direction{}, destination, &state, 10000)

	fmt.Printf("sum: %d\n", state.minSolution-field[0][0])
}

type CostMemoryKey struct {
	place      C
	direction1 Direction
	direction2 Direction
	direction3 Direction
}
type DFSState struct {
	minSolution int
	costMemory  map[CostMemoryKey]int
}

func dfs(field [][]int, path []C, cost int, directions []Direction, destination C, state *DFSState, maxDepth int) {
	if maxDepth == 0 {
		return
	}
	current := path[len(path)-1]
	costSoFar := cost + field[current.I][current.J]
	if costSoFar >= state.minSolution {
		return
	}
	lastDirections := directions[max(len(directions)-3, 0):]
	memoryKey := calculateMemoryKey(current, lastDirections)
	costFromMemory, isMemorySet := state.costMemory[memoryKey]
	if isMemorySet && costFromMemory <= costSoFar {
		return
	}
	state.costMemory[memoryKey] = costSoFar

	if current == destination {
		println("solution found", costSoFar)
		if costSoFar < state.minSolution {
			state.minSolution = costSoFar
		}
		return
	}

	var allowedDirections []Direction
	if current.I < len(field)-1 {
		allowedDirections = append(allowedDirections, Down)
	}
	if current.J < len(field[0])-1 {
		allowedDirections = append(allowedDirections, Right)
	}
	if current.I > 0 {
		allowedDirections = append(allowedDirections, Up)
	}
	if current.J > 0 {
		allowedDirections = append(allowedDirections, Left)
	}
	if len(directions) > 0 {
		lastDirection := directions[len(directions)-1]
		allowedDirections = lo.Without(allowedDirections, lastDirection.Inverse())
		if len(directions) >= 3 {
			lastDirection2 := directions[len(directions)-2]
			lastDirection3 := directions[len(directions)-3]
			if lastDirection == lastDirection2 && lastDirection == lastDirection3 {
				allowedDirections = lo.Without(allowedDirections, lastDirection)
			}
		}
	}
	sort.Slice(allowedDirections, func(i, j int) bool {
		progress1 := allowedDirections[i].I + allowedDirections[i].J
		progress2 := allowedDirections[j].I + allowedDirections[j].J

		point1 := current.Add(allowedDirections[i])
		point2 := current.Add(allowedDirections[j])
		value1 := field[point1.I][point1.J]
		value2 := field[point2.I][point2.J]

		return (5*progress1 - value1) > (5*progress2 - value2)
	})

	for _, direction := range allowedDirections {
		newPosition := current.Add(direction)
		if !lo.Contains(path, newPosition) {
			dfs(field, append(path, newPosition), costSoFar, append(directions, direction), destination, state, maxDepth-1)
		}
	}
}

func calculateMemoryKey(current C, lastDirections []Direction) CostMemoryKey {
	memoryKey := CostMemoryKey{place: current}
	if len(lastDirections) >= 1 {
		memoryKey.direction1 = lastDirections[0]
	}
	if len(lastDirections) >= 2 {
		memoryKey.direction2 = lastDirections[1]
	}
	if len(lastDirections) >= 3 {
		memoryKey.direction3 = lastDirections[2]
	}
	return memoryKey
}

//// 1  function Dijkstra(Graph, source):
//// 2
//// 3      for each vertex v in Graph.Vertices:
//// 4          dist[v] ← INFINITY
//// 5          prev[v] ← UNDEFINED
//// 6          add v to Q
//// 7      dist[source] ← 0
//// 8
//// 9      while Q is not empty:
////10          u ← vertex in Q with min dist[u]
////11          remove u from Q
////12
////13          for each neighbor v of u still in Q:
////14              alt ← dist[u] + Graph.Edges(u, v)
////15              if alt < dist[v]:
////16                  dist[v] ← alt
////17                  prev[v] ← u
////18
////19      return dist[], prev[]
//func calculateLowestSum(field [][]int, source C, destination C) int {
//	undefinedPosition := C{I: -1, J: -1}
//
//	distance := make([][]int, len(field))
//	prev := make([][]C, len(field))
//	queue := make([][]bool, len(field))
//	vertexes := make([]C, len(field)*len(field[0]))
//
//	for i := 0; i < len(field); i++ {
//		distance[i] = make([]int, len(field[0]))
//		prev[i] = make([]C, len(field[0]))
//		queue[i] = make([]bool, len(field[0]))
//		for j := 0; j < len(field[0]); j++ {
//			distance[i][j] = math.MaxInt
//			prev[i][j] = undefinedPosition
//			queue[i][j] = false
//			vertexes = append(vertexes, C{I: i, J: j})
//		}
//	}
//	queue[0][0] = true
//
//	distance[source.I][source.J] = 0
//	for lo.Contains(lo.Flatten(queue), true) {
//		current := lo.MinBy(lo.Filter(vertexes, func(vertex C, _ int) bool {
//			return queue[vertex.I][vertex.J]
//		}), func(c1 C, c2 C) bool {
//			return distance[c2.I][c2.J] > distance[c1.I][c1.J]
//		})
//		print("processing ")
//		PrintObject(current)
//
//		if current == destination {
//			for i := 0; i < len(field); i++ {
//				for j := 0; j < len(field[0]); j++ {
//					if lo.Contains(getPath(source, destination, prev), C{I: i, J: j}) {
//						print("[")
//					} else {
//						print(" ")
//					}
//					print(field[i][j])
//					if lo.Contains(getPath(source, destination, prev), C{I: i, J: j}) {
//						print("]")
//					} else {
//						print(" ")
//					}
//				}
//				println()
//			}
//			totalCost := lo.SumBy(getPath(source, destination, prev), func(c C) int {
//				return field[c.I][c.J]
//			})
//			println("done, cost:", totalCost)
//			//return totalCost
//		}
//
//		queue[current.I][current.J] = false
//		PrintObject(getNeighbours(current, field))
//		for _, neighbour := range getNeighbours(current, field) {
//			// max 3 thing
//			direction := neighbour.Minus(current)
//			prevPlace1 := prev[current.I][current.J]
//			if prevPlace1 != undefinedPosition {
//				prevDirection1 := current.Minus(prevPlace1)
//				if direction == prevDirection1 {
//					prevPlace2 := prev[prevPlace1.I][prevPlace1.J]
//					if prevPlace2 != undefinedPosition {
//						prevDirection2 := prevPlace1.Minus(prevPlace2)
//						if direction == prevDirection2 {
//							prevPlace3 := prev[prevPlace2.I][prevPlace2.J]
//							if prevPlace3 != undefinedPosition {
//								prevDirection3 := prevPlace2.Minus(prevPlace3)
//								if direction == prevDirection3 {
//									continue
//								}
//							}
//						}
//					}
//				}
//			}
//
//			alternativeDistance := distance[current.I][current.J] + field[neighbour.I][neighbour.J]
//			if alternativeDistance < distance[neighbour.I][neighbour.J] { // we might need the longer path so that we are allowed to go to an other direction
//				distance[neighbour.I][neighbour.J] = alternativeDistance
//				prev[neighbour.I][neighbour.J] = current
//				queue[neighbour.I][neighbour.J] = true
//			}
//		}
//	}
//	return -1
//}
//
//func getPath(source C, destination C, prev [][]C) []C {
//	PrintObject(destination)
//	if source == destination {
//		return []C{}
//	}
//
//	return append([]C{destination}, getPath(source, prev[destination.I][destination.J], prev)...)
//}
//
//func getNeighbours[T any](c C, field [][]T) []C {
//	var neighbours []C
//	if c.I > 0 {
//		neighbours = append(neighbours, c.Add(Up))
//	}
//	if c.I < len(field)-1 {
//		neighbours = append(neighbours, c.Add(Down))
//	}
//	if c.J > 0 {
//		neighbours = append(neighbours, c.Add(Left))
//	}
//	if c.J < len(field[0])-1 {
//		neighbours = append(neighbours, c.Add(Right))
//	}
//	return neighbours
//}

func parseInput(path string) [][]int {
	file, _ := os.Open(path)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var inputs [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputs = append(inputs, lo.Map([]rune(line), func(item rune, _ int) int {
			return UnsafeParseInt(string(item))
		}))
	}

	return inputs
}
