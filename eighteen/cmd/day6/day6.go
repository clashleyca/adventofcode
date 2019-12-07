package main

import (
	"eighteen/internal/parser"
	"fmt"
	"math"
	"strconv"
)

const (
	withinDist = 10000
)

func main() {
	//strnarr := parser.ParseStrings("day6test.txt")
	strnarr := parser.ParseStrings("day6.txt")
	//strnarr := parser.ParseStrings("day6test.txt")
	//strnarr := parser.ParseStrings("day5test.txt")
	//strnarr := parser.ParseStrings("day4test.txt")
	//result := part1(strnarr)
	result := part2(strnarr)
	fmt.Println("Day 6 result", result)
}

type Location struct {
	X         int
	Y         int
	ID        string
	GridCount int
}

func manhattanDistance(x1, y1, x2, y2 int) int {

	xdis := math.Abs(float64(x1 - x2))
	ydis := math.Abs(float64(y1 - y2))
	return int(xdis + ydis)
}

func findDistanceToCords(x, y int, allLocs []Location) int {
	distanceToCords := 0
	for _, loc := range allLocs {
		manDistance := manhattanDistance(x, y, loc.X, loc.Y)
		distanceToCords += manDistance
		if distanceToCords >= withinDist {
			distanceToCords = -1
			break
		}
	}
	return distanceToCords
}

func findMinDistance(x, y int, allLocs []Location) string {
	minDistance := math.MaxInt32
	closestID := ""
	for _, loc := range allLocs {
		manDistance := manhattanDistance(x, y, loc.X, loc.Y)
		if manDistance == minDistance {
			closestID = "equal"
		}
		if manDistance < minDistance {
			minDistance = manDistance
			closestID = loc.ID
		}
	}
	return closestID
}

func idToAlpha(id int) (alpha string) {
	offsetUpper := int('A')
	//offsetLower := int('a')
	return string(id + offsetUpper)

	return
}

func part1(inputs []string) (result string) {

	var allLocs []Location
	maxX := 0
	maxY := 0
	for i, s := range inputs {
		var X, Y int
		fmt.Sscanf(s, "%d,%d", &X, &Y)
		//fmt.Println("X",X,"Y",Y)
		if X > maxX {
			maxX = X
		}
		if Y > maxY {
			maxY = Y
		}
		//iid :=  strconv.Itoa(i)
		loc := Location{X: X, Y: Y, ID: idToAlpha(i)}
		allLocs = append(allLocs, loc)
	}
	//fmt.Println(allLocs)

	//------------------

	var grid []Location
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			closestID := findMinDistance(x, y, allLocs)
			gridLoc := Location{X: x, Y: y, ID: closestID}
			grid = append(grid, gridLoc)
		}
	}

	//----

	for i, _ := range allLocs {
		for _, coord := range grid {
			if coord.ID == allLocs[i].ID {
				if coord.X == 0 || coord.Y == 0 || coord.X == maxX || coord.Y == maxY {
					allLocs[i].GridCount = -1
					break
				} else {
					allLocs[i].GridCount++
				}
			}
		}
	}

	maxArea := 0
	maxAreaId := ""
	for _, loc := range allLocs {
		if loc.GridCount > maxArea {
			maxArea = loc.GridCount
			maxAreaId = loc.ID
		}
	}
	fmt.Println(maxAreaId)
	return strconv.Itoa(maxArea)
}

type ValidArea struct {
	distanceToCords int
	numLocs         int
}

func part2(inputs []string) (result string) {

	var allLocs []Location
	maxX := 0
	maxY := 0
	for i, s := range inputs {
		var X, Y int
		fmt.Sscanf(s, "%d,%d", &X, &Y)
		//fmt.Println("X",X,"Y",Y)
		if X > maxX {
			maxX = X
		}
		if Y > maxY {
			maxY = Y
		}
		//iid :=  strconv.Itoa(i)
		loc := Location{X: X, Y: Y, ID: idToAlpha(i)}
		allLocs = append(allLocs, loc)
	}

	//---------------

	validPoints := 0
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			distanceToCoords := findDistanceToCords(x, y, allLocs)
			if distanceToCoords != -1 {
				validPoints++
			}
		}
	}

	return strconv.Itoa(validPoints)
}
