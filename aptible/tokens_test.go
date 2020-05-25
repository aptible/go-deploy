package aptible

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func mockEnv(envVar, value string) func() {
	original := os.Getenv(envVar)
	deferFunc := func() {
		_ = os.Setenv(envVar, original)
	}

	_ = os.Setenv(envVar, value)

	// Caller should defer on this to revert
	return deferFunc
}

func TestGetTokenEnv(t *testing.T) {
	token := "test_token"
	defer mockEnv("APTIBLE_ACCESS_TOKEN", token)()

	t.Run("Use environment variable", func(t *testing.T) {
		ans, err := GetToken()
		if err != nil {
			t.Errorf("Input: %s caused error: %s", token, err)
		}

		if ans != token {
			t.Errorf("got %s, want %s", ans, token)
		}
	})
}

func TestGetToken(t *testing.T) {
	tmpHome, err := ioutil.TempDir(os.TempDir(), "go-deploy")
	if err != nil {
		t.Error("Failed to create the temp directory", err)
	}
	defer os.RemoveAll(tmpHome)

	err = os.Mkdir(filepath.Join(tmpHome, ".aptible"), 0755)
	if err != nil {
		t.Error("Failed to create the temp directory", err)
	}

	defer mockEnv("HOME", tmpHome)()
	defer mockEnv("APTIBLE_AUTH_ROOT_URL", "auth.aptible.com")()

	var tests = []struct {
		authURL, token, json string
		errored              bool
	}{
		{"auth.aptible.com", "token", `{"auth.aptible.com": "token"}`, false},
		{"auth.aptible.com", "", `{"fail.aptible.com": "token"}`, true},
		// Add other test cases if we need in the future
	}

	for _, tc := range tests {
		testName := fmt.Sprintf("%s: %s -- %s", tc.authURL, tc.token, tc.json)
		t.Run(testName, func(t *testing.T) {
			path := filepath.Join(tmpHome, ".aptible", "tokens.json")
			err := ioutil.WriteFile(path, []byte(tc.json), 0644)
			if err != nil {
				t.Error("Failed to create the temp tokens.json", err)
			}
			_ = os.Unsetenv("APTIBLE_ACCESS_TOKEN")
			_ = os.Setenv("APTIBLE_AUTH_ROOT_URL", tc.authURL)
			ans, err := GetToken()
			// Unexpected error encountered
			if (err != nil) != tc.errored {
				t.Errorf("Input: %s caused error: %s", tc.token, err)
			} else if ans != tc.token {
				t.Errorf("got %s, want %s", ans, tc.token)
			}
		})
	}
}
