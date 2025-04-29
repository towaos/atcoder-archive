package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	joinStr := a + b
	joinInt, _ := strconv.Atoi(joinStr)

	root := math.Sqrt(float64(joinInt))

	var result string = "No"
	if root == float64(int(root)) {
		result = "Yes"
	}

	fmt.Println(result)
}
