package main

import (
	"eighteen/internal/parser"
	"fmt"
	"os"
	"strings"
)

func main() {

	marbleString := parser.ParseStrings("day9test.txt")
	expectedCircles := stringsToCircle(marbleString)
	// TODO fix nplayers
	//part1(9, 25, expectedCircles)
	//part1(9, 25* 10, expectedCircles)
	//part1(9, 25* 100, expectedCircles)
	//part1(9, 25* 1000, expectedCircles)
	//part1(9, 25, expectedCircles)
	if part2(9, 25, expectedCircles) != 32 {
		fmt.Println("failed to match", 32)
	}
	if part2(10, 1618, expectedCircles) != 8317 {
		fmt.Println("failed to match", 8317)
	}

	if part2(13, 7999, expectedCircles) != 146373 {
		fmt.Println("failed to match", 146373)
	}
	if part2(17, 1104, expectedCircles) != 2764 {
		fmt.Println("failed to match", 2764)
	}
	if part2(21, 6111, expectedCircles) != 54718 {
		fmt.Println("failed to match", 54718)
	}
	if part2(30, 5807, expectedCircles) != 37305 {
		fmt.Println("failed to match", 37305)
	}
	//if part1(412, 71646, expectedCircles) != 439635 {
	//	fmt.Println("failed to match",439635)
	//}
	//part1(412, 71646* 100 , expectedCircles)
	println(part2(412, 71646*100, expectedCircles))
	fmt.Println("done")
}

type Marble struct {
	ID      int
	placed  bool
	current bool
	removed bool
}

type Circle struct {
	myMarbles       []Marble
	currentMarbleID int
}

func stringsToCircle(marbleString []string) (allcircles []Circle) {

	for _, ms := range marbleString {
		allMarbles := strings.Split(ms, " ")
		// [6]  0 16  8 17  4 18 19  2 (24)20 10 21  5 22 11  1 12  6 13  3 14  7 15
		var nextCircle Circle
		var index int
		_, _ = fmt.Sscanf(allMarbles[0], "[%d]", &index)
		for i, m := range allMarbles {
			var nextM int
			if i == 0 {
				continue
			}
			foundCurrent := false
			if string(m[0]) == "(" {
				_, _ = fmt.Sscanf(allMarbles[i], "(%d)", &nextM)
				foundCurrent = true
			} else {
				_, _ = fmt.Sscanf(allMarbles[i], "%d", &nextM)
			}
			marb := Marble{ID: nextM}
			nextCircle.myMarbles = append(nextCircle.myMarbles, marb)
			if foundCurrent {
				nextCircle.currentMarbleID = marb.ID
			}
		}
		allcircles = append(allcircles, nextCircle)
	}

	return
}

func setupMarbles(nm int) (marbles []Marble) {

	for i := 0; i <= nm; i++ {
		marble := Marble{ID: i}
		marbles = append(marbles, marble)
	}

	return

}

func setupPlayers(nm int) (players []Player) {

	for i := 0; i < nm; i++ {
		player := Player{ID: i}
		players = append(players, player)
	}

	return

}

func findLowestMarbleIndex2(marbles []Marble) (index int) {

	for i, m := range marbles {
		if m.removed {
			continue
		}
		if !m.placed {
			index = i
			return
		}
	}
	// not found
	return -1
}
func findLowestMarbleIndex(marbles []Marble) (index int) {

	for i, m := range marbles {
		if !m.placed {
			index = i
			return
		}
	}
	// not found
	return -1
}

func getCurrentOffset(circle Circle) (offset int) {

	currentMarbleID := circle.currentMarbleID
	for i, m := range circle.myMarbles {
		if m.ID == currentMarbleID {
			return i
		}
	}
	return 0
}

func findMarbleByOffset(currentOffset int, toRotate int, circleSize int, clockwise bool) int {

	if clockwise {
		totalOff := currentOffset + toRotate
		newIndex := totalOff % circleSize
		return newIndex
	}
	realRotate := toRotate % circleSize
	newIndex := currentOffset - realRotate
	if newIndex < 0 {
		newIndex = circleSize + newIndex
	}

	return newIndex
}

