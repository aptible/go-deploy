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

// GetDatabasesDatabaseIDDependentsReader is a Reader for the GetDatabasesDatabaseIDDependents structure.
type GetDatabasesDatabaseIDDependentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDatabasesDatabaseIDDependentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDatabasesDatabaseIDDependentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDatabasesDatabaseIDDependentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDatabasesDatabaseIDDependentsOK creates a GetDatabasesDatabaseIDDependentsOK with default headers values
func NewGetDatabasesDatabaseIDDependentsOK() *GetDatabasesDatabaseIDDependentsOK {
	return &GetDatabasesDatabaseIDDependentsOK{}
}

/*GetDatabasesDatabaseIDDependentsOK handles this case with default header values.

successful
*/
type GetDatabasesDatabaseIDDependentsOK struct {
	Payload *models.InlineResponse20016
}

func (o *GetDatabasesDatabaseIDDependentsOK) Error() string {
	return fmt.Sprintf("[GET /databases/{database_id}/dependents][%d] getDatabasesDatabaseIdDependentsOK  %+v", 200, o.Payload)
}

func (o *GetDatabasesDatabaseIDDependentsOK) GetPayload() *models.InlineResponse20016 {
	return o.Payload
}

func (o *GetDatabasesDatabaseIDDependentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20016)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatabasesDatabaseIDDependentsDefault creates a GetDatabasesDatabaseIDDependentsDefault with default headers values
func NewGetDatabasesDatabaseIDDependentsDefault(code int) *GetDatabasesDatabaseIDDependentsDefault {
	return &GetDatabasesDatabaseIDDependentsDefault{
		_statusCode: code,
	}
}

/*GetDatabasesDatabaseIDDependentsDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetDatabasesDatabaseIDDependentsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get databases database ID dependents default response
func (o *GetDatabasesDatabaseIDDependentsDefault) Code() int {
	return o._statusCode
}

func (o *GetDatabasesDatabaseIDDependentsDefault) Error() string {
	return fmt.Sprintf("[GET /databases/{database_id}/dependents][%d] GetDatabasesDatabaseIDDependents default  %+v", o._statusCode, o.Payload)
}

func (o *GetDatabasesDatabaseIDDependentsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetDatabasesDatabaseIDDependentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
