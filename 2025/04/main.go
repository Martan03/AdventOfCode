package main

import (
	"bufio"
	"fmt"
	"os"
)

type Grid struct {
	rolls         []rune
	width, height int
}

func main() {
	grid, err := readInput("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Part1:", part1(grid))
	fmt.Println("Part2:", part2(grid))
}

func part1(grid Grid) int {
	total := 0
	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			id := y*grid.width + x
			if grid.rolls[id] == '@' && rollsAdjCnt(grid, x, y) < 4 {
				total++
			}
		}
	}
	return total
}

func part2(grid Grid) int {
	total := 0
	change := true
	for change {
		change = false
		for y := 0; y < grid.height; y++ {
			for x := 0; x < grid.width; x++ {
				id := y*grid.width + x
				if grid.rolls[id] == '@' && rollsAdjCnt(grid, x, y) < 4 {
					grid.rolls[id] = 'x'
					change = true
					total++
				}
			}
		}
	}
	return total
}

func rollsAdjCnt(grid Grid, x, y int) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}

			nx := x + dx
			ny := y + dy
			if nx < 0 || nx >= grid.width || ny < 0 || ny >= grid.height {
				continue
			}

			idx := ny*grid.width + nx
			if grid.rolls[idx] == '@' {
				count++
			}
		}
	}
	return count
}

func readInput(filename string) (Grid, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Grid{nil, 0, 0}, err
	}
	defer file.Close()

	var rolls []rune
	width := 0
	height := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		height++
		width = 0
		for _, c := range scanner.Text() {
			width++
			rolls = append(rolls, c)
		}
	}
	return Grid{rolls, width, height}, nil
}
