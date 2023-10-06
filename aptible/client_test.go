package aptible

import (
	"os"
	"testing"

	deploy "github.com/aptible/go-deploy/client"
)

func TestGetHost(t *testing.T) {
	currentHost := os.Getenv("APTIBLE_API_ROOT_URL")
	var tests = []struct {
		host     string
		expected string
		errored  bool
	}{
		{"https://api.aptible.com", "api.aptible.com", false},
		{"http://api.aptible.com", "api.aptible.com", false},
		{"api.aptible.com", "api.aptible.com", false},
		{"api.aptible.in", "api.aptible.in", false},
		{"api.aptible", "api.aptible", false},
		{"www.api.aptible.c", "www.api.aptible.c", false},
		{"www.aptible_api.com", "", true},
		// Add other test cases if we need in the future
	}

	for _, tc := range tests {
		testName := tc.host
		t.Run(testName, func(t *testing.T) {

			_ = os.Setenv("APTIBLE_API_ROOT_URL", tc.host)

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
	_ = os.Setenv("APTIBLE_API_ROOT_URL", currentHost)
}

func TestGetHostWhenNotSet(t *testing.T) {
	currentHost := os.Getenv("APTIBLE_API_ROOT_URL")
	_ = os.Unsetenv("APTIBLE_API_ROOT_URL")

	t.Run("Environment variable not set", func(t *testing.T) {
		host, err := GetHost()
		if err != nil {
			t.Errorf("Error: %s", err)
		}

		if host != deploy.DefaultHost {
			t.Errorf("got %s, want %s", host, deploy.DefaultHost)
		}
	})

	_ = os.Setenv("APTIBLE_API_ROOT_URL", currentHost)
}
