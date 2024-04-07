package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPairSum(t *testing.T) {
	arr := []int{10, 20, 35, 50, 75, 80}

	want := true

	got := isPairSum(arr, 70)

	assert.Equal(t, want, got)
}
