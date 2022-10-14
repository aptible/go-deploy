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
	individualResult, err := client.UNSUPPORTED_GetStackByName("THIS_STACK_SHOULD_NOT_EXIST_EVER")
	if err == nil {
		t.Fatal("Expected GetStackByName to return an error but did not get an err!")
		return
	}
	if individualResult.ID != 0 {
		t.Fatal("Expected to get no results, but got a result!")
		return
	}

	fmt.Println("Testing get stack list")
	results, err := client.UNSUPPORTED_GetStacks()
	if err != nil {
		t.Fatal("Expected GetStacks to not return an error but got", err.Error())
		return
	}
	if len(results) == 0 {
		t.Fatal("Expected results from stacks in payload ")
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
	resultByName, err := client.UNSUPPORTED_GetStackByName(results[0].Name)
	if err != nil {
		t.Fatal("Expected GetStackByName to not return an error but got", err.Error())
		return
	}
	if resultByName.Name == "" {
		t.Fatal("Expected results from stacks in payload ")
		return
	}
}
