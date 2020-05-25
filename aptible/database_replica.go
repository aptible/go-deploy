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

type ReplicaIdentifiers struct {
	ReplicaID     int64
	ProvisionOpID int64
}

func (c *Client) CreateReplica(attrs ReplicateAttrs) (Database, error) {
	operationType := "replicate"

	req := models.AppRequest23{
		Type:          &operationType,
		Handle:        attrs.ReplicaHandle,
		ContainerSize: attrs.ContainerSize,
		DiskSize:      attrs.DiskSize,
	}

	// replicate operation
	operationParams := operations.NewPostDatabasesDatabaseIDOperationsParams().WithDatabaseID(attrs.DatabaseID).WithAppRequest(&req)
	_, err := c.Client.Operations.PostDatabasesDatabaseIDOperations(operationParams, c.Token)
	if err != nil {
		return Database{}, err
	}
	fmt.Println("The replicate operation has started.")

	repl, err := c.GetReplicaFromHandle(attrs.DatabaseID, attrs.ReplicaHandle)
	for repl == nil {
		time.Sleep(5)
		repl, err = c.GetReplicaFromHandle(attrs.DatabaseID, attrs.ReplicaHandle)
	}

	replicaId := repl.ID

	// Wait on provision operation
	operation := repl.Embedded.LastOperation
	if operation == nil {
		return Database{}, fmt.Errorf("last operation is a nil pointer")
	}
	operationId := (*operation).ID
	deleted, err := c.WaitForOperation(operationId)
	if deleted {
		return Database{}, fmt.Errorf("the replica with handle: %s was unexpectedly deleted", attrs.ReplicaHandle)
	}
	if err != nil {
		return Database{}, err
	}

	// Get replica attributes
	r, deleted, err := c.GetReplica(replicaId)
	if deleted {
		return Database{}, fmt.Errorf("the replica with handle: %s was unexpectedly deleted", attrs.ReplicaHandle)
	}
	if err != nil {
		return Database{}, err
	}
	return r, nil
}

func (c *Client) GetReplica(replicaId int64) (Database, bool, error) {
	return c.GetDatabase(replicaId)
}

func (c *Client) UpdateReplica(replicaId int64, updates DBUpdates) error {
	return c.UpdateDatabase(replicaId, updates)
}

func (c *Client) DeleteReplica(replicaId int64) error {
	return c.DeleteDatabase(replicaId)
}

func (c *Client) GetReplicaFromHandle(databaseId int64, handle string) (*models.InlineResponse20014EmbeddedDatabases, error) {
	params := operations.NewGetDatabasesDatabaseIDDependentsParams().WithDatabaseID(databaseId)
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
	return nil, fmt.Errorf("there are no replicas for database %d with handle: `%s`", databaseId, handle)
}
