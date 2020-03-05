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

// PutOperationsIDReader is a Reader for the PutOperationsID structure.
type PutOperationsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOperationsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOperationsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutOperationsIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutOperationsIDOK creates a PutOperationsIDOK with default headers values
func NewPutOperationsIDOK() *PutOperationsIDOK {
	return &PutOperationsIDOK{}
}

/*PutOperationsIDOK handles this case with default header values.

successful
*/
type PutOperationsIDOK struct {
}

func (o *PutOperationsIDOK) Error() string {
	return fmt.Sprintf("[PUT /operations/{id}][%d] putOperationsIdOK ", 200)
}

func (o *PutOperationsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutOperationsIDDefault creates a PutOperationsIDDefault with default headers values
func NewPutOperationsIDDefault(code int) *PutOperationsIDDefault {
	return &PutOperationsIDDefault{
		_statusCode: code,
	}
}

/*PutOperationsIDDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type PutOperationsIDDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the put operations ID default response
func (o *PutOperationsIDDefault) Code() int {
	return o._statusCode
}

func (o *PutOperationsIDDefault) Error() string {
	return fmt.Sprintf("[PUT /operations/{id}][%d] PutOperationsID default  %+v", o._statusCode, o.Payload)
}

func (o *PutOperationsIDDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *PutOperationsIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
