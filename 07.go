package aoc

import (
	"strings"
	"fmt"
)

func daySevenPartOne(input []string) int {
	startIdx := 0
	for i, c := range input[0] {
		if c == 'S' {
			startIdx = i
			break
		}
	}

	current := make([]int, len(input[0]))
	current[startIdx] = 1

	total := 0
	for row := 1; row < len(input); row++ {
		next := make([]int, len(input[0]))

		for col := 0; col < len(input[0]); col++ {
			if current[col] == 0 {
				continue
			}

			switch input[row][col] {
			case '.':
				next[col] = current[col]
			case '^':
				next[col] = 0
				total++
				if col > 0 {
					next[col-1]++
				}
				if col < len(input[0])-1 {
					next[col+1]++
				}
			}
		}

		current = next
	}

	return total
}

func daySevenPartTwo(input []string) int {
	return 0
}

func daySevenTests() (map[string]stringSliceToIntTestConfig, map[string]stringSliceToIntTestConfig, error) {
	fileInput, err := getInput(7)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get input: %v", err)
	}
	input := fileInput.asStringSlice()

	partOne := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				".......S.......",
				"...............",
				".......^.......",
				"...............",
				"......^.^......",
				"...............",
				".....^.^.^.....",
				"...............",
				"....^.^...^....",
				"...............",
				"...^.^...^.^...",
				"...............",
				"..^...^.....^..",
				"...............",
				".^.^.^.^.^...^.",
				"...............",
			},
			expected: 21,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	partTwo := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				".......S.......",
				"...............",
				".......^.......",
				"...............",
				"......^.^......",
				"...............",
				".....^.^.^.....",
				"...............",
				"....^.^...^....",
				"...............",
				"...^.^...^.^...",
				"...............",
				"..^...^.....^..",
				"...............",
				".^.^.^.^.^...^.",
				"...............",
			},
			expected: 40,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	return partOne, partTwo, nil
}
