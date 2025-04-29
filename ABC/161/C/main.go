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

import (
	"fmt"
)

func minN(N, K int64) int64 {
	rem := N % K
	return min(rem, K-rem)
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	var N, K int64
	fmt.Scan(&N, &K)
	fmt.Println(minN(N, K))
}
