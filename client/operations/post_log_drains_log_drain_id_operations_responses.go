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

// PostLogDrainsLogDrainIDOperationsReader is a Reader for the PostLogDrainsLogDrainIDOperations structure.
type PostLogDrainsLogDrainIDOperationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLogDrainsLogDrainIDOperationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostLogDrainsLogDrainIDOperationsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostLogDrainsLogDrainIDOperationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostLogDrainsLogDrainIDOperationsCreated creates a PostLogDrainsLogDrainIDOperationsCreated with default headers values
func NewPostLogDrainsLogDrainIDOperationsCreated() *PostLogDrainsLogDrainIDOperationsCreated {
	return &PostLogDrainsLogDrainIDOperationsCreated{}
}

/* PostLogDrainsLogDrainIDOperationsCreated describes a response with status code 201, with default header values.

successful
*/
type PostLogDrainsLogDrainIDOperationsCreated struct {
	Payload *models.InlineResponse20030
}

func (o *PostLogDrainsLogDrainIDOperationsCreated) Error() string {
	return fmt.Sprintf("[POST /log_drains/{log_drain_id}/operations][%d] postLogDrainsLogDrainIdOperationsCreated  %+v", 201, o.Payload)
}
func (o *PostLogDrainsLogDrainIDOperationsCreated) GetPayload() *models.InlineResponse20030 {
	return o.Payload
}

func (o *PostLogDrainsLogDrainIDOperationsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20030)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLogDrainsLogDrainIDOperationsDefault creates a PostLogDrainsLogDrainIDOperationsDefault with default headers values
func NewPostLogDrainsLogDrainIDOperationsDefault(code int) *PostLogDrainsLogDrainIDOperationsDefault {
	return &PostLogDrainsLogDrainIDOperationsDefault{
		_statusCode: code,
	}
}

/* PostLogDrainsLogDrainIDOperationsDefault describes a response with status code -1, with default header values.

Error response. Often a 4xx or 5xx status code
*/
type PostLogDrainsLogDrainIDOperationsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the post log drains log drain ID operations default response
func (o *PostLogDrainsLogDrainIDOperationsDefault) Code() int {
	return o._statusCode
}

func (o *PostLogDrainsLogDrainIDOperationsDefault) Error() string {
	return fmt.Sprintf("[POST /log_drains/{log_drain_id}/operations][%d] PostLogDrainsLogDrainIDOperations default  %+v", o._statusCode, o.Payload)
}
func (o *PostLogDrainsLogDrainIDOperationsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *PostLogDrainsLogDrainIDOperationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
