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

// GetPermissionsReader is a Reader for the GetPermissions structure.
type GetPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetPermissionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPermissionsOK creates a GetPermissionsOK with default headers values
func NewGetPermissionsOK() *GetPermissionsOK {
	return &GetPermissionsOK{}
}

/*GetPermissionsOK handles this case with default header values.

successful
*/
type GetPermissionsOK struct {
	Payload *models.InlineResponse20030
}

func (o *GetPermissionsOK) Error() string {
	return fmt.Sprintf("[GET /permissions][%d] getPermissionsOK  %+v", 200, o.Payload)
}

func (o *GetPermissionsOK) GetPayload() *models.InlineResponse20030 {
	return o.Payload
}

func (o *GetPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20030)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPermissionsDefault creates a GetPermissionsDefault with default headers values
func NewGetPermissionsDefault(code int) *GetPermissionsDefault {
	return &GetPermissionsDefault{
		_statusCode: code,
	}
}

/*GetPermissionsDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetPermissionsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get permissions default response
func (o *GetPermissionsDefault) Code() int {
	return o._statusCode
}

func (o *GetPermissionsDefault) Error() string {
	return fmt.Sprintf("[GET /permissions][%d] GetPermissions default  %+v", o._statusCode, o.Payload)
}

func (o *GetPermissionsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetPermissionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
