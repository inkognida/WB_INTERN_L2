package main

import (
	"testing"
)

func TestRemoveEmptyLines(t *testing.T) {
	lines := []string{"abc", "", "def", "   ", "ghi"}

	expected := []string{"abc", "def", "ghi"}
	result := removeEmptyLines(lines)

	if len(result) != len(expected) {
		t.Errorf("Expected %d lines, but got %d", len(expected), len(result))
	}

	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			t.Errorf("Expected line '%s', but got '%s'", expected[i], result[i])
		}
	}
}

func TestSortLines(t *testing.T) {
	lines := []string{"abc", "def", "ghi", "123", "456", "789"}

	expected := []string{"123", "456", "789", "abc", "def", "ghi"}

	sortLines(lines)

	if len(lines) != len(expected) {
		t.Errorf("Expected %d lines, but got %d", len(expected), len(lines))
	}

	for i := 0; i < len(lines); i++ {
		if lines[i] != expected[i] {
			t.Errorf("Expected line '%s', but got '%s'", expected[i], lines[i])
		}
	}
}

func TestParseNumericValueWithSuffix(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"123", 123},
		{"1K", 1000},
		{"2.5K", 2500},
		{"1M", 1000000},
		{"0.5M", 500000},
		{"3B", 3000000000},
		{"", 0},
	}

	for _, test := range tests {
		result := parseNumericValueWithSuffix(test.input)

		if result != test.expected {
			t.Errorf("Expected value %f, but got %f", test.expected, result)
		}
	}
}

func TestReverse(t *testing.T) {
	lines := []string{"abc", "def", "ghi"}

	expected := []string{"ghi", "def", "abc"}

	reverse(lines)

	if len(lines) != len(expected) {
		t.Errorf("Expected %d lines, but got %d", len(expected), len(lines))
	}

	for i := 0; i < len(lines); i++ {
		if lines[i] != expected[i] {
			t.Errorf("Expected line '%s', but got '%s'", expected[i], lines[i])
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {
	lines := []string{"abc", "def", "ghi", "abc", "def"}

	expected := []string{"abc", "def", "ghi"}

	result := removeDuplicates(lines)

	if len(result) != len(expected) {
		t.Errorf("Expected %d lines, but got %d", len(expected), len(result))
	}

	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			t.Errorf("Expected line '%s', but got '%s'", expected[i], result[i])
		}
	}
}

func TestIsSorted(t *testing.T) {
	lines := []string{"abc", "def", "ghi"}

	if isSorted(lines) {
		t.Error("Expected lines to be unsorted, but they are sorted")
	}

	sortLines(lines)

	if !isSorted(lines) {
		t.Error("Expected lines to be sorted, but they are unsorted")
	}
}


