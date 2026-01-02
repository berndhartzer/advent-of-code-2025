package aoc

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type stringNum struct {
	value []byte
}

func newStringNum(s string) stringNum {
	b := []byte(s)
	return stringNum{
		value: b,
	}
}

func (s *stringNum) increment() {
	n := len(s.value)

	// Start from the last digit
	i := n - 1

	// Increment from right to left, handling carry
	for i >= 0 {
		if s.value[i] < '9' {
			// just increment this digit
			s.value[i]++
			return
		}
		// set to 0 and carry over
		s.value[i] = '0'
		i--
	}

	// If we're here, we had all 9's (e.g. "999" -> "1000")
	zeros := bytes.Repeat([]byte{'0'}, n)
	s.value = append([]byte{'1'}, zeros...)

}

func (s *stringNum) bytes() []byte {
	return s.value
}

func (s *stringNum) eq(v stringNum) bool {
	return bytes.Equal(s.value, v.bytes())
}

func (s *stringNum) toInt() int {
	n, err := strconv.Atoi(string(s.value))
	if err != nil {
		panic("error converting string to int")
	}

	return n
}

func dayTwo(input []string, isInvalidFn func([]byte) bool) int {
	total := 0

	for _, ids := range input {
		split := strings.Split(ids, "-")
		from, to := newStringNum(split[0]), newStringNum(split[1])

		to.increment()
		for !from.eq(to) {
			id := from.bytes()

			if isInvalidFn(id) {
				total += from.toInt()
			}

			from.increment()
		}
	}

	return total
}

func dayTwoPartOne(input []string) int {
	isInvalidFn := func(id []byte) bool {
		l := len(id)
		if l%2 != 0 {
			return false
		}

		first, second := id[:l/2], id[l/2:]
		if !bytes.Equal(first, second) {
			return false
		}

		return true
	}

	return dayTwo(input, isInvalidFn)
}

// Compute the KMP prefix table (also called failure function or LPS array)
func computeLPS(b []byte) []int {
	n := len(b)
	lps := make([]int, n)
	lps[0] = 0
	length := 0 // length of previous longest prefix-suffix
	i := 1

	for i < n {
		if b[i] == b[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				// Fall back — do NOT increment i
				length = lps[length-1]
			} else {
				// No match and can't fall back
				lps[i] = 0
				i++
			}
		}
	}

	return lps
}

func dayTwoPartTwo(input []string) int {
	isInvalidFn := func(id []byte) bool {
		lps := computeLPS(id)

		n := len(id)
		last := lps[n-1]
		period := n - last

		if last > 0 && n%period == 0 {
			return true
		}

		return false
	}

	return dayTwo(input, isInvalidFn)
}

func dayTwoTests() (map[string]stringSliceToIntTestConfig, map[string]stringSliceToIntTestConfig, error) {
	fileInput, err := getInput(2)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get input: %v", err)
	}
	input := fileInput.asCommaSeparatedStrings()

	partOne := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"11-22",
				"95-115",
				"998-1012",
				"1188511880-1188511890",
				"222220-222224",
				"1698522-1698528",
				"446443-446449",
				"38593856-38593862",
				"565653-565659",
				"824824821-824824827",
				"2121212118-2121212124",
			},
			expected: 1227775554,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	partTwo := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"11-22",
				"95-115",
				"998-1012",
				"1188511880-1188511890",
				"222220-222224",
				"1698522-1698528",
				"446443-446449",
				"38593856-38593862",
				"565653-565659",
				"824824821-824824827",
				"2121212118-2121212124",
			},
			expected: 4174379265,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	return partOne, partTwo, nil
}