func findCircleAndMarbleByOffset2(circle Circle, currentOffset int, toRotate int, circleSize int, clockwise bool) int {

	//seventhCounterClockwiseIndex := findMarbleByOffset2(getCurrentOffset(circle), 7, len(circle.myMarbles), false)
	//removedMarble := circle.myMarbles[seventhCounterClockwiseIndex].ID
	if clockwise {
		totalOff := currentOffset + toRotate
		newIndex := totalOff % circleSize
		for i := 0; i < toRotate; i++ {
			if circle.myMarbles[currentOffset+i].removed == true {
				newIndex++
			}
		}
		return newIndex % circleSize
	}
	realRotate := toRotate % circleSize
	newIndex := currentOffset - realRotate
	if newIndex < 0 {
		newIndex = circleSize + newIndex
	}
	for i := 0; i < toRotate; i++ {
		newNewOffset := currentOffset - i
		if newNewOffset < 0 {
			newNewOffset = circleSize + newNewOffset
		}
		if circle.myMarbles[newNewOffset].removed == true {
			newIndex++
		}
	}

	return newIndex
}

func InsertAtIndex(marbles *[]Marble, newMarble Marble, index int) {

	if index < 0 || index > len(*marbles) {
		return
	}
	// increase size of slice
	*marbles = append(*marbles, Marble{})
	// copy [ID:] to [indxe + 1 ]
	copy((*marbles)[index+1:], (*marbles)[index:])
	newMarble.placed = true
	(*marbles)[index] = newMarble
}

func placeFirstMarble(circle *Circle, marbles *[]Marble, index int) {
	(*marbles)[index].placed = true
	circle.myMarbles = append(circle.myMarbles, (*marbles)[index])
	circle.currentMarbleID = (*marbles)[index].ID

}

func addMarbleAfter(circle *Circle, lowMarble Marble, thisMarble Marble) {

	for i, c := range circle.myMarbles {
		if c.ID == thisMarble.ID {
			InsertAtIndex(&circle.myMarbles, lowMarble, i+1)
			circle.currentMarbleID = lowMarble.ID
			return
		}
	}
}

func addMarbleAfter2(circle *Circle, lowMarble Marble, thisMarble Marble) {

	for i, c := range circle.myMarbles {
		if c.ID == thisMarble.ID {
			InsertAtIndex(&circle.myMarbles, lowMarble, i+1)
			circle.currentMarbleID = lowMarble.ID
			return
		}
	}
}

func printMarble(circle Circle, marble Marble) {
	if marble.removed {
		return
	}
	isCurrent := circle.currentMarbleID == marble.ID
	if isCurrent {
		fmt.Printf("(")
	} else {
		fmt.Printf(" ")
	}
	fmt.Printf("%d", marble.ID)
	if isCurrent {
		fmt.Printf(")")
	} else {
		fmt.Printf(" ")
	}
	fmt.Printf(" ")
}

func printCircle(circle Circle, index int) {

	fmt.Printf("[%3d] ", index+1)

	for _, m := range circle.myMarbles {

		printMarble(circle, m)

	}
	fmt.Println()

}

func compareCircles(expected Circle, found Circle) bool {
	if expected.currentMarbleID != found.currentMarbleID {
		return false
	}
	if len(expected.myMarbles) != len(found.myMarbles) {
		return false
	}
	foundI := 0
	for i, _ := range expected.myMarbles {
		if found.myMarbles[i].removed {
			foundI++
		}
		if expected.myMarbles[i].ID != found.myMarbles[foundI].ID {
			return false
		}
		foundI++
	}
	return true
}

type Player struct {
	ID    int
	score int
}

func removeMarbleAtIndex(circle *Circle, index int) {

	// a = append(a[:i], a[i+1:]...)
	//// or
	//a = a[:i+copy(a[i:], a[i+1:])]

	circle.myMarbles = append(circle.myMarbles[:index], circle.myMarbles[index+1:]...)
	return
}

