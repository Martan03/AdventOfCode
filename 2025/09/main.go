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
	fmt.Println("Part2:", part2(points))
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

func part2(points []Vec2) int {
	points = append(points, points[0:1]...)

	max := 0
	for i := 0; i < len(points)-2; i++ {
		p1 := points[i]
		for j := i + 1; j < len(points)-2; j++ {
			p2 := points[j]
			x := p1.x - p2.x
			if x < 0 {
				x = -x
			}
			y := p1.y - p2.y
			if y < 0 {
				y = -y
			}

			area := (x + 1) * (y + 1)
			if area > max && validArea(points, p1, p2) {
				max = area
			}
		}
	}
	return max
}

func validArea(points []Vec2, p1, p2 Vec2) bool {
	minX, maxX := minMaxInt(p1.x, p2.x)
	minY, maxY := minMaxInt(p1.y, p2.y)
	for i := 1; i < len(points); i++ {
		a := points[i-1]
		b := points[i]
		if p1 == a || p1 == b || p2 == a || p2 == b {
			continue
		}

		lx, hx := minMaxInt(a.x, b.x)
		ly, hy := minMaxInt(a.y, b.y)
		if (a.x == b.x && minX < a.x && a.x < maxX && ly < maxY && minY < hy) ||
			(a.y == b.y && minY < a.y && a.y < maxY && lx < maxX && minX < hx) {
			return false
		}
	}
	return true
}

func minMaxInt(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
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
