package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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

func readInput(filename string) ([]string, int, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, 0, 0, err
	}
	defer file.Close()

	var plan []string
	var guardX, guardY int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for i, c := range line {
			if c != '#' && c != '.' {
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

func visitedChars(plan []string, x, y, dir int) (int, int) {
	vis := make(map[int]int)
	move(plan, vis, x, y, dir)

	return len(vis), block(plan, vis, x, y, dir)
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

func move(plan []string, vis map[int]int, x, y int, dir int) bool {
	dirX, dirY := getDir(dir)
	nextX := x + dirX
	nextY := y + dirY

	width := len(plan[0])
	height := len(plan)
	for nextX >= 0 && nextY >= 0 && nextX < width && nextY < height {
		if d, e := vis[x+y*width]; e && d&(1<<dir) != 0 {
			return true
		}

		vis[x+y*width] |= 1 << dir
		if plan[nextY][nextX] == '#' || plan[nextY][nextX] == 'O' {
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
	vis[x+y*width] |= 1 << dir
	return false
}

func block(plan []string, vis map[int]int, x, y int, dir int) int {
	width := len(plan[0])

	loop_cnt := 0
	for pos, val := range vis {
		cx := pos % width
		cy := pos / width
		if val == 0 || (cx == x && cy == y) {
			continue
		}

		plan[cy] = setCharAt(plan[cy], cx, 'O')
		lvis := make(map[int]int)
		if move(plan, lvis, x, y, dir) {
			loop_cnt++
		}
		plan[cy] = setCharAt(plan[cy], cx, '.')
	}

	return loop_cnt
}

func setCharAt(str string, id int, c rune) string {
	runes := []rune(str)
	runes[id] = c
	return string(runes)
}
