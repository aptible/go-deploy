package aptible

import (
	"github.com/reggregory/go-deploy/client/operations"
	"github.com/reggregory/go-deploy/models"
)

type ReplicateAttrs struct {
	ReplicaHandle string
	ContainerSize int64
	DiskSize      int64
}

func (c *Client) CreateReplica(db_id int64, attrs ReplicateAttrs) (*models.InlineResponse20014EmbeddedDatabases, error) {
	op_type := "replicate"

	req := models.AppRequest23{
		Type:          &op_type,
		Handle:        attrs.ReplicaHandle,
		ContainerSize: attrs.ContainerSize,
		DiskSize:      attrs.DiskSize,
	}

	op_params := operations.NewPostDatabasesDatabaseIDOperationsParams().WithDatabaseID(db_id).WithAppRequest(&req)
	_, err := c.Client.Operations.PostDatabasesDatabaseIDOperations(op_params, c.Token)
	if err != nil {
		return nil, err
	}

	payload, err := c.GetDatabaseFromHandle(4, attrs.ReplicaHandle)
	for payload == nil && err != nil {
		payload, err = c.GetDatabaseFromHandle(4, attrs.ReplicaHandle)
	}
	return payload, nil
}

func (c *Client) GetReplica(replica_id int64) (*models.InlineResponse2014, bool, error) {
	return c.GetDatabase(replica_id)
}

func (c *Client) UpdateReplica(replica_id int64, updates DBUpdates) error {
	return c.UpdateDatabase(replica_id, updates)
}

func (c *Client) DeleteReplica(replica_id int64) error {
	return c.DeleteDatabase(replica_id)
}
