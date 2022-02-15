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

// GetEphemeralContainersIDReader is a Reader for the GetEphemeralContainersID structure.
type GetEphemeralContainersIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEphemeralContainersIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEphemeralContainersIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetEphemeralContainersIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEphemeralContainersIDOK creates a GetEphemeralContainersIDOK with default headers values
func NewGetEphemeralContainersIDOK() *GetEphemeralContainersIDOK {
	return &GetEphemeralContainersIDOK{}
}

/* GetEphemeralContainersIDOK describes a response with status code 200, with default header values.

successful
*/
type GetEphemeralContainersIDOK struct {
	Payload *models.InlineResponse20020
}

func (o *GetEphemeralContainersIDOK) Error() string {
	return fmt.Sprintf("[GET /ephemeral_containers/{id}][%d] getEphemeralContainersIdOK  %+v", 200, o.Payload)
}
func (o *GetEphemeralContainersIDOK) GetPayload() *models.InlineResponse20020 {
	return o.Payload
}

func (o *GetEphemeralContainersIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20020)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEphemeralContainersIDDefault creates a GetEphemeralContainersIDDefault with default headers values
func NewGetEphemeralContainersIDDefault(code int) *GetEphemeralContainersIDDefault {
	return &GetEphemeralContainersIDDefault{
		_statusCode: code,
	}
}

/* GetEphemeralContainersIDDefault describes a response with status code -1, with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetEphemeralContainersIDDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get ephemeral containers ID default response
func (o *GetEphemeralContainersIDDefault) Code() int {
	return o._statusCode
}

func (o *GetEphemeralContainersIDDefault) Error() string {
	return fmt.Sprintf("[GET /ephemeral_containers/{id}][%d] GetEphemeralContainersID default  %+v", o._statusCode, o.Payload)
}
func (o *GetEphemeralContainersIDDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetEphemeralContainersIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
