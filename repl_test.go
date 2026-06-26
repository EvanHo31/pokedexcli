package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "  hello   ",
			expected: []string{"hello"},
		}, {
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		}, {
			input:    "  hellO  world  ",
			expected: []string{"hello", "world"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("failed len test\ninput: '%s'\nexpected: '%v'\nactual  : '%v'\n", c.input, c.expected, actual)
			return
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("failed cleanInput test\ninput: '%s'\nexpected: '%v'\nactual  : '%v'\n", c.input, c.expected, actual)
			}
		}
	}
}
