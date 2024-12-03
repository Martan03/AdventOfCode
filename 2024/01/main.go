package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	first, second, err := readInput("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	dist := calcDist(first, second)
	fmt.Println("Distance:  ", dist)

	similarity := calcSimilarity(first, second)
	fmt.Println("Similarity:", similarity)
}

func readInput(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var first []int
	var second []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid format")
		}

		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			return nil, nil, fmt.Errorf("not a number found")
		}

		first = append(first, num1)
		second = append(second, num2)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return first, second, nil
}

func calcDist(first []int, second []int) int {
	sort.Ints(first)
	sort.Ints(second)

	dist := 0
	for i := 0; i < len(first); i++ {
		dif := first[i] - second[i]
		if dif < 0 {
			dif *= -1
		}
		dist += dif
	}
	return dist
}

func calcSimilarity(first []int, second []int) int {
	cnts := make(map[int]int)
	for _, num := range second {
		cnts[num]++
	}

	similarity := 0
	for _, num := range first {
		cnt := cnts[num]
		similarity += num * cnt
	}

	return similarity
}
