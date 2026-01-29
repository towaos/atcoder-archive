package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	A := make([]int, N)
	Q := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
		A[i]--
		Q[A[i]] = i
	}
	for i := 0; i < N; i++ {
		fmt.Print(Q[i]+1, " ")
	}
	fmt.Println()
}
