package main

import (
	"fmt"
	"image"
	"os"
	"strings"

	"github.com/charlie2clarke/advent-of-code/datastructures"
)

var deltas = map[string]image.Point{
	"L":  {-1, 0},
	"R":  {1, 0},
	"U":  {0, 1},
	"D":  {0, -1},
	"UL": {-1, 1},
	"DL": {-1, -1},
	"UR": {1, 1},
	"DR": {1, -1},
}

func main() {
	file, _ := os.ReadFile("input.txt")
	input := strings.Split(strings.TrimSpace(string(file)), "\n")

	rope := [10]image.Point{}

	gridVisitedPt1, gridVisitedPt2 := datastructures.NewSet[image.Point](), datastructures.NewSet[image.Point]()
	gridVisitedPt1.Add(rope[0])
	gridVisitedPt2.Add(rope[0])

	for _, line := range input {
		var (
			direction string
			numToMove int
		)

		fmt.Sscanf(line, "%s %d", &direction, &numToMove)

		for i := 0; i < numToMove; i++ {
			rope[0] = rope[0].Add(deltas[direction])

			for knot := 1; knot < len(rope); knot++ {
				rope[knot] = follow(rope[knot], rope[knot-1])
			}

			gridVisitedPt1.Add(rope[1])
			gridVisitedPt2.Add(rope[len(rope)-1])
		}
	}

	fmt.Printf("Part 1: %v\n", gridVisitedPt1.Len())
	fmt.Printf("Part 2: %v", gridVisitedPt2.Len())
}

func follow(tailPos, headPos image.Point) image.Point {
	if headPos == tailPos {
		return tailPos
	}

	if isOneSpaceAway(headPos, tailPos) {
		return tailPos
	}

	diff := headPos.Sub(tailPos)
	if abs(diff.X) > 0 || abs(diff.Y) > 0 {
		tailPos = tailPos.Add(image.Point{sgn(diff.X), sgn(diff.Y)})
	}

	return tailPos
}

func isOneSpaceAway(headPos, tailPos image.Point) bool {
	for _, delta := range deltas {
		if tailPos.Add(delta) == headPos {
			return true
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sgn(x int) int {
	if x > 0 {
		return 1
	} else if x == 0 {
		return 0
	}
	return -1
}
