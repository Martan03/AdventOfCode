package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

type Vec2 struct {
	x, y int
}

func main() {
	hmap, err := readMap("input.txt")
	if err != nil {
		log.Fatalln("Error:", err)
		return
	}

	total, totalTrails := trailHeadTotal(hmap)
	fmt.Println("Total:       ", total)
	fmt.Println("Total trails:", totalTrails)
}

func readMap(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var input [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line []int
		for _, c := range scanner.Text() {
			if !unicode.IsDigit(c) {
				return nil, err
			}
			line = append(line, int(c-'0'))
		}
		input = append(input, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return input, nil
}

func trailHeadTotal(hmap [][]int) (int, int) {
	total := 0
	totalTops := 0
	for y, line := range hmap {
		for x, val := range line {
			if val == 0 {
				tops := make(map[Vec2]int)
				total += checkTrail(hmap, tops, -1, x, y)
				totalTops += len(tops)
			}
		}
	}
	return totalTops, total
}

func checkTrail(hmap [][]int, tops map[Vec2]int, prev, curX, curY int) int {
	if !checkBounds(hmap, curX, curY) || prev+1 != hmap[curY][curX] {
		return 0
	}

	cur := hmap[curY][curX]
	if cur == 9 {
		tops[Vec2{curX, curY}]++
		return 1
	}

	total := checkTrail(hmap, tops, cur, curX+1, curY)
	total += checkTrail(hmap, tops, cur, curX-1, curY)
	total += checkTrail(hmap, tops, cur, curX, curY+1)
	total += checkTrail(hmap, tops, cur, curX, curY-1)
	return total
}

func checkBounds(hmap [][]int, x, y int) bool {
	return x >= 0 && y >= 0 && y < len(hmap) && x < len(hmap[0])
}
