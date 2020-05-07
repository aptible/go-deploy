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

// GetCertificatesIDReader is a Reader for the GetCertificatesID structure.
type GetCertificatesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCertificatesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCertificatesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCertificatesIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCertificatesIDOK creates a GetCertificatesIDOK with default headers values
func NewGetCertificatesIDOK() *GetCertificatesIDOK {
	return &GetCertificatesIDOK{}
}

/*GetCertificatesIDOK handles this case with default header values.

successful
*/
type GetCertificatesIDOK struct {
	Payload *models.InlineResponse2012
}

func (o *GetCertificatesIDOK) Error() string {
	return fmt.Sprintf("[GET /certificates/{id}][%d] getCertificatesIdOK  %+v", 200, o.Payload)
}

func (o *GetCertificatesIDOK) GetPayload() *models.InlineResponse2012 {
	return o.Payload
}

func (o *GetCertificatesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse2012)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCertificatesIDDefault creates a GetCertificatesIDDefault with default headers values
func NewGetCertificatesIDDefault(code int) *GetCertificatesIDDefault {
	return &GetCertificatesIDDefault{
		_statusCode: code,
	}
}

/*GetCertificatesIDDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetCertificatesIDDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get certificates ID default response
func (o *GetCertificatesIDDefault) Code() int {
	return o._statusCode
}

func (o *GetCertificatesIDDefault) Error() string {
	return fmt.Sprintf("[GET /certificates/{id}][%d] GetCertificatesID default  %+v", o._statusCode, o.Payload)
}

func (o *GetCertificatesIDDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetCertificatesIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
