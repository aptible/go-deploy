//go:build integration
// +build integration

package integration

import (
	"github.com/aptible/go-deploy/aptible"
	"reflect"
	"testing"
)

func TestMetricDrains(t *testing.T) {
	// Setup
	integrationPreCheck(t)
	envID := getEnvID(t)
	client := getClient(t)
	var metricDrain *aptible.MetricDrain
	var err error

	// Test
	testNotFound := func(id int64) {
		metricDrain, err = client.GetMetricDrain(id)
		if err != nil {
			t.Error("Expected GetMetricDrain to not return an error but got", err.Error())
		}
		if !metricDrain.Deleted {
			t.Error("Expected the drain to be deleted but it is not")
		}
	}

	testNotFound(0)

	metricDrain, err = client.CreateMetricDrain("it-metric-drain", envID, &aptible.MetricDrainCreateAttrs{
		DrainType: "datadog",
		APIKey:    "ABC123",
	})
	if err != nil {
		t.Fatal("Expected CreateMetricDrain to not return an error but got", err.Error())
		return
	}

	drainID := metricDrain.ID
	prevDrain := metricDrain
	metricDrain, err = client.GetMetricDrain(drainID)
	if err != nil {
		t.Error("Expected GetMetricDrain to not return an error but got", err.Error())
	}
	if !reflect.DeepEqual(metricDrain, prevDrain) {
		t.Errorf("Expected the drain to be the same as when it was created: expected %v, got %v", prevDrain, metricDrain)
	}

	success, err := client.DeleteMetricDrain(drainID)
	if err != nil {
		t.Error("Expected DeleteMetricDrain to not return an error but got", err.Error())
	}
	if !success {
		t.Error("Expected DeleteMetricDrain to return true but got false")
	}

	testNotFound(metricDrain.ID)
}
