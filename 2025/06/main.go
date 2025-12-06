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
	filename := "input.txt"
	fmt.Println("Part1:", calc(filename, readInput1))
	fmt.Println("Part2:", calc(filename, readInput2))
}

func readInput1(filename string) ([]string, [][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var columns [][]int
	var ops []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if strings.ContainsAny(parts[0], "*+") {
			ops = parts
			continue
		}
		if len(columns) == 0 {
			columns = make([][]int, len(parts))
		}
		for i, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, nil, fmt.Errorf("not a number found")
			}
			columns[i] = append(columns[i], num)
		}
	}

	return ops, columns, nil
}

func readInput2(filename string) ([]string, [][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var lines []string
	var ops []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.ContainsAny(line, "*+") {
			ops = strings.Fields(line)
			continue
		}

		line_len := len(line)
		if len(lines) == 0 {
			lines = make([]string, line_len)
		}
		for i, c := range line {
			lines[i] += string(c)
		}
	}

	columns := make([][]int, len(ops))
	col_id := 0
	for _, line := range lines {
		tline := strings.TrimSpace(line)
		if len(tline) == 0 {
			col_id++
			continue
		}
		num, err := strconv.Atoi(tline)
		if err != nil {
			return nil, nil, fmt.Errorf("not a number found")
		}
		columns[col_id] = append(columns[col_id], num)
	}

	return ops, columns, nil
}

func calc(filename string, input func(string) ([]string, [][]int, error)) int {
	ops, columns, err := input(filename)
	if err != nil {
		log.Fatalln("Error:", err)
		return 0
	}

	total := 0
	for i, col := range columns {
		res := col[0]
		for j := 1; j < len(col); j++ {
			switch ops[i] {
			case "+":
				res += col[j]
			case "*":
				res *= col[j]
			}
		}
		total += res
	}
	return total
}
