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

// GetStacksIDReader is a Reader for the GetStacksID structure.
type GetStacksIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStacksIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStacksIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetStacksIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStacksIDOK creates a GetStacksIDOK with default headers values
func NewGetStacksIDOK() *GetStacksIDOK {
	return &GetStacksIDOK{}
}

/*GetStacksIDOK handles this case with default header values.

successful
*/
type GetStacksIDOK struct {
	Payload *models.InlineResponse20039
}

func (o *GetStacksIDOK) Error() string {
	return fmt.Sprintf("[GET /stacks/{id}][%d] getStacksIdOK  %+v", 200, o.Payload)
}

func (o *GetStacksIDOK) GetPayload() *models.InlineResponse20039 {
	return o.Payload
}

func (o *GetStacksIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20039)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStacksIDDefault creates a GetStacksIDDefault with default headers values
func NewGetStacksIDDefault(code int) *GetStacksIDDefault {
	return &GetStacksIDDefault{
		_statusCode: code,
	}
}

/*GetStacksIDDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetStacksIDDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get stacks ID default response
func (o *GetStacksIDDefault) Code() int {
	return o._statusCode
}

func (o *GetStacksIDDefault) Error() string {
	return fmt.Sprintf("[GET /stacks/{id}][%d] GetStacksID default  %+v", o._statusCode, o.Payload)
}

func (o *GetStacksIDDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetStacksIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
