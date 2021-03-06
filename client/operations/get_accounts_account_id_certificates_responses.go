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

// GetAccountsAccountIDCertificatesReader is a Reader for the GetAccountsAccountIDCertificates structure.
type GetAccountsAccountIDCertificatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountsAccountIDCertificatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountsAccountIDCertificatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAccountsAccountIDCertificatesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAccountsAccountIDCertificatesOK creates a GetAccountsAccountIDCertificatesOK with default headers values
func NewGetAccountsAccountIDCertificatesOK() *GetAccountsAccountIDCertificatesOK {
	return &GetAccountsAccountIDCertificatesOK{}
}

/*GetAccountsAccountIDCertificatesOK handles this case with default header values.

successful
*/
type GetAccountsAccountIDCertificatesOK struct {
	Payload *models.InlineResponse2008
}

func (o *GetAccountsAccountIDCertificatesOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{account_id}/certificates][%d] getAccountsAccountIdCertificatesOK  %+v", 200, o.Payload)
}

func (o *GetAccountsAccountIDCertificatesOK) GetPayload() *models.InlineResponse2008 {
	return o.Payload
}

func (o *GetAccountsAccountIDCertificatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse2008)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsAccountIDCertificatesDefault creates a GetAccountsAccountIDCertificatesDefault with default headers values
func NewGetAccountsAccountIDCertificatesDefault(code int) *GetAccountsAccountIDCertificatesDefault {
	return &GetAccountsAccountIDCertificatesDefault{
		_statusCode: code,
	}
}

/*GetAccountsAccountIDCertificatesDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetAccountsAccountIDCertificatesDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get accounts account ID certificates default response
func (o *GetAccountsAccountIDCertificatesDefault) Code() int {
	return o._statusCode
}

func (o *GetAccountsAccountIDCertificatesDefault) Error() string {
	return fmt.Sprintf("[GET /accounts/{account_id}/certificates][%d] GetAccountsAccountIDCertificates default  %+v", o._statusCode, o.Payload)
}

func (o *GetAccountsAccountIDCertificatesDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetAccountsAccountIDCertificatesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
