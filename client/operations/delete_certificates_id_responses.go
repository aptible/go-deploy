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

// DeleteCertificatesIDReader is a Reader for the DeleteCertificatesID structure.
type DeleteCertificatesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCertificatesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteCertificatesIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteCertificatesIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCertificatesIDNoContent creates a DeleteCertificatesIDNoContent with default headers values
func NewDeleteCertificatesIDNoContent() *DeleteCertificatesIDNoContent {
	return &DeleteCertificatesIDNoContent{}
}

/*DeleteCertificatesIDNoContent handles this case with default header values.

successful
*/
type DeleteCertificatesIDNoContent struct {
}

func (o *DeleteCertificatesIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /certificates/{id}][%d] deleteCertificatesIdNoContent ", 204)
}

func (o *DeleteCertificatesIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCertificatesIDDefault creates a DeleteCertificatesIDDefault with default headers values
func NewDeleteCertificatesIDDefault(code int) *DeleteCertificatesIDDefault {
	return &DeleteCertificatesIDDefault{
		_statusCode: code,
	}
}

/*DeleteCertificatesIDDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type DeleteCertificatesIDDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the delete certificates ID default response
func (o *DeleteCertificatesIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCertificatesIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /certificates/{id}][%d] DeleteCertificatesID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCertificatesIDDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *DeleteCertificatesIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
