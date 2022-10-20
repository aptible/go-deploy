package aptible

import (
	"fmt"
	"strings"

	"github.com/aptible/go-deploy/client/operations"
)

const MaximumPagesOfStacks = 10

type Stack struct {
	ID             int64
	OrganizationID string
	Name           string
}

func (s Stack) isShared() bool {
	return s.OrganizationID == ""
}

func (c *Client) GetStacks() ([]Stack, error) {
	page := int64(1)
	stacksToReturn := make([]Stack, 0)

	for {
		params := operations.NewGetStacksParams().WithPage(&page)
		stacks, err := c.Client.Operations.GetStacks(params, c.Token)
		if err != nil {
			return nil, err
		} else if len(stacks.GetPayload().Embedded.Stacks) == 0 {
			break
		} else if page >= MaximumPagesOfStacks {
			// infinite loop guard
			return nil, fmt.Errorf("exceeded %d pages of results for stacks in population. "+
				"Something has gone wrong", MaximumPagesOfStacks)
		}

		for _, stack := range stacks.GetPayload().Embedded.Stacks {
			stackToAppend := Stack{
				ID:   stack.ID,
				Name: stack.Name,
			}
			if stack.Links.Organization != nil &&
				stack.Links.Organization.Href != "" {
				orgIdParts := strings.Split(stack.Links.Organization.Href.String(), "/")
				stackToAppend.OrganizationID = orgIdParts[len(orgIdParts)-1]
			}
			stacksToReturn = append(stacksToReturn, stackToAppend)
		}
		page += 1
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

func (c *Client) GetStackByName(name string) (Stack, error) {
	stacks, err := c.GetStacks()
	if err != nil {
		return Stack{}, err
	}
	for _, stack := range stacks {
		if name == stack.Name {
			return Stack{
				ID:             stack.ID,
				Name:           stack.Name,
				OrganizationID: stack.OrganizationID,
			}, nil
		}
	}
	return Stack{}, fmt.Errorf("Error, unable to find stack from the name provided: %s\n", name)
}
