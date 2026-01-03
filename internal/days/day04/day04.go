package day04

import (
	"bufio"
	"os"
)

type Grid [][]byte

type Point struct {
	x, y int
}

func parseGrid(inputFile string) (Grid, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			return
		}
	}(file)

	var grid Grid
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	return grid, scanner.Err()
}

func isRollAccessable(grid Grid, point Point) bool {
	if point.y > len(grid)-1 || point.x > len(grid[0])-1 || point.x < 0 || point.y < 0 {
		return false
	}
	offsets := []Point{
		{-1, -1}, {-1, 0}, {-1, 1}, // Top row
		{0, -1}, {0, 1}, // Middle row (excluding self)
		{1, -1}, {1, 0}, {1, 1}, // Bottom row
	}
	var neighbors []byte
	for _, offset := range offsets {
		neighborX := point.x + offset.x
		neighborY := point.y + offset.y

		// Check for boundary
		if neighborY >= 0 && neighborY < len(grid) && neighborX >= 0 && neighborX < len(grid[0]) {
			neighbor := grid[neighborY][neighborX]
			neighbors = append(neighbors, neighbor)
		}
	}
	rollsCount := 0
	for _, neighbor := range neighbors {
		if neighbor == '@' {
			rollsCount++
		}
		if rollsCount >= 4 {
			return false
		}
	}
	return true
}

func SolvePart1(input string) (int, error) {
	grid, err := parseGrid(input)
	if err != nil {
		return 0, err
	}

	accessableRollCount := 0
	for y, row := range grid {
		for x, _ := range row {
			point := Point{x, y}
			if grid[y][x] == '@' && isRollAccessable(grid, point) {
				accessableRollCount++
			}
		}
	}
	return accessableRollCount, nil
}

func updateGrid(grid Grid) (Grid, int) {
	updatedGrid := grid
	accessableRollCount := 0
	for y, row := range grid {
		for x, _ := range row {
			point := Point{x, y}
			if grid[y][x] == '@' && isRollAccessable(grid, point) {
				accessableRollCount++
				updatedGrid[y][x] = '.'
			}
		}
	}
	return updatedGrid, accessableRollCount
}

func SolvePart2(input string) (int, error) {
	grid, err := parseGrid(input)
	if err != nil {
		return 0, err
	}
	accessableTotalRollCount := 0
	for {
		var accessableRollCount int
		grid, accessableRollCount = updateGrid(grid)
		if accessableRollCount == 0 {
			break
		}
		accessableTotalRollCount += accessableRollCount
	}

	return accessableTotalRollCount, nil
}
