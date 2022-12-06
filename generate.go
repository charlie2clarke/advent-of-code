package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var day = 0

func main() {
	flag.IntVar(&day, "day", 0, "day to generate")
	flag.Parse()

	if day == 0 {
		panic("day must be set")
	}

	if err := generateLaunchJSON(day); err != nil {
		panic(err)
	}

	if err := generateFiles(day); err != nil {
		panic(err)
	}
}

var launchJSONTemplate = `
	{
		"version": "0.2.0",
		"configurations": [
			{
				"name": "Launch Package",
				"type": "go",
				"request": "launch",
				"mode": "auto",
				"program": "${workspaceFolder}/2022/%v/%v.go"
			}
		]
	}
`

func generateLaunchJSON(day int) error {
	wd, err := os.Getwd()
	if err != nil {
		return errors.New("could not get root dir")
	}

	// change to the .vscode dir
	if err := os.Chdir(fmt.Sprintf("%v/.vscode", wd)); err != nil {
		return errors.New("could not change to .vscode dir")
	}

	if err := os.Remove("launch.json"); err != nil {
		return err
	}

	// create a new launch.json file
	launchJSON, err := os.Create("launch.json")
	if err != nil {
		return err
	}

	// if the day is a single digit, add a 0 to the front of the day
	// ex: 1 -> 01, 10 -> 10
	dirName := fmt.Sprintf("%02d", day)
	launchJSON.WriteString(fmt.Sprintf(launchJSONTemplate, dirName, day))

	return nil
}

var pkgmainTemplate = "package main\n \nfunc main() {\n\t\n}"

func generateFiles(day int) error {
	wd, err := os.Getwd()
	if err != nil {
		return errors.New("could not get root dir")
	}
	// get the workspaceFolder
	wd = filepath.Dir(wd)

	if err := os.Chdir(fmt.Sprintf("%v/2022", wd)); err != nil {
		return errors.New("could not change to 2022 dir")
	}

	// if the day is a single digit, add a 0 to the front of the day
	// ex: 1 -> 01, 10 -> 10
	dirName := fmt.Sprintf("%02d", day)

	// create the day dir
	if err := os.Mkdir(dirName, os.ModePerm); err != nil {
		return err
	}

	// change to the day dir
	if err := os.Chdir(dirName); err != nil {
		return errors.New("could not change to day dir")
	}

	// create the day.go file
	dayFile, err := os.Create(fmt.Sprintf("%v.go", day))
	if err != nil {
		return err
	}

	// write a pkgmain to the day.go file
	_, err = dayFile.WriteString(pkgmainTemplate)
	if err != nil {
		return err
	}

	// create the input.txt file
	_, err = os.Create("input.txt")
	if err != nil {
		return err
	}

	return nil
}
