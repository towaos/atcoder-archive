package main

import "fmt"

func main() {
	var N, total int
	fmt.Scan(&N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
		if i%2 == 0 {
			total += A[i]
		}
	}
	fmt.Println(total)
}
