package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "hello joshi",
			expected: []string{
				"hello",
				"joshi",
			},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("The Lenghts are not equal: %v %v", len(actual), len(cs.expected))
			continue
		}

		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Errorf("The words are not equal: %v %v", actualWord, expectedWord)
			}
		}

	}

}
