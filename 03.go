package aoc

import (
	"fmt"
	"strconv"
)

func dayThreePartOne(input []string) int {
	total := 0

	for _, line := range input {
		firstIdx := 0
		secondIdx := 1
		highest := 0

		for i := 0; i < len(line)-1; i++ {
			b := line[i]

			// First num isn't higher
			if b < line[firstIdx] {
				continue
			}

			firstIdx = i

			for j := i+1; j < len(line); j++ {
				c := line[j]

				if secondIdx <= i {
					secondIdx = j
				}

				if c < line[secondIdx] {
					continue
				}

				secondIdx = j
			}

			candidate := int(line[firstIdx]-'0') * 10
			candidate += int(line[secondIdx]-'0')

			if candidate > highest {
				highest = candidate
			}
		}

		total += highest
	}

	return total
}

// greedy algorithm
// For i from 1 to k:
//     remaining = k - (i-1) # how many digits we still need (including current)
//     latest_possible_index = n - remaining
//     search from current_position to latest_possible_index
//     find the maximum digit in that range
//     pick the leftmost one if there are ties
//     append it to answer
//     set current_position = that_index + 1
// End for
func largestKDigits(s string, k int) string {
	n := len(s)
	if k > n || k <= 0 {
		return ""
	}

	result := make([]byte, 0, k)

	// where we start searching for the next digit
	start := 0

	for needed := k; needed > 0; needed-- {
		// We need 'needed' digits total, including the one now
		// So the last position we can pick from is n - needed
		end := n - needed // inclusive

		// Find maximum digit in [start, end]
		bestIdx := start
		for i := start; i <= end; i++ {
			if s[i] > s[bestIdx] {
				bestIdx = i
			}
		}

		result = append(result, s[bestIdx])
		start = bestIdx + 1 // next digit must come after this one
	}

	return string(result)
}

func dayThreePartTwo(input []string) int {
	total := 0

	for _, line := range input {
		largest := largestKDigits(line, 12)
		n, err := strconv.Atoi(largest)
		if err != nil {
			panic("error converting string to int")
		}
		total += n
	}

	return total
}

func dayThreeTests() (map[string]stringSliceToIntTestConfig, map[string]stringSliceToIntTestConfig, error) {
	fileInput, err := getInput(3)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get input: %v", err)
	}
	input := fileInput.asStringSlice()

	partOne := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"987654321111111",
				"811111111111119",
				"234234234234278",
				"818181911112111",
			},
			expected: 357,
		},
		"2": {
			input: []string{
				"1239800009211",
			},
			expected: 99,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	partTwo := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"987654321111111",
				"811111111111119",
				"234234234234278",
				"818181911112111",
			},
			expected: 3121910778619,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	return partOne, partTwo, nil
}
