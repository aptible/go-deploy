package aptible

import (
	"fmt"

	"github.com/aptible/go-deploy/client/operations"
	"github.com/aptible/go-deploy/models"
)

type ReplicateAttrs struct {
	EnvID            int64
	DatabaseID       int64
	ReplicaHandle    string
	ContainerSize    int64
	DiskSize         int64
	ContainerProfile string
}

type ReplicaIdentifiers struct {
	ReplicaID     int64
	ProvisionOpID int64
}

func (c *Client) CreateReplica(attrs ReplicateAttrs) (Database, error) {
	operationType := "replicate"

	req := models.AppRequest24{
		Type:            &operationType,
		Handle:          attrs.ReplicaHandle,
		ContainerSize:   attrs.ContainerSize,
		DiskSize:        attrs.DiskSize,
		InstanceProfile: attrs.ContainerProfile,
		ProvisionedIops: attrs.Iops,
	}

	// replicate operation
	operationParams := operations.NewPostDatabasesDatabaseIDOperationsParams().WithDatabaseID(attrs.DatabaseID).WithAppRequest(&req)
	op, err := c.Client.Operations.PostDatabasesDatabaseIDOperations(operationParams, c.Token)
	if err != nil {
		return Database{}, err
	}
	operationID := *op.Payload.ID
	_, err = c.WaitForOperation(operationID)
	if err != nil {
		return Database{}, err
	}

	repl, err := c.GetReplicaFromHandle(attrs.DatabaseID, attrs.ReplicaHandle)
	if err != nil {
		return Database{}, err
	}
	replicaID := repl.ID

	operation := repl.Embedded.LastOperation
	if operation == nil {
		return Database{}, fmt.Errorf("last operation is a nil pointer")
	}
	operationID = (*operation).ID
	deleted, err := c.WaitForOperation(operationID)
	if deleted {
		return Database{}, fmt.Errorf("the replica with handle: %s was unexpectedly deleted", attrs.ReplicaHandle)
	}
	if err != nil {
		return Database{}, err
	}

	// Get replica attributes
	r, err := c.GetReplica(replicaID)
	if deleted {
		return Database{}, fmt.Errorf("the replica with handle: %s was unexpectedly deleted", attrs.ReplicaHandle)
	}
	if err != nil {
		return Database{}, err
	}
	return r, nil
}

func (c *Client) GetReplica(replicaID int64) (Database, error) {
	return c.GetDatabase(replicaID)
}

func (c *Client) UpdateReplica(replicaID int64, updates DBUpdates) error {
	return c.UpdateDatabase(replicaID, updates)
}

func (c *Client) DeleteReplica(replicaID int64) error {
	return c.DeleteDatabase(replicaID)
}

func (c *Client) GetReplicaFromHandle(databaseID int64, handle string) (*models.InlineResponse20016EmbeddedDatabases, error) {
	params := operations.NewGetDatabasesDatabaseIDDependentsParams().WithDatabaseID(databaseID)
	resp, err := c.Client.Operations.GetDatabasesDatabaseIDDependents(params, c.Token)
	if err != nil {
		return nil, err
	}

	replicas := resp.Payload.Embedded.Databases
	for i := range replicas {
		r := replicas[i]
		if r.Handle == handle && r.Embedded.LastOperation != nil && r.Embedded.LastOperation.Type == "provision" {
			return replicas[i], nil
		}
	}
	return nil, fmt.Errorf("there are no replicas for database %d with handle: `%s`", databaseID, handle)
}
