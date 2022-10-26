//go:build integration
// +build integration

package integration

import (
	"fmt"
	"testing"
)

func TestOrganization(t *testing.T) {
	client := getClient(t)

	resp, err := client.GetOrganization()
	if err != nil {
		t.Fatalf("Error when trying to get organization - %s", err.Error())
		return
	}
	if resp.ID == "" {
		t.Fatal("Error when retrieving organization, id not populated")
		return
	}

	fmt.Println("Organization ID retrieved", resp.ID)
	fmt.Println("Organization Name retrieved", resp.Name)
}
