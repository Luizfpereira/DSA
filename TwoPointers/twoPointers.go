package twopointers

func isPairSum(arr []int, x int) bool {
	i := 0
	j := len(arr)-1

	for i < j {
		if arr[i] + arr[j] == x {
			return true
		}

		if arr[i] + arr[j] > x {
			j--
		} else {
			i++
		}
	}
	return false
}