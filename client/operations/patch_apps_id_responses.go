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

// PatchAppsIDReader is a Reader for the PatchAppsID structure.
type PatchAppsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAppsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAppsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchAppsIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchAppsIDOK creates a PatchAppsIDOK with default headers values
func NewPatchAppsIDOK() *PatchAppsIDOK {
	return &PatchAppsIDOK{}
}

/*PatchAppsIDOK handles this case with default header values.

successful
*/
type PatchAppsIDOK struct {
}

func (o *PatchAppsIDOK) Error() string {
	return fmt.Sprintf("[PATCH /apps/{id}][%d] patchAppsIdOK ", 200)
}

func (o *PatchAppsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchAppsIDDefault creates a PatchAppsIDDefault with default headers values
func NewPatchAppsIDDefault(code int) *PatchAppsIDDefault {
	return &PatchAppsIDDefault{
		_statusCode: code,
	}
}

/*PatchAppsIDDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type PatchAppsIDDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the patch apps ID default response
func (o *PatchAppsIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchAppsIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /apps/{id}][%d] PatchAppsID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchAppsIDDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *PatchAppsIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
