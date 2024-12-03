package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	err := reportSafeCnt("input.txt", isReportSafe1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = reportSafeCnt("input.txt", isReportSafe2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func reportSafeCnt(
	filename string,
	isSafe func(string) (bool, error),
) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cnt := 0
	for scanner.Scan() {
		line := scanner.Text()
		safe, err := isSafe(line)
		if err != nil {
			break
		}

		if safe {
			cnt++
		}
	}

	fmt.Println("Safe count:", cnt)
	return nil
}

func isReportSafe1(report string) (bool, error) {
	levels := strings.Fields(report)
	if len(levels) == 0 {
		return false, nil
	}

	prev, prevErr := strconv.Atoi(levels[0])
	prevDir := 0
	if prevErr != nil {
		return false, prevErr
	}

	for i := 1; i < len(levels); i++ {
		num, numErr := strconv.Atoi(levels[i])
		if numErr != nil {
			return false, prevErr
		}

		dir := num - prev
		if dir*prevDir < 0 || dir < -3 || 3 < dir || num == prev {
			return false, nil
		}

		prev = num
		prevDir = dir
	}
	return true, nil
}

func isReportSafe2(report string) (bool, error) {
	levels := strings.Fields(report)
	if len(levels) == 0 {
		return false, nil
	}

	var nums []int
	for _, level := range levels {
		level, err := strconv.Atoi(level)
		if err != nil {
			return false, fmt.Errorf("not a number found")
		}
		nums = append(nums, level)
	}

	if checkReport(nums) {
		return true, nil
	}

	for i := 0; i < len(nums); i++ {
		level_nums := append([]int{}, nums...)
		level_nums = append(level_nums[:i], level_nums[i+1:]...)
		if checkReport(level_nums) {
			return true, nil
		}
	}

	return false, nil
}

func checkReport(levels []int) bool {
	prev := levels[0]
	prevDir := 0
	for i := 1; i < len(levels); i++ {
		dir := levels[i] - prev
		if dir*prevDir < 0 || dir < -3 || 3 < dir || levels[i] == prev {
			return false
		}

		prev = levels[i]
		prevDir = dir
	}
	return true
}
