package stringcalculator

import "testing"

func TestAdd_EmptyString_ReturnsZero(t *testing.T) {
    result := Add("")
    if result != 0 {
        t.Errorf("Expected 0, got %d", result)
    }
}
