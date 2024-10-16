package stringcalculator

import (
	"fmt"
	"strconv"
	"strings"
)

// Add takes a string of numbers and returns their sum.
func Add(numbers string) int {
	if numbers == "" {
		return 0
	}

	delimiter, remaining := extractDelimiter(numbers)
	remaining = strings.ReplaceAll(remaining, "\n", delimiter)
	parts := strings.Split(remaining, delimiter)

	sum := 0
	negatives := []int{}
	for _, part := range parts {
		num, _ := strconv.Atoi(part)
		if num < 0 {
			negatives = append(negatives, num)
		}
		if num <= 1000 {
			sum += num
		}
	}

	if len(negatives) > 0 {
		panic(fmt.Sprintf("negatives not allowed: %v", negatives))
	}
	return sum
}

// extractDelimiter handles custom delimiter logic.
func extractDelimiter(input string) (string, string) {
	if strings.HasPrefix(input, "//") {
		parts := strings.SplitN(input, "\n", 2)
		return parts[0][2:], parts[1]
	}
	return ",", input
}
