//go:build integration
// +build integration

package integration

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	// Setup
	client := getClient(t)
	var err error

	fmt.Println("Testing non-existent stack get (should err)")
	_, err = client.GetStack(-999)
	if err == nil {
		t.Fatal("Expected GetStackById to return an error but got none!")
		return
	}

	fmt.Println("Testing non-existent stack get by name (should err)")
	individualResult, err := client.GetStackByName("THIS_STACK_SHOULD_NOT_EXIST_EVER")
	if err == nil {
		t.Fatal("Expected GetStackByName to return an error but did not get an err!")
		return
	}
	if individualResult.ID != 0 {
		t.Fatal("Expected to get no results, but got a result!")
		return
	}

	fmt.Println("Testing get stack list")
	results, err := client.GetStacks()
	if err != nil {
		t.Fatal("Expected GetStacks to not return an error but got", err.Error())
		return
	}
	if len(results) == 0 {
		t.Fatal("Expected results from stacks in payload ")
		return
	}

	fmt.Println("Testing get stack by name")
	stackResultByName, err := client.GetStackByName(results[0].Name)
	if err != nil {
		t.Fatal("Expected GetStackByName to function properly without an error but got", err.Error())
		return
	}
	if stackResultByName.Name != results[0].Name {
		t.Fatal("Expected stack name to match the one that was requested to be created but got a mismatch", stackResultByName.Name, results[0].Name)
		return
	}

	// get first item off the list and test getters
	fmt.Println("Testing get by stack id")
	_, err = client.GetStack(results[0].ID)
	if err != nil {
		t.Fatal("Expected GetStackByID to not return an error but got", err.Error())
		return
	}

	fmt.Println("Testing get by stack by name")
	resultByName, err := client.GetStackByName(results[0].Name)
	if err != nil {
		t.Fatal("Expected GetStackByName to not return an error but got", err.Error())
		return
	}
	if resultByName.Name == "" {
		t.Fatal("Expected results from stacks in payload ")
		return
	}
}
