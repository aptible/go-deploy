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

// GetDatabaseImagesIDReader is a Reader for the GetDatabaseImagesID structure.
type GetDatabaseImagesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDatabaseImagesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDatabaseImagesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDatabaseImagesIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDatabaseImagesIDOK creates a GetDatabaseImagesIDOK with default headers values
func NewGetDatabaseImagesIDOK() *GetDatabaseImagesIDOK {
	return &GetDatabaseImagesIDOK{}
}

/* GetDatabaseImagesIDOK describes a response with status code 200, with default header values.

successful
*/
type GetDatabaseImagesIDOK struct {
	Payload *models.InlineResponse20014
}

func (o *GetDatabaseImagesIDOK) Error() string {
	return fmt.Sprintf("[GET /database_images/{id}][%d] getDatabaseImagesIdOK  %+v", 200, o.Payload)
}
func (o *GetDatabaseImagesIDOK) GetPayload() *models.InlineResponse20014 {
	return o.Payload
}

func (o *GetDatabaseImagesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20014)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatabaseImagesIDDefault creates a GetDatabaseImagesIDDefault with default headers values
func NewGetDatabaseImagesIDDefault(code int) *GetDatabaseImagesIDDefault {
	return &GetDatabaseImagesIDDefault{
		_statusCode: code,
	}
}

/* GetDatabaseImagesIDDefault describes a response with status code -1, with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetDatabaseImagesIDDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get database images ID default response
func (o *GetDatabaseImagesIDDefault) Code() int {
	return o._statusCode
}

func (o *GetDatabaseImagesIDDefault) Error() string {
	return fmt.Sprintf("[GET /database_images/{id}][%d] GetDatabaseImagesID default  %+v", o._statusCode, o.Payload)
}
func (o *GetDatabaseImagesIDDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetDatabaseImagesIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
