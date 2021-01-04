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

// GetServicesServiceIDOperationsReader is a Reader for the GetServicesServiceIDOperations structure.
type GetServicesServiceIDOperationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServicesServiceIDOperationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServicesServiceIDOperationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetServicesServiceIDOperationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServicesServiceIDOperationsOK creates a GetServicesServiceIDOperationsOK with default headers values
func NewGetServicesServiceIDOperationsOK() *GetServicesServiceIDOperationsOK {
	return &GetServicesServiceIDOperationsOK{}
}

/*GetServicesServiceIDOperationsOK handles this case with default header values.

successful
*/
type GetServicesServiceIDOperationsOK struct {
	Payload *models.InlineResponse20031
}

func (o *GetServicesServiceIDOperationsOK) Error() string {
	return fmt.Sprintf("[GET /services/{service_id}/operations][%d] getServicesServiceIdOperationsOK  %+v", 200, o.Payload)
}

func (o *GetServicesServiceIDOperationsOK) GetPayload() *models.InlineResponse20031 {
	return o.Payload
}

func (o *GetServicesServiceIDOperationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20031)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServicesServiceIDOperationsDefault creates a GetServicesServiceIDOperationsDefault with default headers values
func NewGetServicesServiceIDOperationsDefault(code int) *GetServicesServiceIDOperationsDefault {
	return &GetServicesServiceIDOperationsDefault{
		_statusCode: code,
	}
}

/*GetServicesServiceIDOperationsDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetServicesServiceIDOperationsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get services service ID operations default response
func (o *GetServicesServiceIDOperationsDefault) Code() int {
	return o._statusCode
}

func (o *GetServicesServiceIDOperationsDefault) Error() string {
	return fmt.Sprintf("[GET /services/{service_id}/operations][%d] GetServicesServiceIDOperations default  %+v", o._statusCode, o.Payload)
}

func (o *GetServicesServiceIDOperationsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetServicesServiceIDOperationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
