package main

import (
	"fmt"
	"sort"
)

func main() {
	var K, N int
	fmt.Scan(&K, &N)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}
	sort.Ints(A)

	var total int
	total = ((K - A[N-1]) + A[0])
	for i := 1; i < N; i++ {
		ans := A[i] - A[i-1]
		if total < ans {
			total = ans
		}
	}

	result := K - total

	fmt.Println(result)
}
