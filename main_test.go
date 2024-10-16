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

func TestAdd_MultipleNumbers_ReturnsSum(t *testing.T) {
    result := Add("1,2,3,4")
    if result != 10 {
        t.Errorf("Expected 10, got %d", result)
    }
}

func TestAdd_TwoNumbers_ReturnsSum(t *testing.T) {
    result := Add("1,2")
    if result != 3 {
        t.Errorf("Expected 3, got %d", result)
    }
}

func TestAdd_NewlineDelimiter_ReturnsSum(t *testing.T) {
    result := Add("1\n2,3")
    if result != 6 {
        t.Errorf("Expected 6, got %d", result)
    }
}

