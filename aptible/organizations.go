package aptible

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type HALOrganization struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type HALOrganizationParent struct {
	Organizations []HALOrganization `json:"organizations"`
}

type HALOrganizationResponse struct {
	Embedded HALOrganizationParent `json:"_embedded"`
}

type Organization struct {
	ID   string
	Name string
}

// getOrganizations - get all organizations a user has access to
func (c *Client) getOrganizations() ([]Organization, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/organizations", GetAuthURL()), nil)
	//Handle Error
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.RawToken))
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		bodyErr := Body.Close()
		if bodyErr != nil {
			log.Fatalf("Error in body reader close of http client %s\n", err.Error())
		}
	}(resp.Body)

	out := HALOrganizationResponse{}
	if err = json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}

	orgs := make([]Organization, len(out.Embedded.Organizations))
	for orgIdx, org := range out.Embedded.Organizations {
		orgs[orgIdx] = Organization{
			ID:   org.ID,
			Name: org.Name,
		}
	}

	return orgs, nil
}

// GetOrganization get organizations by user's token. Attempts to get organizations from auth api, then
// attempts to get the FIRST one. if more than one is found or none are found, error is raised
func (c *Client) GetOrganization() (Organization, error) {
	organizations, err := c.getOrganizations()
	if err != nil {
		return Organization{}, err
	}
	if len(organizations) > 1 {
		return Organization{}, fmt.Errorf("multiple organizations for user, unable to determine" +
			" a default organization in result")
	}

	if len(organizations) == 0 {
		return Organization{}, fmt.Errorf("no organizations found in response")
	}

	return organizations[0], nil
}
