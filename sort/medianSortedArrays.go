package sort

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr := []int{}

	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			arr = append(arr, nums1[i])
			i++
		} else {
			arr = append(arr, nums2[j])
			j++
		}
	}

	for i < len(nums1) {
		arr = append(arr, nums1[i])
		i++
	}

	for j < len(nums2) {
		arr = append(arr, nums2[j])
		j++
	}

	fmt.Println(arr)

	if len(arr)%2 == 0 {
		n1 := len(arr)/2 - 1
		n2 := len(arr) / 2
		return (float64(arr[n1]) + float64(arr[n2])) / 2
	} else {
		return float64(arr[len(arr)/2])
	}
}
