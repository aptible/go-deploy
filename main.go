package main

import (
	"fmt"
	"log"

	"github.com/reggregory/go-deploy/client/operations"

	"github.com/reggregory/go-deploy/helpers"
)

func main() {
	token, err := helpers.GetToken()
	if err != nil {
		log.Fatalf("couldn't do it: %s", err)
	}
	ops, err := getOperations(token)
	if err != nil {
		log.Fatalf("couldn't do it: %s", err)
	}

	fmt.Println("Results:")
	for _, op := range ops {
		fmt.Println(op)
	}
}

func getOperations(token string) ([]string, error) {
	client, bearerTokenAuth := helpers.SetUpClient()
	page := int64(1)
	params := operations.NewGetAccountsAccountIDOperationsParams().WithAccountID(2).WithPage(&page)
	resp, err := client.Operations.GetAccountsAccountIDOperations(params, bearerTokenAuth)
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
