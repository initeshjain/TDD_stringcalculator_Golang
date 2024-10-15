package stringcalculator

import "testing"

func TestAdd_EmptyString_ReturnsZero(t *testing.T) {
    result := Add("")
    if result != 0 {
        t.Errorf("Expected 0, got %d", result)
    }
}

func TestAdd_OneNumber_ReturnsSameNumber(t *testing.T) {
    result := Add("1")
    if result != 1 {
        t.Errorf("Expected 1, got %d", result)
    }
}

func TestAdd_TwoNumbers_ReturnsSum(t *testing.T) {
    result := Add("1,2")
    if result != 3 {
        t.Errorf("Expected 3, got %d", result)
    }
}
