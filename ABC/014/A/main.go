package main

import "fmt"

func main() {
	var A, B, C int
	fmt.Scan(&A, &B, &C)

	count := 0

	for i := 0; ; i++ {
		// fmt.Println(A, B, C)
		if A%2 != 0 || B%2 != 0 || C%2 != 0 {
			break
		} else if A == B && B == C && C == A {
			count = -1
			break
		} else {
			halfA := A / 2
			halfB := B / 2
			halfC := C / 2

			newA := halfB + halfC
			newB := halfA + halfC
			newC := halfA + halfB

			A, B, C = newA, newB, newC

			count++
		}
	}

	fmt.Println(count)
}
