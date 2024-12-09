package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"unicode"
)

const empty = math.MaxInt

func main() {
	blocks, err := readBlocks("input.txt")
	if err != nil {
		log.Fatalln("Error:", err)
	}

	scheme := blocksToScheme(blocks)
	schemeCopy := make([]int, len(scheme))
	copy(schemeCopy, scheme)

	squeezeBlocks(scheme)
	sum := checkSum(scheme)
	fmt.Println("Check sum:  ", sum)

	rearangeBlocks(schemeCopy)
	sum = checkSum(schemeCopy)
	fmt.Println("Check sum 2:", sum)
}

func readBlocks(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var input []int
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		for _, c := range scanner.Text() {
			if !unicode.IsDigit(c) {
				return nil, err
			}
			input = append(input, int(c-'0'))
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return input, nil
}

func blocksToScheme(blocks []int) []int {
	var scheme []int
	id := 0
	for i, cnt := range blocks {
		c := empty
		if i%2 == 0 {
			c = id
			id++
		}

		for j := 0; j < cnt; j++ {
			scheme = append(scheme, c)
		}
	}
	return scheme
}

func squeezeBlocks(scheme []int) {
	start, end := 0, len(scheme)-1
	for start < end {
		if scheme[start] != empty {
			start++
			continue
		}

		if scheme[end] == empty {
			end--
			continue
		}

		temp := scheme[start]
		scheme[start] = scheme[end]
		scheme[end] = temp
	}
}

func rearangeBlocks(scheme []int) {
	for cur := len(scheme) - 1; cur >= 0; cur-- {
		if scheme[cur] == empty {
			continue
		}

		len := blockLen(scheme, cur)
		moveIfPossible(scheme, cur, len)
		cur -= len - 1
	}
}

func blockLen(scheme []int, id int) int {
	len := 1
	for id-len >= 0 && scheme[id] == scheme[id-len] {
		len++
	}
	return len
}

func moveIfPossible(scheme []int, id, len int) {
	spaces := 0
	for i := 0; i <= id-len; i++ {
		if scheme[i] == empty {
			spaces++
		} else {
			spaces = 0
		}

		if spaces == len {
			setVals(scheme, i-len+1, i, scheme[id])
			setVals(scheme, id-len+1, id, empty)
		}
	}
}

func setVals(scheme []int, start, end, val int) {
	for start <= end {
		scheme[start] = val
		start++
	}
}

func checkSum(scheme []int) int {
	total := 0
	for i, val := range scheme {
		if val == empty {
			continue
		}
		total += i * val
	}
	return total
}
