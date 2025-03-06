package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a[i])
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})

	var total int = a[0]

	for i := 1; i < N; i++ {
		if i%2 == 0 {
			total += a[i]
		} else {
			total -= a[i]
		}
	}

	fmt.Println(total)
}
