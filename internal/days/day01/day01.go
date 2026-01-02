package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var currentPosition Position = 50

type Rotation struct {
	direction string
	tics      int
}

type Combination struct {
	positions []Position
}

type Position int

func newRotation(direction string, tics int) *Rotation {
	return &Rotation{direction: direction, tics: tics}
}

func newCombination() *Combination {
	positions := make([]Position, 0)
	return &Combination{positions: positions}
}

func SolvePart1(input string) (int, error) {
	combination := newCombination()
	combination.positions = append(combination.positions, 50)

	file, err := os.Open(input)
	if err != nil {
		return 0, err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	rotations := []Rotation{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0:1]
		if direction != "L" && direction != "R" {
			return 0, fmt.Errorf("%s is not a valid rotation direction", direction)
		}
		tics, err := strconv.Atoi(line[1:])
		if err != nil {
			return 0, err
		}
		rotation := newRotation(direction, tics)
		rotations = append(rotations, *rotation)
	}

	zeroCount := 0
	for _, rotation := range rotations {
		switch rotation.direction {
		case "L":
			currentPosition -= Position(rotation.tics)
			if currentPosition < 0 {
				currentPosition += 100 * ((-currentPosition / 100) + 1)
			}
		case "R":
			currentPosition += Position(rotation.tics)
		}
		currentPosition = currentPosition % 100
		if currentPosition == 0 {
			zeroCount++
		}

		combination.positions = append(combination.positions, currentPosition)
	}
	return zeroCount, nil
}

// floorDiv computes floor(a/b) for integer division (rounds toward negative infinity)
func floorDiv(a, b int) int {
	if a >= 0 {
		return a / b
	}
	return (a - b + 1) / b
}

func SolvePart2(input string) (int, error) {
	currentPosition = 50 // Reset position for Part 2
	combination := newCombination()
	combination.positions = append(combination.positions, 50)

	file, err := os.Open(input)
	if err != nil {
		return 0, err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	rotations := []Rotation{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0:1]
		if direction != "L" && direction != "R" {
			return 0, fmt.Errorf("%s is not a valid rotation direction", direction)
		}
		tics, err := strconv.Atoi(line[1:])
		if err != nil {
			return 0, err
		}
		rotation := newRotation(direction, tics)
		rotations = append(rotations, *rotation)
	}

	zeroCount := 0
	for _, rotation := range rotations {
		p := int(currentPosition)
		t := rotation.tics

		switch rotation.direction {
		case "L":
			// Positions visited: p-1, p-2, ..., p-t
			// Count how many are ≡ 0 (mod 100)
			zeroCount += floorDiv(p-1, 100) - floorDiv(p-t-1, 100)
			currentPosition -= Position(t)
			for currentPosition < 0 {
				currentPosition += 100
			}
		case "R":
			// Positions visited: p+1, p+2, ..., p+t
			// Count how many are ≡ 0 (mod 100)
			zeroCount += floorDiv(p+t, 100) - floorDiv(p, 100)
			currentPosition += Position(t)
		}
		currentPosition = currentPosition % 100

		combination.positions = append(combination.positions, currentPosition)
	}
	return zeroCount, nil
}
