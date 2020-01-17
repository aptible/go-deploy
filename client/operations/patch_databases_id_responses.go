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

// PatchDatabasesIDReader is a Reader for the PatchDatabasesID structure.
type PatchDatabasesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchDatabasesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchDatabasesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchDatabasesIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchDatabasesIDOK creates a PatchDatabasesIDOK with default headers values
func NewPatchDatabasesIDOK() *PatchDatabasesIDOK {
	return &PatchDatabasesIDOK{}
}

/*PatchDatabasesIDOK handles this case with default header values.

successful
*/
type PatchDatabasesIDOK struct {
}

func (o *PatchDatabasesIDOK) Error() string {
	return fmt.Sprintf("[PATCH /databases/{id}][%d] patchDatabasesIdOK ", 200)
}

func (o *PatchDatabasesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchDatabasesIDDefault creates a PatchDatabasesIDDefault with default headers values
func NewPatchDatabasesIDDefault(code int) *PatchDatabasesIDDefault {
	return &PatchDatabasesIDDefault{
		_statusCode: code,
	}
}

/*PatchDatabasesIDDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type PatchDatabasesIDDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the patch databases ID default response
func (o *PatchDatabasesIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchDatabasesIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /databases/{id}][%d] PatchDatabasesID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchDatabasesIDDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *PatchDatabasesIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
