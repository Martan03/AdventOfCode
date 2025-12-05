package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start, end int
}

func main() {
	ranges, ids, err := readInput("input.txt")
	if err != nil {
		log.Fatalln("Error:", err)
		return
	}

	fmt.Println("Part1:", part1(ranges, ids))
}

func part1(ranges []Range, ids []int) int {
	fresh := 0
	for _, id := range ids {
		valid := false
		for _, r := range ranges {
			if r.start <= id && id <= r.end {
				valid = true
				break
			}
		}
		if valid {
			fresh++
		}
	}
	return fresh
}

func readInput(filename string) ([]Range, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var ranges []Range
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		limits := strings.Split(line, "-")
		if len(limits) != 2 {
			return nil, nil, fmt.Errorf("invalid input format")
		}

		num1, err1 := strconv.Atoi(limits[0])
		num2, err2 := strconv.Atoi(limits[1])
		if err1 != nil || err2 != nil {
			return nil, nil, fmt.Errorf("not a number found")
		}
		ranges = append(ranges, Range{num1, num2})
	}

	var ids []int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, nil, fmt.Errorf("not a number found")
		}
		ids = append(ids, num)
	}

	return ranges, ids, nil
}
