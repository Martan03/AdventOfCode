package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	res  int
	nums []int
}

func main() {
	res, resImp, err := checkEquations("input.txt")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	fmt.Println("Total:    ", res)
	fmt.Println("Total imp:", resImp)
}

func checkEquations(filename string) (int, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	totalImp := 0
	for scanner.Scan() {
		eq, err := parseLine(scanner.Text())
		if err != nil {
			return 0, 0, err
		}

		if checkEquation(eq, eq.nums[0], 1) {
			total += eq.res
		}
		if checkEquationImp(eq, eq.nums[0], 1) {
			totalImp += eq.res
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, err
	}

	return total, totalImp, nil
}

func parseLine(line string) (Equation, error) {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return Equation{}, fmt.Errorf("invalid line format:\n%s", line)
	}

	res, err := strconv.Atoi(parts[0])
	if err != nil {
		return Equation{}, err
	}

	fields := strings.Fields(parts[1])
	eq := Equation{res, make([]int, len(fields))}
	for i, field := range fields {
		eq.nums[i], err = strconv.Atoi(field)
		if err != nil {
			return Equation{}, err
		}
	}
	return eq, nil
}

func checkEquation(eq Equation, res, cur int) bool {
	if cur >= len(eq.nums) {
		return eq.res == res
	}

	if checkEquation(eq, res+eq.nums[cur], cur+1) ||
		checkEquation(eq, res*eq.nums[cur], cur+1) {
		return true
	}
	return false
}

func checkEquationImp(eq Equation, res, cur int) bool {
	if cur >= len(eq.nums) {
		return eq.res == res
	}

	if checkEquationImp(eq, res+eq.nums[cur], cur+1) ||
		checkEquationImp(eq, res*eq.nums[cur], cur+1) {
		return true
	}

	num, err := numConcat(res, eq.nums[cur])
	if err == nil && checkEquationImp(eq, num, cur+1) {
		return true
	}

	return false
}

func numConcat(a, b int) (int, error) {
	str := strconv.Itoa(a) + strconv.Itoa(b)
	return strconv.Atoi(str)
}
