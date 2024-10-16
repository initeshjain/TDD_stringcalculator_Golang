package stringcalculator

import (
    "strconv"
    "strings"
		"fmt"
)

func Add(numbers string) int {
    if numbers == "" {
        return 0
    }

    delimiter := ","
    if strings.HasPrefix(numbers, "//") {
        parts := strings.SplitN(numbers, "\n", 2)
        delimiter = parts[0][2:]
        numbers = parts[1]
    }

    numbers = strings.ReplaceAll(numbers, "\n", delimiter)
    parts := strings.Split(numbers, delimiter)

    sum := 0
    negatives := []int{}
    for _, part := range parts {
        num, _ := strconv.Atoi(part)
        if num < 0 {
            negatives = append(negatives, num)
        }
        sum += num
    }

    if len(negatives) > 0 {
        panic(fmt.Sprintf("negatives not allowed: %v", negatives))
    }
    return sum
}
