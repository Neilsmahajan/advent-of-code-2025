package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/neilsmahajan/advent-of-code-2025/internal/days/day01"
	"github.com/neilsmahajan/advent-of-code-2025/internal/days/day02"
	"github.com/neilsmahajan/advent-of-code-2025/internal/days/day03"
	"github.com/neilsmahajan/advent-of-code-2025/internal/days/day04"
	"github.com/neilsmahajan/advent-of-code-2025/internal/days/day05"
	"github.com/neilsmahajan/advent-of-code-2025/internal/days/day06"
	"github.com/neilsmahajan/advent-of-code-2025/internal/days/day07"
	"github.com/neilsmahajan/advent-of-code-2025/internal/days/day08"
	"github.com/neilsmahajan/advent-of-code-2025/internal/days/day09"
	"github.com/neilsmahajan/advent-of-code-2025/internal/days/day10"
	"github.com/neilsmahajan/advent-of-code-2025/internal/days/day11"
	"github.com/neilsmahajan/advent-of-code-2025/internal/days/day12"
)

type Solution struct {
	Part1 func(string) (int, error)
	Part2 func(string) (int, error)
}

var solutions = map[int]Solution{
	1:  {day01.SolvePart1, day01.SolvePart2},
	2:  {day02.SolvePart1, day02.SolvePart2},
	3:  {day03.SolvePart1, day03.SolvePart2},
	4:  {day04.SolvePart1, day04.SolvePart2},
	5:  {day05.SolvePart1, day05.SolvePart2},
	6:  {day06.SolvePart1, day06.SolvePart2},
	7:  {day07.SolvePart1, day07.SolvePart2},
	8:  {day08.SolvePart1, day08.SolvePart2},
	9:  {day09.SolvePart1, day09.SolvePart2},
	10: {day10.SolvePart1, day10.SolvePart2},
	11: {day11.SolvePart1, day11.SolvePart2},
	12: {day12.SolvePart1, day12.SolvePart2},
}

func main() {
	var day int
	var part int
	var all bool
	var inputArg string

	flag.IntVar(&day, "day", 1, "Day to solve (1-25)")
	flag.IntVar(&part, "part", 1, "Part to solve (1 or 2)")
	flag.BoolVar(&all, "all", false, "Run all implemented solutions")
	flag.StringVar(&inputArg, "input", "", "Optional input file name or path (e.g. example_input.txt). If a bare filename, it is resolved inside that day's directory.")
	flag.Parse()

	if all {
		runAllSolutions()
		return
	}

	if day < 1 || day > 25 {
		fmt.Printf("Error: Day must be between 1 and 25, got %d\n", day)
		os.Exit(1)
	}

	if part < 1 || part > 2 {
		fmt.Printf("Error: Part must be 1 or 2, got %d\n", part)
		os.Exit(1)
	}

	inputPath, err := resolveInputPath(day, inputArg)
	if err != nil {
		fmt.Printf("Error resolving input file: %v\n", err)
		os.Exit(1)
	}

	solution, exists := solutions[day]
	if !exists {
		fmt.Printf("Error: Day %d not implemented yet\n", day)
		os.Exit(1)
	}

	var result int
	if part == 1 {
		result, err = solution.Part1(inputPath)
	} else {
		result, err = solution.Part2(inputPath)
	}

	if err != nil {
		fmt.Printf("Error solving Day %d Part %d: %v\n", day, part, err)
		os.Exit(1)
	}

	fmt.Printf("Day %d Part %d (%s): %d\n", day, part, filepath.Base(inputPath), result)
}

// resolveInputPath determines the path to the input file for a given day.
// Rules:
//   - If inputArg is empty -> internal/days/dayXX/input.txt
//   - If inputArg is an absolute path and exists -> use as-is
//   - If inputArg contains a path separator and exists as given -> use as given
//   - Otherwise treat inputArg as a filename inside the day directory
func resolveInputPath(day int, inputArg string) (string, error) {
	dayDir := filepath.Join("internal", "days", fmt.Sprintf("day%02d", day))
	defaultPath := filepath.Join(dayDir, "input.txt")

	if strings.TrimSpace(inputArg) == "" {
		return defaultPath, nil
	}

	candidate := inputArg
	// Absolute path
	if filepath.IsAbs(candidate) {
		if _, err := os.Stat(candidate); err == nil {
			return candidate, nil
		}
		return "", fmt.Errorf("absolute path not found: %s", candidate)
	}

	// Relative path provided (may include separators)
	if strings.ContainsRune(candidate, os.PathSeparator) {
		if _, err := os.Stat(candidate); err == nil {
			return candidate, nil
		}
	}

	// Treat as filename within day directory
	inDay := filepath.Join(dayDir, candidate)
	if _, err := os.Stat(inDay); err == nil {
		return inDay, nil
	}

	// Finally, fallback: user maybe gave relative path from repo root that includes no separator (rare). Already checked.
	return "", fmt.Errorf("could not locate input file '%s' (looked for %s and %s)", candidate, inDay, defaultPath)
}

func runAllSolutions() {
	fmt.Println("Running all implemented solutions:")
	fmt.Println("===================================")

	for day := 1; day <= 25; day++ {
		solution, exists := solutions[day]
		if !exists {
			continue
		}

		inputPath := filepath.Join("internal", "days", fmt.Sprintf("day%02d", day), "input.txt")

		// Check if input file exists
		if _, err := os.Stat(inputPath); os.IsNotExist(err) {
			fmt.Printf("Day %02d: Input file not found\n", day)
			continue
		}

		// Try Part 1
		if result, err := solution.Part1(inputPath); err == nil {
			fmt.Printf("Day %02d Part 1: %d\n", day, result)
		} else {
			fmt.Printf("Day %02d Part 1: Error - %v\n", day, err)
		}

		// Try Part 2
		if result, err := solution.Part2(inputPath); err == nil {
			fmt.Printf("Day %02d Part 2: %d\n", day, result)
		} else {
			fmt.Printf("Day %02d Part 2: Error - %v\n", day, err)
		}

		fmt.Println()
	}
}
