package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileArg := flag.String("file", "input.txt", "file.txt")
	flag.Parse()
	// check if the file exists
	if _, err := os.Stat(*fileArg); err == os.ErrNotExist {
		panic(fmt.Sprintf("file %s does not exists", *fileArg))
	}

	file, err := os.Open(*fileArg)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	// process each line
	for scanner.Scan() {
		line := scanner.Text()
		gameId, highestRedCubes, highestGreenCubes, highestBlueCubes := processLine(line)
		if highestRedCubes <= 12 && highestGreenCubes <= 13 && highestBlueCubes <= 14 {
			sum += gameId
		}
	}

	fmt.Printf("The sum of all possible is %d\n", sum)
}

func processLine(input string) (int, int, int, int) {
	parts := strings.Split(strings.ToLower(input), ":")
	gamePart := parts[0]
	sets := strings.Split(strings.ReplaceAll(parts[1], " ", ""), ";")
	highestRedCubes := 0
	highestGreenCubes := 0
	highestBlueCubes := 0
	for _, set := range sets {
		cubes := strings.Split(set, ",")
		for _, cube := range cubes {
			if strings.HasSuffix(cube, "red") {
				num, err := strconv.Atoi(strings.ReplaceAll(cube, "red", ""))
				if err != nil {
					panic(err)
				}

				if num > highestRedCubes {
					highestRedCubes = num
				}
			} else if strings.HasSuffix(cube, "green") {
				num, err := strconv.Atoi(strings.ReplaceAll(cube, "green", ""))
				if err != nil {
					panic(err)
				}

				if num > highestGreenCubes {
					highestGreenCubes = num
				}
			} else if strings.HasSuffix(cube, "blue") {
				num, err := strconv.Atoi(strings.ReplaceAll(cube, "blue", ""))
				if err != nil {
					panic(err)
				}

				if num > highestBlueCubes {
					highestBlueCubes = num
				}
			}
		}
	}

	gameId, err := strconv.Atoi(strings.ReplaceAll(gamePart, "game ", ""))
	if err != nil {
		panic(err)
	}
	return gameId, highestRedCubes, highestGreenCubes, highestBlueCubes
}
