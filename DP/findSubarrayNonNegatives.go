package dp

import "fmt"

// usind Nested Loop
func subArraySumDp(arr []int, x int) bool {
	m := make(map[int]int)
	currSum := 0

	for i, v := range arr {
		currSum += v
		if currSum == x {
			fmt.Println(currSum)
			return true
		}

		if _, ok := m[currSum-x]; ok {
			return true
		}
		m[currSum] = i
	}
	return false
}
