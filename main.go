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
