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
	models "github.com/t0k4rt/gandi-livedns-go/models"
)

// GetZonesUUIDSnapshotsSnapshotUUIDReader is a Reader for the GetZonesUUIDSnapshotsSnapshotUUID structure.
type GetZonesUUIDSnapshotsSnapshotUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetZonesUUIDSnapshotsSnapshotUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetZonesUUIDSnapshotsSnapshotUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetZonesUUIDSnapshotsSnapshotUUIDOK creates a GetZonesUUIDSnapshotsSnapshotUUIDOK with default headers values
func NewGetZonesUUIDSnapshotsSnapshotUUIDOK() *GetZonesUUIDSnapshotsSnapshotUUIDOK {
	return &GetZonesUUIDSnapshotsSnapshotUUIDOK{}
}

/*GetZonesUUIDSnapshotsSnapshotUUIDOK handles this case with default header values.

OK
*/
type GetZonesUUIDSnapshotsSnapshotUUIDOK struct {
	Payload *GetZonesUUIDSnapshotsSnapshotUUIDOKBody
}

func (o *GetZonesUUIDSnapshotsSnapshotUUIDOK) Error() string {
	return fmt.Sprintf("[GET /zones/{uuid}/snapshots/{snapshot_uuid}][%d] getZonesUuidSnapshotsSnapshotUuidOK  %+v", 200, o.Payload)
}

func (o *GetZonesUUIDSnapshotsSnapshotUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetZonesUUIDSnapshotsSnapshotUUIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetZonesUUIDSnapshotsSnapshotUUIDOKBody get zones UUID snapshots snapshot UUID o k body
swagger:model GetZonesUUIDSnapshotsSnapshotUUIDOKBody
*/
type GetZonesUUIDSnapshotsSnapshotUUIDOKBody struct {

	// date created
	DateCreated string `json:"date_created,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`

	// zone data
	ZoneData []*models.Record `json:"zone_data"`

	// zone uuid
	ZoneUUID string `json:"zone_uuid,omitempty"`
}

// Validate validates this get zones UUID snapshots snapshot UUID o k body
func (o *GetZonesUUIDSnapshotsSnapshotUUIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateZoneData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetZonesUUIDSnapshotsSnapshotUUIDOKBody) validateZoneData(formats strfmt.Registry) error {

	if swag.IsZero(o.ZoneData) { // not required
		return nil
	}

	for i := 0; i < len(o.ZoneData); i++ {
		if swag.IsZero(o.ZoneData[i]) { // not required
			continue
		}

		if o.ZoneData[i] != nil {
			if err := o.ZoneData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getZonesUuidSnapshotsSnapshotUuidOK" + "." + "zone_data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetZonesUUIDSnapshotsSnapshotUUIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetZonesUUIDSnapshotsSnapshotUUIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetZonesUUIDSnapshotsSnapshotUUIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
