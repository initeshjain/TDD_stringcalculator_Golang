package stringcalculator

import (
    "strconv"
    "strings"
)

func Add(numbers string) int {
    if numbers == "" {
        return 0
    }

    // Replace newlines with commas for consistent splitting
    numbers = strings.ReplaceAll(numbers, "\n", ",")
    parts := strings.Split(numbers, ",")
    sum := 0
    for _, part := range parts {
        num, _ := strconv.Atoi(part)
        sum += num
    }
    return sum
}

