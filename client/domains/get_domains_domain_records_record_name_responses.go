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

// GetDomainsDomainRecordsRecordNameReader is a Reader for the GetDomainsDomainRecordsRecordName structure.
type GetDomainsDomainRecordsRecordNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDomainsDomainRecordsRecordNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDomainsDomainRecordsRecordNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetDomainsDomainRecordsRecordNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetDomainsDomainRecordsRecordNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDomainsDomainRecordsRecordNameOK creates a GetDomainsDomainRecordsRecordNameOK with default headers values
func NewGetDomainsDomainRecordsRecordNameOK() *GetDomainsDomainRecordsRecordNameOK {
	return &GetDomainsDomainRecordsRecordNameOK{}
}

/*GetDomainsDomainRecordsRecordNameOK handles this case with default header values.

OK
*/
type GetDomainsDomainRecordsRecordNameOK struct {
	Payload []*models.Record
}

func (o *GetDomainsDomainRecordsRecordNameOK) Error() string {
	return fmt.Sprintf("[GET /domains/{domain}/records/{record_name}][%d] getDomainsDomainRecordsRecordNameOK  %+v", 200, o.Payload)
}

func (o *GetDomainsDomainRecordsRecordNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomainsDomainRecordsRecordNameBadRequest creates a GetDomainsDomainRecordsRecordNameBadRequest with default headers values
func NewGetDomainsDomainRecordsRecordNameBadRequest() *GetDomainsDomainRecordsRecordNameBadRequest {
	return &GetDomainsDomainRecordsRecordNameBadRequest{}
}

/*GetDomainsDomainRecordsRecordNameBadRequest handles this case with default header values.

Not OK
*/
type GetDomainsDomainRecordsRecordNameBadRequest struct {
	Payload *models.Return400
}

func (o *GetDomainsDomainRecordsRecordNameBadRequest) Error() string {
	return fmt.Sprintf("[GET /domains/{domain}/records/{record_name}][%d] getDomainsDomainRecordsRecordNameBadRequest  %+v", 400, o.Payload)
}

func (o *GetDomainsDomainRecordsRecordNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Return400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomainsDomainRecordsRecordNameDefault creates a GetDomainsDomainRecordsRecordNameDefault with default headers values
func NewGetDomainsDomainRecordsRecordNameDefault(code int) *GetDomainsDomainRecordsRecordNameDefault {
	return &GetDomainsDomainRecordsRecordNameDefault{
		_statusCode: code,
	}
}

/*GetDomainsDomainRecordsRecordNameDefault handles this case with default header values.

Unexpected error
*/
type GetDomainsDomainRecordsRecordNameDefault struct {
	_statusCode int

	Payload *models.Return40x
}

// Code gets the status code for the get domains domain records record name default response
func (o *GetDomainsDomainRecordsRecordNameDefault) Code() int {
	return o._statusCode
}

func (o *GetDomainsDomainRecordsRecordNameDefault) Error() string {
	return fmt.Sprintf("[GET /domains/{domain}/records/{record_name}][%d] GetDomainsDomainRecordsRecordName default  %+v", o._statusCode, o.Payload)
}

func (o *GetDomainsDomainRecordsRecordNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Return40x)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
