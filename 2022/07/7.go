package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"unicode"
)

func main() {
	file, _ := os.ReadFile("input.txt")
	input := strings.Split(string(file), "\n")

	dirAndSize := map[string]int{}
	currentPath := ""

	isCdCmd := func(line string) bool {
		return strings.HasPrefix(line, "$ cd")
	}

	isLsCmd := func(line string) bool {
		return strings.HasPrefix(line, "$ ls")
	}

	isFile := func(line string) bool {
		return unicode.IsDigit(rune(line[0]))
	}

	for _, line := range input {
		if isCdCmd(line) {
			currentPath = path.Join(currentPath, strings.TrimPrefix(line, "$ cd "))
		} else if isLsCmd(line) {
			continue
		} else if isFile(line) {
			var size, name = 0, ""
			parses, err := fmt.Sscanf(line, "%d %s", &size, &name)
			if err != nil || parses != 2 {
				panic("parse error")
			}

			for dir := currentPath; dir != "/"; dir = path.Dir(dir) {
				dirAndSize[dir] += size
			}
			dirAndSize["/"] += size
		}
	}

	pt1, pt2 := part1(dirAndSize), part2(dirAndSize)

	fmt.Printf("Part 1: %v\n", pt1)
	fmt.Printf("Part 2: %v\n", pt2)
}

func part1(dirAndSize map[string]int) int {
	var pt1 int
	for _, size := range dirAndSize {
		if size <= 10000000 {
			pt1 += size
		}
	}
	return pt1
}

func part2(dirAndSize map[string]int) int {
	maxMemAllowed := 70000000 - 30000000
	memUsed := dirAndSize["/"]
	minDirSize := memUsed

	for dir, size := range dirAndSize {
		if dir == "/" {
			continue
		}

		if memUsed-size < maxMemAllowed && size < minDirSize {
			minDirSize = size
		}
	}

	return minDirSize
}
