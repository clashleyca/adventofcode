package main

import (
	"eighteen/internal/parser"
	"fmt"
)

type fabricPixel struct {
	across, down int
}

func main() {
	strnarr := parser.ParseStrings("day3.txt")
	fmt.Println(testPartTwp(strnarr))
}

// Contains tells whether a contains x.
func contains(a *[]square, row int, col int) {
	for i, _ := range *a {
		if (*a)[i].Row == row && (*a)[i].Col == col {
			(*a)[i].FoundOverlap = true
			return
		}
	}
	newSquare := square{Row: row, Col: col}
	*a = append(*a, newSquare)
}

func addSquare(grid *[]square, across int, along int) {
}

func markSquare(leftEdge, topEdge, width, height int, grid *[]square) {
	//fabricMap = make(map[string]int)
	for across := leftEdge; across < leftEdge+width; across++ {
		for along := topEdge; along < topEdge+height; along++ {
			contains(grid, across, along)
		}
	}

}

type square struct {
	Row          int
	Col          int
	FoundOverlap bool
}

func partOne(strnarr []string) int {

	var grid []square
	for _, claim := range strnarr {
		var claimNum, leftEdge, topEdge, width, height int
		fmt.Sscanf(claim, "#%d @ %d,%d:%dx%d", &claimNum, &leftEdge, &topEdge, &width, &height)
		//fmt.Println(claimNum, leftEdge, topEdge, width, height)
		markSquare(leftEdge, topEdge, width, height, &grid)
	}
	var overlapCount int

	for _, s := range grid {
		if s.FoundOverlap == true {
			overlapCount++
		}
	}
	return overlapCount
}

func testPartOne(strnarr []string) int {
	//var testStrings = []string{
	//	"#1 @ 1,3: 4x4",
	//	"#2 @ 3,1: 4x4",
	//	"#3 @ 5,5: 2x2",
	//}
	return partOne(strnarr)
}
func testPartTwp(strnarr []string) int {
	//var testStrings = []string{
	//	"#1 @ 1,3: 4x4",
	//	"#2 @ 3,1: 4x4",
	//	"#3 @ 5,5: 2x2",
	//}
	return newPartTwo(strnarr)
}

//func partTwo(strnarr [] string ) int {
//
//	var grid []square
//	for _,claim := range strnarr{
//		var claimNum, leftEdge, topEdge, width, height int
//		fmt.Sscanf(claim, "#%d @ %d,%d:%dx%d", &claimNum, &leftEdge, &topEdge, &width, &height)
//		//fmt.Println(claimNum, leftEdge, topEdge, width, height)
//		findNoOverlap(claimNum, leftEdge, topEdge, width, height,&grid)
//	}
//	var overlapCount int
//
//	for _,s := range grid {
//		if s.FoundOverlap == true {
//			overlapCount ++
//		}
//	}
//	return overlapCount
//}

type partTwoSquare struct {
	Row          int
	Col          int
	FoundOverlap bool
	IDs          []int
}

//func findNoOverlap(claimNum int, leftEdge, topEdge, width, height int, grid *[]square ){
//	//fabricMap = make(map[string]int)
//	for across := leftEdge ; across < leftEdge + width ; across ++ {
//		for along := topEdge; along < topEdge + height ; along ++ {
//			(grid, claimNum, across, along)
//		}
//	}
//
//}

// Contains tells whether a contains x.
func findInSquares(ptSquares []partTwoSquare, row int, col int) (index int) {
	for i, _ := range ptSquares {
		if (ptSquares)[i].Row == row && (ptSquares)[i].Col == col {
			return i
		}
	}
	return -1
}

// if a square with this <row,col> is already in our list, append to IDs
// otherwise, append square to ptSquareqs
func trackASquare(ptSquares *[]partTwoSquare, claim int, row int, col int) {
	// return index of square in ptSquares that coorosponds with <row,col>
	index := findInSquares(*ptSquares, row, col)
	var ptSquare partTwoSquare
	if index < 0 {
		ptSquare = partTwoSquare{Row: row, Col: col}
		*ptSquares = append(*ptSquares, ptSquare)
		index = len(*ptSquares) - 1
	} else {
		ptSquare = (*ptSquares)[index]
	}
	(*ptSquares)[index].IDs = append((*ptSquares)[index].IDs, claim)
}

// populate array ptSquares.  ptSaures.IDs should list each claim#s that sit on <row,col>
func buildTheGrid(ptSquares *[]partTwoSquare, claimNum int, leftEdge, topEdge, width, height int) {
	for row := leftEdge; row < leftEdge+width; row++ {
		for col := topEdge; col < topEdge+height; col++ {
			trackASquare(ptSquares, claimNum, row, col)
		}
	}
}

type oneClaim struct {
	claimNum  int
	leftEdge  int
	topEdge   int
	width     int
	height    int
	mySquares []partTwoSquare
}

func newPartTwo(strnarr []string) int {

	var alltheSquares []partTwoSquare
	var allTheClaims []oneClaim
	for _, claim := range strnarr {
		var claimNum, leftEdge, topEdge, width, height int
		fmt.Sscanf(claim, "#%d @ %d,%d:%dx%d", &claimNum, &leftEdge, &topEdge, &width, &height)
		newClaim := oneClaim{claimNum: claimNum, leftEdge: leftEdge, topEdge: topEdge, width: width, height: height}
		allTheClaims = append(allTheClaims, newClaim)
		//fmt.Println(claimNum, leftEdge, topEdge, width, height)
		// create array of partTwoSquares.  Each square contains {row},{col}, and list of claim#s that sit on <row,col>
		buildTheGrid(&alltheSquares, claimNum, leftEdge, topEdge, width, height)
	}

	var onlyOneClaimSquares []partTwoSquare
	var aHappyClaim []oneClaim

	for _, square := range alltheSquares {
		if len(square.IDs) == 1 {
			onlyOneClaimSquares = append(onlyOneClaimSquares, square)
		}
	}
	for _, claim := range allTheClaims {
		for _, square := range onlyOneClaimSquares {
			if claim.claimNum == square.IDs[0] {
				claim.mySquares = append(claim.mySquares, square)
				if len(claim.mySquares) == claim.height*claim.width {
					aHappyClaim = append(aHappyClaim, claim)
				}
			}

		}
		// we already know theres only one ID
		//if square.IDs[0] != claim.claimNum {
		//	continue
		//}
		//if foundClaim(square, claim) {
		//	return 1
		//}
	}
	return aHappyClaim[0].claimNum
}

//func foundClaim(square partTwoSquare, claim oneClaim) bool {
//	if square.Row != claim.leftEdge {
//		return false
//	}
//	if square.Col != claim.topEdge {
//		return false
//	}
//
//
//	return false
//}
