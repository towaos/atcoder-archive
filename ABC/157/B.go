package main

import "fmt"

func main() {
	var bingoTable = make([][]int, 3)
	for i := 0; i < 3; i++ {
		var bingoNum = make([]int, 3)
		fmt.Scan(&bingoNum[0], &bingoNum[1], &bingoNum[2])
		bingoTable[i] = bingoNum
	}

	var N int
	fmt.Scan(&N)

	var b int
	for i := 0; i < N; i++ {
		fmt.Scan(&b)
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				if b == bingoTable[x][y] {
					bingoTable[x][y] = 0
				}
			}
		}
	}

	var result string = "No"
	for i := 0; i < 3; i++ {
		if bingoTable[i][0] == 0 && bingoTable[i][1] == 0 && bingoTable[i][2] == 0 {
			result = "Yes"
		}
	}
	for i := 0; i < 3; i++ {
		if bingoTable[0][i] == 0 && bingoTable[1][i] == 0 && bingoTable[2][i] == 0 {
			result = "Yes"
		}
	}
	if bingoTable[0][0] == 0 && bingoTable[1][1] == 0 && bingoTable[2][2] == 0 || bingoTable[0][2] == 0 && bingoTable[1][1] == 0 && bingoTable[2][0] == 0 {
		result = "Yes"
	}

	fmt.Println(result)
}
