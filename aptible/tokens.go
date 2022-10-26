package aptible

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

// loginWithUsernameAndPassword - specifically requests APTIBLE_AUTH_ROOT_URL for token with a username/password
func loginWithUsernameAndPassword(authUrl, user, password string) (string, error) {
	authCreds := map[string]interface{}{
		"expires":    43200,
		"username":   user,
		"password":   password,
		"grant_type": "password",
		"scope":      "manage",
	}
	authPayload, err := json.Marshal(authCreds)
	if err != nil {
		return "", fmt.Errorf("Unable to encode username and password to login. Error: %v", err.Error())
	}
	request, err := http.NewRequest("POST", fmt.Sprintf("%s/tokens", authUrl), bytes.NewBuffer(authPayload))
	if err != nil {
		return "", fmt.Errorf("Unable to construct request to login with username and password. Error: %v", err.Error())
	}

	// set required headers for auth to function
	request.Header.Set("accept", "application/json, text/javascript, */*; q=0.01")
	request.Header.Set("content-type", "application/json")
	response, err := (&http.Client{}).Do(request)
	if err != nil {
		return "", fmt.Errorf("Unable to login with username and password. Error: %v", err.Error())
	}

	// read the payload
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("Unable to read response from logging in with username and password: %v", err.Error())
	}

	// unmarshal to extract tokens
	output := make(map[string]interface{})
	if err := json.Unmarshal(resp, &output); err != nil {
		return "", fmt.Errorf("Unable to unmarshal raw response from logging in with username and password: %v", err.Error())
	}

	// if token not found or if it's not a string (undefined/null for example)
	if _, ok := output["access_token"]; !ok || fmt.Sprintf("%T", output["access_token"]) != "string" {
		return "", fmt.Errorf("Token not found in payload when logging in with username and password")
	}

	return output["access_token"].(string), nil
}

func GetAuthURL() string {
	// Gets auth server
	auth := os.Getenv("APTIBLE_AUTH_ROOT_URL")
	if auth == "" {
		auth = "https://auth.aptible.com"
	}
	return auth
}

// GetToken - gets a token (or error) for use from username password, an environment variable or filesystem
func GetToken() (string, error) {
	auth := GetAuthURL()

	// use auth-server to get tokens
	user := os.Getenv("APTIBLE_USERNAME")
	password := os.Getenv("APTIBLE_PASSWORD")
	if user != "" && password != "" {
		return loginWithUsernameAndPassword(auth, user, password)
	}

	// Tries to get token via environment variable
	token := os.Getenv("APTIBLE_ACCESS_TOKEN")
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

	// Checks if there is a token.
	token = tokens[auth]
	if token != "" {
		return token, nil
	}

	return "", fmt.Errorf("No token can be found for %v. Are you logged in?", auth)
}
