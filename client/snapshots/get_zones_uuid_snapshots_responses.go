// Code generated by go-swagger; DO NOT EDIT.

package snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetZonesUUIDSnapshotsReader is a Reader for the GetZonesUUIDSnapshots structure.
type GetZonesUUIDSnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetZonesUUIDSnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetZonesUUIDSnapshotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetZonesUUIDSnapshotsOK creates a GetZonesUUIDSnapshotsOK with default headers values
func NewGetZonesUUIDSnapshotsOK() *GetZonesUUIDSnapshotsOK {
	return &GetZonesUUIDSnapshotsOK{}
}

/*GetZonesUUIDSnapshotsOK handles this case with default header values.

OK
*/
type GetZonesUUIDSnapshotsOK struct {
	Payload []*GetZonesUUIDSnapshotsOKBodyItems0
}

func (o *GetZonesUUIDSnapshotsOK) Error() string {
	return fmt.Sprintf("[GET /zones/{uuid}/snapshots][%d] getZonesUuidSnapshotsOK  %+v", 200, o.Payload)
}

func (o *GetZonesUUIDSnapshotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetZonesUUIDSnapshotsOKBodyItems0 get zones UUID snapshots o k body items0
swagger:model GetZonesUUIDSnapshotsOKBodyItems0
*/
type GetZonesUUIDSnapshotsOKBodyItems0 struct {

	// date created
	DateCreated string `json:"date_created,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this get zones UUID snapshots o k body items0
func (o *GetZonesUUIDSnapshotsOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetZonesUUIDSnapshotsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetZonesUUIDSnapshotsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetZonesUUIDSnapshotsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}