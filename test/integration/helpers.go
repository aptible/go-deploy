//go:build integration
// +build integration

package integration

import (
	"github.com/aptible/go-deploy/aptible"
	"os"
	"regexp"
	"strconv"
	"testing"
)

func integrationPreCheck(t *testing.T) {
	testInt := os.Getenv("TEST_INTEGRATION")
	if testInt != "true" {
		t.Skipf("TEST_INTEGRATION is %s must be \"true\" to run integration tests", testInt)
	}

	apiURL := os.Getenv("APTIBLE_API_ROOT_URL")
	if matched, _ := regexp.Match("aptible-sandbox\\.com$", []byte(apiURL)); !matched {
		t.Fatal("APTIBLE_API_ROOT_URL must be a URL for a sandbox stack, got", apiURL)
	}
	authURL := os.Getenv("APTIBLE_AUTH_ROOT_URL")
	if matched, _ := regexp.Match("aptible-sandbox\\.com$", []byte(authURL)); !matched {
		t.Fatal("APTIBLE_AUTH_ROOT_URL must be a URL for a sandbox stack, got", authURL)
	}
}

func getClient(t *testing.T) aptible.Client {
	client, err := aptible.SetUpClient()
	if err != nil {
		t.Fatal("Error setting up client:", err.Error())
	}

	return *client
}

func getEnvID(t *testing.T) int64 {
	idStr := os.Getenv("APTIBLE_ENVIRONMENT_ID")
	if idStr == "" {
		t.Fatal("APTIBLE_ENVIRONMENT_ID is not set")
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		t.Fatal("Error retrieving APTIBLE_ENVIRONMENT_ID:", err.Error())
	}

	return id
}
