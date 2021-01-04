// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aptible/go-deploy/models"
)

// GetAccountsAccountIDBackupRetentionPoliciesReader is a Reader for the GetAccountsAccountIDBackupRetentionPolicies structure.
type GetAccountsAccountIDBackupRetentionPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountsAccountIDBackupRetentionPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountsAccountIDBackupRetentionPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAccountsAccountIDBackupRetentionPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAccountsAccountIDBackupRetentionPoliciesOK creates a GetAccountsAccountIDBackupRetentionPoliciesOK with default headers values
func NewGetAccountsAccountIDBackupRetentionPoliciesOK() *GetAccountsAccountIDBackupRetentionPoliciesOK {
	return &GetAccountsAccountIDBackupRetentionPoliciesOK{}
}

/*GetAccountsAccountIDBackupRetentionPoliciesOK handles this case with default header values.

successful
*/
type GetAccountsAccountIDBackupRetentionPoliciesOK struct {
	Payload *models.InlineResponse2004
}

func (o *GetAccountsAccountIDBackupRetentionPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{account_id}/backup_retention_policies][%d] getAccountsAccountIdBackupRetentionPoliciesOK  %+v", 200, o.Payload)
}

func (o *GetAccountsAccountIDBackupRetentionPoliciesOK) GetPayload() *models.InlineResponse2004 {
	return o.Payload
}

func (o *GetAccountsAccountIDBackupRetentionPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse2004)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsAccountIDBackupRetentionPoliciesDefault creates a GetAccountsAccountIDBackupRetentionPoliciesDefault with default headers values
func NewGetAccountsAccountIDBackupRetentionPoliciesDefault(code int) *GetAccountsAccountIDBackupRetentionPoliciesDefault {
	return &GetAccountsAccountIDBackupRetentionPoliciesDefault{
		_statusCode: code,
	}
}

/*GetAccountsAccountIDBackupRetentionPoliciesDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetAccountsAccountIDBackupRetentionPoliciesDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get accounts account ID backup retention policies default response
func (o *GetAccountsAccountIDBackupRetentionPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *GetAccountsAccountIDBackupRetentionPoliciesDefault) Error() string {
	return fmt.Sprintf("[GET /accounts/{account_id}/backup_retention_policies][%d] GetAccountsAccountIDBackupRetentionPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *GetAccountsAccountIDBackupRetentionPoliciesDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetAccountsAccountIDBackupRetentionPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
