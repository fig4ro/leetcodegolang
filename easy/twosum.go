package easy

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

/* Another way to implement the function, not fully comprehend yet
func twoSum(nums []int, target int) []int {
    // Space: O(n)
    s := make(map[int]int)

    // Time: O(n)
    for idx, num := range nums {
        // Time: O(1)
        if pos, ok := s[target-num]; ok {
            return []int{pos, idx}
        }
        s[num] = idx
    }
    return []int{}
}
*/
