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

// GetIntrusionDetectionReportsIntrusionReportIDDownloadCsvReader is a Reader for the GetIntrusionDetectionReportsIntrusionReportIDDownloadCsv structure.
type GetIntrusionDetectionReportsIntrusionReportIDDownloadCsvReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntrusionDetectionReportsIntrusionReportIDDownloadCsvReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntrusionDetectionReportsIntrusionReportIDDownloadCsvOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetIntrusionDetectionReportsIntrusionReportIDDownloadCsvDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIntrusionDetectionReportsIntrusionReportIDDownloadCsvOK creates a GetIntrusionDetectionReportsIntrusionReportIDDownloadCsvOK with default headers values
func NewGetIntrusionDetectionReportsIntrusionReportIDDownloadCsvOK() *GetIntrusionDetectionReportsIntrusionReportIDDownloadCsvOK {
	return &GetIntrusionDetectionReportsIntrusionReportIDDownloadCsvOK{}
}

/* GetIntrusionDetectionReportsIntrusionReportIDDownloadCsvOK describes a response with status code 200, with default header values.

Presigned URL
*/
type GetIntrusionDetectionReportsIntrusionReportIDDownloadCsvOK struct {
}

func (o *GetIntrusionDetectionReportsIntrusionReportIDDownloadCsvOK) Error() string {
	return fmt.Sprintf("[GET /intrusion_detection_reports/{intrusion_report_id}/download_csv][%d] getIntrusionDetectionReportsIntrusionReportIdDownloadCsvOK ", 200)
}

func (o *GetIntrusionDetectionReportsIntrusionReportIDDownloadCsvOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIntrusionDetectionReportsIntrusionReportIDDownloadCsvDefault creates a GetIntrusionDetectionReportsIntrusionReportIDDownloadCsvDefault with default headers values
func NewGetIntrusionDetectionReportsIntrusionReportIDDownloadCsvDefault(code int) *GetIntrusionDetectionReportsIntrusionReportIDDownloadCsvDefault {
	return &GetIntrusionDetectionReportsIntrusionReportIDDownloadCsvDefault{
		_statusCode: code,
	}
}

/* GetIntrusionDetectionReportsIntrusionReportIDDownloadCsvDefault describes a response with status code -1, with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetIntrusionDetectionReportsIntrusionReportIDDownloadCsvDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get intrusion detection reports intrusion report ID download csv default response
func (o *GetIntrusionDetectionReportsIntrusionReportIDDownloadCsvDefault) Code() int {
	return o._statusCode
}

func (o *GetIntrusionDetectionReportsIntrusionReportIDDownloadCsvDefault) Error() string {
	return fmt.Sprintf("[GET /intrusion_detection_reports/{intrusion_report_id}/download_csv][%d] GetIntrusionDetectionReportsIntrusionReportIDDownloadCsv default  %+v", o._statusCode, o.Payload)
}
func (o *GetIntrusionDetectionReportsIntrusionReportIDDownloadCsvDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetIntrusionDetectionReportsIntrusionReportIDDownloadCsvDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
