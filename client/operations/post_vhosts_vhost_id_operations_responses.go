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

// PostVhostsVhostIDOperationsReader is a Reader for the PostVhostsVhostIDOperations structure.
type PostVhostsVhostIDOperationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostVhostsVhostIDOperationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostVhostsVhostIDOperationsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostVhostsVhostIDOperationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostVhostsVhostIDOperationsCreated creates a PostVhostsVhostIDOperationsCreated with default headers values
func NewPostVhostsVhostIDOperationsCreated() *PostVhostsVhostIDOperationsCreated {
	return &PostVhostsVhostIDOperationsCreated{}
}

/*PostVhostsVhostIDOperationsCreated handles this case with default header values.

successful
*/
type PostVhostsVhostIDOperationsCreated struct {
	Payload *models.InlineResponse20030
}

func (o *PostVhostsVhostIDOperationsCreated) Error() string {
	return fmt.Sprintf("[POST /vhosts/{vhost_id}/operations][%d] postVhostsVhostIdOperationsCreated  %+v", 201, o.Payload)
}

func (o *PostVhostsVhostIDOperationsCreated) GetPayload() *models.InlineResponse20030 {
	return o.Payload
}

func (o *PostVhostsVhostIDOperationsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20030)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVhostsVhostIDOperationsDefault creates a PostVhostsVhostIDOperationsDefault with default headers values
func NewPostVhostsVhostIDOperationsDefault(code int) *PostVhostsVhostIDOperationsDefault {
	return &PostVhostsVhostIDOperationsDefault{
		_statusCode: code,
	}
}

/*PostVhostsVhostIDOperationsDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type PostVhostsVhostIDOperationsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the post vhosts vhost ID operations default response
func (o *PostVhostsVhostIDOperationsDefault) Code() int {
	return o._statusCode
}

func (o *PostVhostsVhostIDOperationsDefault) Error() string {
	return fmt.Sprintf("[POST /vhosts/{vhost_id}/operations][%d] PostVhostsVhostIDOperations default  %+v", o._statusCode, o.Payload)
}

func (o *PostVhostsVhostIDOperationsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *PostVhostsVhostIDOperationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
