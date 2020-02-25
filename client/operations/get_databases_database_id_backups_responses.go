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

// GetDatabasesDatabaseIDBackupsReader is a Reader for the GetDatabasesDatabaseIDBackups structure.
type GetDatabasesDatabaseIDBackupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDatabasesDatabaseIDBackupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDatabasesDatabaseIDBackupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDatabasesDatabaseIDBackupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDatabasesDatabaseIDBackupsOK creates a GetDatabasesDatabaseIDBackupsOK with default headers values
func NewGetDatabasesDatabaseIDBackupsOK() *GetDatabasesDatabaseIDBackupsOK {
	return &GetDatabasesDatabaseIDBackupsOK{}
}

/*GetDatabasesDatabaseIDBackupsOK handles this case with default header values.

successful
*/
type GetDatabasesDatabaseIDBackupsOK struct {
	Payload *models.InlineResponse2004
}

func (o *GetDatabasesDatabaseIDBackupsOK) Error() string {
	return fmt.Sprintf("[GET /databases/{database_id}/backups][%d] getDatabasesDatabaseIdBackupsOK  %+v", 200, o.Payload)
}

func (o *GetDatabasesDatabaseIDBackupsOK) GetPayload() *models.InlineResponse2004 {
	return o.Payload
}

func (o *GetDatabasesDatabaseIDBackupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse2004)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatabasesDatabaseIDBackupsDefault creates a GetDatabasesDatabaseIDBackupsDefault with default headers values
func NewGetDatabasesDatabaseIDBackupsDefault(code int) *GetDatabasesDatabaseIDBackupsDefault {
	return &GetDatabasesDatabaseIDBackupsDefault{
		_statusCode: code,
	}
}

/*GetDatabasesDatabaseIDBackupsDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetDatabasesDatabaseIDBackupsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get databases database ID backups default response
func (o *GetDatabasesDatabaseIDBackupsDefault) Code() int {
	return o._statusCode
}

func (o *GetDatabasesDatabaseIDBackupsDefault) Error() string {
	return fmt.Sprintf("[GET /databases/{database_id}/backups][%d] GetDatabasesDatabaseIDBackups default  %+v", o._statusCode, o.Payload)
}

func (o *GetDatabasesDatabaseIDBackupsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetDatabasesDatabaseIDBackupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
