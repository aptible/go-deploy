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

// GetVpcPeersIDReader is a Reader for the GetVpcPeersID structure.
type GetVpcPeersIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVpcPeersIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVpcPeersIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetVpcPeersIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVpcPeersIDOK creates a GetVpcPeersIDOK with default headers values
func NewGetVpcPeersIDOK() *GetVpcPeersIDOK {
	return &GetVpcPeersIDOK{}
}

/* GetVpcPeersIDOK describes a response with status code 200, with default header values.

successful
*/
type GetVpcPeersIDOK struct {
	Payload *models.InlineResponse20042
}

func (o *GetVpcPeersIDOK) Error() string {
	return fmt.Sprintf("[GET /vpc_peers/{id}][%d] getVpcPeersIdOK  %+v", 200, o.Payload)
}
func (o *GetVpcPeersIDOK) GetPayload() *models.InlineResponse20042 {
	return o.Payload
}

func (o *GetVpcPeersIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20042)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVpcPeersIDDefault creates a GetVpcPeersIDDefault with default headers values
func NewGetVpcPeersIDDefault(code int) *GetVpcPeersIDDefault {
	return &GetVpcPeersIDDefault{
		_statusCode: code,
	}
}

/* GetVpcPeersIDDefault describes a response with status code -1, with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetVpcPeersIDDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get vpc peers ID default response
func (o *GetVpcPeersIDDefault) Code() int {
	return o._statusCode
}

func (o *GetVpcPeersIDDefault) Error() string {
	return fmt.Sprintf("[GET /vpc_peers/{id}][%d] GetVpcPeersID default  %+v", o._statusCode, o.Payload)
}
func (o *GetVpcPeersIDDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetVpcPeersIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
