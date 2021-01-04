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

// GetEphemeralSessionsEphemeralSessionIDEphemeralContainersReader is a Reader for the GetEphemeralSessionsEphemeralSessionIDEphemeralContainers structure.
type GetEphemeralSessionsEphemeralSessionIDEphemeralContainersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEphemeralSessionsEphemeralSessionIDEphemeralContainersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEphemeralSessionsEphemeralSessionIDEphemeralContainersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetEphemeralSessionsEphemeralSessionIDEphemeralContainersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEphemeralSessionsEphemeralSessionIDEphemeralContainersOK creates a GetEphemeralSessionsEphemeralSessionIDEphemeralContainersOK with default headers values
func NewGetEphemeralSessionsEphemeralSessionIDEphemeralContainersOK() *GetEphemeralSessionsEphemeralSessionIDEphemeralContainersOK {
	return &GetEphemeralSessionsEphemeralSessionIDEphemeralContainersOK{}
}

/*GetEphemeralSessionsEphemeralSessionIDEphemeralContainersOK handles this case with default header values.

successful
*/
type GetEphemeralSessionsEphemeralSessionIDEphemeralContainersOK struct {
	Payload *models.InlineResponse20019
}

func (o *GetEphemeralSessionsEphemeralSessionIDEphemeralContainersOK) Error() string {
	return fmt.Sprintf("[GET /ephemeral_sessions/{ephemeral_session_id}/ephemeral_containers][%d] getEphemeralSessionsEphemeralSessionIdEphemeralContainersOK  %+v", 200, o.Payload)
}

func (o *GetEphemeralSessionsEphemeralSessionIDEphemeralContainersOK) GetPayload() *models.InlineResponse20019 {
	return o.Payload
}

func (o *GetEphemeralSessionsEphemeralSessionIDEphemeralContainersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20019)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEphemeralSessionsEphemeralSessionIDEphemeralContainersDefault creates a GetEphemeralSessionsEphemeralSessionIDEphemeralContainersDefault with default headers values
func NewGetEphemeralSessionsEphemeralSessionIDEphemeralContainersDefault(code int) *GetEphemeralSessionsEphemeralSessionIDEphemeralContainersDefault {
	return &GetEphemeralSessionsEphemeralSessionIDEphemeralContainersDefault{
		_statusCode: code,
	}
}

/*GetEphemeralSessionsEphemeralSessionIDEphemeralContainersDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetEphemeralSessionsEphemeralSessionIDEphemeralContainersDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get ephemeral sessions ephemeral session ID ephemeral containers default response
func (o *GetEphemeralSessionsEphemeralSessionIDEphemeralContainersDefault) Code() int {
	return o._statusCode
}

func (o *GetEphemeralSessionsEphemeralSessionIDEphemeralContainersDefault) Error() string {
	return fmt.Sprintf("[GET /ephemeral_sessions/{ephemeral_session_id}/ephemeral_containers][%d] GetEphemeralSessionsEphemeralSessionIDEphemeralContainers default  %+v", o._statusCode, o.Payload)
}

func (o *GetEphemeralSessionsEphemeralSessionIDEphemeralContainersDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetEphemeralSessionsEphemeralSessionIDEphemeralContainersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
