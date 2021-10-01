//go:build api
// +build api

// These tests call SetUpClient and require a valid client
// setup for API access. By default they are skipped on CI

package aptible

import (
	"testing"
)

func TestWaitForOperation(t *testing.T) {
	var tests = []struct {
		name        string
		operationID int64
		deleted     bool
		errored     bool
	}{
		{"test_404", 0, true, false},
		// Add other test cases if we need in the future
	}

	for _, tc := range tests {
		c, err := SetUpClient()
		if err != nil {
			t.Errorf("Unable to set up client due to error. \nERROR -- %s", err)
		}
		testName := tc.name
		t.Run(testName, func(t *testing.T) {

			deleted, err := c.WaitForOperation(tc.operationID)

			if deleted != tc.deleted {
				t.Errorf("Input: %d should have resulted in a 404.", tc.operationID)
			}

			if (err != nil) != tc.errored {
				t.Errorf("Input: %d caused error: %s", tc.operationID, err)
			}
		})
	}
}
