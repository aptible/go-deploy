package main

import (
	"fmt"
	"log"

	"github.com/aptible/go-deploy/aptible"
	"github.com/aptible/go-deploy/client/operations"
)

func main() {
	ops, err := getOperations()
	if err != nil {
		log.Fatalf("couldn't do it: %s", err)
	}

	fmt.Println("Results:")
	for _, op := range ops {
		fmt.Println(op)
	}

	c, err := aptible.SetUpClient()
	if err != nil {
		log.Fatalf("couldn't do it: %s", err)
	}

	prov_op, err := c.GetLatestProvisionOperation(33)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Op ID: ", prov_op.ID)
		fmt.Println("Container size: ", *prov_op.ContainerSize)
		fmt.Println("Disk size: ", prov_op.DiskSize)
	}

	// attrs := aptible.ReplicateAttrs{
	// 	EnvID:         4,
	// 	DatabaseID:    25,
	// 	ReplicaHandle: "replica-wait",
	// }

	// // tests getting database from handle
	// database, err := c.GetDatabaseFromHandle(4, "repli")
	// if err != nil {
	// 	fmt.Println("couldn't do it: ", err)
	// } else {
	// 	fmt.Println("database handle + ID:", database.Handle, database.ID)
	// }

	// // tests creating replica
	// replica, err := c.CreateReplica(attrs)
	// if replica != nil {
	// 	id := replica.ID

	// 	time.Sleep(20 * time.Second)

	// 	err = deleteReplica(id)
	// 	if err != nil {
	// 		log.Fatalf("couldn't do it: %s", err)
	// 	} else {
	// 		fmt.Println("Successfully deleted replica!")
	// 	}

	// } else {
	// 	fmt.Println("Replication failed.")
	// }
}

func getOperations() ([]string, error) {
	c, err := aptible.SetUpClient()
	if err != nil {
		return []string{}, err
	}
	page := int64(1)
	params := operations.NewGetAccountsAccountIDOperationsParams().WithAccountID(2).WithPage(&page)
	resp, err := c.Client.Operations.GetAccountsAccountIDOperations(params, c.Token)
	if err != nil {
		return []string{}, err
	}

	var lines []string
	for _, op := range resp.Payload.Embedded.Operations {
		line := fmt.Sprintf("%v -- %v %s - %s -- %s (%s)",
			op.CreatedAt, op.ID, op.Type, op.Status, op.UserName, op.UserEmail)
		lines = append(lines, line)
	}
	return lines, nil
}

func deleteReplica(replica_id int64) error {
	c, err := aptible.SetUpClient()
	if err != nil {
		return err
	}

	return c.DeleteReplica(replica_id)
}
