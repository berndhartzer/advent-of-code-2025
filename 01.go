package aoc

import (
	"fmt"
	"strconv"
)

func dayOnePartOne(input []string) int {
	pos := 50
	zeros := 0

	for _, line := range input {
		n, err := strconv.Atoi(string(line[1:]))
		if err != nil {
			panic("could not convert str to number")
		}

		move := n % 100

		switch line[0] {
		case 'L':
			pos -= move
			if pos < 0 {
				pos = 100 + pos
			}

		case 'R':
			pos += move
			if pos > 99 {
				pos = pos - 100
			}
		}

		if pos == 0 {
			zeros++
		}
	}

	return zeros
}

func dayOnePartTwo(input []string) int {
	pos := 50
	zeros := 0

	for _, line := range input {
		n, err := strconv.Atoi(string(line[1:]))
		if err != nil {
			panic("could not convert str to number")
		}

		move := n % 100
		zeros += (n / 100)

		startedOnZero := false
		if pos == 0 {
			startedOnZero = true
		}

		hitZero := false

		switch line[0] {
		case 'L':
			pos -= move
			if pos < 0 {
				hitZero = true
				pos = 100 + pos
			}

		case 'R':
			pos += move
			if pos > 99 {
				hitZero = true
				pos = pos - 100
			}
		}

		if startedOnZero {
			continue
		}

		if pos == 0 {
			hitZero = true
		}

		if hitZero {
			zeros++
		}
	}

	return zeros
}

func dayOneTests() (map[string]stringSliceToIntTestConfig, map[string]stringSliceToIntTestConfig, error) {
	fileInput, err := getInput(1)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get input: %v", err)
	}
	input := fileInput.asStringSlice()

	partOne := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"L68",
				"L30",
				"R48",
				"L5",
				"R60",
				"L55",
				"L1",
				"L99",
				"R14",
				"L82",
			},
			expected: 3,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	partTwo := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"L68",
				"L30",
				"R48",
				"L5",
				"R60",
				"L55",
				"L1",
				"L99",
				"R14",
				"L82",
			},
			expected: 6,
		},
		"2": {
			input: []string{
				"L50",
				"L50",
			},
			expected: 6,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	return partOne, partTwo, nil
}
