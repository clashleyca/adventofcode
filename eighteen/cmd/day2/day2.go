package main

import (
	"eighteen/internal/parser"
	"fmt"
	"math"
)

func main() {
	strnarr := parser.ParseStrings("day2.txt")
	fmt.Println(partTwo(strnarr))
}

func testPartOne() {
	var testStrings = []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}
	partOne(testStrings)

}

func hasRepeatedLetters(repeats int, boxID string) (has bool) {
	var letters = make(map[rune]int)
	for _, c := range boxID {
		count := letters[c]
		letters[c] = count + 1
	}
	for _, v := range letters {
		if v == repeats {
			return true
		}
	}
	return
}

func partOne(boxIDs []string) {
	var twoCount int
	var threeCount int
	for _, boxID := range boxIDs {
		if hasRepeatedLetters(2, boxID) {
			//fmt.Println(boxID,"has exactly 2 matching letters")
			twoCount = twoCount + 1
		}
		if hasRepeatedLetters(3, boxID) {
			//fmt.Println(boxID,"has exactly 3 matching letters")
			threeCount = threeCount + 1
		}
	}
	fmt.Println("Box Checksum", twoCount*threeCount)
}

//func testPartTwo() {
//	var testStrings = []string {
//		"abcdef",
//		"bababc",
//		"abbcde",
//		"abcccd",
//		"aabcdd",
//		"abcdee",
//		"ababab",
//	}
//	partOne(testStrings)
//
//}

func testPartTwo() {
	var testStrings = []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}
	fmt.Println(partTwo(testStrings))
}

func checkForMatch(currentBoxID string, thisBoxID string) (matchingRunes []rune, found bool) {

	if currentBoxID == thisBoxID {
		return
	}
	currentRunes := []rune(currentBoxID)
	theseRunes := []rune(thisBoxID)

	var foundOneOff bool
	for i := 0; i < len(currentRunes); i++ {
		runeC := currentRunes[i]
		runtT := theseRunes[i]
		diff := math.Abs(float64(runeC - runtT))
		if diff != 0 {
			if foundOneOff {
				return nil, false
			}
			foundOneOff = true
		} else if diff == 0 {
			matchingRunes = append(matchingRunes, currentRunes[i])
		}
	}
	if !foundOneOff {
		return nil, false
	}
	found = len(matchingRunes) > 0
	return matchingRunes, found
}

func findAMatch(currentBoxID string, boxIDs []string) (matchString string, found bool) {
	for _, boxID := range boxIDs {
		var fRunes []rune
		fRunes, found := checkForMatch(currentBoxID, boxID)
		if found {
			return string(fRunes), true
		}
	}
	return "", false
}

func partTwo(boxIDs []string) string {
	var foundStr string
	for _, boxID := range boxIDs {
		foundStr, found := findAMatch(boxID, boxIDs)
		if found {
			return foundStr
		}
	}
	return foundStr
}
