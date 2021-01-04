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

// GetStacksStackIDIntrusionDetectionReportsReader is a Reader for the GetStacksStackIDIntrusionDetectionReports structure.
type GetStacksStackIDIntrusionDetectionReportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStacksStackIDIntrusionDetectionReportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStacksStackIDIntrusionDetectionReportsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetStacksStackIDIntrusionDetectionReportsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStacksStackIDIntrusionDetectionReportsOK creates a GetStacksStackIDIntrusionDetectionReportsOK with default headers values
func NewGetStacksStackIDIntrusionDetectionReportsOK() *GetStacksStackIDIntrusionDetectionReportsOK {
	return &GetStacksStackIDIntrusionDetectionReportsOK{}
}

/*GetStacksStackIDIntrusionDetectionReportsOK handles this case with default header values.

successful
*/
type GetStacksStackIDIntrusionDetectionReportsOK struct {
	Payload *models.InlineResponse20026
}

func (o *GetStacksStackIDIntrusionDetectionReportsOK) Error() string {
	return fmt.Sprintf("[GET /stacks/{stack_id}/intrusion_detection_reports][%d] getStacksStackIdIntrusionDetectionReportsOK  %+v", 200, o.Payload)
}

func (o *GetStacksStackIDIntrusionDetectionReportsOK) GetPayload() *models.InlineResponse20026 {
	return o.Payload
}

func (o *GetStacksStackIDIntrusionDetectionReportsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20026)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStacksStackIDIntrusionDetectionReportsDefault creates a GetStacksStackIDIntrusionDetectionReportsDefault with default headers values
func NewGetStacksStackIDIntrusionDetectionReportsDefault(code int) *GetStacksStackIDIntrusionDetectionReportsDefault {
	return &GetStacksStackIDIntrusionDetectionReportsDefault{
		_statusCode: code,
	}
}

/*GetStacksStackIDIntrusionDetectionReportsDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetStacksStackIDIntrusionDetectionReportsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get stacks stack ID intrusion detection reports default response
func (o *GetStacksStackIDIntrusionDetectionReportsDefault) Code() int {
	return o._statusCode
}

func (o *GetStacksStackIDIntrusionDetectionReportsDefault) Error() string {
	return fmt.Sprintf("[GET /stacks/{stack_id}/intrusion_detection_reports][%d] GetStacksStackIDIntrusionDetectionReports default  %+v", o._statusCode, o.Payload)
}

func (o *GetStacksStackIDIntrusionDetectionReportsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetStacksStackIDIntrusionDetectionReportsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
