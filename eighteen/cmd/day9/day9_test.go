package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNthCWIndex(t *testing.T) {
	ast := assert.New(t)

	var cwTests = []struct {
		current    int
		offset     int
		circlesize int
		clockwise  bool
		expected   int
	}{
		{4, 3, 12, true, 7},
		{4, 4, 12, true, 8},
		{4, 5, 12, true, 9},
		{4, 6, 12, true, 10},
		{4, 7, 12, true, 11},
		{4, 8, 12, true, 0},
		{4, 9, 12, true, 1},
		{4, 10, 12, true, 2},
		{4, 11, 12, true, 3},

		{3, 1, 4, true, 0},
		{3, 2, 4, true, 1},
		{3, 3, 4, true, 2},
		{3, 4, 4, true, 3},
		{3, 5, 4, true, 0},

		{4, 4, 12, true, 8},
		{4, 5, 12, true, 9},

		{4, 3, 12, false, 1},
		{4, 4, 12, false, 0},

		{4, 5, 12, false, 11},
		{4, 6, 12, false, 10},
		{4, 7, 12, false, 9},
		{4, 8, 12, false, 8},
		{4, 9, 12, false, 7},
		{4, 10, 12, false, 6},
		{4, 11, 12, false, 5},

		{3, 1, 4, false, 2},
		{3, 2, 4, false, 1},
		{3, 3, 4, false, 0},
		{3, 4, 4, false, 3},

		{4, 3, 12, false, 1},
		{4, 4, 12, false, 0},

		{4, 5, 12, false, 11},

		// loop round
		{4, 36, 12, false, 4},
	}

	for _, cw := range cwTests {
		found := findMarbleByOffset(cw.current, cw.offset, cw.circlesize, cw.clockwise)
		ast.Equal(cw.expected, found, "The two offsets should be the same.", cw.current, cw.offset, cw.circlesize, cw.clockwise)
	}
}
