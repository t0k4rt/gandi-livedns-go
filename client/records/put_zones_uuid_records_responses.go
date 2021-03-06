// Code generated by go-swagger; DO NOT EDIT.

package records

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/t0k4rt/gandi-livedns-go/models"
)

// PutZonesUUIDRecordsReader is a Reader for the PutZonesUUIDRecords structure.
type PutZonesUUIDRecordsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutZonesUUIDRecordsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPutZonesUUIDRecordsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutZonesUUIDRecordsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutZonesUUIDRecordsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutZonesUUIDRecordsCreated creates a PutZonesUUIDRecordsCreated with default headers values
func NewPutZonesUUIDRecordsCreated() *PutZonesUUIDRecordsCreated {
	return &PutZonesUUIDRecordsCreated{}
}

/*PutZonesUUIDRecordsCreated handles this case with default header values.

OK
*/
type PutZonesUUIDRecordsCreated struct {
	Payload *models.Return200
}

func (o *PutZonesUUIDRecordsCreated) Error() string {
	return fmt.Sprintf("[PUT /zones/{uuid}/records][%d] putZonesUuidRecordsCreated  %+v", 201, o.Payload)
}

func (o *PutZonesUUIDRecordsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Return200)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutZonesUUIDRecordsBadRequest creates a PutZonesUUIDRecordsBadRequest with default headers values
func NewPutZonesUUIDRecordsBadRequest() *PutZonesUUIDRecordsBadRequest {
	return &PutZonesUUIDRecordsBadRequest{}
}

/*PutZonesUUIDRecordsBadRequest handles this case with default header values.

Not OK
*/
type PutZonesUUIDRecordsBadRequest struct {
	Payload *models.Return400
}

func (o *PutZonesUUIDRecordsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /zones/{uuid}/records][%d] putZonesUuidRecordsBadRequest  %+v", 400, o.Payload)
}

func (o *PutZonesUUIDRecordsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Return400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutZonesUUIDRecordsDefault creates a PutZonesUUIDRecordsDefault with default headers values
func NewPutZonesUUIDRecordsDefault(code int) *PutZonesUUIDRecordsDefault {
	return &PutZonesUUIDRecordsDefault{
		_statusCode: code,
	}
}

/*PutZonesUUIDRecordsDefault handles this case with default header values.

Unexpected error
*/
type PutZonesUUIDRecordsDefault struct {
	_statusCode int

	Payload *models.Return40x
}

// Code gets the status code for the put zones UUID records default response
func (o *PutZonesUUIDRecordsDefault) Code() int {
	return o._statusCode
}

func (o *PutZonesUUIDRecordsDefault) Error() string {
	return fmt.Sprintf("[PUT /zones/{uuid}/records][%d] PutZonesUUIDRecords default  %+v", o._statusCode, o.Payload)
}

func (o *PutZonesUUIDRecordsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Return40x)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
