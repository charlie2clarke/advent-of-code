package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type CMD struct {
	name   string
	target string
}

type Dir struct {
	name      string
	parent    *Dir
	children  []Dir
	files     []File
	totalSize int
}

func (d *Dir) size() int {
	if d.totalSize != 0 {
		return d.totalSize
	}

	for _, file := range d.files {
		d.totalSize += file.size
	}

	for _, child := range d.children {
		d.totalSize += child.size()
	}

	return d.totalSize
}

type File struct {
	name string
	size int
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	dirs := make([]*Dir, 0)
	var currentDir *Dir

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "$") {
			cmd := &CMD{}
			splitLine := strings.Split(line, " ")

			switch splitLine[1] {
			case "cd":
				cmd.name = splitLine[1]
				cmd.target = splitLine[2]

				if cmd.target == ".." {
					currentDir = currentDir.parent
					continue
				}

				currentDir = &Dir{name: cmd.target, parent: currentDir}
			case "ls":
				cmd.name = splitLine[1]

				contents := getLsContents(scanner)
				if contents == nil {
					panic("couldn't get ls contents")
				}

				currentDir.children = append(currentDir.children, contents.dirs...)
				currentDir.files = append(currentDir.files, contents.files...)
			}

			currentDir.totalSize = currentDir.size()
			dirs = append(dirs, currentDir)
		}
	}
}

type LsContents struct {
	dirs  []Dir
	files []File
}

func getLsContents(scanner *bufio.Scanner) *LsContents {
	contents := &LsContents{
		dirs:  make([]Dir, 0),
		files: make([]File, 0),
	}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" || strings.HasPrefix(line, "$") {
			break
		}

		splitLine := strings.Split(line, " ")
		if splitLine[0] == "dir" {
			contents.dirs = append(contents.dirs, Dir{name: splitLine[1]})
		} else {
			size, err := strconv.Atoi(splitLine[0])
			if err != nil {
				panic("couldn't convert size to an int")
			}

			contents.files = append(contents.files, File{name: splitLine[1], size: size})
		}
	}

	return contents
}
