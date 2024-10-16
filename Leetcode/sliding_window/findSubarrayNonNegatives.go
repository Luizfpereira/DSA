package slidingwindow

import "fmt"

// usind Nested Loop
func subArraySumNestedLoop(arr []int, x int) bool {
	for i := range arr {
		currSum := 0
		for j := i; j < len(arr); j++ {
			currSum += arr[j]
			if currSum == x {
				fmt.Println(currSum)
				return true
			}
		}
	}
	return false
}

func subArraySumSlidingWindow(arr []int, x int) bool {
	currSum := arr[0]
	start := 0

	for i := 1; i < len(arr); i++ {

		for currSum > x && start < i-1 {
			currSum -= arr[start]
			start++
		}

		if currSum == x {
			fmt.Println(currSum)
			return true
		}

		if currSum < x {
			currSum += arr[i]
		}
	}
	return false
}
