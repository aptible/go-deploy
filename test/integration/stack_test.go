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

	individualResult, err := client.GetStackByName("THIS_STACK_SHOULD_NOT_EXIST_EVER")
	if err == nil {
		t.Fatal("Expected GetStacksByName to NOT return an error but got an err!", err.Error())
		return
	}
	if individualResult.ID != 0 {
		t.Fatal("Expected to get no results, but got a result!")
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

	resultByName, err := client.GetStackByName(results[0].Name)
	if err != nil {
		t.Fatal("Expected GetStacks by ID to not return an error but got", err.Error())
		return
	}
	if resultByName.Name == "" {
		t.Fatal("Expected results from stacks in payload ")
	}
}
