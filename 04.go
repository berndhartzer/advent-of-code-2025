package aoc

import (
	"fmt"
)

func dayFour(input [][]byte) (int, [][]byte) {
	total := 0

	updated := make([][]byte, len(input))
	for i, line := range input {
		updated[i] = make([]byte, len(line))
	}

	for y, line := range input {
		for x, point := range line {
			if point == '.' {
				updated[y][x] = '.'
				continue
			}

			adjacent := 0

			if y > 0 {
				if x > 0 {
					if input[y-1][x-1] == '@' {
						adjacent++
					}
				}

				if input[y-1][x] == '@' {
					adjacent++
				}

				if x < len(line)-1 {
					if input[y-1][x+1] == '@' {
						adjacent++
					}
				}
			}

			if x < len(line)-1 {
				if input[y][x+1] == '@' {
					adjacent++
				}

				if y < len(input)-1 {
					if input[y+1][x+1] == '@' {
						adjacent++
					}
				}
			}

			if y < len(input)-1 {
				if input[y+1][x] == '@' {
					adjacent++
				}

				if x > 0 {
					if input [y+1][x-1] == '@' {
						adjacent++
					}
				}
			}

			if x > 0 {
				if input[y][x-1] == '@' {
					adjacent++
				}
			}

			if adjacent >= 4 {
				updated[y][x] = '@'
				continue
			}

			// Change state
			updated[y][x] = '.'

			total++
		}
	}

	return total, updated
}

func dayFourPartOne(input []string) int {
	inputBytes := make([][]byte, len(input))
	for i, line := range input {
		inputBytes[i] = []byte(line)
	}

	total, _ := dayFour(inputBytes)

	return total
}

func dayFourPartTwo(input []string) int {
	inputBytes := make([][]byte, len(input))
	for i, line := range input {
		inputBytes[i] = []byte(line)
	}

	total := 0
	for {
		n, updated := dayFour(inputBytes)
		total += n
		if n == 0 {
			break
		}

		inputBytes = updated


	}
	return total
}

func dayFourTests() (map[string]stringSliceToIntTestConfig, map[string]stringSliceToIntTestConfig, error) {
	fileInput, err := getInput(4)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get input: %v", err)
	}
	input := fileInput.asStringSlice()

	partOne := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"..@@.@@@@.",
				"@@@.@.@.@@",
				"@@@@@.@.@@",
				"@.@@@@..@.",
				"@@.@@@@.@@",
				".@@@@@@@.@",
				".@.@.@.@@@",
				"@.@@@.@@@@",
				".@@@@@@@@.",
				"@.@.@@@.@.",
			},
			expected: 13,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	partTwo := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"..@@.@@@@.",
				"@@@.@.@.@@",
				"@@@@@.@.@@",
				"@.@@@@..@.",
				"@@.@@@@.@@",
				".@@@@@@@.@",
				".@.@.@.@@@",
				"@.@@@.@@@@",
				".@@@@@@@@.",
				"@.@.@@@.@.",
			},
			expected: 43,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	return partOne, partTwo, nil
}
