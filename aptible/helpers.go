package aptible

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/aptible/go-deploy/client/operations"
)

// Waits for operation to succeed.
func (c *Client) WaitForOperation(operationID int64) (bool, error) {
	if c == nil {
		return false, fmt.Errorf("client is nil")
	}
	params := operations.NewGetOperationsIDParams().WithID(operationID)
	op, err := c.Client.Operations.GetOperationsID(params, c.Token)
	if err != nil {
		errStruct := err.(*operations.GetOperationsIDDefault)
		switch errStruct.Code() {
		case 404:
			// If deleted, then the resource needs to be removed from Terraform.
			return true, nil
		default:
			e := fmt.Errorf("there was an error when getting the operation for op id: %d \n[ERROR] -%s", operationID, err)
			return false, e
		}
	}
	status := *op.Payload.Status

	for status != "succeeded" {
		if status == "failed" {
			return false, fmt.Errorf("[ERROR] - operation failed for op id: %d", operationID)
		}
		time.Sleep(5 * time.Second)
		op, err = c.Client.Operations.GetOperationsID(params, c.Token)
		if err != nil {
			errStruct := err.(*operations.GetOperationsIDDefault)
			switch errStruct.Code() {
			case 404:
				// If deleted, then the resource needs to be removed from Terraform.
				return true, nil
			default:
				e := fmt.Errorf("there was an error when getting the operation for op id: %d \n[ERROR] -%s", operationID, err)
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
			return 0, fmt.Errorf("failed to convert string to int for href = %s \n[ERROR] %s", href, err)
		}
		return int64(id), nil
	}
	return 0, fmt.Errorf("href is shorter than expected")
}

// makes a string slice out of a slice of type interface
func MakeStringSlice(interfaceSlice []interface{}) ([]string, error) {
	strSlice := make([]string, len(interfaceSlice))
	for i := 0; i < len(interfaceSlice); i++ {
		if (reflect.TypeOf(interfaceSlice[i]).Kind()) != reflect.String {
			return []string{}, fmt.Errorf("slice contains non-string elements")
		}
		strSlice[i] = interfaceSlice[i].(string)
	}
	return strSlice, nil
}

// makes a int64 slice out of a slice of type interface
func MakeInt64Slice(interfaceSlice []interface{}) ([]int64, error) {
	int64Slice := make([]int64, len(interfaceSlice))
	for i := 0; i < len(interfaceSlice); i++ {
		if (reflect.TypeOf(interfaceSlice[i]).Kind()) != reflect.Int64 {
			return []int64{}, fmt.Errorf("slice contains non-string elements")
		}
		int64Slice[i] = interfaceSlice[i].(int64)
	}
	return int64Slice, nil
}
