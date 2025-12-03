package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	banks, err := readInput("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Part1:", joltage(banks, 2))
	fmt.Println("Part2:", joltage(banks, 12))
}

func joltage(banks [][]int, n int) int {
	total := 0
	for _, bank := range banks {
		max := 0
		start := 0
		end := len(bank) - n + 1
		for range n {
			lmax, maxId := getMax(bank, start, end)
			max = max*10 + lmax
			start = maxId + 1
			end++
		}

		total += max
	}
	return total
}

func getMax(bank []int, start, end int) (int, int) {
	max := 0
	maxId := -1
	for i := start; i < end; i++ {
		if max < bank[i] {
			max = bank[i]
			maxId = i
		}
	}
	return max, maxId
}

func readInput(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var banks [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var bank []int
		for _, c := range scanner.Text() {
			bank = append(bank, int(c-'0'))
		}
		banks = append(banks, bank)
	}
	return banks, nil
}
