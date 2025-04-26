package main

import "fmt"

func main() {
	var S1, S2 string
	fmt.Scan(&S1, &S2)

	result := 4

	if S1 == "sick" && S2 == "sick" {
		result = 1
	} else if S1 == "sick" {
		result = 2
	} else if S2 == "sick" {
		result = 3
	}

	fmt.Println(result)

}
