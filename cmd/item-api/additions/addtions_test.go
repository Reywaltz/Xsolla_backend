package additions_test

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/Reywaltz/backend_xsolla/cmd/item-api/additions"
	"github.com/stretchr/testify/assert"
)

func TestParamValidation(t *testing.T) {
	type testCase struct {
		Name        string
		In          []string
		ExpectedErr bool
	}

	const (
		URL = `/items?limit=%s&offset=%s&type=%s&min=%s&max=%s`
	)

	stringToIfs := func(in []string) []interface{} {
		out := make([]interface{}, len(in))
		for i, v := range in {
			out[i] = v
		}

		return out
	}

	limitOffsetcases := []testCase{
		{Name: "Literal limit", In: []string{"qwe", "", "", "", ""}, ExpectedErr: true},
		{Name: "Negative limit", In: []string{"-2", "", "", "", ""}, ExpectedErr: true},
		{Name: "OK limit", In: []string{"5", "", "", "", ""}, ExpectedErr: false},
		{Name: "Literal offset", In: []string{"", "asd", "", "", ""}, ExpectedErr: true},
		{Name: "Negative offset", In: []string{"", "-2", "", "", ""}, ExpectedErr: true},
		{Name: "OK offset", In: []string{"", "1", "", "", ""}, ExpectedErr: false},
		{Name: "OK type", In: []string{"", "", "1", "", ""}, ExpectedErr: false},
		{Name: "Negative min cost", In: []string{"", "", "", "-1", ""}, ExpectedErr: true},
		{Name: "Negative max cost", In: []string{"", "", "", "", "-1"}, ExpectedErr: true},
		{Name: "OK min cost", In: []string{"", "", "", "1", ""}, ExpectedErr: false},
		{Name: "OK max cost", In: []string{"", "", "", "", "1"}, ExpectedErr: false},
	}
	for _, tc := range limitOffsetcases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			testURL := fmt.Sprintf(URL, stringToIfs(tc.In)...)
			r := httptest.NewRequest("GET", testURL, nil)
			var req additions.Query

			err := req.HandleURLQueries(r)
			assert.Equal(t, tc.ExpectedErr, err != nil)
		})
	}
}

func TestHandleMinCost(t *testing.T) {
	type testCase struct {
		Name          string
		In            []string
		ExpectedErr   bool
		ExpectedValue string
	}

	const (
		URL = `/items?&min=%s`
	)

	limitOffsetcases := []testCase{
		{Name: "Literal min cost", In: []string{"ax"}, ExpectedErr: true},
		{Name: "Negative min cost", In: []string{"-2"}, ExpectedErr: true},
		{Name: "Default min cost", In: []string{""}, ExpectedErr: false, ExpectedValue: additions.DefaultMinCost},
	}

	stringToIfs := func(in []string) []interface{} {
		out := make([]interface{}, len(in))
		for i, v := range in {
			out[i] = v
		}

		return out
	}
	for _, tc := range limitOffsetcases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			testURL := fmt.Sprintf(URL, stringToIfs(tc.In)...)
			r := httptest.NewRequest("GET", testURL, nil)
			var req additions.Query

			err := req.HandleURLQueries(r)
			assert.Equal(t, tc.ExpectedErr, err != nil)
			if err == nil {
				assert.Equal(t, tc.ExpectedValue, req.MinCost)
			}
		})
	}
}

func TestHandleMaxCost(t *testing.T) {
	type testCase struct {
		Name          string
		In            []string
		ExpectedErr   bool
		ExpectedValue string
	}

	const (
		URL = `/items?&max=%s`
	)

	limitOffsetcases := []testCase{
		{Name: "Literal max cost", In: []string{"ax"}, ExpectedErr: true},
		{Name: "Negative max cost", In: []string{"-2"}, ExpectedErr: true},
		{Name: "Default max cost", In: []string{""}, ExpectedErr: false, ExpectedValue: additions.DefaultMaxCost},
	}

	stringToIfs := func(in []string) []interface{} {
		out := make([]interface{}, len(in))
		for i, v := range in {
			out[i] = v
		}

		return out
	}
	for _, tc := range limitOffsetcases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			testURL := fmt.Sprintf(URL, stringToIfs(tc.In)...)
			r := httptest.NewRequest("GET", testURL, nil)
			var req additions.Query

			err := req.HandleURLQueries(r)
			assert.Equal(t, tc.ExpectedErr, err != nil)
			if err == nil {
				assert.Equal(t, tc.ExpectedValue, req.MaxCost)
			}
		})
	}
}
