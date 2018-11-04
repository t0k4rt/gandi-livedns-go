// Code generated by go-swagger; DO NOT EDIT.

package domains

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/t0k4rt/gandi-livedns-go/models"
)

// GetDomainsDomainRecordsReader is a Reader for the GetDomainsDomainRecords structure.
type GetDomainsDomainRecordsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDomainsDomainRecordsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDomainsDomainRecordsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDomainsDomainRecordsOK creates a GetDomainsDomainRecordsOK with default headers values
func NewGetDomainsDomainRecordsOK() *GetDomainsDomainRecordsOK {
	return &GetDomainsDomainRecordsOK{}
}

/*GetDomainsDomainRecordsOK handles this case with default header values.

OK
*/
type GetDomainsDomainRecordsOK struct {
	Payload []*models.Record
}

func (o *GetDomainsDomainRecordsOK) Error() string {
	return fmt.Sprintf("[GET /domains/{domain}/records][%d] getDomainsDomainRecordsOK  %+v", 200, o.Payload)
}

func (o *GetDomainsDomainRecordsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}