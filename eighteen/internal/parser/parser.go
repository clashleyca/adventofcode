package parser

import (
	"io/ioutil"
	"strconv"
	"strings"
)

const inputDir = "/home/ssgroot/projects/go/src/adventofcode/eighteen/input/"

func ParseStrings(filename string) (allstrs []string) {
	nums, _ := ioutil.ReadFile(inputDir + filename)
	strnums := string(nums)
	allstrs = strings.Split(strnums, "\n")
	return

}

func ParseSpaces(filename string) (allstrs []string) {
	nums, _ := ioutil.ReadFile(inputDir + filename)
	strnums := string(nums)
	strnums = strings.TrimSuffix(strnums, "\n")
	allstrs = strings.Split(strnums, " ")
	return

}

func GetString(filename string) (response string) {
	resp, _ := ioutil.ReadFile(inputDir + filename)
	response = string(resp)
	return

}

func ParseNums(filename string) (allnums []int64) {
	strnarr := ParseStrings(filename)
	for _, snum := range strnarr {
		inum, err := strconv.ParseInt(snum, 10, 0)
		if err == nil {
			allnums = append(allnums, inum)
		}
	}
	return
}

func ParseSpacesdNums(filename string) (allnums []int) {
	strnarr := ParseSpaces(filename)
	for _, snum := range strnarr {
		inum, err := strconv.ParseInt(snum, 10, 0)
		if err == nil {
			allnums = append(allnums, int(inum))
		}
	}
	return
}
