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

// GetDatabasesDatabaseIDConfigurationsReader is a Reader for the GetDatabasesDatabaseIDConfigurations structure.
type GetDatabasesDatabaseIDConfigurationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDatabasesDatabaseIDConfigurationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDatabasesDatabaseIDConfigurationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDatabasesDatabaseIDConfigurationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDatabasesDatabaseIDConfigurationsOK creates a GetDatabasesDatabaseIDConfigurationsOK with default headers values
func NewGetDatabasesDatabaseIDConfigurationsOK() *GetDatabasesDatabaseIDConfigurationsOK {
	return &GetDatabasesDatabaseIDConfigurationsOK{}
}

/* GetDatabasesDatabaseIDConfigurationsOK describes a response with status code 200, with default header values.

successful
*/
type GetDatabasesDatabaseIDConfigurationsOK struct {
	Payload *models.InlineResponse2009
}

func (o *GetDatabasesDatabaseIDConfigurationsOK) Error() string {
	return fmt.Sprintf("[GET /databases/{database_id}/configurations][%d] getDatabasesDatabaseIdConfigurationsOK  %+v", 200, o.Payload)
}
func (o *GetDatabasesDatabaseIDConfigurationsOK) GetPayload() *models.InlineResponse2009 {
	return o.Payload
}

func (o *GetDatabasesDatabaseIDConfigurationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse2009)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatabasesDatabaseIDConfigurationsDefault creates a GetDatabasesDatabaseIDConfigurationsDefault with default headers values
func NewGetDatabasesDatabaseIDConfigurationsDefault(code int) *GetDatabasesDatabaseIDConfigurationsDefault {
	return &GetDatabasesDatabaseIDConfigurationsDefault{
		_statusCode: code,
	}
}

/* GetDatabasesDatabaseIDConfigurationsDefault describes a response with status code -1, with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetDatabasesDatabaseIDConfigurationsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get databases database ID configurations default response
func (o *GetDatabasesDatabaseIDConfigurationsDefault) Code() int {
	return o._statusCode
}

func (o *GetDatabasesDatabaseIDConfigurationsDefault) Error() string {
	return fmt.Sprintf("[GET /databases/{database_id}/configurations][%d] GetDatabasesDatabaseIDConfigurations default  %+v", o._statusCode, o.Payload)
}
func (o *GetDatabasesDatabaseIDConfigurationsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetDatabasesDatabaseIDConfigurationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
