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
	op_type := "replicate"

	req := models.AppRequest23{
		Type:          &op_type,
		Handle:        attrs.ReplicaHandle,
		ContainerSize: attrs.ContainerSize,
		DiskSize:      attrs.DiskSize,
	}

	// replicate operation
	op_params := operations.NewPostDatabasesDatabaseIDOperationsParams().WithDatabaseID(attrs.DatabaseID).WithAppRequest(&req)
	_, err := c.Client.Operations.PostDatabasesDatabaseIDOperations(op_params, c.Token)
	if err != nil {
		return Database{}, err
	}
	fmt.Println("The replicate operation has started.")

	repl, err := c.GetReplicaFromHandle(attrs.DatabaseID, attrs.ReplicaHandle)
	for repl == nil {
		time.Sleep(5)
		repl, err = c.GetReplicaFromHandle(attrs.DatabaseID, attrs.ReplicaHandle)
	}

	repl_id := repl.ID

	// Wait on provision operation
	op := repl.Embedded.LastOperation
	if op == nil {
		return Database{}, fmt.Errorf("Last operation is a nil pointer.")
	}
	op_id := (*op).ID
	deleted, err := c.WaitForOperation(op_id)
	if deleted {
		return Database{}, fmt.Errorf("The replica with handle: %s was unexpectedly deleted.", attrs.ReplicaHandle)
	}
	if err != nil {
		return Database{}, err
	}

	// Get replica attributes
	r, deleted, err := c.GetReplica(repl_id)
	if deleted {
		return Database{}, fmt.Errorf("The replica with handle: %s was unexpectedly deleted.", attrs.ReplicaHandle)
	}
	if err != nil {
		return Database{}, err
	}
	return r, nil
}

func (c *Client) GetReplica(replica_id int64) (Database, bool, error) {
	return c.GetDatabase(replica_id)
}

func (c *Client) UpdateReplica(replica_id int64, updates DBUpdates) error {
	return c.UpdateDatabase(replica_id, updates)
}

func (c *Client) DeleteReplica(replica_id int64) error {
	return c.DeleteDatabase(replica_id)
}

func (c *Client) GetReplicaFromHandle(db_id int64, handle string) (*models.InlineResponse20014EmbeddedDatabases, error) {
	params := operations.NewGetDatabasesDatabaseIDDependentsParams().WithDatabaseID(db_id)
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
	return nil, fmt.Errorf("There are no replicas for database %d with handle: `%s`.", db_id, handle)
}
