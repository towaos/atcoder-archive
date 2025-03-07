package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	sc.Split(bufio.ScanWords)
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func main() {
	_, M := scanInt(), scanInt()

	u := make(map[string]bool)
	duplicateCount := 0

	for i := 0; i < M; i++ {
		a, b := scanInt(), scanInt()

		if a == b {
			duplicateCount++
			continue
		}

		if a > b {
			a, b = b, a
		}

		key := fmt.Sprintf("%d_%d", a, b)

		if u[key] {
			duplicateCount++
		} else {
			u[key] = true
		}
	}

	fmt.Println(duplicateCount)
}
