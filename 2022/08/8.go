package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input.txt")
	strInput := strings.Split(string(file), "\n")
	var input [][]int
	for _, line := range strInput {
		var row []int
		for _, char := range line {
			row = append(row, int(char-'0'))
		}
		input = append(input, row)
	}

	visibleCount := part1(input)
	scenicScore := part2(input)

	fmt.Printf("Part 1: %d\n", visibleCount)
	fmt.Printf("Part 2: %d\n", scenicScore)
}

func part1(input [][]int) int {
	edgeCount := func(input [][]int) int {
		edgeCount := 0
		for i := 0; i < len(input); i++ {
			for j := 0; j < len(input[i]); j++ {
				if i == 0 || i == len(input)-1 || j == 0 || j == len(input[i])-1 {
					edgeCount++
				}
			}
		}
		return edgeCount
	}

	visibleCount := edgeCount(input)

	isVisibleFromLeft := func(currentTree int, input [][]int, i, j int) bool {
		for k := j - 1; k >= 0; k-- {
			if input[i][k] >= currentTree {
				return false
			}
		}
		return true
	}

	isVisibleFromRight := func(currentTree int, input [][]int, i, j int) bool {
		for k := j + 1; k < len(input[i]); k++ {
			if input[i][k] >= currentTree {
				return false
			}
		}
		return true
	}

	isVisibleFromTop := func(currentTree int, input [][]int, i, j int) bool {
		for k := i - 1; k >= 0; k-- {
			if input[k][j] >= currentTree {
				return false
			}
		}
		return true
	}

	isVisibleFromBottom := func(currentTree int, input [][]int, i, j int) bool {
		for k := i + 1; k < len(input); k++ {
			if input[k][j] >= currentTree {
				return false
			}
		}
		return true
	}

	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[i])-1; j++ {
			if isVisibleFromLeft(input[i][j], input, i, j) || isVisibleFromRight(input[i][j], input, i, j) || isVisibleFromTop(input[i][j], input, i, j) || isVisibleFromBottom(input[i][j], input, i, j) {
				visibleCount++
			}
		}
	}

	return visibleCount
}

func part2(input [][]int) int {
	highestScenicScore := 0

	calcScenicScore := func(leftDistance int, rightDistance int, topDistance int, bottomDistance int) int {
		return leftDistance * rightDistance * topDistance * bottomDistance
	}

	leftDistance := func(currentTree int, input [][]int, i, j int) int {
		distance := 0
		for k := j - 1; k >= 0; k-- {
			distance++
			if input[i][k] >= currentTree {
				return distance
			}
		}
		return distance
	}

	rightDistance := func(currentTree int, input [][]int, i, j int) int {
		distance := 0
		for k := j + 1; k < len(input[i]); k++ {
			distance++
			if input[i][k] >= currentTree {
				return distance
			}
		}
		return distance
	}

	topDistance := func(currentTree int, input [][]int, i, j int) int {
		distance := 0
		for k := i - 1; k >= 0; k-- {
			distance++
			if input[k][j] >= currentTree {
				return distance
			}
		}
		return distance
	}

	bottomDistance := func(currentTree int, input [][]int, i, j int) int {
		distance := 0
		for k := i + 1; k < len(input); k++ {
			distance++
			if input[k][j] >= currentTree {
				return distance
			}
		}
		return distance
	}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			leftDistance, rightDistance, topDistance, bottomDistance := leftDistance(input[i][j], input, i, j), rightDistance(input[i][j], input, i, j), topDistance(input[i][j], input, i, j), bottomDistance(input[i][j], input, i, j)

			scenicScore := calcScenicScore(leftDistance, rightDistance, topDistance, bottomDistance)
			if scenicScore > highestScenicScore {
				highestScenicScore = scenicScore
			}
		}
	}

	return highestScenicScore
}
