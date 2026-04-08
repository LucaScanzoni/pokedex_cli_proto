package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input		string
		expected 	[]string
	}{
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  HellO  World  ",
			expected: []string{"hello", "world"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}

func TestCleanInputOutputLength(t *testing.T) {
	cases := []struct {
		input		string
		expected	int
	}{
		{
			input:		"",
			expected:	0,
		},
		{
			input:		" hello ",
			expected:	1,
		},
		{
			input:		"hello world",
			expected:	2,
		},
	}
	for _, c := range cases {
		actual_output := cleanInput(c.input)
		actual_len := len(actual_output)
		if actual_len != c.expected {
			t.Fatalf("expected len: %v, got: %v", c.expected, actual_len)
		}
	}
}