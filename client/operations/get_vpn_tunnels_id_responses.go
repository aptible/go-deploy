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

// GetVpnTunnelsIDReader is a Reader for the GetVpnTunnelsID structure.
type GetVpnTunnelsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVpnTunnelsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVpnTunnelsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetVpnTunnelsIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVpnTunnelsIDOK creates a GetVpnTunnelsIDOK with default headers values
func NewGetVpnTunnelsIDOK() *GetVpnTunnelsIDOK {
	return &GetVpnTunnelsIDOK{}
}

/*GetVpnTunnelsIDOK handles this case with default header values.

successful
*/
type GetVpnTunnelsIDOK struct {
	Payload *models.InlineResponse20044
}

func (o *GetVpnTunnelsIDOK) Error() string {
	return fmt.Sprintf("[GET /vpn_tunnels/{id}][%d] getVpnTunnelsIdOK  %+v", 200, o.Payload)
}

func (o *GetVpnTunnelsIDOK) GetPayload() *models.InlineResponse20044 {
	return o.Payload
}

func (o *GetVpnTunnelsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20044)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVpnTunnelsIDDefault creates a GetVpnTunnelsIDDefault with default headers values
func NewGetVpnTunnelsIDDefault(code int) *GetVpnTunnelsIDDefault {
	return &GetVpnTunnelsIDDefault{
		_statusCode: code,
	}
}

/*GetVpnTunnelsIDDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetVpnTunnelsIDDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get vpn tunnels ID default response
func (o *GetVpnTunnelsIDDefault) Code() int {
	return o._statusCode
}

func (o *GetVpnTunnelsIDDefault) Error() string {
	return fmt.Sprintf("[GET /vpn_tunnels/{id}][%d] GetVpnTunnelsID default  %+v", o._statusCode, o.Payload)
}

func (o *GetVpnTunnelsIDDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetVpnTunnelsIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
