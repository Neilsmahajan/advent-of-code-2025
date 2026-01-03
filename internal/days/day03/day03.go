package day03

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parseBanks(inputFile string) ([][]int, error) {
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

	scanner := bufio.NewScanner(file)
	var banks [][]int
	for scanner.Scan() {
		line := scanner.Text()
		bank := make([]int, len(line))
		for i, r := range line {
			bank[i] = int(r - '0')
		}
		banks = append(banks, bank)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return banks, nil
}

func SolvePart1(input string) (int, error) {
	banks, err := parseBanks(input)
	if err != nil {
		return 0, err
	}

	totalJolts := 0
	for _, bank := range banks {
		maxBankJolts := -1
		left := 0
		for right := 1; right < len(bank); right++ {
			currentJolts := bank[left]*10 + bank[right]
			if currentJolts > maxBankJolts {
				maxBankJolts = currentJolts
			}
			if bank[right] > bank[left] {
				left = right
			}
		}
		totalJolts += maxBankJolts
	}
	return totalJolts, nil
}

func sliceToInt(slice []int) (int, error) {
	stringSlice := make([]string, len(slice))
	for index, value := range slice {
		stringSlice[index] = strconv.Itoa(value)
	}
	combinedString := strings.Join(stringSlice, "")
	result, err := strconv.Atoi(combinedString)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func SolvePart2(input string) (int, error) {
	banks, err := parseBanks(input)
	if err != nil {
		return 0, err
	}

	totalJolts := 0
	for _, bank := range banks {
		k := len(bank) - 12
		var stack []int
		for _, jolt := range bank {
			for k > 0 && len(stack) > 0 && jolt > stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
				k--
			}
			stack = append(stack, jolt)
		}
		for k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}
		maxJoltsInBank, err := sliceToInt(stack)
		if err != nil {
			return 0, err
		}
		totalJolts += maxJoltsInBank
	}
	return totalJolts, nil
}
