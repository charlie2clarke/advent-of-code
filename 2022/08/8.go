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

	visibleCount, scenicScore := part1(input), part2(input)

	fmt.Printf("Part 1: %d\n", visibleCount)
	fmt.Printf("Part 2: %d\n", scenicScore)
}

func part1(trees [][]int) int {
	edgeCount := func(trees [][]int) int {
		edgeCount := 0
		for i := 0; i < len(trees); i++ {
			for j := 0; j < len(trees[i]); j++ {
				if i == 0 || i == len(trees)-1 || j == 0 || j == len(trees[i])-1 {
					edgeCount++
				}
			}
		}
		return edgeCount
	}

	visibleCount := edgeCount(trees)

	isVisibleFromLeft := func(currentTree int, trees [][]int, i, j int) bool {
		for k := j - 1; k >= 0; k-- {
			if trees[i][k] >= currentTree {
				return false
			}
		}
		return true
	}

	isVisibleFromRight := func(currentTree int, trees [][]int, i, j int) bool {
		for k := j + 1; k < len(trees[i]); k++ {
			if trees[i][k] >= currentTree {
				return false
			}
		}
		return true
	}

	isVisibleFromTop := func(currentTree int, trees [][]int, i, j int) bool {
		for k := i - 1; k >= 0; k-- {
			if trees[k][j] >= currentTree {
				return false
			}
		}
		return true
	}

	isVisibleFromBottom := func(currentTree int, trees [][]int, i, j int) bool {
		for k := i + 1; k < len(trees); k++ {
			if trees[k][j] >= currentTree {
				return false
			}
		}
		return true
	}

	for i := 1; i < len(trees)-1; i++ {
		for j := 1; j < len(trees[i])-1; j++ {
			if isVisibleFromLeft(trees[i][j], trees, i, j) || isVisibleFromRight(trees[i][j], trees, i, j) || isVisibleFromTop(trees[i][j], trees, i, j) || isVisibleFromBottom(trees[i][j], trees, i, j) {
				visibleCount++
			}
		}
	}

	return visibleCount
}

func part2(trees [][]int) int {
	highestScenicScore := 0

	leftDistance := func(currentTree int, trees [][]int, i, j int) int {
		distance := 0
		for k := j - 1; k >= 0; k-- {
			distance++
			if trees[i][k] >= currentTree {
				return distance
			}
		}
		return distance
	}

	rightDistance := func(currentTree int, trees [][]int, i, j int) int {
		distance := 0
		for k := j + 1; k < len(trees[i]); k++ {
			distance++
			if trees[i][k] >= currentTree {
				return distance
			}
		}
		return distance
	}

	topDistance := func(currentTree int, trees [][]int, i, j int) int {
		distance := 0
		for k := i - 1; k >= 0; k-- {
			distance++
			if trees[k][j] >= currentTree {
				return distance
			}
		}
		return distance
	}

	bottomDistance := func(currentTree int, trees [][]int, i, j int) int {
		distance := 0
		for k := i + 1; k < len(trees); k++ {
			distance++
			if trees[k][j] >= currentTree {
				return distance
			}
		}
		return distance
	}

	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[i]); j++ {
			scenicScore := leftDistance(trees[i][j], trees, i, j) * rightDistance(trees[i][j], trees, i, j) * topDistance(trees[i][j], trees, i, j) * bottomDistance(trees[i][j], trees, i, j)

			if scenicScore > highestScenicScore {
				highestScenicScore = scenicScore
			}
		}
	}

	return highestScenicScore
}
