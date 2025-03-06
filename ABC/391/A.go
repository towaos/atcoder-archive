package main

import (
	"fmt"
)

func main() {
	var D string
	fmt.Scan(&D)

	directions := map[string]string{
		"N":  "S",
		"S":  "N",
		"E":  "W",
		"W":  "E",
		"NE": "SW",
		"SE": "NW",
		"NW": "SE",
		"SW": "NE",
	}

	if opposite, exists := directions[D]; exists {
		fmt.Println(opposite)
	} else {
		fmt.Println("Invalid direction")
	}
}
