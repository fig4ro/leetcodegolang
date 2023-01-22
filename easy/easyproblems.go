package easy

import (
	"fmt"
)

func EasyDifficulty() {
	// Variable to store problem number
	var num int
	
	// Reads user input
	fmt.Print("twoSum(1): returns two numbers, in arrys [15, 2, 11, 7] that sum is equal to target(9)\nrunningSum(2): Descriptione\nmaximumWealth(3): Description\nWhich problem do you wanna solve? ")
	fmt.Scanf("%d", &num)
	
	// Runs problem selected
	switch num {
		case 1:
			result := twoSum([]int{15, 2, 11, 7}, 9)
			fmt.Printf("Output: %d\n", result)
		case 2:
			result := runningSum([]int{1, 2, 3, 4})
			fmt.Printf("Output: %d\n", result)
		case 3:
			i := maximumWealth(
				[][]int{
					{1,5},{7,3},{3,5},
				},
			)
			fmt.Println(i)
		case 4:
			fizzBuzz(5)
		default:
			fmt.Println("Not implemented.")
	}
}
