package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Vec2 struct {
	x, y int
}

func main() {
	points, err := readInput("input.txt")
	if err != nil {
		log.Fatalln("Error:", err)
		return
	}
	fmt.Println("Part1:", part1(points))
}

func part1(points []Vec2) int {
	max := 0
	for _, p1 := range points {
		for _, p2 := range points {
			x := p1.x - p2.x
			if x < 0 {
				x = -x
			}
			y := p1.y - p2.y
			if y < 0 {
				y = -y
			}

			area := (x + 1) * (y + 1)
			if area > max {
				max = area
			}
		}
	}
	return max
}

func readInput(filename string) ([]Vec2, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var points []Vec2
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		if len(parts) != 2 {
			return nil, fmt.Errorf("expects 2D point coordinates")
		}
		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("invalid point format")
		}

		points = append(points, Vec2{num1, num2})
	}
	return points, nil
}
