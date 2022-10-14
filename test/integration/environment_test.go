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

	environment, err = client.CreateEnvironment(orgID, stackID, aptible.EnvironmentCreateAttrs{
		Handle: "1234-testing-handle",
	})
	if err != nil {
		t.Fatal("Expected CreateEnvironment to not return an error but got", err.Error())
		return
	}

	// get it

	envID := environment.ID
	prevEnvironment := environment
	environment, err = client.GetEnvironment(envID)
	if err != nil {
		t.Error("Expected GetEnvironment to not return an error but got", err.Error())
		return
	}
	if !reflect.DeepEqual(environment, prevEnvironment) {
		t.Errorf("Expected the environment to be the same as when it was created: expected %v, got %v", prevEnvironment, environment)
		return
	}

	// update it
	updateEnvironmentAttributes := aptible.EnvironmentUpdates{
		Handle: "2345-testing-handle-update",
	}
	err = client.UpdateEnvironment(environment.ID, updateEnvironmentAttributes)
	if err != nil {
		t.Error("Expected UpdateEnvironment to not return an error but got", err.Error())
		return
	}

	environment, err = client.GetEnvironment(envID)
	if err != nil {
		t.Error("Expected GetEnvironment to not return an error but got", err.Error())
		return
	}
	if !reflect.DeepEqual(environment.Handle, updateEnvironmentAttributes.Handle) {
		t.Errorf("Expected the environment to have some handle changes: expected %v, got %v", updateEnvironmentAttributes.Handle, environment.Handle)
		return
	}

	// delete it
	err = client.DeleteEnvironment(envID)
	if err != nil {
		t.Error("Expected DeleteEnvironment to not return an error but got", err.Error())
		return
	}
	environment, _ = client.GetEnvironment(envID)
	if environment.Deleted == false {
		t.Error("Expected GetEnvironment to return deleted but got false", err.Error())
		return
	}

	testNotFound(environment.ID)
}
