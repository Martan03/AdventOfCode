package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

const obstacle = math.MaxInt

func main() {
	input, x, y, err := readInput("input.txt")
	if err != nil {
		log.Fatalln("Error:", err)
		return
	}

	vis, loops := visitedChars(input, x, y, 0)
	fmt.Println("Visited:", vis)
	fmt.Println("Loops:  ", loops)
}

func readInput(filename string) ([][]int, int, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, 0, 0, err
	}
	defer file.Close()

	var plan [][]int
	var guardX, guardY int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line []int
		for i, c := range scanner.Text() {
			line = append(line, 0)
			if c == '#' {
				line[i] = obstacle
			}

			if c == '^' {
				guardX = i
				guardY = len(plan)
			}
		}

		plan = append(plan, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, 0, 0, err
	}
	return plan, guardX, guardY, nil
}

func visitedChars(plan [][]int, x, y, dir int) (int, int) {
	move(plan, x, y, dir)

	vis := getVisited(plan)
	part1 := reset(plan)

	return part1, block(plan, vis, x, y, dir)
}

func getDir(dir int) (int, int) {
	switch dir {
	case 0:
		return 0, -1
	case 1:
		return 1, 0
	case 2:
		return 0, 1
	case 3:
		return -1, 0
	default:
		return 0, 0
	}
}

func move(plan [][]int, x, y int, dir int) bool {
	dirX, dirY := getDir(dir)
	nextX := x + dirX
	nextY := y + dirY

	width := len(plan[0])
	height := len(plan)
	for nextX >= 0 && nextY >= 0 && nextX < width && nextY < height {
		if plan[y][x]&(1<<dir) != 0 {
			return true
		}

		plan[y][x] |= 1 << dir
		if plan[nextY][nextX] == obstacle {
			dir = (dir + 1) % 4
			dirX, dirY = getDir(dir)
			nextX = x + dirX
			nextY = y + dirY
			continue
		}
		x = nextX
		y = nextY

		nextX += dirX
		nextY += dirY
	}
	plan[y][x] |= 1 << dir
	return false
}

func block(plan [][]int, vis []int, x, y int, dir int) int {
	width := len(plan[0])

	loop_cnt := 0
	for _, val := range vis {
		cx := val % width
		cy := val / width
		if cx == x && cy == y {
			continue
		}

		plan[cy][cx] = obstacle
		if move(plan, x, y, dir) {
			loop_cnt++
		}
		plan[cy][cx] = 0
		reset(plan)
	}

	return loop_cnt
}

func getVisited(plan [][]int) []int {
	var visited []int
	width := len(plan[0])
	for y, r := range plan {
		for x, c := range r {
			if c > 0 && c != obstacle {
				visited = append(visited, x+y*width)
			}
		}
	}
	return visited
}

func reset(plan [][]int) int {
	visited := 0
	for y, r := range plan {
		for x, c := range r {
			if c > 0 && c != obstacle {
				plan[y][x] = 0
				visited++
			}
		}
	}
	return visited
}
