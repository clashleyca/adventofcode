package main

import (
	"eighteen/internal/parser"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	strnarr := parser.ParseStrings("day5.txt")
	//strnarr := parser.ParseStrings("day5test.txt")
	//strnarr := parser.ParseStrings("day4test.txt")
	result := part2(strnarr[0])
	fmt.Println("Day 5 result", result)
}

func part1() {
	strnarr := parser.ParseStrings("day5.txt")
	//strnarr := parser.ParseStrings("day5test.txt")
	//strnarr := parser.ParseStrings("day4test.txt")
	result := collapse(strnarr[0])
	fmt.Println("Day 5 result", len(result))
}

func makeCharMap(input string) (charMap map[string]int) {
	inChars := strings.Split(input, "")
	for _, r := range inChars {
		charMap[strings.ToLower(r)] = charMap[strings.ToLower(r)] + 1
	}
	var mostChars = 0
	for _, c := range charMap {
		if c > mostChars {
			mostChars = c
		}
	}
	return
}

func part2(input string) (result string) {

	letDiff := 'a' - 'A'

	minLen := len(input)
	fmt.Println("input len", minLen)

	for i := 'A'; i < 'Z'+1; i++ {
		upper := string(i)
		lower := string(i + letDiff)
		if !strings.Contains(input, upper) || !strings.Contains(input, lower) {
			continue
		}
		fmt.Println("Stripping", upper, "and", lower)
		stripped := strings.ReplaceAll(input, upper, "")
		stripped = strings.ReplaceAll(stripped, lower, "")
		//fmt.Println("stripped", stripped)
		collapsed := collapse(stripped)
		colLen := len(collapsed)
		fmt.Println("collapsed len", colLen)
		if colLen < minLen {
			minLen = colLen
		}
	}

	return strconv.Itoa(minLen)
}

func removeAPair(polymer string) (result string, finished bool) {
	p := []rune(polymer)

	result = polymer
	//fmt.Println(polymer)
	for i, current := range p {
		if i+1 == len(p) {
			// we didn't find anything
			return result, true
		}
		next := p[i+1]
		//fmt.Println(string(current), string(next))
		distance := math.Abs(float64('A' - 'a'))
		diff := current - next
		adiff := math.Abs(float64(diff))
		if adiff == distance {
			newr := p[0:i]
			nexr := p[i+2:]
			result = string(newr) + string(nexr)
			return result, false
		}
	}
	return result, true
}

func collapse(input string) (result string) {

	substr, finished := removeAPair(input)
	for finished == false {
		substr, finished = removeAPair(substr)
	}
	return substr
}
