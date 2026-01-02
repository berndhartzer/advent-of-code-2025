package aoc

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func dayFivePartOne(input []string) int {
	ranges := [][2]int{}

	i := 0
	for ; i < len(input); i++ {
		if input[i] == "" {
			i++
			break
		}

		split := strings.Split(input[i], "-")

		newRange := [2]int{}
		for j := 0; j < 2; j++ {
			n, err := strconv.Atoi(split[j])
			if err != nil {
				panic("failed to convert string to int")
			}

			newRange[j] = n
		}

		ranges = append(ranges, newRange)
	}

	slices.SortFunc(ranges, func(a, b [2]int) int {
		return a[0] - b[0]
	})

	merged := [][2]int{ranges[0]}

	// merge ranges
	for j := 1; j < len(ranges); j++ {
		curr := ranges[j]
		prev := merged[len(merged)-1]

		// if current range overlaps or is adjacent to last merged range
		if curr[0] <= prev[1]+1 {
			// if the current range extends the last merged range
			// update it
			if curr[1] > prev[1] {
				// prev[1] = curr[1]
				merged[len(merged)-1][1] = curr[1]
			}
		} else {
			merged = append(merged, curr)
		}
	}

	total := 0

	for ; i < len(input); i++ {
		n, err := strconv.Atoi(input[i])
		if err != nil {
			panic("failed to convert string to int")
		}

		for _, idRange := range merged {
			if n >= idRange[0] && n <= idRange[1] {
				total++
				break
			}
		}
	}

	return total
}

func dayFivePartTwo(input []string) int {
	ranges := [][2]int{}

	i := 0
	for ; i < len(input); i++ {
		if input[i] == "" {
			i++
			break
		}

		split := strings.Split(input[i], "-")

		newRange := [2]int{}
		for j := 0; j < 2; j++ {
			n, err := strconv.Atoi(split[j])
			if err != nil {
				panic("failed to convert string to int")
			}

			newRange[j] = n
		}

		ranges = append(ranges, newRange)
	}

	slices.SortFunc(ranges, func(a, b [2]int) int {
		return a[0] - b[0]
	})

	merged := [][2]int{ranges[0]}

	// merge ranges
	for j := 1; j < len(ranges); j++ {
		curr := ranges[j]
		prev := merged[len(merged)-1]

		// if current range overlaps or is adjacent to last merged range
		if curr[0] <= prev[1]+1 {
			// if the current range extends the last merged range
			// update it
			if curr[1] > prev[1] {
				// prev[1] = curr[1]
				merged[len(merged)-1][1] = curr[1]
			}
		} else {
			merged = append(merged, curr)
		}
	}

	total := 0
	for _, idRange := range merged {
		total += (idRange[1]-idRange[0]) + 1
	}

	return total
}

func dayFiveTests() (map[string]stringSliceToIntTestConfig, map[string]stringSliceToIntTestConfig, error) {
	fileInput, err := getInput(5)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get input: %v", err)
	}
	input := fileInput.asStringSlice()

	partOne := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"3-5",
				"10-14",
				"16-20",
				"12-18",
				"",
				"1",
				"5",
				"8",
				"11",
				"17",
				"32",
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
				"3-5",
				"10-14",
				"16-20",
				"12-18",
				"",
				"1",
				"5",
				"8",
				"11",
				"17",
				"32",
			},
			expected: 14,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	return partOne, partTwo, nil
}
