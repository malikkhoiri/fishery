package fishery

import (
	"strings"
	"testing"
)

func TestClient(t *testing.T) {
	cases := []struct {
		url      string
		desc     string
		expected string
	}{
		{
			url:      "ftp://example.com",
			desc:     "invalid non http/https url",
			expected: "scheme must be use http or https",
		},
		{
			url:      "example.com",
			desc:     "invalid url without scheme",
			expected: "http or https scheme not specified",
		},
		{
			url:  "http://example.com",
			desc: "valid url",
		},
		{
			url:  "https://example.com",
			desc: "valid https url",
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			_, err := NewClient(tc.url, "apiKey")

			if err != nil {
				if tc.expected != "" {
					if !strings.Contains(err.Error(), tc.expected) {
						t.Fatalf("unexpected error: %s", err)
					}
				} else {
					t.Fatalf("unexpected error: %s", err)
				}
			}
		})
	}
}
