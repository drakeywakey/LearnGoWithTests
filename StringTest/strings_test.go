package stringTest

import (
	"strings"
	"testing"
)

func TestContains(t *testing.T) {
	result := strings.Contains("abc", "a")
	expected := true

	if result != expected {
		t.Errorf("expected '%v' but got '%v'", expected, result)
	}
}
