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

// GetStacksReader is a Reader for the GetStacks structure.
type GetStacksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStacksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStacksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetStacksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStacksOK creates a GetStacksOK with default headers values
func NewGetStacksOK() *GetStacksOK {
	return &GetStacksOK{}
}

/*GetStacksOK handles this case with default header values.

successful
*/
type GetStacksOK struct {
	Payload *models.InlineResponse20038
}

func (o *GetStacksOK) Error() string {
	return fmt.Sprintf("[GET /stacks][%d] getStacksOK  %+v", 200, o.Payload)
}

func (o *GetStacksOK) GetPayload() *models.InlineResponse20038 {
	return o.Payload
}

func (o *GetStacksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20038)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStacksDefault creates a GetStacksDefault with default headers values
func NewGetStacksDefault(code int) *GetStacksDefault {
	return &GetStacksDefault{
		_statusCode: code,
	}
}

/*GetStacksDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetStacksDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get stacks default response
func (o *GetStacksDefault) Code() int {
	return o._statusCode
}

func (o *GetStacksDefault) Error() string {
	return fmt.Sprintf("[GET /stacks][%d] GetStacks default  %+v", o._statusCode, o.Payload)
}

func (o *GetStacksDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetStacksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
