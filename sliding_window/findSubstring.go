package slidingwindow

import (
	"math"
)

func findSubstring(str string) string {
	if len(str) < 2 {
		return str
	}

	var distCount int

	visited := make([]bool, 256)
	for _, s := range str {
		if !visited[s] {
			visited[s] = true
			distCount++
		}
	}

	count := 0
	start := 0
	startIdx := 0
	curr := make([]int, 256)
	minLen := math.MaxInt
	var windowLen int

	for i, s := range str {
		curr[s]++

		if curr[s] == 1 {
			count++
		}

		if count == distCount {
			for curr[str[start]] > 1 {
				if curr[str[start]] > 1 {
					curr[str[start]]--
					start++
				}
			}

			windowLen = i - start + 1

			if minLen > windowLen {
				minLen = windowLen
				startIdx = start
			}
		}
	}
	return str[startIdx : startIdx+minLen]
}

func lengthOfLongestSubstring(s string) int {
	// Initialize a map to store the last seen index of each character
	lastSeen := make(map[byte]int)

	// Initialize pointers and max length variables
	left, maxLen := 0, 0

	// Iterate over the string
	for right := 0; right < len(s); right++ {
		// If the character is already in the map and its index is greater than or equal to the left pointer,
		// update the left pointer to the next index of the repeating character
		if idx, found := lastSeen[s[right]]; found && idx >= left {
			left = idx + 1
		}

		// Update the last seen index of the current character
		lastSeen[s[right]] = right

		// Update the maximum length if the current window is longer
		if currLen := right - left + 1; currLen > maxLen {
			maxLen = currLen
		}
	}

	return maxLen
}
