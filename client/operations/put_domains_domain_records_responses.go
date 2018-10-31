// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	models "github.com/tokart/gandi-livedns-go/models"
)

// PutDomainsDomainRecordsReader is a Reader for the PutDomainsDomainRecords structure.
type PutDomainsDomainRecordsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDomainsDomainRecordsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutDomainsDomainRecordsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutDomainsDomainRecordsOK creates a PutDomainsDomainRecordsOK with default headers values
func NewPutDomainsDomainRecordsOK() *PutDomainsDomainRecordsOK {
	return &PutDomainsDomainRecordsOK{}
}

/*PutDomainsDomainRecordsOK handles this case with default header values.

OK
*/
type PutDomainsDomainRecordsOK struct {
	Payload *models.ReturnMessage
}

func (o *PutDomainsDomainRecordsOK) Error() string {
	return fmt.Sprintf("[PUT /domains/{domain}/records][%d] putDomainsDomainRecordsOK  %+v", 200, o.Payload)
}

func (o *PutDomainsDomainRecordsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReturnMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PutDomainsDomainRecordsBody put domains domain records body
swagger:model PutDomainsDomainRecordsBody
*/
type PutDomainsDomainRecordsBody struct {

	// items
	Items []*models.Record `json:"items"`
}

// Validate validates this put domains domain records body
func (o *PutDomainsDomainRecordsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutDomainsDomainRecordsBody) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(o.Items) { // not required
		return nil
	}

	for i := 0; i < len(o.Items); i++ {
		if swag.IsZero(o.Items[i]) { // not required
			continue
		}

		if o.Items[i] != nil {
			if err := o.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("record" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutDomainsDomainRecordsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutDomainsDomainRecordsBody) UnmarshalBinary(b []byte) error {
	var res PutDomainsDomainRecordsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
