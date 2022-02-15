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

// PostAppsAppIDOperationsReader is a Reader for the PostAppsAppIDOperations structure.
type PostAppsAppIDOperationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAppsAppIDOperationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostAppsAppIDOperationsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostAppsAppIDOperationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostAppsAppIDOperationsCreated creates a PostAppsAppIDOperationsCreated with default headers values
func NewPostAppsAppIDOperationsCreated() *PostAppsAppIDOperationsCreated {
	return &PostAppsAppIDOperationsCreated{}
}

/* PostAppsAppIDOperationsCreated describes a response with status code 201, with default header values.

successful
*/
type PostAppsAppIDOperationsCreated struct {
	Payload *models.InlineResponse20030
}

func (o *PostAppsAppIDOperationsCreated) Error() string {
	return fmt.Sprintf("[POST /apps/{app_id}/operations][%d] postAppsAppIdOperationsCreated  %+v", 201, o.Payload)
}
func (o *PostAppsAppIDOperationsCreated) GetPayload() *models.InlineResponse20030 {
	return o.Payload
}

func (o *PostAppsAppIDOperationsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20030)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAppsAppIDOperationsDefault creates a PostAppsAppIDOperationsDefault with default headers values
func NewPostAppsAppIDOperationsDefault(code int) *PostAppsAppIDOperationsDefault {
	return &PostAppsAppIDOperationsDefault{
		_statusCode: code,
	}
}

/* PostAppsAppIDOperationsDefault describes a response with status code -1, with default header values.

Error response. Often a 4xx or 5xx status code
*/
type PostAppsAppIDOperationsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the post apps app ID operations default response
func (o *PostAppsAppIDOperationsDefault) Code() int {
	return o._statusCode
}

func (o *PostAppsAppIDOperationsDefault) Error() string {
	return fmt.Sprintf("[POST /apps/{app_id}/operations][%d] PostAppsAppIDOperations default  %+v", o._statusCode, o.Payload)
}
func (o *PostAppsAppIDOperationsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *PostAppsAppIDOperationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
