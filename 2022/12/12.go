package main

import (
	"fmt"
	"image"
	"os"
	"strings"

	"github.com/charlie2clarke/advent-of-code/datastructures"
)

func main() {
	file, _ := os.ReadFile("input.txt")
	grid := make(map[image.Point]rune)
	deltas := map[string]image.Point{
		"left":  {-1, 0},
		"right": {1, 0},
		"up":    {0, -1},
		"down":  {0, 1},
	}

	var start, end image.Point
	for y, line := range strings.Fields(string(file)) {
		for x, char := range line {
			if char == 'S' {
				start = image.Point{x, y}
			} else if char == 'E' {
				end = image.Point{x, y}
			}

			grid[image.Point{x, y}] = char
		}
	}
	grid[start] = 'a'
	grid[end] = 'z'

	visited := make(map[image.Point]bool)
	queue := datastructures.Queue[image.Point]{end}
	stack := datastructures.NewStack[image.Point]()
	stack.Push(end)
	distance := map[image.Point]int{end: 0}

	for queue.Size() > 0 {
		current := queue.Dequeue()
		visited[current] = true

		for _, delta := range deltas {
			next := current.Add(delta)
			_, inGrid := grid[next]

			if !visited[next] && inGrid && grid[current] <= grid[next]+1 {
				distance[next] = distance[current] + 1
				queue.Enqueue(next)
			}
		}
	}

	fmt.Printf("Part 1: %v", grid[start])
}
