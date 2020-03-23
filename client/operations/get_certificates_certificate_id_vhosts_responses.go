// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/reggregory/go-deploy/models"
)

// GetCertificatesCertificateIDVhostsReader is a Reader for the GetCertificatesCertificateIDVhosts structure.
type GetCertificatesCertificateIDVhostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCertificatesCertificateIDVhostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCertificatesCertificateIDVhostsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCertificatesCertificateIDVhostsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCertificatesCertificateIDVhostsOK creates a GetCertificatesCertificateIDVhostsOK with default headers values
func NewGetCertificatesCertificateIDVhostsOK() *GetCertificatesCertificateIDVhostsOK {
	return &GetCertificatesCertificateIDVhostsOK{}
}

/*GetCertificatesCertificateIDVhostsOK handles this case with default header values.

successful
*/
type GetCertificatesCertificateIDVhostsOK struct {
	Payload *models.InlineResponse20038
}

func (o *GetCertificatesCertificateIDVhostsOK) Error() string {
	return fmt.Sprintf("[GET /certificates/{certificate_id}/vhosts][%d] getCertificatesCertificateIdVhostsOK  %+v", 200, o.Payload)
}

func (o *GetCertificatesCertificateIDVhostsOK) GetPayload() *models.InlineResponse20038 {
	return o.Payload
}

func (o *GetCertificatesCertificateIDVhostsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20038)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCertificatesCertificateIDVhostsDefault creates a GetCertificatesCertificateIDVhostsDefault with default headers values
func NewGetCertificatesCertificateIDVhostsDefault(code int) *GetCertificatesCertificateIDVhostsDefault {
	return &GetCertificatesCertificateIDVhostsDefault{
		_statusCode: code,
	}
}

/*GetCertificatesCertificateIDVhostsDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetCertificatesCertificateIDVhostsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get certificates certificate ID vhosts default response
func (o *GetCertificatesCertificateIDVhostsDefault) Code() int {
	return o._statusCode
}

func (o *GetCertificatesCertificateIDVhostsDefault) Error() string {
	return fmt.Sprintf("[GET /certificates/{certificate_id}/vhosts][%d] GetCertificatesCertificateIDVhosts default  %+v", o._statusCode, o.Payload)
}

func (o *GetCertificatesCertificateIDVhostsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetCertificatesCertificateIDVhostsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
