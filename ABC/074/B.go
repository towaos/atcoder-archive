package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	x := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&x[i])
	}

	move := make([][2]int, N)

	for i := 0; i < N; i++ {
		move[i][0] = 2 * x[i]
		move[i][1] = 2 * (K - x[i])
	}

	var total = make([]int, N)

	for i := 0; i < N; i++ {
		total[i] = min(move[i][0], move[i][1])
	}

	// fmt.Println(move, total)

	fmt.Println(sum(total))

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func sum(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	return total
}
