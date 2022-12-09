package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input.txt")
	input := strings.Split(strings.TrimSpace(string(file)), "\n")

	part1 := part1(input)

	fmt.Printf("Part 1: %v\n", part1)
}

func part1(input []string) int {
	headPos, tailPos := image.Point{0, 0}, image.Point{0, 0}
	gridVisited := map[image.Point]bool{
		tailPos: true,
	}

	var (
		LEFT  = image.Point{-1, 0}
		RIGHT = image.Point{1, 0}
		DOWN  = image.Point{0, -1}
		UP    = image.Point{0, 1}
	)

	doMove := func(tailPos, delta image.Point, numToMove int) image.Point {
		for newPos, i := tailPos.Add(delta), 1; i < numToMove; i++ {
			newPos = newPos.Add(delta)
			if i+1 == numToMove {
				return newPos
			}
		}
		return image.Point{}
	}

	isAdjacent := func(tailPos, headPos image.Point) bool {
		for _, direction := range []image.Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			if tailPos.Add(direction) == headPos {
				return true
			}
		}
		return false
	}

	for _, line := range input {
		var direction string
		var num int

		parses, err := fmt.Sscanf(line, "%s %d", &direction, &num)
		if parses != 2 || err != nil {
			panic("err parsing")
		}

		switch direction {
		case "R":
			tailPos = doMove(tailPos, RIGHT, num)
		case "L":
			tailPos = doMove(tailPos, LEFT, num)
		case "U":
			tailPos = doMove(tailPos, UP, num)
		case "D":
			tailPos = doMove(tailPos, DOWN, num)
		}
		gridVisited[tailPos] = true
		fmt.Println("oi")
	}

	return len(gridVisited)
}
