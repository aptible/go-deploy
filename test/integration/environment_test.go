//go:build integration
// +build integration

package integration

import (
	"reflect"
	"testing"

	"github.com/aptible/go-deploy/aptible"
)

func TestEnvironments(t *testing.T) {
	// Setup
	orgID := getOrgID(t)
	stackID := getStackID(t)
	client := getClient(t)
	var environment aptible.Environment
	var err error

	// Test
	testNotFound := func(id int64) {
		environment, err = client.GetEnvironment(id)
		if err != nil {
			t.Error("Expected GetEnvironment to not return an error but got", err.Error())
		}
	}

	testNotFound(0)

	// create it

	environment, err = client.CreateEnvironment(orgID, aptible.EnvironmentCreateAttrs{
		Handle:  "123-testing-handle",
		Type:    "development",
		StackID: stackID,
	})
	if err != nil {
		t.Fatal("Expected CreateEnvironment to not return an error but got", err.Error())
		return
	}

	// get it

	envID := environment.ID
	prevEnvironment := environment
	environment, err = client.GetEnvironment(*envID)
	if err != nil {
		t.Error("Expected GetEnvironment to not return an error but got", err.Error())
	}
	if !reflect.DeepEqual(environment, prevEnvironment) {
		t.Errorf("Expected the environment to be the same as when it was created: expected %v, got %v", prevEnvironment, environment)
	}

	// update it
	updateEnvironmentAttributes := client.EnvironmentUpdates{
		Handle: "234-testing-handle-update",
	}
	err = client.UpdateEnvironment(environment.ID, updateEnvironmentAttributes)
	if err != nil {
		t.Error("Expected UpdateEnvironment to not return an error but got", err.Error())
	}

	environment, err = client.GetEnvironment(*envID)
	if err != nil {
		t.Error("Expected GetEnvironment to not return an error but got", err.Error())
	}
	if !reflect.DeepEqual(environment.Handle, updateEnvironmentAttributes.Handle) {
		t.Errorf("Expected the environment to have some handle changes: expected %v, got %v", updateEnvironmentAttributes.Handle, environment.Handle)
	}

	// delete it
	err = client.DeleteEnvironment(*envID)
	if err != nil {
		t.Error("Expected DeleteEnvironment to not return an error but got", err.Error())
	}

	testNotFound(*environment.ID)
}
