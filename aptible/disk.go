package aptible

import (
	"fmt"
	"github.com/aptible/go-deploy/client/operations"
)

type Disk struct {
	ID               int64
	Size             int64
	AvailabilityZone string
}

func (c *Client) GetDisk(diskID int64) (Disk, error) {
	disk := Disk{}

	params := operations.NewGetDisksIDParams().WithID(diskID)
	response, err := c.Client.Operations.GetDisksID(params, c.Token)
	if err != nil {
		return disk, err
	}

	if response.Payload.Size == nil {
		return disk, fmt.Errorf("disk size is a nil pointer")
	}
	disk.Size = *response.Payload.Size

	if response.Payload.AvailabilityZone == nil {
		return disk, fmt.Errorf("disk availability zone is a nil pointer")
	}
	disk.AvailabilityZone = *response.Payload.AvailabilityZone

	return disk, nil
}

func (c *Client) GetDiskFromHref(href string) (Disk, error) {
	diskID, err := GetIDFromHref(href)
	if err != nil {
		return Disk{}, err
	}
	return c.GetDisk(diskID)
}