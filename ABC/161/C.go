// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var N, K int
// 	fmt.Scan(&N, &K)

// 	ans := []int{}

// 	if K == 1 {
// 		ans = append(ans, 0)
// 	} else {
// 		for i := 0; i < 2; i++ {
// 			subtractN := N - K

// 			if subtractN < 0 {
// 				N = -subtractN
// 			} else {
// 				N = subtractN
// 			}

// 			ans = append(ans, N)
// 		}
// 	}

// 	fmt.Println(min(ans))
// }

// func min(arr []int) int {
// 	minVal := arr[0]
// 	for _, v := range arr {
// 		if v < minVal {
// 			minVal = v
// 		}
// 	}
// 	return minVal
// }

package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	n := N % K

	if K == 1 {
		fmt.Println(0)
		return
	}

	if n < abs(n, K) {
		fmt.Println(n)
	} else {
		fmt.Println(abs(n, K))
	}

}

func abs(a, b int) int {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}
