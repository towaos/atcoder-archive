package utils

import "fmt"

// Standard String and Numeric Acceptance
func standard() {
	var A int
	fmt.Scan(&A)

	var B, C string
	fmt.Scan(&B, &C)
}

// How to receive 2-dimensional arrays
func twoArr() {
	var D = make([][]int, 3)
	for i := 0; i < 3; i++ {
		var E = make([]int, 3)
		fmt.Scan(&E[0], &E[1], &E[2])
		D[i] = E
	}
}

// How to receive arrays
func oneArr() {
	var F = make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Scan(F[i])
	}
}
