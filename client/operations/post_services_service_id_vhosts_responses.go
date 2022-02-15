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

// PostServicesServiceIDVhostsReader is a Reader for the PostServicesServiceIDVhosts structure.
type PostServicesServiceIDVhostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostServicesServiceIDVhostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostServicesServiceIDVhostsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostServicesServiceIDVhostsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostServicesServiceIDVhostsCreated creates a PostServicesServiceIDVhostsCreated with default headers values
func NewPostServicesServiceIDVhostsCreated() *PostServicesServiceIDVhostsCreated {
	return &PostServicesServiceIDVhostsCreated{}
}

/* PostServicesServiceIDVhostsCreated describes a response with status code 201, with default header values.

successful
*/
type PostServicesServiceIDVhostsCreated struct {
	Payload *models.InlineResponse2019
}

func (o *PostServicesServiceIDVhostsCreated) Error() string {
	return fmt.Sprintf("[POST /services/{service_id}/vhosts][%d] postServicesServiceIdVhostsCreated  %+v", 201, o.Payload)
}
func (o *PostServicesServiceIDVhostsCreated) GetPayload() *models.InlineResponse2019 {
	return o.Payload
}

func (o *PostServicesServiceIDVhostsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse2019)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostServicesServiceIDVhostsDefault creates a PostServicesServiceIDVhostsDefault with default headers values
func NewPostServicesServiceIDVhostsDefault(code int) *PostServicesServiceIDVhostsDefault {
	return &PostServicesServiceIDVhostsDefault{
		_statusCode: code,
	}
}

/* PostServicesServiceIDVhostsDefault describes a response with status code -1, with default header values.

Error response. Often a 4xx or 5xx status code
*/
type PostServicesServiceIDVhostsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the post services service ID vhosts default response
func (o *PostServicesServiceIDVhostsDefault) Code() int {
	return o._statusCode
}

func (o *PostServicesServiceIDVhostsDefault) Error() string {
	return fmt.Sprintf("[POST /services/{service_id}/vhosts][%d] PostServicesServiceIDVhosts default  %+v", o._statusCode, o.Payload)
}
func (o *PostServicesServiceIDVhostsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *PostServicesServiceIDVhostsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
