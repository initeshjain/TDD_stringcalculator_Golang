package stringcalculator

import (
	"fmt"
	"strconv"
	"strings"
	"regexp"
)

// Helper function to convert string to integer with error handling.
func toInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("Invalid number: %s", str))
	}
	return num
}

// Add takes a string of numbers and returns their sum.
func Add(numbers string) int {
	if numbers == "" {
		return 0
	}

	// Extract the delimiter and remaining part of the input
	delimiter, remaining := extractDelimiter(numbers)

	// Handle both custom delimiters and newlines by adding \n to the delimiter set
	delimiterRegex := fmt.Sprintf("%s|\n", delimiter)
	re := regexp.MustCompile(delimiterRegex)

	parts := re.Split(remaining, -1)
	sum := 0
	negatives := []int{}

	for _, part := range parts {
		if part == "" {
			continue
		}

		num := toInt(part)
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

// extractDelimiter supports multi-character and multiple delimiters.
func extractDelimiter(input string) (string, string) {
	if strings.HasPrefix(input, "//[") {
		re := regexp.MustCompile(`\[(.*?)\]`)
		matches := re.FindAllStringSubmatch(input, -1)
		delimiters := []string{}
		for _, match := range matches {
			delimiters = append(delimiters, regexp.QuoteMeta(match[1]))
		}
		parts := strings.SplitN(input, "\n", 2)
		return strings.Join(delimiters, "|"), parts[1]
	} else if strings.HasPrefix(input, "//") {
		parts := strings.SplitN(input, "\n", 2)
		return regexp.QuoteMeta(parts[0][2:]), parts[1]
	}
	return ",", input
}