package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const Max = 99

func main() {
	rots, err := readInput("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	part1 := part1(rots, 50)
	fmt.Println("Part 1:", part1)

	part2 := part2(rots, 50)
	fmt.Println("Part 2:", part2)
}

func part1(rots []int, rot int) int {
	zeros := 0
	for _, r := range rots {
		rot = (rot + r) % (Max + 1)
		if rot == 0 {
			zeros += 1
		}
	}
	return zeros
}

func part2(rots []int, rot int) int {
	zeros := 0
	for _, r := range rots {
		new := rot + r
		if new <= 0 {
			zeros += -new / (Max + 1)
			if rot > 0 {
				zeros += 1
			}
		} else {
			zeros += new / (Max + 1)
		}
		rot = new % (Max + 1)
		if rot < 0 {
			rot += Max + 1
		}
	}
	return zeros
}

func readInput(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rots []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		num, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, fmt.Errorf("expected number after rotation direction")
		}

		char := line[0]
		if char == 'L' {
			num = -num
		}
		rots = append(rots, num)
	}
	return rots, nil
}
