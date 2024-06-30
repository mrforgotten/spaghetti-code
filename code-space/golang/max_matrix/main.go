package main

import (
	"fmt"
)

func getMaxValue(val ...int32) int32 {
	max := val[0]
	for _, v := range val {
		if v > max {
			max = v
		}
	}
	return max
}

func flippingMatrix(matrix [][]int32) int32 {
	// Write your code here
	var matrixLen int = len(matrix)

	var result int32

	// loop row
	for i := 0; i < matrixLen/2; i++ {
		var maxVal int32
		// loop col
		for j := 0; j < matrixLen/2; j++ {
			//loop corner
			maxVal = getMaxValue(matrix[i][j], matrix[i][matrixLen-j-1], matrix[matrixLen-i-1][j], matrix[matrixLen-i-1][matrixLen-j-1])
			result += maxVal
		}
	}
	return result
}

/**
112 42 83 119
56 125 56 49
15 78 101 43
62 98 114 108

expect output 414
*/

/**
107 54 128 15
12 75 110 138
100 96 34 85
75 15 28 112

expect output 488
*/

func main() {
	matrix := [][]int32{{112, 42, 83, 119}, {56, 125, 56, 49}, {15, 78, 101, 43}, {62, 98, 114, 108}} // expect 414
	// matrix := [][]int32{{107, 54, 128, 15}, {12, 75, 110, 138}, {100, 96, 34, 85}, {75, 15, 28, 112}} // expect 488
	// matrix := [][]int32{{1, 2}, {3, 4}} // exptected 4

	printMatrix := func(matrix [][]int32) {
		matrixLen := len(matrix)
		fmt.Println("-----------")
		for i := 0; i < matrixLen; i++ {
			for j := 0; j < matrixLen; j++ {
				fmt.Print(matrix[i][j], " ")
			}
			fmt.Println()
		}
		fmt.Println("-----------")
	}

	printMatrix(matrix)
	// fmt.Println(matrix)
	result := flippingMatrix(matrix)
	println(result)
}
