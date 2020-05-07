package aptible

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/reggregory/go-deploy/client/operations"
)

// Waits for operation to succeed.
func (c *Client) WaitForOperation(op_id int64) (bool, error) {
	if c == nil {
		return false, fmt.Errorf("Client is nil!")
	}
	params := operations.NewGetOperationsIDParams().WithID(op_id)
	op, err := c.Client.Operations.GetOperationsID(params, c.Token)
	if err != nil {
		err_struct := err.(*operations.GetOperationsIDDefault)
		switch err_struct.Code() {
		case 404:
			// If deleted, then the resource needs to be removed from Terraform.
			return true, nil
		default:
			e := fmt.Errorf("There was an error when getting the operation. \n[ERROR] -%s", err)
			return false, e
		}
	}
	status := *op.Payload.Status

	for status != "succeeded" {
		if status == "failed" {
			return false, fmt.Errorf("[ERROR] - Operation failed!")
		}
		time.Sleep(5 * time.Second)
		op, err = c.Client.Operations.GetOperationsID(params, c.Token)
		if err != nil {
			err_struct := err.(*operations.GetOperationsIDDefault)
			switch err_struct.Code() {
			case 404:
				// If deleted, then the resource needs to be removed from Terraform.
				return true, nil
			default:
				e := fmt.Errorf("There was an error when getting the operation. \n[ERROR] -%s", err)
				return false, e
			}
		}
		status = *op.Payload.Status
		fmt.Println("Still creating...")
	}
	fmt.Println("Done!")

	return false, nil
}

// Gets ID from an href
func GetIDFromHref(href string) (int64, error) {
	str := ""
	splits := strings.Split(href, "/")
	if len(splits) == 5 {
		str = splits[4]
		id, err := strconv.Atoi(str)
		if err != nil {
			return 0, fmt.Errorf("Failed to convert string to int for href = %s \n[ERROR] %s", href, err)
		}
		return int64(id), nil
	}
	return 0, fmt.Errorf("Href is shorter than expected. Better parsing is needed.")
}

// makes a string slice out of a slice of type interface
func MakeStringSlice(if_slice []interface{}) ([]string, error) {
	str_slice := make([]string, len(if_slice))
	for i := 0; i < len(if_slice); i++ {
		if (reflect.TypeOf(if_slice[i]).Kind()) != reflect.String {
			return []string{}, fmt.Errorf("Slice contains non-string elements.")
		}
		str_slice[i] = (if_slice[i].(string))
	}
	return str_slice, nil
}
