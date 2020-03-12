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

// GetDatabasesIDReader is a Reader for the GetDatabasesID structure.
type GetDatabasesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDatabasesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDatabasesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDatabasesIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDatabasesIDOK creates a GetDatabasesIDOK with default headers values
func NewGetDatabasesIDOK() *GetDatabasesIDOK {
	return &GetDatabasesIDOK{}
}

/*GetDatabasesIDOK handles this case with default header values.

successful
*/
type GetDatabasesIDOK struct {
	Payload *models.InlineResponse2014
}

func (o *GetDatabasesIDOK) Error() string {
	return fmt.Sprintf("[GET /databases/{id}][%d] getDatabasesIdOK  %+v", 200, o.Payload)
}

func (o *GetDatabasesIDOK) GetPayload() *models.InlineResponse2014 {
	return o.Payload
}

func (o *GetDatabasesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse2014)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatabasesIDDefault creates a GetDatabasesIDDefault with default headers values
func NewGetDatabasesIDDefault(code int) *GetDatabasesIDDefault {
	return &GetDatabasesIDDefault{
		_statusCode: code,
	}
}

/*GetDatabasesIDDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetDatabasesIDDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get databases ID default response
func (o *GetDatabasesIDDefault) Code() int {
	return o._statusCode
}

func (o *GetDatabasesIDDefault) Error() string {
	return fmt.Sprintf("[GET /databases/{id}][%d] GetDatabasesID default  %+v", o._statusCode, o.Payload)
}

func (o *GetDatabasesIDDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetDatabasesIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
