package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var dirs = [][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
	{1, 1},
	{-1, -1},
	{1, -1},
	{-1, 1},
}

func main() {
	lines, err := readInput("input.txt")
	if err != nil {
		log.Fatal("Error:", err)
		return
	}

	total := findWordCnt(lines, "XMAS")
	fmt.Println("XMAS cnt: ", total)

	total = findXMasCnt(lines)
	fmt.Println("X-MAS cnt:", total)
}

func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return input, nil
}

func findWordCnt(lines []string, word string) int {
	total := 0
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			total += checkAllDirs(lines, word, x, y)
		}
	}
	return total
}

func checkAllDirs(lines []string, word string, x, y int) int {
	total := 0
	for _, dir := range dirs {
		if checkDir(lines, word, x, y, dir[0], dir[1]) {
			total++
		}
	}
	return total
}

func checkDir(lines []string, word string, x, y, xd, yd int) bool {
	for i := 0; i < len(word); i++ {
		if x >= len(lines[0]) || x < 0 || y >= len(lines) || y < 0 {
			return false
		}

		if word[i] != lines[y][x] {
			return false
		}
		x += xd
		y += yd
	}

	return true
}

func findXMasCnt(lines []string) int {
	total := 0
	for y := 0; y < len(lines)-2; y++ {
		for x := 0; x < len(lines[0])-2; x++ {
			if checkXMas(lines, x, y) {
				total++
			}
		}
	}
	return total
}

func checkXMas(lines []string, x, y int) bool {
	return (checkDir(lines, "MAS", x, y, 1, 1) ||
		checkDir(lines, "SAM", x, y, 1, 1)) &&
		(checkDir(lines, "MAS", x+2, y, -1, 1) ||
			checkDir(lines, "SAM", x+2, y, -1, 1))
}
