package main

import (
	"fmt"
	"leetcode/easy"
)

func main() {
	
	var difficulty string

	fmt.Print("Difficulty level?(e/m/h): ")
	fmt.Scanf("%s", &difficulty)
	switch difficulty {
	case "e":
		fmt.Println("Easy difficulty selected.")
		easy.EasyDifficulty()
	case "m":
		fmt.Println("Not implemented yet.")
	case "h":
		fmt.Println("Not implemented yet.")
	default:
		fmt.Println("Only accepts e(easy), m(medium) or h(hard).")
	}
}
