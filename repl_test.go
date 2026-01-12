package main

import (
	"testing"
)

func TestCleanInput(t *testing.T){
	cases := []struct {
	input    string
	expected []string
	}{
	{
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	},
	{
		input:    "Charmander Bulbasaur PIKACHU",
		expected: []string{"charmander", "bulbasaur", "pikachu"},
	},
	{
		input:    "  hello  ",
		expected: []string{"hello"},
	},
	{
		input:    "hElLO  WorLd  ",
		expected: []string{"hello", "world"},
	},
	{
		input:    "  hello  WORLD",
		expected: []string{"hello", "world"},
	},
	// add more cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("failed === expected: %v - actual: %v", actual, c.expected)
			t.Fail()
		}
	// Check the length of the actual slice against the expected slice
	// if they don't match, use t.Errorf to print an error message
	// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("failed = index %v - expected: %v - actual: %v =", i, expectedWord, word)
				t.Fail()
			}
		}
	}
	return
}

