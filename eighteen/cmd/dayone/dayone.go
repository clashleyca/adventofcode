package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	nums, _ := ioutil.ReadFile("/home/ssgroot/projects/go/src/adventofcode/eighteen/cmd/dayone/input.txt")
	strnums := string(nums)
	strnarr := strings.Split(strnums, "\n")
	var allnums []int64
	var allsum int64
	for _, snum := range strnarr {
		inum, err := strconv.ParseInt(snum, 10, 0)
		if err == nil {
			allnums = append(allnums, inum)
			allsum = allsum + inum
		}
	}
	//fmt.Println(allnums)
	fmt.Println("Part 1 - sum:", allsum)
	doubleFreq := partTwo(allnums)
	fmt.Println("Part 1 - freq:", doubleFreq)
	//fileLogger := logrus.New()

}

func testPartTwo() {
	testArr := []int64{1, -2, 3, 1}
	partTwo(testArr)
}

func findInArray(value int64, values []int64) (found bool) {
	for _, v := range values {
		if v == value {
			found = true
			return
		}
	}
	return
}

func partTwo(allnums []int64) (allsum int64) {

	var allsums []int64
	for {
		for _, inum := range allnums {
			allsum = allsum + inum
			foundSum := findInArray(allsum, allsums)
			if foundSum {
				return
			}
			allsums = append(allsums, allsum)
		}
	}
	return
}
