// Code generated by go-swagger; DO NOT EDIT.

package domains

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
	models "github.com/tokart/gandi-livedns-go/models"
)

// PatchDomainsDomainReader is a Reader for the PatchDomainsDomain structure.
type PatchDomainsDomainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchDomainsDomainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchDomainsDomainOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchDomainsDomainOK creates a PatchDomainsDomainOK with default headers values
func NewPatchDomainsDomainOK() *PatchDomainsDomainOK {
	return &PatchDomainsDomainOK{}
}

/*PatchDomainsDomainOK handles this case with default header values.

OK
*/
type PatchDomainsDomainOK struct {
	Payload []*models.ReturnMessage
}

func (o *PatchDomainsDomainOK) Error() string {
	return fmt.Sprintf("[PATCH /domains/{domain}][%d] patchDomainsDomainOK  %+v", 200, o.Payload)
}

func (o *PatchDomainsDomainOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PatchDomainsDomainBody patch domains domain body
swagger:model PatchDomainsDomainBody
*/
type PatchDomainsDomainBody struct {

	// zone uuid
	// Required: true
	ZoneUUID *string `json:"zone_uuid"`
}

// Validate validates this patch domains domain body
func (o *PatchDomainsDomainBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateZoneUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchDomainsDomainBody) validateZoneUUID(formats strfmt.Registry) error {

	if err := validate.Required("zone_uuid"+"."+"zone_uuid", "body", o.ZoneUUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchDomainsDomainBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchDomainsDomainBody) UnmarshalBinary(b []byte) error {
	var res PatchDomainsDomainBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
