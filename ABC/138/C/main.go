package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	v := make([]float64, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&v[i])
	}

	sort.Float64s(v)
	cur := v[0]
	for i := 1; i < N; i++ {
		cur = (cur + float64(v[i])) / 2.0
	}
	ans := cur
	fmt.Println(ans)
}
