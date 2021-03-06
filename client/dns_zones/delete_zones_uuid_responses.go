// Code generated by go-swagger; DO NOT EDIT.

package dns_zones

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/t0k4rt/gandi-livedns-go/models"
)

// DeleteZonesUUIDReader is a Reader for the DeleteZonesUUID structure.
type DeleteZonesUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteZonesUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteZonesUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteZonesUUIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteZonesUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteZonesUUIDOK creates a DeleteZonesUUIDOK with default headers values
func NewDeleteZonesUUIDOK() *DeleteZonesUUIDOK {
	return &DeleteZonesUUIDOK{}
}

/*DeleteZonesUUIDOK handles this case with default header values.

OK
*/
type DeleteZonesUUIDOK struct {
	Payload *models.Zone
}

func (o *DeleteZonesUUIDOK) Error() string {
	return fmt.Sprintf("[DELETE /zones/{uuid}][%d] deleteZonesUuidOK  %+v", 200, o.Payload)
}

func (o *DeleteZonesUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Zone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteZonesUUIDBadRequest creates a DeleteZonesUUIDBadRequest with default headers values
func NewDeleteZonesUUIDBadRequest() *DeleteZonesUUIDBadRequest {
	return &DeleteZonesUUIDBadRequest{}
}

/*DeleteZonesUUIDBadRequest handles this case with default header values.

Not OK
*/
type DeleteZonesUUIDBadRequest struct {
	Payload *models.Return400
}

func (o *DeleteZonesUUIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /zones/{uuid}][%d] deleteZonesUuidBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteZonesUUIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Return400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteZonesUUIDDefault creates a DeleteZonesUUIDDefault with default headers values
func NewDeleteZonesUUIDDefault(code int) *DeleteZonesUUIDDefault {
	return &DeleteZonesUUIDDefault{
		_statusCode: code,
	}
}

/*DeleteZonesUUIDDefault handles this case with default header values.

Unexpected error
*/
type DeleteZonesUUIDDefault struct {
	_statusCode int

	Payload *models.Return40x
}

// Code gets the status code for the delete zones UUID default response
func (o *DeleteZonesUUIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteZonesUUIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /zones/{uuid}][%d] DeleteZonesUUID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteZonesUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Return40x)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
