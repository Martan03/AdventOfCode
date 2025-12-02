package main

import (
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
	ranges, err := readInput("input.txt")
	if err != nil {
		log.Fatalln("Error:", err)
		return
	}

	part1, part2 := sillyCnt(ranges)
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}

func sillyCnt(ranges []Range) (int, int) {
	silly1 := 0
	silly2 := 0
	for _, r := range ranges {
		for n := r.start; n <= r.end; n++ {
			if isSilly(n, 2) {
				silly1 += n
				continue
			}

			d := 3
			for i := 100; i <= n; i *= 10 {
				if isSilly(n, d) {
					silly2 += n
					break
				}
				d++
			}
		}
	}
	return silly1, silly1 + silly2
}

func isSilly(num int, div int) bool {
	tnum := strconv.Itoa(num)
	nlen := len(tnum)
	if nlen%div != 0 {
		return false
	}

	step := nlen / div
	first := tnum[0:step]

	offset := step
	for i := 1; i < div; i++ {
		if first != tnum[offset:offset+step] {
			return false
		}
		offset += step
	}
	return true
}

func readInput(filename string) ([]Range, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var ranges []Range

	data := strings.TrimSpace(string(content))
	for r := range strings.SplitSeq(data, ",") {
		limits := strings.Split(r, "-")
		if len(limits) != 2 {
			return nil, fmt.Errorf("invalid input format")
		}

		num1, err1 := strconv.Atoi(limits[0])
		num2, err2 := strconv.Atoi(limits[1])
		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("not a number found")
		}
		ranges = append(ranges, Range{num1, num2})
	}
	return ranges, nil
}
