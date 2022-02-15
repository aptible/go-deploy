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

// GetMetricDrainsMetricDrainIDOperationsReader is a Reader for the GetMetricDrainsMetricDrainIDOperations structure.
type GetMetricDrainsMetricDrainIDOperationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMetricDrainsMetricDrainIDOperationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMetricDrainsMetricDrainIDOperationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMetricDrainsMetricDrainIDOperationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMetricDrainsMetricDrainIDOperationsOK creates a GetMetricDrainsMetricDrainIDOperationsOK with default headers values
func NewGetMetricDrainsMetricDrainIDOperationsOK() *GetMetricDrainsMetricDrainIDOperationsOK {
	return &GetMetricDrainsMetricDrainIDOperationsOK{}
}

/* GetMetricDrainsMetricDrainIDOperationsOK describes a response with status code 200, with default header values.

successful
*/
type GetMetricDrainsMetricDrainIDOperationsOK struct {
	Payload *models.InlineResponse20031
}

func (o *GetMetricDrainsMetricDrainIDOperationsOK) Error() string {
	return fmt.Sprintf("[GET /metric_drains/{metric_drain_id}/operations][%d] getMetricDrainsMetricDrainIdOperationsOK  %+v", 200, o.Payload)
}
func (o *GetMetricDrainsMetricDrainIDOperationsOK) GetPayload() *models.InlineResponse20031 {
	return o.Payload
}

func (o *GetMetricDrainsMetricDrainIDOperationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20031)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMetricDrainsMetricDrainIDOperationsDefault creates a GetMetricDrainsMetricDrainIDOperationsDefault with default headers values
func NewGetMetricDrainsMetricDrainIDOperationsDefault(code int) *GetMetricDrainsMetricDrainIDOperationsDefault {
	return &GetMetricDrainsMetricDrainIDOperationsDefault{
		_statusCode: code,
	}
}

/* GetMetricDrainsMetricDrainIDOperationsDefault describes a response with status code -1, with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetMetricDrainsMetricDrainIDOperationsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get metric drains metric drain ID operations default response
func (o *GetMetricDrainsMetricDrainIDOperationsDefault) Code() int {
	return o._statusCode
}

func (o *GetMetricDrainsMetricDrainIDOperationsDefault) Error() string {
	return fmt.Sprintf("[GET /metric_drains/{metric_drain_id}/operations][%d] GetMetricDrainsMetricDrainIDOperations default  %+v", o._statusCode, o.Payload)
}
func (o *GetMetricDrainsMetricDrainIDOperationsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetMetricDrainsMetricDrainIDOperationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
