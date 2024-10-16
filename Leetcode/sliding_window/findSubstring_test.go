package slidingwindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSubstring(t *testing.T) {
	t.Run("aabcbcdbca", func(t *testing.T) {
		want := "dbca"
		got := findSubstring("aabcbcdbca")
		assert.Equal(t, want, got)
	})
}
