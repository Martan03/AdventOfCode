package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	stones, err := readStones("input.txt")
	if err != nil {
		log.Fatalln("Error:", err)
	}

	total := applyBlinks(stones, 75)
	fmt.Println("Total:", total)
}

func readStones(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var input []int
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		stones := strings.Fields(scanner.Text())
		for _, stone := range stones {
			num, err := strconv.Atoi(stone)
			if err != nil {
				return nil, err
			}
			input = append(input, num)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return input, nil
}

func applyBlinks(stones []int, blinks int) int {
	total := 0
	cache := make(map[int]map[int]int)
	for _, stone := range stones {
		total += applyBlink(stone, blinks, cache)
	}
	return total
}

func applyBlink(stone int, blinks int, cache map[int]map[int]int) int {
	if blinks == 0 {
		return 1
	}

	if stoneCache, exists := cache[stone]; exists {
		if res, exists := stoneCache[blinks]; exists {
			return res
		}
	} else {
		cache[stone] = make(map[int]int)
	}

	blinks--
	res := 0
	if stone == 0 {
		res = applyBlink(1, blinks, cache)
	} else if len := numLen(stone); len%2 == 0 {
		first, second := splitNum(stone, len)
		res = applyBlink(first, blinks, cache)
		res += applyBlink(second, blinks, cache)
	} else {
		res = applyBlink(stone*2024, blinks, cache)
	}
	cache[stone][blinks+1] = res
	return res
}

func numLen(num int) int {
	len := 0
	for num > 0 {
		num /= 10
		len++
	}
	return len
}

func splitNum(num, len int) (int, int) {
	len /= 2
	div := 1
	for len > 0 {
		div *= 10
		len--
	}
	return num % div, num / div
}