func part2(nPlayers int, nMarbles int, expectedCircles []Circle) (max int) {

	var circle Circle
	marbles := setupMarbles(nMarbles)
	players := setupPlayers(nPlayers)
	placeFirstMarble(&circle, &marbles, 0)
	marbles[0].placed = true
	printCircle(circle, -1)

	//var happyPlayer Player
	playerID := 0
	for i := 0; i < nMarbles; i++ {
		currentId := circle.currentMarbleID
		lowMarbleIndex := findLowestMarbleIndex2(marbles)
		if currentId > 0 && (currentId+1)%23 == 0 {
			players[playerID].score = players[playerID].score + currentId + 1
			seventhCounterClockwiseIndex := findCircleAndMarbleByOffset2(circle, getCurrentOffset(circle), 7, len(circle.myMarbles), false)
			players[playerID].score = players[playerID].score + circle.myMarbles[seventhCounterClockwiseIndex].ID
			circle.myMarbles[seventhCounterClockwiseIndex].removed = true

			nextIndex := seventhCounterClockwiseIndex + 1
			for circle.myMarbles[nextIndex].removed == true {
				nextIndex++
			}
			circle.currentMarbleID = circle.myMarbles[seventhCounterClockwiseIndex+1].ID
			marbles[lowMarbleIndex].placed = true
			max = max + circle.myMarbles[seventhCounterClockwiseIndex].ID

		} else {

			nextClockwiseIndex := findCircleAndMarbleByOffset2(circle, getCurrentOffset(circle), 1, len(circle.myMarbles), true)
			if nextClockwiseIndex >= len(circle.myMarbles) {
				fmt.Println("clockwise too large")
			} else if lowMarbleIndex >= len(marbles) {
				fmt.Println("low too large")
			} else {
				addMarbleAfter2(&circle, marbles[lowMarbleIndex], circle.myMarbles[nextClockwiseIndex])
				marbles[lowMarbleIndex].placed = true
			}
			max = nextClockwiseIndex + 1
		}
		if false {
			//if len(expectedCircles) <= i+1 {
			//	return string(happyPlayer.score)
			//}
			expectedCircle := expectedCircles[i+1]
			if !compareCircles(expectedCircle, circle) {
				fmt.Printf("found\t")
				printCircle(circle, i)
				fmt.Printf("\nexpect\t")
				printCircle(expectedCircle, i)
				fmt.Println("exiting")
				os.Exit(0)
			} else {
				fmt.Println("match")
				printCircle(expectedCircle, i)

			}
		}
		playerID = playerID + 1
		if playerID%nPlayers == 0 {
			playerID = 0
		}
	}

	winningElfID := 0
	maxScore := 0
	for _, p := range players {
		if p.score > maxScore {
			winningElfID = p.ID
			maxScore = p.score
		}
	}
	fmt.Println("ID", winningElfID+1, "score", maxScore)
	return maxScore

}

func part1(nPlayers int, nMarbles int, expectedCircles []Circle) (max int) {

	var circle Circle
	marbles := setupMarbles(nMarbles)
	players := setupPlayers(nPlayers)
	placeFirstMarble(&circle, &marbles, 0)
	marbles[0].placed = true
	printCircle(circle, -1)

	doTest := false
	//var happyPlayer Player
	playerID := 0
	for i := 0; i < nMarbles; i++ {
		//players[playerID]
		currentId := circle.currentMarbleID
		lowMarbleIndex := findLowestMarbleIndex(marbles)
		//fmt.Println("id mod", currentId, (currentId+1)%23)
		if currentId > 0 && (currentId+1)%23 == 0 {
			players[playerID].score = players[playerID].score + currentId + 1
			seventhCounterClockwiseIndex := findMarbleByOffset(getCurrentOffset(circle), 7, len(circle.myMarbles), false)
			removedMarble := circle.myMarbles[seventhCounterClockwiseIndex].ID
			players[playerID].score = players[playerID].score + removedMarble
			removeMarbleAtIndex(&circle, seventhCounterClockwiseIndex)
			circle.currentMarbleID = circle.myMarbles[seventhCounterClockwiseIndex].ID
			marbles[lowMarbleIndex].placed = true

		} else {

			nextClockwiseIndex := findMarbleByOffset(getCurrentOffset(circle), 1, len(circle.myMarbles), true)
			if nextClockwiseIndex >= len(circle.myMarbles) {
				fmt.Println("clockwise too large")
			} else if lowMarbleIndex >= len(marbles) {
				fmt.Println("low too large")
			} else {
				addMarbleAfter(&circle, marbles[lowMarbleIndex], circle.myMarbles[nextClockwiseIndex])
				marbles[lowMarbleIndex].placed = true
			}
		}
		if doTest {
			//if len(expectedCircles) <= i+1 {
			//	return string(happyPlayer.score)
			//}
			expectedCircle := expectedCircles[i+1]
			if !compareCircles(expectedCircle, circle) {
				fmt.Printf("found\t")
				printCircle(circle, i)
				fmt.Printf("\nexpect\t")
				printCircle(expectedCircle, i)
				fmt.Println("exiting")
				os.Exit(0)
			} else {
				fmt.Println("match")
				printCircle(expectedCircle, i)

			}
		}
		playerID = playerID + 1
		if playerID%nPlayers == 0 {
			playerID = 0
		}
	}

	winningElfID := 0
	maxScore := 0
	for _, p := range players {
		if p.score > maxScore {
			winningElfID = p.ID
			maxScore = p.score
		}
	}
	fmt.Println("ID", winningElfID+1, "score", maxScore)
	return maxScore
}
