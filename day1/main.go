package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var replacementMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

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
		processedLine := ""
		// process each character
		for _, c := range line {
			// first add the character to the processed line
			processedLine += string(c)
			for k, v := range replacementMap {
				// if the line contains a written number, replace it with the number, but keep the last character
				// because sometimes it could be twone, which should be 2, 1 and not only 2
				if strings.Contains(processedLine, k) {
					processedLine = strings.Replace(processedLine, k[:len(k)-1], v, 1)
					processedLine += string(k[len(k)-1])
				}
			}
		}

		// convert the processed line to a slice of ints
		nums := make([]int, 0)
		for _, r := range processedLine {
			if num, err := strconv.Atoi(string(r)); err == nil {
				nums = append(nums, num)
			}
		}
		if len(nums) == 0 {
			continue
		}

		// the first number is the last number of the line
		firstNum := nums[0]
		lastNum := nums[len(nums)-1]
		// combine the first and last number to get the calibration value
		numStr := fmt.Sprintf("%d%d", firstNum, lastNum)
		// convert the string to an int
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}

		sum += num
	}

	fmt.Printf("The sum of all calibration values is %d\n", sum)
}
