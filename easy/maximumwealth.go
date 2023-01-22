package easy

func maximumWealth(accounts [][]int) int {

	wealth := make([]int, len(accounts))

	for i := 0; i < len(accounts); i++ {
		wealth[i] = 0
		for j := 0; j < len(accounts[i]); j++ {
			wealth[i] += accounts[i][j]
		}
	}

	// Could be implemented into the loop above
	max := wealth[0]
	for _, bank := range wealth {
		if bank > max {
			max = bank
		}
	}

	return max
}

/* A better looking
func maximumWealth(accounts [][]int) int {
	max := 0
	for _, a := range accounts {
		current := 0
		for _, v := range a {
			current += v
		}
		if max < current {
			max = current
		}
	}
	return max
}
*/