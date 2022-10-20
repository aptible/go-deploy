//go:build integration
// +build integration

package integration

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"github.com/aptible/go-deploy/aptible"
	"github.com/aptible/go-deploy/client/operations"
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

	_, err = client.GetEnvironmentFromHandle(fmt.Sprintf("THIS_HANDLE_SHOULD_NOT_EXIST_%d", rand.Int()))
	if err == nil {
		t.Error("Expected no environment to be found and magically one was found!", err.Error())
		return
	}

	// create it
	fmt.Println("Testing environment creation")
	environment, err = client.CreateEnvironment(orgID, stackID, aptible.EnvironmentCreateAttrs{
		Handle: "1234-testing-handle",
	})
	if err != nil {
		t.Fatal("Expected CreateEnvironment to not return an error but got", err.Error())
		return
	}

	// Check pagination of GetEnvironmentFromHandle
	var envs []aptible.Environment

	defer func() {
		for _, env := range envs {
			if err := client.DeleteEnvironment(env.ID); err != nil {
				t.Errorf("Error deleteing Environment %s (%d): %s", env.Handle, env.ID, err.Error())
			}
		}
	}()

	for i := 0; i < 40; i++ {
		handle := fmt.Sprintf("pagination-test-%d", i)
		env, err := client.CreateEnvironment(orgID, stackID, aptible.EnvironmentCreateAttrs{
			Handle: handle,
		})
		if err != nil {
			t.Fatalf("Error creating environment %s: %s", handle, err.Error())
			return
		}
		envs = append(envs, env)
	}

	page := int64(1)
	params := operations.NewGetAccountsParams().WithPage(&page)
	environments, err := client.Client.Operations.GetAccounts(params, client.Token)
	if err != nil {
		t.Fatal("Error getting Environments:", err.Error())
		return
	}

	for _, env := range environments.GetPayload().Embedded.Accounts {
		if env.ID == environment.ID {
			t.Fatalf("Expected envionment %d to not be in the first page of results, but it was", environment.ID)
			return
		}
	}

	environmentByHandle, err := client.GetEnvironmentFromHandle(environment.Handle)
	if err != nil {
		t.Fatal("Expected GetEnvironmentFromHandle to find an environment with handle provided, but it errored", err.Error())
		return
	}
	if environmentByHandle.Handle != environment.Handle {
		t.Fatal("Expected found environment handle (environmentByHandle) to match the one that was provided yet (environment) there is a mismatch", environmentByHandle.Handle, environment.Handle)
		return
	}

	// get it
	fmt.Println("Testing environment getting")
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
	if !reflect.DeepEqual(environment.OrganizationID, orgID) {
		t.Errorf("Expected the organization id to be the same as when it was created: expected %v, got %v", environment.OrganizationID, orgID)
		return
	}
	if !reflect.DeepEqual(environment.StackID, stackID) {
		t.Errorf("Expected the stack id to be the same as when it was created: expected %v, got %v", environment.StackID, stackID)
		return
	}

	// update it
	fmt.Println("Testing environment update")
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
	fmt.Println("Testing environment deletion")
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
