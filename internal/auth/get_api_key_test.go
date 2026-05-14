package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	test := []struct {
		input string
		want  string
	}{
		{input: "ApiKey 1234566", want: "ApiKey 1234566"},
	}

	for _, ts := range test {
		header := http.Header{}
		header.Set("Authorization", ts.input)
		_, err := GetAPIKey(header)
		if err != nil {
			t.Fatalf("[TEST] - Input: '%s', Test Error: %s\n", ts.input, err)
		}
	}
}
