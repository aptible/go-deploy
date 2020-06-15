package aptible

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func GetToken() (string, error) {
	// Tries to get token via environment variable
	token := os.Getenv("APTIBLE_ACCESS_TOKEN")
	if token != "" {
		return token, nil
	}

	home, err := homedir.Dir()
	fmt.Println(home)
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

	fmt.Println(auth)
	fmt.Printf("%+v\n", tokens)

	// Checks if there is a token.
	token = tokens[auth]
	if token != "" {
		return token, nil
	}

	return "", fmt.Errorf("no token found for %s", auth)
}
