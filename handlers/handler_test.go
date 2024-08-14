package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"gofr.dev/pkg/gofr"
	gofrHTTP "gofr.dev/pkg/gofr/http"
	"gofr.dev/pkg/gofr/http/response"

	"find_pairs/models"
)

func Test_FindPairs_BindError(t *testing.T) {
	r, _ := http.NewRequest(http.MethodPost, "/find-pairs", bytes.NewBuffer([]byte(`invalid body`)))
	r.Header.Set("content-type", "application/json")
	req := gofrHTTP.NewRequest(r)

	output, err := FindPairs(&gofr.Context{Context: context.Background(), Request: req})
	require.Error(t, err, "TEST Failed.\n%s", "FindPairs_BindError")

	require.Nil(t, output, "TEST Failed.\n%s", "FindPairs_BindError")
}

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
		body, err := json.Marshal(tc.input)
		require.NoError(t, err, "TEST[%d], Failed.\n%s", i, tc.desc)

		r, _ := http.NewRequest(http.MethodPost, "/find-pairs", bytes.NewBuffer(body))
		r.Header.Set("content-type", "application/json")
		req := gofrHTTP.NewRequest(r)

		output, err := FindPairs(&gofr.Context{Context: context.Background(), Request: req})
		require.NoError(t, err, "TEST[%d], Failed.\n%s", i, tc.desc)

		resp := response.Raw{models.Response{tc.output}}
		require.Equal(t, resp, output, "TEST[%d], Failed.\n%s", i, tc.desc)
	}
}
