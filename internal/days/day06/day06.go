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

func SolvePart2(input string) (int, error) {
	return 0, nil
}
