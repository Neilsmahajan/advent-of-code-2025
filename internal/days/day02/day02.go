package day02

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ProductRange struct {
	Start, End int
}

func getProductRanges(inputFile string) ([]ProductRange, error) {
	file, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, err
	}

	rangesSlice := strings.Split(string(file), ",")

	var productRanges []ProductRange
	for _, rangeStr := range rangesSlice {
		ranges := strings.Split(rangeStr, "-")
		if len(ranges) != 2 {
			return nil, fmt.Errorf("invalid ranges: %s", rangeStr)
		}
		start, err2 := strconv.Atoi(ranges[0])
		if err2 != nil {
			return nil, err2
		}
		end, err2 := strconv.Atoi(ranges[1])
		if err2 != nil {
			return nil, err2
		}
		productRanges = append(productRanges, ProductRange{start, end})
	}
	return productRanges, nil
}

func isInvalidNumber(number int) bool {
	numberStr := strconv.Itoa(number)
	digits := len(numberStr)
	if digits%2 != 0 {
		return false
	}

	firstHalf, err := strconv.Atoi(numberStr[0 : digits/2])
	if err != nil {
		return false
	}

	lastHalf, err := strconv.Atoi(numberStr[digits/2:])
	if err != nil {
		return false
	}
	return firstHalf == lastHalf
}

func SolvePart1(input string) (int, error) {
	productRanges, err := getProductRanges(input)
	if err != nil {
		return 0, err
	}

	invalidSum := 0
	for _, productRange := range productRanges {
		for number := productRange.Start; number <= productRange.End; number++ {
			if isInvalidNumber(number) {
				invalidSum += number
			}
		}
	}
	return invalidSum, nil
}

func findAllDivisors(number int) []int {
	if number <= 0 {
		return []int{}
	}

	var divisors []int
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			divisors = append(divisors, i)
			if i*i != number {
				divisors = append(divisors, number/i)
			}
		}
	}
	return divisors
}

func hasRepeatedSequenceWithSize(numberString string, size int) bool {
	// Pattern must appear at least twice
	if len(numberString) < 2*size {
		return false
	}

	substringSet := map[string]bool{
		numberString[0:size]: true,
	}

	for i := size; i < len(numberString); i += size {
		substring := numberString[i : i+size]
		if _, ok := substringSet[substring]; !ok {
			return false
		}
	}
	return true
}

func hasRepeatedSequence(number int) bool {
	numberString := strconv.Itoa(number)
	divisors := findAllDivisors(len(numberString))
	divisors = append(divisors, 1)
	for _, divisor := range divisors {
		if hasRepeatedSequenceWithSize(numberString, divisor) {
			return true
		}
	}
	return false
}

func SolvePart2(input string) (int, error) {
	productRanges, err := getProductRanges(input)
	if err != nil {
		return 0, err
	}

	invalidSum := 0
	for _, productRange := range productRanges {
		for number := productRange.Start; number <= productRange.End; number++ {
			if hasRepeatedSequence(number) {
				invalidSum += number
			}
		}
	}
	return invalidSum, nil
}
