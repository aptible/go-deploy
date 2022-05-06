package aptible

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func GetToken(attrs ClientAttrs) (string, error) {
	// Tries to get token via ClientAttrs
	token := attrs.TokenString
	if token != "" {
		return token, nil
	}

	// Tries to get token via environment variable
	token = os.Getenv("APTIBLE_ACCESS_TOKEN")
	if token != "" {
		return token, nil
	}

	home, err := homedir.Dir()
	if err != nil {
		return "", fmt.Errorf("Your token is invalid. Are you logged in? Error: %v", err.Error())
	}

	dat, err := ioutil.ReadFile(filepath.Join(home, ".aptible", "tokens.json"))
	if err != nil {
		return "", fmt.Errorf("Your token is invalid. Are you logged in? Error: %v", err.Error())
	}

	// Contains tokens from the tokens.json file.
	var tokens map[string]string
	if err := json.Unmarshal(dat, &tokens); err != nil {
		return "", fmt.Errorf("Your token is invalid. Are you logged in? Error: %v", err.Error())
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

	return "", fmt.Errorf("No token can be found for %v. Are you logged in?", auth)
}
