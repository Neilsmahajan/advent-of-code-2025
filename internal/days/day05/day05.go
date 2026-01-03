package day05

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

func getTotalIngrediantRange(ingrediantRanges []IngrediantRange) IngrediantRange {
	minimumIngrediant := math.MaxInt
	maximumIngrediant := math.MinInt
	for _, ingrediantRange := range ingrediantRanges {
		if ingrediantRange.Start < minimumIngrediant {
			minimumIngrediant = ingrediantRange.Start
		}
		if ingrediantRange.End > maximumIngrediant {
			maximumIngrediant = ingrediantRange.End
		}
	}
	return IngrediantRange{
		Start: minimumIngrediant,
		End:   maximumIngrediant,
	}
}

func SolvePart2(input string) (int, error) {
	ingrediantRanges, _, err := parseRangesAndIngredientIDs(input)
	if err != nil {
		return 0, err
	}

	sort.Slice(ingrediantRanges, func(i, j int) bool {
		if ingrediantRanges[i].Start == ingrediantRanges[j].Start {
			return ingrediantRanges[i].End < ingrediantRanges[j].End
		}
		return ingrediantRanges[i].Start < ingrediantRanges[j].Start
	})

	numberOfFreshIngredients := 0
	currentStart := ingrediantRanges[0].Start
	currentEnd := ingrediantRanges[0].End

	for _, ingrediantRange := range ingrediantRanges {
		if ingrediantRange.Start <= currentEnd+1 {
			if ingrediantRange.End > currentEnd {
				currentEnd = ingrediantRange.End
			}
		} else {
			numberOfFreshIngredients += currentEnd - currentStart + 1
			currentStart = ingrediantRange.Start
			currentEnd = ingrediantRange.End
		}
	}

	numberOfFreshIngredients += currentEnd - currentStart + 1

	return numberOfFreshIngredients, nil
}
