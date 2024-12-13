package main

import (
	"fmt"
	"log"
	"os"
)

type Vec2 struct {
	x, y int
}

func (v Vec2) add(other Vec2) Vec2 {
	return Vec2{v.x + other.x, v.y + other.y}
}

func (v Vec2) sub(other Vec2) Vec2 {
	return Vec2{v.x - other.x, v.y - other.y}
}

type Machine struct {
	buttonA Vec2
	buttonB Vec2
	price   Vec2
}

func main() {
	machines, err := readInput("input.txt")
	if err != nil {
		log.Fatalln("Error:", err)
		return
	}

	tokens1, tokens2 := tokensToWin(machines)
	fmt.Println("Tokens 1:", tokens1)
	fmt.Println("Tokens 2:", tokens2)
}

func readInput(filename string) ([]Machine, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var a, b, p Vec2
	var machines []Machine
	for {
		_, err := fmt.Fscanf(file, "Button A: X+%d, Y+%d\n", &a.x, &a.y)
		if err != nil {
			break
		}
		_, err = fmt.Fscanf(file, "Button B: X+%d, Y+%d\n", &b.x, &b.y)
		if err != nil {
			break
		}
		_, err = fmt.Fscanf(file, "Prize: X=%d, Y=%d\n\n", &p.x, &p.y)
		if err != nil {
			break
		}

		machines = append(machines, Machine{a, b, p})
	}
	return machines, nil
}

func tokensToWin(machines []Machine) (int, int) {
	tokens1, tokens2 := 0, 0
	for _, machine := range machines {
		tokens1 += findCheapest(machine)
		machine.price.x += 10000000000000
		machine.price.y += 10000000000000
		tokens2 += findCheapest(machine)
	}
	return tokens1, tokens2
}

func findCheapest(m Machine) int {
	y := m.buttonB.x*m.buttonA.y - m.buttonB.y*m.buttonA.x
	res := m.price.x*m.buttonA.y - m.price.y*m.buttonA.x

	if res%y != 0 {
		return 0
	}
	y = res / y

	x := m.price.x - (m.buttonB.x * y)
	if x%m.buttonA.x != 0 {
		return 0
	}
	x /= m.buttonA.x

	if x < 0 || y < 0 {
		return 0
	}
	return x*3 + y
}
