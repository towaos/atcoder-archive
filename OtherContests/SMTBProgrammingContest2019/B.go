package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	minPrice := int(math.Floor(float64(n) / 1.08))

	var result string = ":("

	for i := minPrice; i <= minPrice+1; i++ {
		if int(float64(i)*1.08) == n {
			result = fmt.Sprintf("%d", i)
			break
		}
	}

	fmt.Println(result)
}
