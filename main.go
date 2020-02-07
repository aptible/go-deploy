package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-openapi/strfmt"
	"github.com/reggregory/go-deploy/client/operations"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	deploy "github.com/reggregory/go-deploy/client"
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
	rt := httptransport.New(
		os.Getenv("APTIBLE_API_ROOT_URL"),
		deploy.DefaultBasePath,
		deploy.DefaultSchemes)
	rt.Consumers["application/hal+json"] = runtime.JSONConsumer()
	rt.Producers["application/hal+json"] = runtime.JSONProducer()
	client := deploy.New(rt, strfmt.Default)

	bearerTokenAuth := httptransport.BearerToken(token)

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
