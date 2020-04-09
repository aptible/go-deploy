package aptible

import (
	"fmt"
	"time"

	"github.com/aptible/go-deploy/client/operations"
	"github.com/aptible/go-deploy/models"
)

type ReplicateAttrs struct {
	EnvID         int64
	DatabaseID    int64
	ReplicaHandle string
	ContainerSize int64
	DiskSize      int64
}

func (c *Client) CreateReplica(attrs ReplicateAttrs) (*models.InlineResponse20014EmbeddedDatabases, error) {
	op_type := "replicate"

	req := models.AppRequest23{
		Type:          &op_type,
		Handle:        attrs.ReplicaHandle,
		ContainerSize: attrs.ContainerSize,
		DiskSize:      attrs.DiskSize,
	}

	op_params := operations.NewPostDatabasesDatabaseIDOperationsParams().WithDatabaseID(attrs.DatabaseID).WithAppRequest(&req)
	_, err := c.Client.Operations.PostDatabasesDatabaseIDOperations(op_params, c.Token)
	if err != nil {
		return nil, err
	}

	// waiting for provision operation to start...
	payload, deleted, err := c.GetDatabaseFromHandle(attrs.EnvID, attrs.ReplicaHandle)
	for payload == nil && err != nil && !deleted {
		payload, deleted, err = c.GetDatabaseFromHandle(attrs.EnvID, attrs.ReplicaHandle)
	}

	// waiting for provision operation to complete...
	for payload.Status != "provisioned" {
		if payload.Status == "failed" {
			return nil, fmt.Errorf("The provision failed.")
		}
		payload, _, err = c.GetDatabaseFromHandle(attrs.EnvID, attrs.ReplicaHandle)
		if err != nil {
			return nil, err
		}
		time.Sleep(3 * time.Second)
		fmt.Println("Still creating...")
	}
	fmt.Println("Done!")
	return payload, nil
}

func (c *Client) GetReplica(replica_id int64) (DBUpdates, bool, error) {
	return c.GetDatabase(replica_id)
}

func (c *Client) UpdateReplica(replica_id int64, updates DBUpdates) error {
	return c.UpdateDatabase(replica_id, updates)
}

func (c *Client) DeleteReplica(replica_id int64) error {
	return c.DeleteDatabase(replica_id)
}
