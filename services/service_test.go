package services

import (
	"testing"

	"github.com/stretchr/testify/require"

	"find_pairs/models"
)

func Test_FindPairs(t *testing.T) {
	tests := []struct {
		desc   string
		input  models.Input
		output models.Solutions
	}{
		{"empty input array", models.Input{[]int{}, 6}, models.Solutions(nil)},
		{"target not found", models.Input{[]int{1, 2, 3, 4, 5}, 10}, models.Solutions(nil)},
		{"single pair", models.Input{[]int{1, 2, 3}, 5}, models.Solutions{{1, 2}}},
		{"multiple pairs", models.Input{[]int{1, 2, 3, 4, 5, 6}, 6}, models.Solutions{{1, 3}, {0, 4}}},
		{"array with negative number", models.Input{[]int{-1, 2, 3, 4}, 3}, models.Solutions{{0, 3}}},
	}

	for i, tc := range tests {
		output, err := FindPairs(tc.input)

		require.NoError(t, err, "TEST[%d], Failed.\n%s", i, tc.desc)

		require.Equal(t, tc.output, output, "TEST[%d], Failed.\n%s", i, tc.desc)
	}
}
