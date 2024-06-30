package main

import "testing"

func TestDoFib(t *testing.T) {
	type testStrc struct {
		input    uint
		expected uint
	}

	testTable := []testStrc{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {8, 21}, {14, 377}, {17, 1597}, {18, 2584},
	}
	for _, test := range testTable {
		actual := doFib(test.input)
		t.Logf("Test case input: %v, expected: %v", test.input, test.expected)
		if actual != test.expected {
			t.Errorf("Test failed, input: %v expected: '%v', got: '%v'", test.input, test.expected, actual)
		}
	}
}
