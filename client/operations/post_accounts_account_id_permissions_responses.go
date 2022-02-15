// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/aptible/go-deploy/models"
)

// PostAccountsAccountIDPermissionsReader is a Reader for the PostAccountsAccountIDPermissions structure.
type PostAccountsAccountIDPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAccountsAccountIDPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostAccountsAccountIDPermissionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostAccountsAccountIDPermissionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostAccountsAccountIDPermissionsCreated creates a PostAccountsAccountIDPermissionsCreated with default headers values
func NewPostAccountsAccountIDPermissionsCreated() *PostAccountsAccountIDPermissionsCreated {
	return &PostAccountsAccountIDPermissionsCreated{}
}

/* PostAccountsAccountIDPermissionsCreated describes a response with status code 201, with default header values.

successful
*/
type PostAccountsAccountIDPermissionsCreated struct {
	Payload *models.InlineResponse2017
}

func (o *PostAccountsAccountIDPermissionsCreated) Error() string {
	return fmt.Sprintf("[POST /accounts/{account_id}/permissions][%d] postAccountsAccountIdPermissionsCreated  %+v", 201, o.Payload)
}
func (o *PostAccountsAccountIDPermissionsCreated) GetPayload() *models.InlineResponse2017 {
	return o.Payload
}

func (o *PostAccountsAccountIDPermissionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse2017)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAccountsAccountIDPermissionsDefault creates a PostAccountsAccountIDPermissionsDefault with default headers values
func NewPostAccountsAccountIDPermissionsDefault(code int) *PostAccountsAccountIDPermissionsDefault {
	return &PostAccountsAccountIDPermissionsDefault{
		_statusCode: code,
	}
}

/* PostAccountsAccountIDPermissionsDefault describes a response with status code -1, with default header values.

Error response. Often a 4xx or 5xx status code
*/
type PostAccountsAccountIDPermissionsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the post accounts account ID permissions default response
func (o *PostAccountsAccountIDPermissionsDefault) Code() int {
	return o._statusCode
}

func (o *PostAccountsAccountIDPermissionsDefault) Error() string {
	return fmt.Sprintf("[POST /accounts/{account_id}/permissions][%d] PostAccountsAccountIDPermissions default  %+v", o._statusCode, o.Payload)
}
func (o *PostAccountsAccountIDPermissionsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *PostAccountsAccountIDPermissionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
