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

// DeleteDatabasesIDReader is a Reader for the DeleteDatabasesID structure.
type DeleteDatabasesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDatabasesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteDatabasesIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteDatabasesIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteDatabasesIDNoContent creates a DeleteDatabasesIDNoContent with default headers values
func NewDeleteDatabasesIDNoContent() *DeleteDatabasesIDNoContent {
	return &DeleteDatabasesIDNoContent{}
}

/*DeleteDatabasesIDNoContent handles this case with default header values.

successful
*/
type DeleteDatabasesIDNoContent struct {
}

func (o *DeleteDatabasesIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /databases/{id}][%d] deleteDatabasesIdNoContent ", 204)
}

func (o *DeleteDatabasesIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDatabasesIDDefault creates a DeleteDatabasesIDDefault with default headers values
func NewDeleteDatabasesIDDefault(code int) *DeleteDatabasesIDDefault {
	return &DeleteDatabasesIDDefault{
		_statusCode: code,
	}
}

/*DeleteDatabasesIDDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type DeleteDatabasesIDDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the delete databases ID default response
func (o *DeleteDatabasesIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteDatabasesIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /databases/{id}][%d] DeleteDatabasesID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDatabasesIDDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *DeleteDatabasesIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
