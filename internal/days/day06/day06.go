package day06

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Operation byte

const (
	Multiply Operation = '*'
	Add      Operation = '+'
)

type Problem struct {
	ProblemOperation Operation
	ProblemNumbers   []int
}

func isOperationByte(char byte) bool {
	charOperation := Operation(char)
	if charOperation == Multiply || charOperation == Add {
		return true
	}
	return false
}

func parseProblems(input string) ([]Problem, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err2 := file.Close()
		if err2 != nil {
			log.Fatal(err2)
		}
	}(file)

	var problems []Problem
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if isOperationByte(parts[0][0]) {
			if len(parts) != len(problems) {
				return nil, fmt.Errorf("expected %d parts, got %d", len(problems), len(parts))
			}
			for index, part := range parts {
				partOperation := Operation(part[0])
				problems[index].ProblemOperation = partOperation
			}
			break
		}

		problemNumbers := make([]int, len(parts))
		for index, part := range parts {
			problemNumbers[index], err = strconv.Atoi(part)
			if err != nil {
				return nil, err
			}
		}

		if len(problems) == 0 {
			for _, problemNumber := range problemNumbers {
				problem := Problem{
					ProblemNumbers: []int{
						problemNumber,
					},
				}
				problems = append(problems, problem)
			}
			continue
		}

		if len(problemNumbers) != len(problems) {
			return nil, fmt.Errorf("expected %d parts, got %d", len(problems), len(problems))
		}
		for index, problemNumber := range problemNumbers {
			problems[index].ProblemNumbers = append(problems[index].ProblemNumbers, problemNumber)
		}
	}

	return problems, scanner.Err()
}

func SolvePart1(input string) (int, error) {
	problems, err := parseProblems(input)
	if err != nil {
		return 0, err
	}

	grandTotal := 0
	for _, problem := range problems {
		problemTotal := 0
		switch problem.ProblemOperation {
		case Multiply:
			for _, number := range problem.ProblemNumbers {
				if problemTotal == 0 {
					problemTotal = 1
				}
				problemTotal *= number
			}
		case Add:
			for _, number := range problem.ProblemNumbers {
				problemTotal += number
			}
		}
		grandTotal += problemTotal
	}
	return grandTotal, nil
}

func parseOperations(lastLine string) ([]Operation, []int, error) {
	var operations []Operation
	var columnWidths []int
	var operatorPositions []int

	for index, char := range lastLine {
		if isOperationByte(byte(char)) {
			operations = append(operations, Operation(char))
			operatorPositions = append(operatorPositions, index)
		}
	}

	// Calculate column widths based on operator positions
	// Each column ends at the operator position, starts after previous operator + 1 (for space separator)
	for i := 0; i < len(operatorPositions); i++ {
		var width int
		if i == 0 {
			// First column starts at 0
			width = operatorPositions[i]
		} else {
			// Subsequent columns start after previous operator + 1 space
			width = operatorPositions[i] - operatorPositions[i-1] - 1
		}
		columnWidths = append(columnWidths, width)
	}

	return operations, columnWidths, nil
}

func parseNumbers(grid []string, columnWidths []int) ([][]int, error) {
	var numbersGrid [][]string
	for _, row := range grid {
		currentIndex := 0
		var numbersGridRow []string
		for _, width := range columnWidths {
			numbersGridRow = append(numbersGridRow, row[currentIndex:currentIndex+width])
			currentIndex += width + 1
		}
		numbersGrid = append(numbersGrid, numbersGridRow)
	}

	numberOfRows := len(numbersGrid)
	numberOfColumns := len(numbersGrid[0])
	var problemNumbersList [][]int
	// For each problem (column in the grid)
	for column := 0; column < numberOfColumns; column++ {
		// Find the width of this problem's section
		width := len(numbersGrid[0][column])
		var problemNumbers []int
		// For each character position (right to left = each number)
		for charPos := width - 1; charPos >= 0; charPos-- {
			number := 0
			// For each row (top to bottom = most significant to least significant digit)
			for row := 0; row < numberOfRows; row++ {
				numberString := numbersGrid[row][column]
				if charPos < len(numberString) && numberString[charPos] != ' ' {
					digit := int(numberString[charPos] - '0')
					number = number*10 + digit
				}
			}
			problemNumbers = append(problemNumbers, number)
		}
		problemNumbersList = append(problemNumbersList, problemNumbers)
	}
	return problemNumbersList, nil
}

func parseGridPart2(input string) ([]Problem, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err2 := file.Close()
		if err2 != nil {
			log.Fatal(err2)
		}
	}(file)

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	operations, columnWidths, err := parseOperations(grid[len(grid)-1])
	if err != nil {
		return nil, err
	}
	problemNumbersList, err := parseNumbers(grid[:len(grid)-1], columnWidths)
	if err != nil {
		return nil, err
	}
	var result []Problem
	for i, problemNumbers := range problemNumbersList {
		result = append(result, Problem{
			ProblemOperation: operations[i],
			ProblemNumbers:   problemNumbers,
		})
	}
	return result, nil
}

func SolvePart2(input string) (int, error) {
	problems, err := parseGridPart2(input)
	if err != nil {
		return 0, err
	}

	total := 0
	for _, problem := range problems {
		switch problem.ProblemOperation {
		case Multiply:
			if total == 0 {
				total = 1
			}
			for _, number := range problem.ProblemNumbers {
				total *= number
			}
		case Add:
			for _, number := range problem.ProblemNumbers {
				total += number
			}
		}
	}
	return total, nil
}
