package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	lights  string
	buttons [][]int
	joltage []int
}

type QueueItem struct {
	steps  int
	lights []bool
}

func main() {
	machines, err := readInput("input.txt")
	if err != nil {
		log.Fatalln("Error:", err)
		return
	}
	fmt.Println("Part1:", part1(machines))
}

func part1(machines []Machine) int {
	total := 0
	for _, machine := range machines {
		total += fixLights(machine)
	}
	return total
}

func fixLights(machine Machine) int {
	state := make([]bool, len(machine.lights))
	stateKey := lightsKey(state)
	if stateKey == machine.lights {
		return 0
	}

	queue := []QueueItem{{0, state}}
	visited := map[string]bool{stateKey: true}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, button := range machine.buttons {
			next := make([]bool, len(cur.lights))
			copy(next, cur.lights)
			for _, l := range button {
				next[l] = !next[l]
			}

			key := lightsKey(next)
			if key == machine.lights {
				return cur.steps + 1
			}

			if visited[key] {
				continue
			}

			visited[key] = true
			queue = append(queue, QueueItem{
				lights: next,
				steps:  cur.steps + 1,
			})
		}
	}
	return 0
}

func lightsKey(lights []bool) string {
	var sb strings.Builder
	for _, l := range lights {
		if l {
			sb.WriteRune('#')
		} else {
			sb.WriteRune('.')
		}
	}
	return sb.String()
}

func readInput(filename string) ([]Machine, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var machines []Machine
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		machine := Machine{}
		for part := range strings.FieldsSeq(scanner.Text()) {
			val := part[1 : len(part)-1]
			switch part[0] {
			case '[':
				machine.lights = val
			case '(':
				machine.buttons = append(machine.buttons, parseNums(val))
			case '{':
				machine.joltage = parseNums(val)
			default:
				return nil, fmt.Errorf("invalid machine field found")
			}
		}
		machines = append(machines, machine)
	}
	return machines, nil
}

func parseNums(value string) []int {
	var res []int
	for part := range strings.SplitSeq(value, ",") {
		num, err := strconv.Atoi(part)
		if err != nil {
			continue
		}
		res = append(res, num)
	}
	return res
}
