package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var pattern = regexp.MustCompile(`mul\((\d+),(\d+)\)|don't\(\)|do\(\)`)

func main() {
	totalCollect, totalSkip, err := multiple_sum("input.txt")
	if err != nil {
		log.Fatal("Error:", err)
		return
	}

	fmt.Println("Part 1:", totalCollect+totalSkip)
	fmt.Println("Part 2:", totalCollect)
}

func multiple_sum(filename string) (int, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	var totalCollect, totalSkip int
	collecting := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		processLine(scanner.Text(), &totalCollect, &totalSkip, &collecting)
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, err
	}
	return totalCollect, totalSkip, nil
}

func processLine(line string, totalCollect, totalSkip *int, collecting *bool) {
	for _, match := range pattern.FindAllStringSubmatch(line, -1) {
		switch match[0] {
		case "don't()":
			*collecting = false
		case "do()":
			*collecting = true
		default:
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])

			if *collecting {
				*totalCollect += num1 * num2
			} else {
				*totalSkip += num1 * num2
			}
		}
	}
}
