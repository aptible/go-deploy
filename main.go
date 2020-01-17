package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/aptible/go-deploy/client/operations"
	"github.com/go-openapi/strfmt"

	deploy "github.com/aptible/go-deploy/client"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/mitchellh/go-homedir"
)

func main() {
	token, err := getToken()
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

func getToken() (string, error) {
	// Tries to get token via environment variable
	token := os.Getenv("APTIBLE_ACCESS_TOKEN")
	if token != "" {
		return token, nil
	}

	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	dat, err := ioutil.ReadFile(filepath.Join(home, ".aptible", "tokens.json"))
	if err != nil {
		return "", err
	}

	// Contains tokens from the tokens.json file.
	var tokens map[string]string
	if err := json.Unmarshal(dat, &tokens); err != nil {
		panic(err)
	}

	// Gets auth server
	auth := os.Getenv("APTIBLE_AUTH_ROOT_URL")
	if auth == "" {
		auth = "https://auth.aptible.com"
	}

	// Checks if there is a token.
	token = tokens[auth]
	if token != "" {
		return token, nil
	}

	return "", fmt.Errorf("No token found for %s", auth)
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
	params := operations.NewGetAccountsAccountIDOperationsParams().WithAccountID(1).WithPage(&page)
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
