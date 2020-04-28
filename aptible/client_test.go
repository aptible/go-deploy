package aptible

import (
	"fmt"
	"os"
	"testing"

	deploy "github.com/reggregory/go-deploy/client"
)

func TestGetHost(t *testing.T) {

	var tests = []struct {
		host     string
		expected string
		errored  bool
	}{
		{"https://api.aptible.com", "api.aptible.com", false},
		{"http://api.aptible.com", "api.aptible.com", false},
		{"api.aptible.com", "api.aptible.com", false},
		{"api.aptible", "", true},
		{"www.api.aptible.c", "", true},
		// Add other test cases if we need in the future
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("%s", tc.host)
		t.Run(testname, func(t *testing.T) {

			os.Setenv("APTIBLE_API_ROOT_URL", tc.host)

			host, err := GetHost()

			if (err != nil) != tc.errored {
				t.Errorf("Input: %s caused error: %s", tc.host, err)
			}
			// Unexpected error encountered
			if host != tc.expected {
				t.Errorf("got %s, want %s", host, tc.expected)
			}
		})
	}
}

func TestGetHostWhenNotSet(t *testing.T) {
	curr_host := os.Getenv("APTIBLE_API_ROOT_URL")
	os.Unsetenv("APTIBLE_API_ROOT_URL")

	t.Run("Environment variable not set", func(t *testing.T) {
		host, err := GetHost()
		if err != nil {
			t.Errorf("Error: %s", err)
		}

		if host != deploy.DefaultHost {
			t.Errorf("got %s, want %s", host, deploy.DefaultHost)
		}
	})

	os.Setenv("APTIBLE_API_ROOT_URL", curr_host)
}
