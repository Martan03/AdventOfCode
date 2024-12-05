package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	totalValid, totalCorrected, err := middleTotal("input.txt")
	if err != nil {
		log.Fatal("Error:", err)
		return
	}
	fmt.Println("Total valid:    ", totalValid)
	fmt.Println("Total corrected:", totalCorrected)
}

func middleTotal(filename string) (int, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	order, orderErr := readOrder(scanner)
	if orderErr != nil {
		return 0, 0, orderErr
	}

	return checkUpdates(scanner, order)
}

func readOrder(scanner *bufio.Scanner) (map[int][]int, error) {
	order := make(map[int][]int)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "|")
		if len(parts) != 2 {
			break
		}

		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("not a number found")
		}

		order[num1] = append(order[num1], num2)
	}
	return order, nil
}

func checkUpdates(scan *bufio.Scanner, order map[int][]int) (int, int, error) {
	totalValid := 0
	totalCorrected := 0
	for {
		update, err := readUpdate(scan)
		if update == nil {
			break
		}

		if err != nil {
			return 0, 0, err
		}

		if checkUpdate(update, order) {
			totalValid += update[len(update)/2]
		} else if cor := correctUpdate(update, order); cor != nil {
			totalCorrected += cor[len(cor)/2]
		}
	}

	return totalValid, totalCorrected, nil
}

func readUpdate(scanner *bufio.Scanner) ([]int, error) {
	var update []int
	if !scanner.Scan() {
		return nil, nil
	}

	parts := strings.Split(scanner.Text(), ",")

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("not a number found")
		}
		update = append(update, num)
	}

	return update, nil
}

func checkUpdate(update []int, order map[int][]int) bool {
	seen := make(map[int]bool)
	for _, num := range update {
		for _, after := range order[num] {
			if seen[after] {
				return false
			}
		}
		seen[num] = true
	}
	return true
}

func correctUpdate(update []int, order map[int][]int) []int {
	subOrder := make(map[int][]int)
	degree := make(map[int]int)
	for _, num := range update {
		for _, neighbor := range order[num] {
			if contains(update, neighbor) {
				subOrder[num] = append(subOrder[num], neighbor)
				degree[neighbor]++
			}
		}
		if _, ok := degree[num]; !ok {
			degree[num] = 0
		}
	}

	validOrder := degreeSort(subOrder, degree)
	if validOrder == nil {
		return nil
	}

	orderIndex := make(map[int]int)
	for i, val := range validOrder {
		orderIndex[val] = i
	}

	sort.Slice(update, func(i, j int) bool {
		return orderIndex[update[i]] < orderIndex[update[j]]
	})

	return update
}

func contains(values []int, value int) bool {
	for _, val := range values {
		if val == value {
			return true
		}
	}
	return false
}

func degreeSort(order map[int][]int, degree map[int]int) []int {
	var queue []int
	for num, degree := range degree {
		if degree == 0 {
			queue = append(queue, num)
		}
	}

	var result []int
	for len(queue) > 0 {
		num := queue[0]
		queue = queue[1:]
		result = append(result, num)

		for _, after := range order[num] {
			degree[after]--
			if degree[after] == 0 {
				queue = append(queue, after)
			}
		}
	}

	if len(result) != len(degree) {
		return nil
	}

	return result
}
