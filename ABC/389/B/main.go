package main

import "fmt"

func main() {
	var X int
	fmt.Scan(&X)

	count := 0
	// result := 1

	// for X > result {
	// 	count++
	// 	result *= count
	// }

	// This is more optional in terms of processing.
	for i := 2; X%i == 0; i++ {
		X /= i
		count = i
	}

	fmt.Println(count)
}
