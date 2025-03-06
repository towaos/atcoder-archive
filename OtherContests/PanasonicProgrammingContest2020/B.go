package main

import "fmt"

func main() {
	var H, W int
	fmt.Scan(&H, &W)

	var result int

	// if H%2 == 0 && W%2 == 0 {
	// 	result = (H * W) / 2
	// } else {
	// 	result = (H*W + 1) / 2
	// }

	if H == 1 || W == 1 {
		result = 1
	} else {
		result = (H*W + 1) / 2
	}

	fmt.Println(result)
}

// (row+column)mod2 = You can only go to an even number of squares.
