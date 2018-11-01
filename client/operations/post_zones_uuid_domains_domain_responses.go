// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/t0k4rt/gandi-livedns-go/models"
)

// PostZonesUUIDDomainsDomainReader is a Reader for the PostZonesUUIDDomainsDomain structure.
type PostZonesUUIDDomainsDomainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostZonesUUIDDomainsDomainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostZonesUUIDDomainsDomainOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostZonesUUIDDomainsDomainOK creates a PostZonesUUIDDomainsDomainOK with default headers values
func NewPostZonesUUIDDomainsDomainOK() *PostZonesUUIDDomainsDomainOK {
	return &PostZonesUUIDDomainsDomainOK{}
}

/*PostZonesUUIDDomainsDomainOK handles this case with default header values.

OK
*/
type PostZonesUUIDDomainsDomainOK struct {
	Payload *models.ReturnMessage
}

func (o *PostZonesUUIDDomainsDomainOK) Error() string {
	return fmt.Sprintf("[POST /zones/{uuid}/domains/{domain}][%d] postZonesUuidDomainsDomainOK  %+v", 200, o.Payload)
}

func (o *PostZonesUUIDDomainsDomainOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReturnMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
