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

// PatchOperationsIDReader is a Reader for the PatchOperationsID structure.
type PatchOperationsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchOperationsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchOperationsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchOperationsIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchOperationsIDOK creates a PatchOperationsIDOK with default headers values
func NewPatchOperationsIDOK() *PatchOperationsIDOK {
	return &PatchOperationsIDOK{}
}

/* PatchOperationsIDOK describes a response with status code 200, with default header values.

successful
*/
type PatchOperationsIDOK struct {
}

func (o *PatchOperationsIDOK) Error() string {
	return fmt.Sprintf("[PATCH /operations/{id}][%d] patchOperationsIdOK ", 200)
}

func (o *PatchOperationsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchOperationsIDDefault creates a PatchOperationsIDDefault with default headers values
func NewPatchOperationsIDDefault(code int) *PatchOperationsIDDefault {
	return &PatchOperationsIDDefault{
		_statusCode: code,
	}
}

/* PatchOperationsIDDefault describes a response with status code -1, with default header values.

Error response. Often a 4xx or 5xx status code
*/
type PatchOperationsIDDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the patch operations ID default response
func (o *PatchOperationsIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchOperationsIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /operations/{id}][%d] PatchOperationsID default  %+v", o._statusCode, o.Payload)
}
func (o *PatchOperationsIDDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *PatchOperationsIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
