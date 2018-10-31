// Code generated by go-swagger; DO NOT EDIT.

package domains

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetDomainsReader is a Reader for the GetDomains structure.
type GetDomainsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDomainsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDomainsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDomainsOK creates a GetDomainsOK with default headers values
func NewGetDomainsOK() *GetDomainsOK {
	return &GetDomainsOK{}
}

/*GetDomainsOK handles this case with default header values.

OK
*/
type GetDomainsOK struct {
	Payload []*GetDomainsOKBodyItems0
}

func (o *GetDomainsOK) Error() string {
	return fmt.Sprintf("[GET /domains][%d] getDomainsOK  %+v", 200, o.Payload)
}

func (o *GetDomainsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetDomainsOKBodyItems0 get domains o k body items0
swagger:model GetDomainsOKBodyItems0
*/
type GetDomainsOKBodyItems0 struct {

	// fqdn
	Fqdn string `json:"fqdn,omitempty"`
}

// Validate validates this get domains o k body items0
func (o *GetDomainsOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetDomainsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDomainsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetDomainsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
