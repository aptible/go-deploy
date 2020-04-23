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

// PostAccountsAccountIDLogDrainsReader is a Reader for the PostAccountsAccountIDLogDrains structure.
type PostAccountsAccountIDLogDrainsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAccountsAccountIDLogDrainsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostAccountsAccountIDLogDrainsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostAccountsAccountIDLogDrainsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostAccountsAccountIDLogDrainsCreated creates a PostAccountsAccountIDLogDrainsCreated with default headers values
func NewPostAccountsAccountIDLogDrainsCreated() *PostAccountsAccountIDLogDrainsCreated {
	return &PostAccountsAccountIDLogDrainsCreated{}
}

/*PostAccountsAccountIDLogDrainsCreated handles this case with default header values.

successful
*/
type PostAccountsAccountIDLogDrainsCreated struct {
	Payload *models.InlineResponse2015
}

func (o *PostAccountsAccountIDLogDrainsCreated) Error() string {
	return fmt.Sprintf("[POST /accounts/{account_id}/log_drains][%d] postAccountsAccountIdLogDrainsCreated  %+v", 201, o.Payload)
}

func (o *PostAccountsAccountIDLogDrainsCreated) GetPayload() *models.InlineResponse2015 {
	return o.Payload
}

func (o *PostAccountsAccountIDLogDrainsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse2015)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAccountsAccountIDLogDrainsDefault creates a PostAccountsAccountIDLogDrainsDefault with default headers values
func NewPostAccountsAccountIDLogDrainsDefault(code int) *PostAccountsAccountIDLogDrainsDefault {
	return &PostAccountsAccountIDLogDrainsDefault{
		_statusCode: code,
	}
}

/*PostAccountsAccountIDLogDrainsDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type PostAccountsAccountIDLogDrainsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the post accounts account ID log drains default response
func (o *PostAccountsAccountIDLogDrainsDefault) Code() int {
	return o._statusCode
}

func (o *PostAccountsAccountIDLogDrainsDefault) Error() string {
	return fmt.Sprintf("[POST /accounts/{account_id}/log_drains][%d] PostAccountsAccountIDLogDrains default  %+v", o._statusCode, o.Payload)
}

func (o *PostAccountsAccountIDLogDrainsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *PostAccountsAccountIDLogDrainsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
