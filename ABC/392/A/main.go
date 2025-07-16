package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	var numbers []int
	for _, s := range strings.Fields(scanner.Text()) {
		num, _ := strconv.Atoi(s)
		numbers = append(numbers, num)
	}

	sort.Ints(numbers)

	result := "No"
	ans := numbers[2]
	total := numbers[0] * numbers[1]

	if ans == total {
		result = "Yes"
	}

	fmt.Println(result)
}
