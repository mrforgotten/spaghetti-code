package main

import "testing"

func TestGetMaxValue(t *testing.T) {
	type testStrc struct {
		arr      []int32
		expected int32
	}
	testTable := []testStrc{
		{[]int32{4, 2, 3, 1}, 4},
		{[]int32{112, 42, 83, 119}, 119},
		{[]int32{56, 125, 56, 49}, 125},
		{[]int32{15, 78, 101, 43}, 101},
		{[]int32{62, 98, 114, 108}, 114},
	}
	for _, test := range testTable {
		actual := getMaxValue(test.arr...)
		if actual != test.expected {
			t.Errorf("Test failed, expected: '%v', got: '%v'", test.expected, actual)
		}

	}
}

func TestFlippingMatrix(t *testing.T) {
	type testStrc struct {
		matrix   [][]int32
		expected int32
	}

	testTable := []testStrc{
		{[][]int32{{1, 2}, {3, 4}}, 4},
		{[][]int32{{112, 42, 83, 119}, {56, 125, 56, 49}, {15, 78, 101, 43}, {62, 98, 114, 108}}, 414},
		{[][]int32{{107, 54, 128, 15}, {12, 75, 110, 138}, {100, 96, 34, 85}, {75, 15, 28, 112}}, 488},
	}

	for _, test := range testTable {
		actual := flippingMatrix(test.matrix)
		if actual != test.expected {
			t.Errorf("Test failed, expected: '%v', got: '%v'", test.expected, actual)
		}
	}
}
