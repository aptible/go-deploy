package aptible

import (
	"fmt"

	"github.com/aptible/go-deploy/client/operations"
)

type DatabaseImage struct {
	ID      int64
	Type    string
	Version string
}

func (c *Client) GetDatabaseImage(imageID int64) (DatabaseImage, error) {
	image := DatabaseImage{
		ID: imageID,
	}

	params := operations.NewGetDatabaseImagesIDParams().WithID(imageID)
	response, err := c.Client.Operations.GetDatabaseImagesID(params, c.Token)
	if err != nil {
		return image, err
	}

	imageType := response.Payload.Type
	if imageType == nil {
		return image, fmt.Errorf("type is a nil pointer")
	}
	image.Type = *imageType

	version := response.Payload.Version
	if version == nil {
		return image, fmt.Errorf("version is a nil pointer")
	}
	image.Version = *version

	return image, nil
}

func (c *Client) GetImageFromHref(href string) (DatabaseImage, error) {
	imageID, err := GetIDFromHref(href)
	if err != nil {
		return DatabaseImage{}, err
	}
	return c.GetDatabaseImage(imageID)
}

func (c *Client) GetDatabaseImageByTypeAndVersion(imageType string, version string) (DatabaseImage, error) {
	image := DatabaseImage{
		Type:    imageType,
		Version: version,
	}

	params := operations.NewGetDatabaseImagesParams()
	response, err := c.Client.Operations.GetDatabaseImages(params, c.Token)
	if err != nil {
		return image, err
	}

	if response.Payload.TotalCount == nil {
		return image, fmt.Errorf("TotalCount is a nil pointer")
	}
	numDatabaseImages := *response.Payload.TotalCount

	if response.Payload.PerPage == nil {
		return image, fmt.Errorf("PerPage is a nil pointer")
	}
	perPage := *response.Payload.PerPage

	if response.Payload.TotalCount == nil {
		return image, fmt.Errorf("CurrentPage is a nil pointer")
	}
	page := *response.Payload.CurrentPage

	for numDatabaseImages > 0 {
		databaseImages := response.Payload.Embedded.DatabaseImages
		for _, i := range databaseImages {
			if i.Version == version && i.Type == imageType {
				image.ID = i.ID
				return image, nil
			}
		}

		if numDatabaseImages-perPage > 0 {
			numDatabaseImages -= perPage
			page += 1
		} else {
			return image, fmt.Errorf("database image cannot be found")
		}

		params := operations.NewGetDatabaseImagesParams().WithPage(&page)
		response, err = c.Client.Operations.GetDatabaseImages(params, c.Token)
		if err != nil {
			return image, err
		}
	}

	return image, fmt.Errorf("database image cannot be found")
}
