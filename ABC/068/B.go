package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	var ans int
	for i := 0; ; i++ {
		power := 1 << i
		if N < power {
			ans = 1 << (i - 1)
			break
		}
	}

	fmt.Println(ans)
}
