package aptible

import (
	"fmt"
	"github.com/aptible/go-deploy/client/operations"
	"strings"
)

type Stack struct {
	ID             int64
	OrganizationID string
	Name           string
	Type           string
}

func (s Stack) isShared() bool {
	return s.OrganizationID == ""
}

func (c *Client) GetStacks() ([]Stack, error) {
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
			Type: stacks.Payload.Embedded.Stacks[idx].ResourceType,
		}
		if stacks.Payload.Embedded.Stacks[idx].Links.Organization != nil &&
			stacks.Payload.Embedded.Stacks[idx].Links.Organization.Href != "" {
			orgIdParts := strings.Split(stacks.Payload.Embedded.Stacks[idx].Links.Organization.Href.String(), "/")
			stacksToReturn[idx].OrganizationID = orgIdParts[len(orgIdParts)-1]
		}
	}

	return stacksToReturn, nil
}

func (c *Client) GetStackById(id int64) (Stack, error) {
	stacks, err := c.GetStacks()
	if err != nil {
		return Stack{}, err
	}
	stackToReturn := Stack{
		ID: id,
	}
	for idx, stack := range stacks {
		if id == stack.ID {
			stackToReturn.Name = stack.Name
			break
		}
		if idx == len(stacks)-1 {
			return Stack{}, fmt.Errorf("Error, unable to find stack from the id provided: %d\n", id)
		}
	}
	return stackToReturn, nil
}

func (c *Client) GetStacksByName(name string) ([]Stack, error) {
	stacks, err := c.GetStacks()
	if err != nil {
		return nil, err
	}
	var stacksToReturn []Stack
	for idx, stack := range stacks {
		if name == stack.Name {
			stacksToReturn = append(stacksToReturn, Stack{
				ID:   stack.ID,
				Name: stack.Name,
			})
		}
		if idx == len(stacks)-1 {
			return nil, fmt.Errorf("Error, unable to find stack from the name provided: %s\n", name)
		}
	}
	return stacksToReturn, nil
}
