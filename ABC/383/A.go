package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	var n int
	fmt.Sscanf(scanner.Text(), "%d", &n)

	water := 0
	prevTime := 0

	for i := 0; i < n; i++ {
		scanner.Scan()
		var currentTime int
		fmt.Sscanf(scanner.Text(), "%d", &currentTime)

		scanner.Scan()
		var volume int
		fmt.Sscanf(scanner.Text(), "%d", &volume)

		timePassed := currentTime - prevTime
		water = max(0, water-timePassed)

		water += volume

		prevTime = currentTime
	}

	fmt.Println(water)
}
