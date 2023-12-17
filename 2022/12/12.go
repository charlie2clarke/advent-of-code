package main

import (
	"fmt"
	"image"
	"os"
	"strings"

	"github.com/charlie2clarke/advent-of-code/datastructures"
)

func loadGrid(file string) (map[image.Point]rune, image.Point, image.Point) {
	grid, start, end := make(map[image.Point]rune), image.Point{}, image.Point{}

	for x, line := range strings.Fields(string(file)) {
		for y, height := range line {
			grid[image.Point{x, y}] = height

			if height == 'S' {
				start = image.Point{x, y}
			} else if height == 'E' {
				end = image.Point{x, y}
			}
		}
	}
	grid[start], grid[end] = 'a', 'z'

	return grid, start, end
}

func main() {
	file, _ := os.ReadFile("input.txt")
	grid, start, end := loadGrid(string(file))

	queue, distance := datastructures.Queue[image.Point]{end}, map[image.Point]int{end: 0}
	var shortest *image.Point

	for !queue.IsEmpty() {
		current := queue.Dequeue()

		if grid[current] == 'a' && shortest == nil {
			shortest = &current
		}

		for _, delta := range []image.Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			next := current.Add(delta)
			_, isVisited := distance[next]
			_, inGrid := grid[next]
			isValidHeight := grid[current] <= grid[next]+1

			if inGrid && !isVisited && isValidHeight {
				distance[next] = distance[current] + 1
				queue.Enqueue(next)
			}
		}
	}

	fmt.Printf("Part 1: %d\n", distance[start])
	fmt.Printf("Part 2: %d\n", distance[*shortest])
}
