package homework

import (
	"reflect"
	"testing"
)

func TestSortMapValues(t *testing.T) {
	// Input -> {2: "a", 0: "b", 1: "c"}
	// Output -> ["b", "c", "a"]
	// Input -> {10: "aa", 0: "bb", 500: "cc"}
	// Output -> ["bb", "aa", "cc"]

	tests := []struct {
		input    map[int]string
		expected []string
	}{
		{
			input:    map[int]string{2: "a", 0: "b", 1: "c"},
			expected: []string{"b", "c", "a"},
		},
		{
			input:    map[int]string{10: "aa", 0: "bb", 500: "cc"},
			expected: []string{"bb", "aa", "cc"},
		},
	}

	for _, test := range tests {
		if result := sortMapValues(test.input); !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Result: %v != Expected: %v\n", result, test.expected)
		}

	}

}
