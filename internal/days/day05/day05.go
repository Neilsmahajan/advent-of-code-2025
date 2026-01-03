package day05

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IngrediantRange struct {
	Start, End int
}

func parseRangesAndIngredientIDs(inputFile string) ([]IngrediantRange, []int, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, nil, err
	}
	defer func(file *os.File) {
		err2 := file.Close()
		if err2 != nil {
			return
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var ingrediantRanges []IngrediantRange
	var index = 0
	line := lines[index]
	for line != "" {
		ingrediantStartString, ingrediantEndString, found := strings.Cut(line, "-")
		if !found {
			return nil, nil, fmt.Errorf("could not parse line: %s", line)
		}
		ingrediantStart, err2 := strconv.Atoi(ingrediantStartString)
		if err2 != nil {
			return nil, nil, fmt.Errorf("could not parse ingrediant start: %s", ingrediantStartString)
		}
		ingrediantEnd, err2 := strconv.Atoi(ingrediantEndString)
		if err2 != nil {
			return nil, nil, fmt.Errorf("could not parse ingrediant end: %s", ingrediantEndString)
		}
		ingrediantRanges = append(ingrediantRanges, IngrediantRange{
			Start: ingrediantStart,
			End:   ingrediantEnd,
		})
		index++
		line = lines[index]
	}
	var ingrediantIDs []int
	for i := index + 1; i < len(lines); i++ {
		line = lines[i]
		ingrediantID, err2 := strconv.Atoi(line)
		if err2 != nil {
			return nil, nil, fmt.Errorf("could not parse ingrediant id: %s", line)
		}
		ingrediantIDs = append(ingrediantIDs, ingrediantID)
	}
	return ingrediantRanges, ingrediantIDs, nil
}

func ingrediantIsFresh(ingrediantRanges []IngrediantRange, ingrediantID int) bool {
	for _, ingrediantRange := range ingrediantRanges {
		if ingrediantID >= ingrediantRange.Start && ingrediantID <= ingrediantRange.End {
			return true
		}
	}
	return false
}

func SolvePart1(input string) (int, error) {
	ingrediantRanges, ingrediantIDs, err := parseRangesAndIngredientIDs(input)
	if err != nil {
		return 0, err
	}

	numberOfFreshIngredients := 0
	for _, ingrediantID := range ingrediantIDs {
		if ingrediantIsFresh(ingrediantRanges, ingrediantID) {
			numberOfFreshIngredients++
		}
	}
	return numberOfFreshIngredients, nil
}

func SolvePart2(input string) (int, error) {
	return 0, nil
}
