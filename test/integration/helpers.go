//go:build integration
// +build integration

package integration

import (
	"github.com/aptible/go-deploy/aptible"
	"os"
	"strconv"
	"testing"
)

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
