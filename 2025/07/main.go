package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	lines, err := readInput("input.txt")
	if err != nil {
		log.Fatalln("Error:", err)
		return
	}

	part1, part2 := taychomBeams(lines)
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}

func taychomBeams(lines []string) (int, int) {
	start := strings.IndexRune(lines[0], 'S')
	if start == -1 {
		return 0, 0
	}

	beams := make(map[int]int)
	beams[start] = 1
	
	splits := 0
	for _, line := range lines[1:] {
		nextBeams := make(map[int]int)
		for beam, cnt := range beams {
			if line[beam] != '^' {
				nextBeams[beam] += cnt
				continue
			}
			splits++
			nextBeams[beam-1] += cnt
			nextBeams[beam+1] += cnt
		}
		beams = nextBeams
	}
	total := 0
	for _, cnt := range beams {
		total += cnt
	}
	return splits, total
}

func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}
