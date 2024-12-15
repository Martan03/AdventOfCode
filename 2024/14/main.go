package main

import (
	"fmt"
	"log"
	"os"
)

const WIDTH = 101
const HEIGHT = 103

type Vec2 struct {
	x, y int
}

func (v Vec2) add(other Vec2) Vec2 {
	return Vec2{v.x + other.x, v.y + other.y}
}

func (v Vec2) sub(other Vec2) Vec2 {
	return Vec2{v.x - other.x, v.y - other.y}
}

type Robot struct {
	pos Vec2
	vel Vec2
}

func main() {
	robots, err := readInput("input.txt")
	if err != nil {
		log.Fatalln("Error:", err)
		return
	}
	robotsClone := make([]Robot, len(robots))
	copy(robotsClone, robots)
	safety := simulate(robotsClone, 100)
	fmt.Println("Safety:", safety)

	cnt := easterEgg(robots)
	fmt.Println("Easter egg:", cnt)
}

func readInput(filename string) ([]Robot, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var p, v Vec2
	var robots []Robot
	for {
		_, err := fmt.Fscanf(file, "p=%d,%d v=%d,%d\n", &p.x, &p.y, &v.x, &v.y)
		if err != nil {
			break
		}

		robots = append(robots, Robot{p, v})
	}
	return robots, nil
}

func simulate(robots []Robot, seconds int) int {
	for i := 0; i < seconds; i++ {
		moveRobots(robots)
	}
	return safetyFactor(robots)
}

func moveRobots(robots []Robot) {
	for i := range robots {
		robots[i].pos.x = (robots[i].pos.x + robots[i].vel.x + WIDTH) % WIDTH
		robots[i].pos.y = (robots[i].pos.y + robots[i].vel.y + HEIGHT) % HEIGHT
	}
}

func safetyFactor(robots []Robot) int {
	var tl, tr, bl, br int
	midx := WIDTH / 2
	midy := HEIGHT / 2
	for _, robot := range robots {
		if robot.pos.x == midx || robot.pos.y == midy {
			continue
		}

		fx := robot.pos.x < midx
		fy := robot.pos.y < midy
		if fx && fy {
			tl++
		} else if fx && !fy {
			bl++
		} else if !fx && fy {
			tr++
		} else if !fx && !fy {
			br++
		}
	}
	return tl * tr * bl * br
}

func easterEgg(robots []Robot) int {
	cnt := 0
	for areOverlapping(robots) {
		moveRobots(robots)
		cnt++
	}
	return cnt
}

func areOverlapping(robots []Robot) bool {
	positions := make(map[Vec2]bool)
	for _, robot := range robots {
		if positions[robot.pos] {
			return true
		}
		positions[robot.pos] = true
	}
	return false
}
