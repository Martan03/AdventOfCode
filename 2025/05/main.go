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
	fmt.Println("Part2:", part2(ranges))
}

func part1(ranges []Range, ids []int) int {
	fresh := 0
	for _, id := range ids {
		for _, r := range ranges {
			if r.start <= id && id <= r.end {
				fresh++
				break
			}
		}
	}
	return fresh
}

func part2(ranges []Range) int {
	change := true
	for change {
		ranges, change = mergeRanges(ranges)
	}

	total := 0
	for _, r := range ranges {
		total += r.end - r.start + 1
	}
	return total
}

func mergeRanges(ranges []Range) ([]Range, bool) {
	var jranges []Range
	change := false
	for _, cr := range ranges {
		found := false
		for i := 0; i < len(jranges); i++ {
			found = true
			if jranges[i].start <= cr.start && cr.end <= jranges[i].end {
				break
			}

			if (cr.start <= jranges[i].end && jranges[i].start <= cr.end) ||
				(jranges[i].start <= cr.end && cr.start <= jranges[i].start) {
				change = true
				jranges[i].start = min(jranges[i].start, cr.start)
				jranges[i].end = max(jranges[i].end, cr.end)
				break
			}

			found = false
		}

		if !found {
			jranges = append(jranges, cr)
		}
	}
	return jranges, change
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
