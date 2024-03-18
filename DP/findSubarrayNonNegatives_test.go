package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubArraySumDp(t *testing.T) {
	t.Run("first case", func(t *testing.T) {
		arr := []int{1, 4, 20, 3, 10, 5}
		sum := 33

		exists := subArraySumDp(arr, sum)

		assert.True(t, exists)
	})

	t.Run("second case", func(t *testing.T) {
		arr := []int{1, 4, 0, 0, 3, 10, 5}
		sum := 7

		exists := subArraySumDp(arr, sum)

		assert.True(t, exists)
	})

	t.Run("third case", func(t *testing.T) {
		arr := []int{1, 4}
		sum := 0

		exists := subArraySumDp(arr, sum)

		assert.False(t, exists)
	})

	t.Run("fourth case", func(t *testing.T) {
		arr := []int{15, 2, 4, 8, 5, 10, 23}
		sum := 21

		exists := subArraySumDp(arr, sum)

		assert.True(t, exists)
	})
}
