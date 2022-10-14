package aptible

import (
	"fmt"
	"strings"

	"github.com/aptible/go-deploy/client/operations"
)

type Stack struct {
	ID             int64
	OrganizationID string
	Name           string
}

func (s Stack) isShared() bool {
	return s.OrganizationID == ""
}

func (c *Client) UNSUPPORTED_GetStacks() ([]Stack, error) {
	page := int64(0)
	params := operations.NewGetStacksParams().WithPage(&page)

	stacks, err := c.Client.Operations.GetStacks(params, c.Token)
	if err != nil {
		return nil, err
	}

	stacksToReturn := make([]Stack, len(stacks.Payload.Embedded.Stacks))
	for idx := range stacksToReturn {
		stacksToReturn[idx] = Stack{
			ID:   stacks.Payload.Embedded.Stacks[idx].ID,
			Name: stacks.Payload.Embedded.Stacks[idx].Name,
		}
		if stacks.Payload.Embedded.Stacks[idx].Links.Organization != nil &&
			stacks.Payload.Embedded.Stacks[idx].Links.Organization.Href != "" {
			orgIdParts := strings.Split(stacks.Payload.Embedded.Stacks[idx].Links.Organization.Href.String(), "/")
			stacksToReturn[idx].OrganizationID = orgIdParts[len(orgIdParts)-1]
		}
	}

	return stacksToReturn, nil
}

func (c *Client) GetStack(id int64) (Stack, error) {
	params := operations.NewGetStacksIDParams().WithID(id)
	result, err := c.Client.Operations.GetStacksID(params, c.Token)
	if err != nil {
		return Stack{}, err
	}
	organizationId := ""
	if result.Payload.Links.Organization != nil &&
		result.Payload.Links.Organization.Href != "" {
		orgIdParts := strings.Split(result.Payload.Links.Organization.Href.String(), "/")
		organizationId = orgIdParts[len(orgIdParts)-1]
	}
	stackToReturn := Stack{
		ID:             id,
		OrganizationID: organizationId,
		Name:           *result.Payload.Name,
	}
	return stackToReturn, nil
}

func (c *Client) UNSUPPORTED_GetStackByName(name string) (Stack, error) {
	stacks, err := c.UNSUPPORTED_GetStacks()
	if err != nil {
		return Stack{}, err
	}
	var stacksToReturn Stack
	for _, stack := range stacks {
		if name == stack.Name {
			return Stack{
				ID:             stack.ID,
				Name:           stack.Name,
				OrganizationID: stack.OrganizationID,
			}, nil
		}
	}
	return stacksToReturn, fmt.Errorf("Error, unable to find stack from the name provided: %s\n", name)
}
