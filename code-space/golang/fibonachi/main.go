package main

import "fmt"

func doFib(n uint) uint {
	return fib(n)
}

// loop version
func fib(n uint) uint {
	arr := make([]uint, n+1)
	arr[1], arr[2] = 1, 1

	var i uint = 2

	for ; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr[n]
}

// simple resurcive version
// func fib(n uint) uint {
// 	if n == 0 {
// 		return 0
// 	} else if n == 1 || n == 2 {
// 		return 1
// 	} else {
// 		return fib(n-1) + fib(n-2)
// 	}
// }

/**
sequence -> 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144
*/

func main() {
	// var n uint = 10 //expect 55
	// var n uint = 30 //expect 832040
	// var n uint = 40 //expect 102334155
	var n uint = 50 //12586269025
	// var n uint = 100 //expect 3736710778780434371
	ans := doFib(n)

	fmt.Println("fib of n:", n, "is", ans)
}
