package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type Vec2 struct {
	x, y int
}

func main() {
	grid, err := readInput("input.txt")
	if err != nil {
		log.Fatalln("Error:", err)
		return
	}

	areas := getAreas(grid)
	total1, total2 := getPrice(grid, areas)
	fmt.Println("Total 1:", total1)
	fmt.Println("Total 1:", total2)
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

func getAreas(grid []string) [][]Vec2 {
	parsed := make(map[Vec2]bool)
	var areas [][]Vec2
	for y, line := range grid {
		for x := range line {
			pos := Vec2{x, y}
			if parsed[Vec2{x, y}] {
				continue
			}
			parsed[pos] = true
			areas = append(areas, getArea(grid, pos, parsed))
		}
	}
	return areas
}

func getArea(grid []string, pos Vec2, seen map[Vec2]bool) []Vec2 {
	c := grid[pos.y][pos.x]

	var queue []Vec2
	queue = append(queue, pos)
	for i := 0; i < len(queue); i++ {
		cpos := queue[i]

		if len(grid[0]) > cpos.x+1 && grid[cpos.y][cpos.x+1] == c {
			npos := Vec2{cpos.x + 1, cpos.y}
			if seen[npos] == false {
				queue = append(queue, Vec2{cpos.x + 1, cpos.y})
				seen[npos] = true
			}
		}
		if len(grid) > cpos.y+1 && grid[cpos.y+1][cpos.x] == c {
			npos := Vec2{cpos.x, cpos.y + 1}
			if seen[npos] == false {
				queue = append(queue, Vec2{cpos.x, cpos.y + 1})
				seen[npos] = true
			}
		}
		if cpos.y-1 >= 0 && grid[cpos.y-1][cpos.x] == c {
			npos := Vec2{cpos.x, cpos.y - 1}
			if seen[npos] == false {
				queue = append(queue, Vec2{cpos.x, cpos.y - 1})
				seen[npos] = true
			}
		}
		if cpos.x-1 >= 0 && grid[cpos.y][cpos.x-1] == c {
			npos := Vec2{cpos.x - 1, cpos.y}
			if seen[npos] == false {
				queue = append(queue, Vec2{cpos.x - 1, cpos.y})
				seen[npos] = true
			}
		}
	}
	return queue
}

func getPrice(grid []string, areas [][]Vec2) (int, int) {
	total1 := 0
	total2 := 0
	for _, area := range areas {
		total1 += getPerimeter(grid, area) * len(area)
		total2 += countSides(grid, area) * len(area)
	}
	return total1, total2
}

func getPerimeter(grid []string, area []Vec2) int {
	total := 0
	c := grid[area[0].y][area[0].x]

	w, h := len(grid[0]), len(grid)
	for _, pos := range area {
		if pos.x+1 >= w || grid[pos.y][pos.x+1] != c {
			total++
		}
		if pos.x-1 < 0 || grid[pos.y][pos.x-1] != c {
			total++
		}
		if pos.y+1 >= h || grid[pos.y+1][pos.x] != c {
			total++
		}
		if pos.y-1 < 0 || grid[pos.y-1][pos.x] != c {
			total++
		}
	}
	return total
}

func countSides(grid []string, area []Vec2) int {
	total := 0
	c := grid[area[0].y][area[0].x]

	sort.Slice(area, func(i, j int) bool {
		if area[i].x == area[j].x {
			return area[i].y < area[j].y
		}
		return area[i].x < area[j].x
	})

	prev := Vec2{-1, -1}
	prevSide := Vec2{0, 0}
	w, h := len(grid[0]), len(grid)
	for _, pos := range area {
		cond1 := prev.x == pos.x && prev.y+1 == pos.y && prevSide.x == 1
		cond2 := prev.x == pos.x && prev.y+1 == pos.y && prevSide.y == 1
		prevSide = Vec2{0, 0}

		if pos.x+1 >= w || grid[pos.y][pos.x+1] != c {
			prevSide.x = 1
			if !cond1 {
				total++
			}
		}
		if pos.x-1 < 0 || grid[pos.y][pos.x-1] != c {
			prevSide.y = 1
			if !cond2 {
				total++
			}
		}
		prev = pos
	}

	sort.Slice(area, func(i, j int) bool {
		if area[i].y == area[j].y {
			return area[i].x < area[j].x
		}
		return area[i].y < area[j].y
	})
	prev = Vec2{-1, -1}
	prevSide = Vec2{0, 0}
	for _, pos := range area {
		cond1 := prev.x+1 == pos.x && prev.y == pos.y && prevSide.x == 1
		cond2 := prev.x+1 == pos.x && prev.y == pos.y && prevSide.y == 1
		prevSide = Vec2{0, 0}

		if pos.y+1 >= h || grid[pos.y+1][pos.x] != c {
			prevSide.x = 1
			if !cond1 {
				total++
			}
		}
		if pos.y-1 < 0 || grid[pos.y-1][pos.x] != c {
			prevSide.y = 1
			if !cond2 {
				total++
			}
		}
		prev = pos
	}

	return total
}

func checkPos(grid []string, pos Vec2) bool {
	return pos.y < len(grid) && pos.y >= 0 &&
		pos.x >= 0 && pos.x < len(grid[0])
}
