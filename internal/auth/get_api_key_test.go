package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	test := []struct {
		input string
		want  string
	}{
		{input: "ApiKey 1234566", want: "ApiKey 1234566"},
		// {input: "1234566", want: "ApiKey 1234566"},
		// {input: "", want: "ApiKey 1234566"},
		// {input: "ApiKey", want: "ApiKey 1234566"},
	}

	for _, ts := range test {
		header := http.Header{}
		header.Set("Authorization", ts.input)
		_, err := GetAPIKey(header)
		if err != nil {
			fmt.Printf("[TEST] - Input: '%s', Test Error: %s\n", ts.input, err)
			t.Fail()
		}
	}
}
