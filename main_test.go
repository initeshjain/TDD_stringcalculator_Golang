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

func TestAdd_CustomDelimiter_ReturnsSum(t *testing.T) {
    result := Add("//;\n1;2")
    if result != 3 {
        t.Errorf("Expected 3, got %d", result)
    }
}

func TestAdd_NegativeNumbers_ThrowsException(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("Expected panic for negative numbers")
        }
    }()
    Add("1,-2")
}

func TestAdd_NumbersGreaterThan1000_Ignored(t *testing.T) {
    result := Add("2,1001")
    if result != 2 {
        t.Errorf("Expected 2, got %d", result)
    }
}

func TestAdd_MultiCharacterDelimiter_ReturnsSum(t *testing.T) {
	result := Add("//[***]\n1***2***3")
	if result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}
}

func TestAdd_MultipleDelimiters_ReturnsSum(t *testing.T) {
	result := Add("//[*][%]\n1*2%3")
	if result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}
}
