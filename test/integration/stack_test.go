//go:build integration
// +build integration

package integration

import (
	"testing"
)

func TestStack(t *testing.T) {
	// Setup
	client := getClient(t)
	var err error

	_, err = client.GetStackById(-999)
	if err == nil {
		t.Fatal("Expected GetStackById to return an error but got none!")
		return
	}

	_, err = client.GetStacksByName("THIS_STACK_SHOULD_NOT_EXIST_EVER")
	if err == nil {
		t.Fatal("Expected GetStacksByName to return an error but got none!")
		return
	}

	results, err := client.GetStacks()
	if err != nil {
		t.Fatal("Expected GetStacks to not return an error but got", err.Error())
		return
	}
	if len(results) == 0 {
		t.Fatal("Expected results from stacks in payload ")
	}

	// get first item off the list and test getters
	_, err = client.GetStackById(results[0].ID)
	if err != nil {
		t.Fatal("Expected GetStacks by ID to not return an error but got", err.Error())
		return
	}

	resultsByName, err := client.GetStacksByName(results[0].Name)
	if err != nil {
		t.Fatal("Expected GetStacks by ID to not return an error but got", err.Error())
		return
	}
	if len(resultsByName) == 0 {
		t.Fatal("Expected results from stacks in payload ")
	}
}
