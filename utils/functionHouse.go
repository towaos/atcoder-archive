package utils

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Sum(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	return total

}
