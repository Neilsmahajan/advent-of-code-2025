# Advent of Code 2024

This repository contains my solutions for [Advent of Code 2024](https://adventofcode.com/2024) written in Go.

## Project Structure

```
advent-of-code-2025/
├── cmd/aoc/main.go          # Main entry point
├── internal/
│   ├── days/                # Solutions for each day
│   │   ├── day01/
│   │   │   ├── day01.go     # Solution implementation
│   │   │   ├── input.txt    # Puzzle input
│   │   │   └── example_input.txt (example / sample input)
│   │   ├── day02/
│   │   │   ├── day02.go
│   │   │   └── input.txt
│   │   └── ...              # Days 3-25
│   └── utils/
│       └── utils.go         # Common utility functions
├── bin/                     # Built binaries
└── go.mod
```

## Usage

### Running Solutions with Makefile (Recommended)

The project includes a Makefile for easier command execution:

```bash
# Run a specific day and part (default input.txt)
make run DAY=1 PART=1

# Run with an alternate input file in the day directory
make run DAY=1 PART=1 INPUT=example_input.txt

# Run with a custom path (relative or absolute)
make run DAY=1 PART=2 INPUT=internal/days/day01/example_input.txt

# Quick shortcuts for common days (uncomment in Makefile if desired)
make day1
make day2

# Run all implemented solutions
make all

# Build the project
make build

# Run tests
make test

# Format code
make fmt

# Clean build artifacts
make clean

# Show all available targets
make help
```

INPUT is optional. If omitted, `input.txt` inside that day's directory is used.

### Running Solutions Directly with Go

You can also run solutions using command line flags directly:

```bash
# Run a specific day and part (default input.txt)
go run ./cmd/aoc/main.go -day=1 -part=1

# Use an alternate input file in the same day directory
go run ./cmd/aoc/main.go -day=1 -part=1 -input=example_input.txt

# Provide a relative or absolute path to any file
go run ./cmd/aoc/main.go -day=1 -part=2 -input=internal/days/day01/example_input.txt

# Run all implemented solutions
go run ./cmd/aoc/main.go -all

# Default behavior (day 1, part 1)
go run ./cmd/aoc/main.go
```

`-input` rules:
- Empty / omitted: uses `internal/days/dayXX/input.txt`
- Absolute path: used directly if it exists
- Path with separator: used as-is if it exists
- Bare filename: resolved relative to the specific day directory

### Building

```bash
# Build the binary (using Makefile)
make build

# Or build directly with Go
go build -o bin/aoc ./cmd/aoc/main.go

# Run the built binary
./bin/aoc -day=1 -part=1
./bin/aoc -day=1 -part=1 -input=example_input.txt
```

### Adding New Day Solutions

1. **Copy your input**: Paste your puzzle input into `internal/days/dayXX/input.txt`
2. **Optional**: Add sample inputs like `example_input.txt` for testing.
3. **Implement the solution**: Edit `internal/days/dayXX/dayXX.go` and implement the `SolvePart1` and `SolvePart2` functions
4. **Add to main.go**: Add the import and entry to the solutions map in `cmd/aoc/main.go`:

```go
package main

import (
    // ... existing imports
    "github.com/neilsmahajan/advent-of-code-2024/internal/days/dayXX"
)

var solutions = map[int]Solution{
    // ... existing solutions
    XX: {dayXX.SolvePart1, dayXX.SolvePart2},
}
```

5. **Test your solution**:

```bash
go run ./cmd/aoc/main.go -day=XX -part=1
go run ./cmd/aoc/main.go -day=XX -part=2
go run ./cmd/aoc/main.go -day=XX -part=1 -input=example_input.txt
```

### Utility Functions

The `internal/utils` package provides common functions:

- `ReadLines(filename)` - Read all lines from a file
- `ReadInts(filename)` - Read integers from a file (one per line)
- `SplitInts(s, delimiter)` - Split string and convert to integers
- `AbsInt(x)` - Absolute value
- `MinInt(a, b)`, `MaxInt(a, b)` - Min/Max functions
- `SumInts(nums)` - Sum of integer slice
- `ParseGrid(lines)` - Parse 2D character grid
- `InBounds(grid, row, col)` - Check grid bounds

Example usage:

```go
package dayXX

import "github.com/neilsmahajan/advent-of-code-2024/internal/utils"

func SolvePart1(input string) (int, error) {
    lines, err := utils.ReadLines(input)
    if err != nil {
        return 0, err
    }

    // Process lines...
    return result, nil
}
```

## Development

### Running Tests

```bash
go test ./...
```

### Project Setup

The project structure is already set up with template files for all 25 days. Each day has a placeholder implementation that you can fill in with your solution.

## Solutions

- [x] Day 1: Secret Entrance
- [x] Day 2: Gift Shop
- [x] Day 3: Lobby
- [x] Day 4: Printing Department
- [x] Day 5: Cafeteria
- [ ] Day 6: Trash Compactor
- [ ] Day 7: Laboratories
- [ ] Day 8: Playground
- [ ] Day 9: Movie Theater
- [ ] Day 10: Factory
- [ ] Day 11: Reactor
- [ ] Day 12: Christmas Tree Farm