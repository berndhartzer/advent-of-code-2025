package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

func daySixPartOne(input []string) int {
	operators := strings.Fields(input[len(input)-1])
	operatorFns := make([]func(int, int) int, len(operators))
	for col, op := range operators {
		switch op {
		case "*":
			operatorFns[col] = func(a, b int) int {
				if a == 0 {
					a = 1
				}
				if b == 0 {
					b = 1
				}
				return a * b
			}
		case "+":
			operatorFns[col] = func(a, b int) int {
				return a + b
			}
		}
	}

	totals := make([]int, len(operators))

	for i, line := range input {
		if i == len(input)-1 {
			break
		}

		nums := strings.Fields(line)
		for col, numStr := range nums {
			n, err := strconv.Atoi(numStr)
			if err != nil {
				panic("failed to convert string to int")
			}

			totals[col] = operatorFns[col](totals[col], n)
		}
	}

	total := 0
	for _, n := range totals {
		total += n
	}

	return total
}

func daySixPartTwo(input []string) int {
	operators := strings.Fields(input[len(input)-1])
	opIdx := len(operators)-1

	newNum := ""
	total := 0
	colTotal := 0

	for col := len(input[0])-1; col >= 0; col-- {
		for row := 0; row < len(input)-1; row++ {
			if input[row][col] == ' ' {
				continue
			}
			newNum += string(input[row][col])
		}

		if newNum == "" {
			total += colTotal
			colTotal = 0
			opIdx--
			continue
		}

		n, err := strconv.Atoi(newNum)
		if err != nil {
			panic("failed to convert string to int")
		}

		switch operators[opIdx] {
		case "+":
			colTotal += n
		case "*":
			if colTotal == 0 {
				colTotal = 1
			}
			colTotal *= n
		}

		newNum = ""
	}
	total += colTotal

	return total
}

func daySixTests() (map[string]stringSliceToIntTestConfig, map[string]stringSliceToIntTestConfig, error) {
	fileInput, err := getInput(6)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get input: %v", err)
	}
	input := fileInput.asStringSlice()

	partOne := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"123 328  51 64 ",
				" 45 64  387 23 ",
				"  6 98  215 314",
				"*   +   *   +  ",
			},
			expected: 4277556,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	partTwo := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"123 328  51 64 ",
				" 45 64  387 23 ",
				"  6 98  215 314",
				"*   +   *   +  ",
			},
			expected: 3263827,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	return partOne, partTwo, nil
}
