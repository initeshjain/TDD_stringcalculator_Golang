import (
    "strconv"
    "strings"
)

func Add(numbers string) int {
    if numbers == "" {
        return 0
    }

    parts := strings.Split(numbers, ",")
    sum := 0
    for _, part := range parts {
        num, _ := strconv.Atoi(part)
        sum += num
    }
    return sum
}
