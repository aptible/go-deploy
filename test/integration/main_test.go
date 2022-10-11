// This file does not have a build scope so that if any tests are missing the
// scope the pre-checks will still run

package integration

import (
	"fmt"
	"os"
	"regexp"
	"testing"
)

func TestMain(m *testing.M) {
	testInt := os.Getenv("TEST_INTEGRATION")
	if testInt != "true" {
		fmt.Printf("TEST_INTEGRATION is %s must be \"true\" to run integration tests: skipping\n", testInt)
		return
	}

	if err := integrationPreCheck(); err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
		return
	}

	os.Exit(m.Run())
}

func integrationPreCheck() error {
	apiURL := os.Getenv("APTIBLE_API_ROOT_URL")
	if matched, _ := regexp.Match("aptible-sandbox\\.com$", []byte(apiURL)); !matched {
		return fmt.Errorf("APTIBLE_API_ROOT_URL must be a URL for a sandbox stack, got %s\n", apiURL)
	}
	authURL := os.Getenv("APTIBLE_AUTH_ROOT_URL")
	if matched, _ := regexp.Match("aptible-sandbox\\.com$", []byte(authURL)); !matched {
		return fmt.Errorf("APTIBLE_AUTH_ROOT_URL must be a URL for a sandbox stack, got %s\n", authURL)
	}
	token := os.Getenv("APTIBLE_ACCESS_TOKEN")
	if token == "" {
		return fmt.Errorf("APTIBLE_ACCESS_TOKEN must be set for integration tests\n")
	}

	return nil
}
