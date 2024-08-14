package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_Main(t *testing.T) {
	const endpoint = "http://localhost:8080/find-pairs"

	go main()
	time.Sleep(100 * time.Millisecond) // giving some time for server to start

	tests := []struct {
		desc       string
		body       []byte
		statusCode int
		output     [][]int
	}{
		{"empty input", nil, http.StatusBadRequest, nil},
		{"success case", []byte(`{"numbers": [1, 2, 3, 4, 5, 6], "target": 6}`),
			http.StatusCreated, [][]int{{1, 3}, {0, 4}}},
	}

	for i, tc := range tests {
		req, _ := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(tc.body))

		req.Header.Set("content-type", "application/json")

		c := http.Client{}

		resp, err := c.Do(req)
		require.NoError(t, err, "TEST[%d], Failed.\n%s", i, tc.desc)

		data := struct {
			Solutions [][]int `json:"solutions"`
		}{}

		b, err := io.ReadAll(resp.Body)
		require.NoError(t, err, "TEST[%d], Failed.\n%s", i, tc.desc)

		err = json.Unmarshal(b, &data)
		require.NoError(t, err, "TEST[%d], Failed.\n%s", i, tc.desc)

		require.Equal(t, tc.output, data.Solutions, "TEST[%d], Failed.\n%s", i, tc.desc)
		require.Equal(t, tc.statusCode, resp.StatusCode, "TEST[%d], Failed.\n%s", i, tc.desc)

		resp.Body.Close()
	}
}
