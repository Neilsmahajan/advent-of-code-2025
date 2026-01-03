package day03

import (
	"bufio"
	"os"
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

func SolvePart2(input string) (int, error) {
	return 0, nil
}
