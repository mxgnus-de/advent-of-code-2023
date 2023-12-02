package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fileArg := flag.String("file", "input.txt", "file.txt")
	flag.Parse()
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
	for scanner.Scan() {
		line := scanner.Text()
		nums := make([]int, 0)
		for _, r := range line {
			if num, err := strconv.Atoi(string(r)); err == nil {
				nums = append(nums, num)
			}
		}
		if len(nums) == 0 {
			continue
		}

		firstNum := nums[0]
		lastNum := nums[len(nums)-1]
		numStr := fmt.Sprintf("%d%d", firstNum, lastNum)
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}

		sum += num
	}

	fmt.Printf("The sum of all calibration values is %d\n", sum)
}
