package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Vec2 struct {
	x, y int
}

func (v Vec2) add(other Vec2) Vec2 {
	return Vec2{v.x + other.x, v.y + other.y}
}

func (v Vec2) sub(other Vec2) Vec2 {
	return Vec2{v.x - other.x, v.y - other.y}
}

func (v Vec2) checkBounds(grid []string) bool {
	return v.x >= 0 && v.y >= 0 && v.x < len(grid[0]) && v.y < len(grid)
}

func main() {
	grid, err := readInput("input.txt")
	if err != nil {
		log.Fatalln("Error:", err)
	}

	total := countAntinodes(grid, checkAnti)
	total2 := countAntinodes(grid, checkAnti2)
	fmt.Println("Total: ", total)
	fmt.Println("Total2:", total2)
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

func countAntinodes(
	grid []string,
	check func([]string, map[Vec2]bool, rune, int, int),
) int {
	locs := make(map[Vec2]bool)
	for y, line := range grid {
		for x, node := range line {
			if node == '.' {
				continue
			}
			check(grid, locs, node, x, y)
		}
	}
	return len(locs)
}

func checkAnti(grid []string, locs map[Vec2]bool, cmp rune, x, y int) {
	for yi := y; yi < len(grid); yi++ {
		for xi, node := range grid[yi] {
			if (yi <= y && xi <= x) || cmp != node {
				continue
			}

			xd := xi - x
			yd := yi - y
			if checkBounds(grid, x-xd, y-yd) {
				locs[Vec2{x - xd, y - yd}] = true
			}
			if checkBounds(grid, xi+xd, yi+yd) {
				locs[Vec2{xi + xd, yi + yd}] = true
			}
		}
	}
}

func checkAnti2(grid []string, locs map[Vec2]bool, cmp rune, x, y int) {
	for yi := y; yi < len(grid); yi++ {
		for xi, node := range grid[yi] {
			if (yi <= y && xi <= x) || cmp != node {
				continue
			}

			dir := Vec2{xi - x, yi - y}
			for v := (Vec2{xi, yi}); v.checkBounds(grid); v = v.add(dir) {
				locs[v] = true
			}
			for v := (Vec2{x, y}); v.checkBounds(grid); v = v.sub(dir) {
				locs[v] = true
			}
		}
	}
}

func checkBounds(grid []string, x int, y int) bool {
	return y >= 0 && y < len(grid) && x >= 0 && x < len(grid[0])
}
