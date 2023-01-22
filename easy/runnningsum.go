package easy

func runningSum(nums []int) []int {
    sumed := make([]int, len(nums))

    sumed[0] = nums[0]

    for i := 1; i < len(nums); i++ {
        sumed[i] = sumed[i-1] + nums[i]
    }

    return sumed
}

/* Smarter way to solve but loses nums original numbers
func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}
 */